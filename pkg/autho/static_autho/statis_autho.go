package static_autho

import (
	"github.com/sirupsen/logrus"
	"l0calh0st.cn/registry-auth-server/api"
)

type staticAclAuthorization struct {

}

func NewStaticAclAuthorization()*staticAclAuthorization{
	logrus.Info("Load Authorization Controller Successfully\n")
	return &staticAclAuthorization{}
}

func(sa *staticAclAuthorization)Authorize(req *api.AuthRequestInfo)([]string,error){
	return nil,nil
}
