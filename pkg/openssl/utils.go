package openssl

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)


func certFileToCertificateHelper(certPath string)(*x509.Certificate,error){
	caFile,err := ioutil.ReadFile(certPath)
	if err != nil{
		return nil, err
	}
	caBlock,_:= pem.Decode(caFile)

	cert,err := x509.ParseCertificate(caBlock.Bytes)
	if err != nil {
		return nil,fmt.Errorf("x509.ParseCertificated failed: %s",err.Error())
	}
	return cert, nil

}

