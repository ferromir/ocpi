package ocpi221

type Role string

const (
	RoleCPO   Role = "CPO"
	RoleEMSP  Role = "EMSP"
	RoleHUB   Role = "HUB"
	RoleNAP   Role = "NAP"
	RoleNSP   Role = "NSP"
	RoleOTHER Role = "OTHER"
	RoleSCSP  Role = "SCSP"
)

type CredentialsRole struct {
	Role            Role            `json:"role"`
	BusinessDetails BusinessDetails `json:"business_details"`
	PartyID         string          `json:"party_id"`
	CountryCode     string          `json:"country_code"`
}

type Credentials struct {
	Token    string            `json:"token"`
	Url      string            `json:"url"`
	Roles    []Role            `json:"roles"`
	Metadata map[string]string `json:"metadata,omitempty"`
}
