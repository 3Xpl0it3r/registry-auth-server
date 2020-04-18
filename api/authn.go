package api



type Authenticator interface {
	Authenticate(username,password string)(bool,error)
}
