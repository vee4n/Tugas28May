package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
)

func main() {
	url := "https://localhost:8081"

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	tlsConnectionState := resp.TLS
	if tlsConnectionState == nil {
		fmt.Println("Failed to retrieve TLS connection")
		os.Exit(1)
	}

	tlsVersion := tlsConnectionState.Version
	tlsVersionString := tlsVersionToString(tlsVersion)

	cipherSuite := tls.CipherSuiteName(tlsConnectionState.CipherSuite)

	var issuerOrg string
	if len(tlsConnectionState.PeerCertificates) > 0 {
		cert := tlsConnectionState.PeerCertificates[0]
		if len(cert.Issuer.Organization) > 0 {
			issuerOrg = cert.Issuer.Organization[0]
		} else {
			issuerOrg = "Unknown"
		}
	} else {
		issuerOrg = "Unknown"
	}

	fmt.Printf("TLS Version: %s\n", tlsVersionString)
	fmt.Printf("CipherSuite: %s\n", cipherSuite)
	fmt.Printf("Issuer Organization: %s\n", issuerOrg)
}

func tlsVersionToString(version uint16) string {
	switch version {
	case tls.VersionTLS10:
		return "TLS 1.0"
	case tls.VersionTLS11:
		return "TLS 1.1"
	case tls.VersionTLS12:
		return "TLS 1.2"
	case tls.VersionTLS13:
		return "TLS 1.3"
	default:
		return "Unknown"
	}
}
