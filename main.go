package main

import (
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"flag"
	"fmt"
	"os"
)

//     sys.stdout.writelines(b64encode(cert.fingerprint(algorithm=SHA256())).decode("utf-8"))

func main() {
	certPath := flag.String("cert_path", "", "Path to the certificate file")

	flag.Parse()

	file, err := os.ReadFile(*certPath)
	if err != nil {
		panic(err)
	}

	block, _ := pem.Decode(file)
	if block == nil {
		panic("failed to decode PEM block containing the certificate")
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		panic(err)
	}

	fingerPrint := sha256.Sum256(cert.Raw)

	stdencoding := base64.StdEncoding.EncodeToString(fingerPrint[:])
	fmt.Println(stdencoding)
}
