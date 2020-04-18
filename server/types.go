package server


type authRequestInfo struct {
	// for authenticate
	Username,Password string
	Account string
	// for authorization
	Type string
	ResourceName string			// resource name
	Actions []string
}
