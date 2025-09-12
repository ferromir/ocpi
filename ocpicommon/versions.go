package ocpicommon

type VersionNumber string

const (
	VersionNumber2_0   VersionNumber = "2.0"
	VersionNumber2_1   VersionNumber = "2.1"
	VersionNumber2_1_1 VersionNumber = "2.1.1"
	VersionNumber2_2   VersionNumber = "2.2"
	VersionNumber2_2_1 VersionNumber = "2.2.1"
	VersionNumber2_3   VersionNumber = "2.3"
)

type InterfaceRole string

const (
	InterfaceRoleSender   InterfaceRole = "SENDER"
	InterfaceRoleReceiver InterfaceRole = "RECEIVER"
)

type ModuleID string

const (
	ModuleIDCDRs                       ModuleID = "cdrs"
	ModuleIDChargingProfiles           ModuleID = "chargingprofiles"
	ModuleIDCommands                   ModuleID = "commands"
	ModuleIDCredentialsAndRegistration ModuleID = "credentials"
	ModuleIDHubClientInfo              ModuleID = "hubclientinfo"
	ModuleIDLocations                  ModuleID = "locations"
	ModuleIDSessions                   ModuleID = "sessions"
	ModuleIDTariffs                    ModuleID = "tariffs"
	ModuleIDTokens                     ModuleID = "tokens"
)

type Version struct {
	Version VersionNumber `json:"version"`
	Url     string        `json:"url"`
}

type VersionDetails struct {
	Version   VersionNumber `json:"version"`
	Endpoints []Endpoint    `json:"endpoints"`
}

type Endpoint struct {
	Identifier ModuleID      `json:"identifier"`
	Role       InterfaceRole `json:"role"`
	Url        string        `json:"url"`
}
