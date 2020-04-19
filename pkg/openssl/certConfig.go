package openssl

type SimpleCertConfig struct {
	CommonName                                string
	Country, Organization, OrganizationalUnit []string
	Expiration                                int64
	DNSName                                   []string
	IPAddress                                 []string
}

func NewDefaultSimpleCertConfig() *SimpleCertConfig {
	return &SimpleCertConfig{
		CommonName:         "l0calh0st.cn",
		Country:            []string{"China"},
		Organization:       []string{"docker"},
		OrganizationalUnit: []string{"docker"},
		Expiration:         3600,
		DNSName:            nil,
		IPAddress:          nil,
	}
}
