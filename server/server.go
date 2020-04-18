package server

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"l0calh0st.cn/registry-auth-server/api"
	"l0calh0st.cn/registry-auth-server/configs"
	"l0calh0st.cn/registry-auth-server/pkg/authn/static_authn"
	"l0calh0st.cn/registry-auth-server/pkg/autho/static_autho"
	"l0calh0st.cn/registry-auth-server/pkg/token"
	"log"
	"net/http"
)

type registryAuthServer struct {
	Authenticator api.Authenticator
	Authorization api.Authorization
	Token         api.IToken

	//
	address string
	pem,key string
}

// NewRegistryAuthServer return a instance of registry auth server
func NewRegistryAuthServer(cfg *configs.Configs)*registryAuthServer{
	tcfg := &api.TokenConfig{
		CertFile: cfg.Tls.Cert,
		KeyFile: cfg.Tls.Key,
		Issuer: cfg.Token.Issuer,
	}
	authnController := static_authn.NewStaticBasicAuthenticator()
	if authnController == nil{
		logrus.WithField("State", "Build Authenticator Failed").Errorf("Authenticator is nil")
		return nil
	}
	authoController := static_autho.NewStaticAclAuthorization()
	if authoController == nil{
		logrus.WithField("State", "Build NewStaticAclAuthorization Failed").Errorf("NewStaticAclAuthorization is nil")
		return nil
	}
	tokenController := token.NewTokenController(tcfg)
	if tokenController == nil{
		logrus.WithField("State", "Build NewTokenController Failed").Errorf("NewTokenController is nil")
		return nil
	}
	return &registryAuthServer{
		Authenticator: authnController,
		Authorization: authoController,
		Token: tokenController,
		pem:           cfg.Tls.Cert,
		key:           cfg.Tls.Key,
		address:	   cfg.Server.Address + ":" + cfg.Server.Port,
	}
}

//
func loggingMiddleware(next  http.Handler)http.Handler{
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Println(request.RequestURI)
		if request.Method == "OPTION"{
			// todo to deal cors
			return
		}
		next.ServeHTTP(writer,request)
	})
}

// run registry server
func(rs *registryAuthServer)Run(ctx context.Context)error{
	logrus.Info("Docker registry token server begin running.....\n")
	route := mux.NewRouter()
	route.Use(loggingMiddleware)
	var errchan chan error
	defer close(errchan)
	var err error
	if rs.key != "" && rs.pem != ""{
		logrus.Infof("Docker Registry Auth server Run as TLS Module\n")
		err = http.ListenAndServeTLS(rs.address,rs.pem, rs.key, route)
	} else {
		logrus.Infof("Docker registry auth server run as insecure module\n")
		err = http.ListenAndServe(rs.address, route)
	}
	select {
	case errchan <- err:
		return <- errchan
	case <-ctx.Done():
		return ctx.Err()
	}
}

func(rs *registryAuthServer)ServeHTTP(w http.ResponseWriter,r *http.Request){

	// parse request
	authReq,err := parseRequest(r)
	if err != nil{
		http.Error(w, fmt.Sprintf("Parse Request authRequest Failed:%s\n", err.Error()), http.StatusBadRequest)
		return
	} else {
		logrus.Debugf("[1].parse request Successfully\n")
	}

	// auth username and password
	if ok,err := rs.Authenticator.Authenticate(authReq.Username, authReq.Password);!ok{
		http.Error(w, fmt.Sprintf("[!]Authenticated Failed:%s\n",err.Error()), http.StatusUnauthorized)
		return
	} else {
		logrus.Debugf("[2]Auth username and password Successfully\n")
	}
	// Acl
	_,err = rs.Authorization.Authorize(authRequestHandler(authReq))
	if err != nil{
		http.Error(w, fmt.Sprintf("Authorization User Acl Faile: %s\n",err.Error()), http.StatusForbidden)
		return
	} else {
		logrus.Debugf("[3]AUthorization action successfully\n")
	}

	// token
	tokenstring,err := rs.Token.GenerateToken(scopeHandler(authReq))
	if err != nil{
		http.Error(w, fmt.Sprintf("Generate token for %s Failed: %s", authReq.Username, err), http.StatusInternalServerError)
		return
	} else {
		logrus.Debugf("[4]Token Generation Successfully:  %s\n", tokenstring)
	}
	w.Write([]byte(tokenstring))
	return
}