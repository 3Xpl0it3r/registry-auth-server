package openssl

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
)

func parseCertAndKeyHandler(certPath, keyPath string) (*x509.Certificate, *rsa.PrivateKey, error) {
	caRaw, err := ioutil.ReadFile(certPath)
	if err != nil {

	}
	caBlock, _ := pem.Decode(caRaw)
	cert, err := x509.ParseCertificate(caBlock.Bytes)
	if err != nil {

	}

	// parse private key
	keyRaw, err := ioutil.ReadFile(keyPath)
	if err != nil {

	}
	keyBlock, _ := pem.Decode(keyRaw)
	privateKey, err := x509.ParsePKCS1PrivateKey(keyBlock.Bytes)
	if err != nil {

	}

	return cert, privateKey, nil

}
