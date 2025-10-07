package field_type_encrypted_config

import (
	"encoding/json"

	"github.com/gofrs/uuid"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"

	"time"
)

type FieldTypeEncryptedConfig struct {
	MinSize         int64          `json:"min_size"`
	MaxSize         int64          `json:"max_size"`
	RegexValidation *string        `json:"regex_validation"`
	EncryptionType  EncryptionType `json:"encryption_type"`
	UseSalt         bool           `json:"use_salt"`
	SaltFieldUUIDs  []uuid.UUID    `json:"salt_field_uuids"`
}

func (e FieldTypeEncryptedConfig) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e FieldTypeEncryptedConfig) EntityIdentifier() string {
	return "field_type_encrypted_config"
}

func (e FieldTypeEncryptedConfig) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["min_size"] = types.IntFieldType
	res["max_size"] = types.IntFieldType
	res["regex_validation"] = types.StringFieldType
	res["encryption_type"] = types.SingleEnumFieldType
	res["use_salt"] = types.BooleanFieldType
	res["salt_field_uuids"] = types.ArrayFieldType
	return res
}

func (e FieldTypeEncryptedConfig) OrderedFieldIdentifiers() []string {
	res := []string{}
	res = append(res, "min_size")
	res = append(res, "max_size")
	res = append(res, "regex_validation")
	res = append(res, "encryption_type")
	res = append(res, "use_salt")
	res = append(res, "salt_field_uuids")

	return res
}

func (e FieldTypeEncryptedConfig) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["salt_field_uuids"] = types.UUIDFieldType
	return res
}

func (e FieldTypeEncryptedConfig) IsDependant() bool {
	return true
}

func FieldTypeEncryptedConfigFromJSON(data json.RawMessage) FieldTypeEncryptedConfig {
	entity := FieldTypeEncryptedConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func FieldTypeEncryptedConfigSliceFromJSON(data json.RawMessage) []FieldTypeEncryptedConfig {
	entity := []FieldTypeEncryptedConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e FieldTypeEncryptedConfig) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldTypeEncryptedConfigToJSON", err)
	}
	return res
}

func FieldTypeEncryptedConfigToJSON(e FieldTypeEncryptedConfig) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldTypeEncryptedConfigToJSON", err)
	}
	return res
}

func FieldTypeEncryptedConfigSliceToJSON(e []FieldTypeEncryptedConfig) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldTypeEncryptedConfigSliceToJSON", err)
	}
	return res
}

func NewFieldTypeEncryptedConfigWithRandomValues() FieldTypeEncryptedConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return FieldTypeEncryptedConfig{
		MinSize:         randomvalues.GetRandomIntValue(),
		MaxSize:         randomvalues.GetRandomIntValue(),
		RegexValidation: randomvalues.GetRandomStringValuePtr(),
		EncryptionType:  randomvalues.GetRandomOptionValue[EncryptionType](2),
		UseSalt:         randomvalues.GetRandomBoolValue(),
		SaltFieldUUIDs:  []uuid.UUID{},
	}
}

func NewFieldTypeEncryptedConfigSliceWithRandomValues(n int) []FieldTypeEncryptedConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []FieldTypeEncryptedConfig{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewFieldTypeEncryptedConfigWithRandomValues())
	}
	return res
}
