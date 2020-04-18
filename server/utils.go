package server

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"l0calh0st.cn/registry-auth-server/api"
	"l0calh0st.cn/registry-auth-server/configs"
	"net/http"
	"strings"
)

func parseRequest(request *http.Request) (*authRequestInfo, error) {
	username, password, ok := request.BasicAuth()
	if !ok {
		return nil, fmt.Errorf("Username or password must be supplied\n")
	}
	account := request.FormValue("account")
	if account != username {
		logrus.Warningf("Username:%s and Account:%s is not same", username, account)
	}

	service := request.FormValue("service")
	authReq := &authRequestInfo{
		Username:     username,
		Password:     password,
		Service:      service,
		Account:      account,
		Actions:      nil,
		Type:         "",
		ResourceName: "",
	}
	parts := strings.Split(request.URL.Query().Get("scope"), ":")

	if len(parts) > 0 {
		authReq.Type = parts[0]
	}
	if len(parts) > 1 {
		authReq.ResourceName = parts[1]
	}
	if len(parts) > 2 {
		authReq.Actions = strings.Split(parts[2], ",")
	}
	return authReq, nil
}

func authRequestHandler(info *authRequestInfo) *api.AuthRequestInfo {
	authReq := &api.AuthRequestInfo{
		Account:      info.Account,
		Actions:      info.Actions,
		ResourceName: info.ResourceName,
		Type:         info.Type,
	}
	return authReq
}

func generateTokenClaimHandler(info *authRequestInfo) *api.TokenClaim {
	sr := &api.TokenClaim{
		Type:    info.Type,
		Account: info.Account,
		Name:    info.ResourceName,
		Actions: info.Actions,
		Service: info.Service,
	}
	return sr
}

func defaultCertAndKey(cfg *configs.Configs) {

	//if cfg.Server.Domain == ""{
	//	cfg.Server.Domain = "reg.example.com"
	//}
	//// create ca
	//
	//
	//if err != nil {
	//	return
	//}
	//
	//caCertPath := "../static/ca.pem"
	//caKeyPath := "../static/ca.key"
	//ioutil.WriteFile(caCertPath, caCertByte, 0777)
	//ioutil.WriteFile(caKeyPath, caKeyByte, 0777)
	//
	//serverCert := &x509.Certificate{
	//	SerialNumber: big.NewInt(1653),
	//	Subject: pkix.Name{
	//		Country: []string{"China"},
	//		Organization: []string{"example"},
	//		OrganizationalUnit: []string{"container"},
	//	},
	//	NotBefore: time.Now(),
	//	NotAfter: time.Now().AddDate(2, 0,0),
	//	SubjectKeyId: []byte{1,2,3,4,5},
	//	ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
	//	KeyUsage: x509.KeyUsageDigitalSignature| x509.KeyUsageCertSign,
	//	IsCA: false,
	//}
	//
	//serverPrivateKey,_ := rsa.GenerateKey(rand.Reader, 1024)
	//serverPublickey := &serverPrivateKey.PublicKey
	//serverCertByte,err := x509.CreateCertificate(rand.Reader,ca,serverCert, serverPublickey, serverPrivateKey)
	//if err != nil{
	//	return
	//}
	//serverPrivateKeyByte := x509.MarshalPKCS1PrivateKey(serverPrivateKey)
	//
	//serverCertPath := "../static/server.pem"
	//serverPriKeyParh := "../static/server.key"
	//ioutil.WriteFile(serverCertPath, serverCertByte, 0777)
	//ioutil.WriteFile(serverPriKeyParh, serverPrivateKeyByte, 0777)
}
