package ocpi221

type CredentialsRole struct {
	Role            Role            `json:"role"`
	BusinessDetails BusinessDetails `json:"business_details"`
	PartyID         string          `json:"party_id"`
	CountryCode     string          `json:"country_code"`
}

type Credentials struct {
	Token    string            `json:"token"`
	Url      string            `json:"url"`
	Roles    []CredentialsRole `json:"roles"`
	Metadata map[string]string `json:"metadata,omitempty"`
}
