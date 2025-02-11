package tcp_ip_connection_type_config

import (
	"encoding/json"

	"fmt"

	"nem/core/entity/types"

	"nem/core/randomvalues"

	"math/rand"

	"time"
)

type TcpIpConnectionTypeConfig struct {
	Hostname string `json:"hostname"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (e TcpIpConnectionTypeConfig) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e TcpIpConnectionTypeConfig) EntityIdentifier() string {
	return "tcp_ip_connection_type_config"
}

func (e TcpIpConnectionTypeConfig) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["hostname"] = types.StringFieldType
	res["port"] = types.StringFieldType
	res["username"] = types.StringFieldType
	res["password"] = types.StringFieldType
	return res
}

func (e TcpIpConnectionTypeConfig) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e TcpIpConnectionTypeConfig) IsDependant() bool {
	return true
}

func TcpIpConnectionTypeConfigFromJSON(data json.RawMessage) TcpIpConnectionTypeConfig {
	entity := TcpIpConnectionTypeConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func TcpIpConnectionTypeConfigSliceFromJSON(data json.RawMessage) []TcpIpConnectionTypeConfig {
	entity := []TcpIpConnectionTypeConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e TcpIpConnectionTypeConfig) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error TcpIpConnectionTypeConfigToJSON", err)
	}
	return res
}

func TcpIpConnectionTypeConfigToJSON(e TcpIpConnectionTypeConfig) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error TcpIpConnectionTypeConfigToJSON", err)
	}
	return res
}

func TcpIpConnectionTypeConfigSliceToJSON(e []TcpIpConnectionTypeConfig) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error TcpIpConnectionTypeConfigSliceToJSON", err)
	}
	return res
}

func NewTcpIpConnectionTypeConfigWithRandomValues() TcpIpConnectionTypeConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return TcpIpConnectionTypeConfig{
		Hostname: randomvalues.GetRandomStringValue(),
		Port:     randomvalues.GetRandomStringValue(),
		Username: randomvalues.GetRandomStringValue(),
		Password: randomvalues.GetRandomStringValue(),
	}
}

func NewTcpIpConnectionTypeConfigSliceWithRandomValues(n int) []TcpIpConnectionTypeConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []TcpIpConnectionTypeConfig{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewTcpIpConnectionTypeConfigWithRandomValues())
	}
	return res
}
