package openssl

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"github.com/sirupsen/logrus"
	"math/big"
	"os"
	"time"
)

type AbstractOpenssl interface {
	Generate()error
	SaveToFile(cert,key string)error

}

type simpleRootCert struct {
	cert []byte
	key []byte

	config *SimpleCertConfig
}

func NewSimpleRootCa(cfg *SimpleCertConfig)*simpleRootCert{
	return &simpleRootCert{config: cfg}
}

func(c *simpleRootCert)Generate()error{
	ca := &x509.Certificate{
		SerialNumber: big.NewInt(1653),
		Subject: pkix.Name{
			Country: c.config.Country,
			Organization: c.config.Organization,
			OrganizationalUnit: c.config.OrganizationalUnit,
		},
		NotBefore: time.Now(),
		NotAfter: time.Now().AddDate(2,0,0),
		SubjectKeyId: []byte{1,2,3,4,5},
		BasicConstraintsValid: true,
		IsCA: true,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage: x509.KeyUsageDigitalSignature|x509.KeyUsageCertSign,
	}
	caPriKey,_ := rsa.GenerateKey(rand.Reader, 1024)

	caPubKey := &caPriKey.PublicKey
	caCertByte,err := x509.CreateCertificate(rand.Reader,ca, ca, caPubKey, caPriKey)
	if err != nil{
		return err
	}
	caKeyByte := x509.MarshalPKCS1PrivateKey(caPriKey)
	c.key =caKeyByte
	c.cert = caCertByte
	logrus.Infof("Generate Ca Cert and Key Successfully\n")

	return nil
}

func(c *simpleRootCert)SaveToFile(cert,key string)error{
	if c.key == nil || c.cert	 == nil{
		return fmt.Errorf("simple root ca is not craeted")
	}
	if _,err := os.Stat(cert);err == nil{
		os.Rename(key, key+".bak")
	}
	certFp,err := os.Create(cert)
	if err != nil {
		return fmt.Errorf("cannot open %s Reason: %s\n", cert, err.Error())
	}
	defer certFp.Close()

	if err := pem.Encode(certFp, &pem.Block{Type: "CERTIFICATE", Bytes:   c.cert});err != nil{
		return fmt.Errorf("pem.Enocee certificate %s failed: %s", cert, err.Error())
	}

	if _,err := os.Stat(key);err == nil{
		os.Rename(key, key+".bak")
	}
	keyFp,err := os.Create(key)
	if err != nil {
		return fmt.Errorf("cannot open %s Reason: %s", key, err.Error())
	}
	defer keyFp.Close()
	if err := pem.Encode(keyFp, &pem.Block{Type: "PRIVATE KEY",Bytes: c.key});err != nil{
		return fmt.Errorf("pem.Encode private key %s failed:%s",key, err.Error())
	}
	logrus.Infof("Save Ca cert:%s\tKey:%s\tSuccessfully\n", cert, key)
	return nil
}