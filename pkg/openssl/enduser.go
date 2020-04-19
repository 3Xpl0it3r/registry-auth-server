package openssl

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"math/big"
	"net"
	"os"
	"time"
)

type simpleEndOfUserCert struct {
	config *SimpleCertConfig
	cert   []byte
	key    []byte
	//
	caCert *x509.Certificate
	caKey  *rsa.PrivateKey
}

func NewSimpleEndOfUserCert(config *SimpleCertConfig, caCert, caKey string) *simpleEndOfUserCert {
	cert, priK, err := parseCertAndKeyHandler(caCert, caKey)

	if err != nil {
		logrus.WithField("Stage", "Load CaFile").Errorln(err.Error())
		return nil
	}
	return &simpleEndOfUserCert{config: config, caCert: cert, caKey: priK}
}

func (c *simpleEndOfUserCert) Generate() error {
	ips := []net.IP{net.ParseIP("127.0.0.1")}
	if len(c.config.IPAddress) > 0 {
		for _, ip := range c.config.IPAddress {
			ips = append(ips, net.ParseIP(ip))
		}
	}
	serialNumber, _ := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128))
	cert := &x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			CommonName:         "docker registry",
			Country:            c.config.Country,
			Organization:       c.config.Organization,
			OrganizationalUnit: c.config.OrganizationalUnit,
		},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(10, 0, 0),
		SubjectKeyId: []byte{1, 2, 3, 4, 6},
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:     x509.KeyUsageDigitalSignature,
		DNSNames:     c.config.DNSName,
		IPAddresses:  ips,
	}
	privateKey, _ := rsa.GenerateKey(rand.Reader, 1024)
	publicKey := &privateKey.PublicKey

	certByte, err := x509.CreateCertificate(rand.Reader, cert, c.caCert, publicKey, c.caKey)
	if err != nil {
		log.Println("create cert2failed")
		return err
	}
	privateKeyByte := x509.MarshalPKCS1PrivateKey(privateKey)

	c.key = privateKeyByte
	c.cert = certByte
	return nil
}

func (c *simpleEndOfUserCert) SaveToFile(cert, key string) error {
	if c.key == nil || c.cert == nil {
		return fmt.Errorf("simple EndOfUser cert is not craeted")
	}
	if _, err := os.Stat(cert); err == nil {
		os.Rename(cert, cert+".bak")
	}
	if _, err := os.Stat(cert); err == nil {
		os.Rename(key, key+".bak")
	}
	certFp, err := os.Create(cert)
	if err != nil {
		return fmt.Errorf("cannot open %s Reason: %s\n", cert, err.Error())
	}
	defer certFp.Close()

	if err := pem.Encode(certFp, &pem.Block{Type: "CERTIFICATE", Bytes: c.cert}); err != nil {
		return fmt.Errorf("pem.Enocee certificate %s failed: %s", cert, err.Error())
	}

	if _, err := os.Stat(key); err == nil {
		os.Rename(key, key+".bak")
	}
	keyFp, err := os.Create(key)
	if err != nil {
		return fmt.Errorf("cannot open %s Reason: %s", key, err.Error())
	}
	defer keyFp.Close()
	if err := pem.Encode(keyFp, &pem.Block{Type: "PRIVATE KEY", Bytes: c.key}); err != nil {
		return fmt.Errorf("pem.Encode private key %s failed:%s", key, err.Error())
	}
	logrus.Infof("Save Ca cert:%s\tKey:%s\tSuccessfully\n", cert, key)
	return nil
}
