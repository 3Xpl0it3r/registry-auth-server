package api

type AuthRequestInfo struct {
	Account string		// request user
	Type 	string    // type in scope
	ResourceName		string   // the name of project/resource requested
	Actions []string	// actions in scope
}

type Authorization interface {
	Authorize(info *AuthRequestInfo)([]string,error)
}
