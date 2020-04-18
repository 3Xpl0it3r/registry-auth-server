package openssl

type SimpleCertConfig struct {
	Country, Organization, OrganizationalUnit []string
	Expiration int64
	DNSName []string
	IPAddress []string
}


func NewDefaultSimpleCertConfig()*SimpleCertConfig{
	return &SimpleCertConfig{
		Country:            []string{"China"},
		Organization:       []string{"docker"},
		OrganizationalUnit: []string{"docker"},
		Expiration:         3600,
		DNSName:            nil,
		IPAddress:          nil,
	}
}
