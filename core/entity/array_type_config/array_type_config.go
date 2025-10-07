package array_type_config

import (
	"encoding/json"

	"github.com/nuzur/nem/core/entity/field_type_char_config"
	"github.com/nuzur/nem/core/entity/field_type_date_config"
	"github.com/nuzur/nem/core/entity/field_type_datetime_config"
	"github.com/nuzur/nem/core/entity/field_type_decimal_config"
	"github.com/nuzur/nem/core/entity/field_type_email_config"
	"github.com/nuzur/nem/core/entity/field_type_encrypted_config"
	"github.com/nuzur/nem/core/entity/field_type_enum_config"
	"github.com/nuzur/nem/core/entity/field_type_float_config"
	"github.com/nuzur/nem/core/entity/field_type_integer_config"
	"github.com/nuzur/nem/core/entity/field_type_phone_config"
	"github.com/nuzur/nem/core/entity/field_type_url_config"
	"github.com/nuzur/nem/core/entity/field_type_varchar_config"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"math/rand"

	"time"
)

type ArrayTypeConfig struct {
	Integer   field_type_integer_config.FieldTypeIntegerConfig     `json:"integer"`
	Float     field_type_float_config.FieldTypeFloatConfig         `json:"float"`
	Decimal   field_type_decimal_config.FieldTypeDecimalConfig     `json:"decimal"`
	Char      field_type_char_config.FieldTypeCharConfig           `json:"char"`
	Varchar   field_type_varchar_config.FieldTypeVarcharConfig     `json:"varchar"`
	Encrypted field_type_encrypted_config.FieldTypeEncryptedConfig `json:"encrypted"`
	URL       field_type_url_config.FieldTypeURLConfig             `json:"url"`
	Email     field_type_email_config.FieldTypeEmailConfig         `json:"email"`
	Phone     field_type_phone_config.FieldTypePhoneConfig         `json:"phone"`
	Date      field_type_date_config.FieldTypeDateConfig           `json:"date"`
	Datetime  field_type_datetime_config.FieldTypeDatetimeConfig   `json:"datetime"`
	Enum      field_type_enum_config.FieldTypeEnumConfig           `json:"enum"`
}

func (e ArrayTypeConfig) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e ArrayTypeConfig) EntityIdentifier() string {
	return "array_type_config"
}

func (e ArrayTypeConfig) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["integer"] = types.SingleDependantEntityFieldType
	res["float"] = types.SingleDependantEntityFieldType
	res["decimal"] = types.SingleDependantEntityFieldType
	res["char"] = types.SingleDependantEntityFieldType
	res["varchar"] = types.SingleDependantEntityFieldType
	res["encrypted"] = types.SingleDependantEntityFieldType
	res["url"] = types.SingleDependantEntityFieldType
	res["email"] = types.SingleDependantEntityFieldType
	res["phone"] = types.SingleDependantEntityFieldType
	res["date"] = types.SingleDependantEntityFieldType
	res["datetime"] = types.SingleDependantEntityFieldType
	res["enum"] = types.SingleDependantEntityFieldType
	return res
}

func (e ArrayTypeConfig) OrderedFieldIdentifiers() []string {
	res := []string{}
	res = append(res, "integer")
	res = append(res, "float")
	res = append(res, "decimal")
	res = append(res, "char")
	res = append(res, "varchar")
	res = append(res, "encrypted")
	res = append(res, "url")
	res = append(res, "email")
	res = append(res, "phone")
	res = append(res, "date")
	res = append(res, "datetime")
	res = append(res, "enum")

	return res
}

func (e ArrayTypeConfig) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e ArrayTypeConfig) IsDependant() bool {
	return true
}

func ArrayTypeConfigFromJSON(data json.RawMessage) ArrayTypeConfig {
	entity := ArrayTypeConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func ArrayTypeConfigSliceFromJSON(data json.RawMessage) []ArrayTypeConfig {
	entity := []ArrayTypeConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e ArrayTypeConfig) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ArrayTypeConfigToJSON", err)
	}
	return res
}

func ArrayTypeConfigToJSON(e ArrayTypeConfig) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ArrayTypeConfigToJSON", err)
	}
	return res
}

func ArrayTypeConfigSliceToJSON(e []ArrayTypeConfig) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ArrayTypeConfigSliceToJSON", err)
	}
	return res
}

func NewArrayTypeConfigWithRandomValues() ArrayTypeConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return ArrayTypeConfig{
		Integer:   field_type_integer_config.NewFieldTypeIntegerConfigWithRandomValues(),
		Float:     field_type_float_config.NewFieldTypeFloatConfigWithRandomValues(),
		Decimal:   field_type_decimal_config.NewFieldTypeDecimalConfigWithRandomValues(),
		Char:      field_type_char_config.NewFieldTypeCharConfigWithRandomValues(),
		Varchar:   field_type_varchar_config.NewFieldTypeVarcharConfigWithRandomValues(),
		Encrypted: field_type_encrypted_config.NewFieldTypeEncryptedConfigWithRandomValues(),
		URL:       field_type_url_config.NewFieldTypeURLConfigWithRandomValues(),
		Email:     field_type_email_config.NewFieldTypeEmailConfigWithRandomValues(),
		Phone:     field_type_phone_config.NewFieldTypePhoneConfigWithRandomValues(),
		Date:      field_type_date_config.NewFieldTypeDateConfigWithRandomValues(),
		Datetime:  field_type_datetime_config.NewFieldTypeDatetimeConfigWithRandomValues(),
		Enum:      field_type_enum_config.NewFieldTypeEnumConfigWithRandomValues(),
	}
}

func NewArrayTypeConfigSliceWithRandomValues(n int) []ArrayTypeConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []ArrayTypeConfig{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewArrayTypeConfigWithRandomValues())
	}
	return res
}
