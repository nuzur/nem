package tcp_ip_ssh_connection_type_config

import (
	"encoding/json"

	"fmt"

	"nem/core/entity/types"

	"nem/core/randomvalues"

	"math/rand"

	"time"
)

type TcpIpSshConnectionTypeConfig struct {
	SshHostname   string `json:"ssh_hostname"`
	SshUsername   string `json:"ssh_username"`
	SshPassword   string `json:"ssh_password"`
	SshKeyFile    string `json:"ssh_key_file"`
	UseSshKeyFile bool   `json:"use_ssh_key_file"`
	Hostname      string `json:"hostname"`
	Port          string `json:"port"`
	Username      string `json:"username"`
	Password      string `json:"password"`
}

func (e TcpIpSshConnectionTypeConfig) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e TcpIpSshConnectionTypeConfig) EntityIdentifier() string {
	return "tcp_ip_ssh_connection_type_config"
}

func (e TcpIpSshConnectionTypeConfig) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["ssh_hostname"] = types.StringFieldType
	res["ssh_username"] = types.StringFieldType
	res["ssh_password"] = types.StringFieldType
	res["ssh_key_file"] = types.StringFieldType
	res["use_ssh_key_file"] = types.BooleanFieldType
	res["hostname"] = types.StringFieldType
	res["port"] = types.StringFieldType
	res["username"] = types.StringFieldType
	res["password"] = types.StringFieldType
	return res
}

func (e TcpIpSshConnectionTypeConfig) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e TcpIpSshConnectionTypeConfig) IsDependant() bool {
	return true
}

func TcpIpSshConnectionTypeConfigFromJSON(data json.RawMessage) TcpIpSshConnectionTypeConfig {
	entity := TcpIpSshConnectionTypeConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func TcpIpSshConnectionTypeConfigSliceFromJSON(data json.RawMessage) []TcpIpSshConnectionTypeConfig {
	entity := []TcpIpSshConnectionTypeConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e TcpIpSshConnectionTypeConfig) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error TcpIpSshConnectionTypeConfigToJSON", err)
	}
	return res
}

func TcpIpSshConnectionTypeConfigToJSON(e TcpIpSshConnectionTypeConfig) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error TcpIpSshConnectionTypeConfigToJSON", err)
	}
	return res
}

func TcpIpSshConnectionTypeConfigSliceToJSON(e []TcpIpSshConnectionTypeConfig) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error TcpIpSshConnectionTypeConfigSliceToJSON", err)
	}
	return res
}

func NewTcpIpSshConnectionTypeConfigWithRandomValues() TcpIpSshConnectionTypeConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return TcpIpSshConnectionTypeConfig{
		SshHostname:   randomvalues.GetRandomStringValue(),
		SshUsername:   randomvalues.GetRandomStringValue(),
		SshPassword:   randomvalues.GetRandomStringValue(),
		SshKeyFile:    randomvalues.GetRandomStringValue(),
		UseSshKeyFile: randomvalues.GetRandomBoolValue(),
		Hostname:      randomvalues.GetRandomStringValue(),
		Port:          randomvalues.GetRandomStringValue(),
		Username:      randomvalues.GetRandomStringValue(),
		Password:      randomvalues.GetRandomStringValue(),
	}
}

func NewTcpIpSshConnectionTypeConfigSliceWithRandomValues(n int) []TcpIpSshConnectionTypeConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []TcpIpSshConnectionTypeConfig{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewTcpIpSshConnectionTypeConfigWithRandomValues())
	}
	return res
}
