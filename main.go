package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"net"
	"time"
)

func main() {
	domainPtr := flag.String("d", "", "domain to inspect")
	flag.Parse()

	if domainPtr == nil {
		panic("domain cannot be empty")
	}

	domain := *domainPtr

	conn, err := tls.Dial(
		"tcp",
		net.JoinHostPort(domain, "443"),
		&tls.Config{InsecureSkipVerify: true},
	)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	s := conn.ConnectionState()
	fmt.Printf("\nTLS version: %v\n", s.Version)
	fmt.Printf("Preffered cipher suites: %v\n", s.CipherSuite)
	fmt.Printf("\nCertificate chain:\n")

	for _, c := range s.PeerCertificates {
		printCertificate(c)
		fmt.Println()
	}
}

func printCertificate(c *x509.Certificate) {
	fmt.Printf("Subject: %v\n", c.Subject)
	fmt.Printf("Subject alternatime names (SAN): %v\n", c.DNSNames)
	fmt.Printf("Issuer: %v\n", c.Issuer)
	fmt.Printf("Valid until: %v\n", c.NotAfter.Format(time.RFC1123))
	fmt.Printf("Issued at: %v\n", c.NotBefore.Format(time.RFC1123))
	fmt.Printf("Email addresses: %v\n", c.EmailAddresses)
	fmt.Printf("Version: %v\n", c.Version)
	fmt.Printf("Serial Number: %v\n", c.SerialNumber)
	fmt.Printf("Public key algorithm: %v\n", c.PublicKeyAlgorithm)
	fmt.Printf("Is CA: %v\n", c.IsCA)
}
