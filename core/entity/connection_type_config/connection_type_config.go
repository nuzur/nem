package connection_type_config

import (
	"encoding/json"

	"github.com/nuzur/nem/core/entity/tcp_ip_connection_type_config"
	"github.com/nuzur/nem/core/entity/tcp_ip_ssh_connection_type_config"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"math/rand"

	"time"
)

type ConnectionTypeConfig struct {
	TcpIp    tcp_ip_connection_type_config.TcpIpConnectionTypeConfig        `json:"tcp_ip"`
	TcpIpSsh tcp_ip_ssh_connection_type_config.TcpIpSshConnectionTypeConfig `json:"tcp_ip_ssh"`
}

func (e ConnectionTypeConfig) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e ConnectionTypeConfig) EntityIdentifier() string {
	return "connection_type_config"
}

func (e ConnectionTypeConfig) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["tcp_ip"] = types.SingleDependantEntityFieldType
	res["tcp_ip_ssh"] = types.SingleDependantEntityFieldType
	return res
}

func (e ConnectionTypeConfig) OrderedFieldIdentifiers() []string {
	res := []string{}
	res = append(res, "tcp_ip")
	res = append(res, "tcp_ip_ssh")

	return res
}

func (e ConnectionTypeConfig) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e ConnectionTypeConfig) IsDependant() bool {
	return true
}

func ConnectionTypeConfigFromJSON(data json.RawMessage) ConnectionTypeConfig {
	entity := ConnectionTypeConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func ConnectionTypeConfigSliceFromJSON(data json.RawMessage) []ConnectionTypeConfig {
	entity := []ConnectionTypeConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e ConnectionTypeConfig) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ConnectionTypeConfigToJSON", err)
	}
	return res
}

func ConnectionTypeConfigToJSON(e ConnectionTypeConfig) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ConnectionTypeConfigToJSON", err)
	}
	return res
}

func ConnectionTypeConfigSliceToJSON(e []ConnectionTypeConfig) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ConnectionTypeConfigSliceToJSON", err)
	}
	return res
}

func NewConnectionTypeConfigWithRandomValues() ConnectionTypeConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return ConnectionTypeConfig{
		TcpIp:    tcp_ip_connection_type_config.NewTcpIpConnectionTypeConfigWithRandomValues(),
		TcpIpSsh: tcp_ip_ssh_connection_type_config.NewTcpIpSshConnectionTypeConfigWithRandomValues(),
	}
}

func NewConnectionTypeConfigSliceWithRandomValues(n int) []ConnectionTypeConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []ConnectionTypeConfig{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewConnectionTypeConfigWithRandomValues())
	}
	return res
}
