package ocpi211

type Credentials struct {
	Token           string            `json:"token"`
	Url             string            `json:"url"`
	BusinessDetails BusinessDetails   `json:"business_details"`
	PartyID         string            `json:"party_id"`
	CountryCode     string            `json:"country_code"`
	Metadata        map[string]string `json:"metadata,omitempty"`
}
