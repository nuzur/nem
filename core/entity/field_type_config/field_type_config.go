package field_type_config

import (
	"encoding/json"

	"nem/core/entity/field_type_array_config"
	"nem/core/entity/field_type_char_config"
	"nem/core/entity/field_type_date_config"
	"nem/core/entity/field_type_datetime_config"
	"nem/core/entity/field_type_decimal_config"
	"nem/core/entity/field_type_email_config"
	"nem/core/entity/field_type_encrypted_config"
	"nem/core/entity/field_type_enum_config"
	"nem/core/entity/field_type_file_config"
	"nem/core/entity/field_type_float_config"
	"nem/core/entity/field_type_integer_config"
	"nem/core/entity/field_type_json_config"
	"nem/core/entity/field_type_phone_config"
	"nem/core/entity/field_type_slug_config"
	"nem/core/entity/field_type_text_config"
	"nem/core/entity/field_type_url_config"
	"nem/core/entity/field_type_varchar_config"

	"fmt"

	"nem/core/entity/types"

	"nem/core/randomvalues"

	"math/rand"

	"time"
)

type FieldTypeConfig struct {
	UUID      string                                               `json:"uuid"`
	Integer   field_type_integer_config.FieldTypeIntegerConfig     `json:"integer"`
	Float     field_type_float_config.FieldTypeFloatConfig         `json:"float"`
	Decimal   field_type_decimal_config.FieldTypeDecimalConfig     `json:"decimal"`
	Boolean   string                                               `json:"boolean"`
	Char      field_type_char_config.FieldTypeCharConfig           `json:"char"`
	Varchar   field_type_varchar_config.FieldTypeVarcharConfig     `json:"varchar"`
	Text      field_type_text_config.FieldTypeTextConfig           `json:"text"`
	Encrypted field_type_encrypted_config.FieldTypeEncryptedConfig `json:"encrypted"`
	Email     field_type_email_config.FieldTypeEmailConfig         `json:"email"`
	Phone     field_type_phone_config.FieldTypePhoneConfig         `json:"phone"`
	URL       field_type_url_config.FieldTypeURLConfig             `json:"url"`
	Location  string                                               `json:"location"`
	Color     string                                               `json:"color"`
	RichText  string                                               `json:"rich_text"`
	Code      string                                               `json:"code"`
	Markdown  string                                               `json:"markdown"`
	File      field_type_file_config.FieldTypeFileConfig           `json:"file"`
	Image     field_type_file_config.FieldTypeFileConfig           `json:"image"`
	Video     field_type_file_config.FieldTypeFileConfig           `json:"video"`
	Audio     field_type_file_config.FieldTypeFileConfig           `json:"audio"`
	Enum      field_type_enum_config.FieldTypeEnumConfig           `json:"enum"`
	JSON      field_type_json_config.FieldTypeJSONConfig           `json:"json"`
	Array     field_type_array_config.FieldTypeArrayConfig         `json:"array"`
	Date      field_type_date_config.FieldTypeDateConfig           `json:"date"`
	Datetime  field_type_datetime_config.FieldTypeDatetimeConfig   `json:"datetime"`
	Time      string                                               `json:"time"`
	Slug      field_type_slug_config.FieldTypeSlugConfig           `json:"slug"`
}

func (e FieldTypeConfig) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e FieldTypeConfig) EntityIdentifier() string {
	return "field_type_config"
}

func (e FieldTypeConfig) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["uuid"] = types.RawJSONFieldType
	res["integer"] = types.SingleDependantEntityFieldType
	res["float"] = types.SingleDependantEntityFieldType
	res["decimal"] = types.SingleDependantEntityFieldType
	res["boolean"] = types.RawJSONFieldType
	res["char"] = types.SingleDependantEntityFieldType
	res["varchar"] = types.SingleDependantEntityFieldType
	res["text"] = types.SingleDependantEntityFieldType
	res["encrypted"] = types.SingleDependantEntityFieldType
	res["email"] = types.SingleDependantEntityFieldType
	res["phone"] = types.SingleDependantEntityFieldType
	res["url"] = types.SingleDependantEntityFieldType
	res["location"] = types.RawJSONFieldType
	res["color"] = types.RawJSONFieldType
	res["rich_text"] = types.RawJSONFieldType
	res["code"] = types.RawJSONFieldType
	res["markdown"] = types.RawJSONFieldType
	res["file"] = types.SingleDependantEntityFieldType
	res["image"] = types.SingleDependantEntityFieldType
	res["video"] = types.SingleDependantEntityFieldType
	res["audio"] = types.SingleDependantEntityFieldType
	res["enum"] = types.SingleDependantEntityFieldType
	res["json"] = types.SingleDependantEntityFieldType
	res["array"] = types.SingleDependantEntityFieldType
	res["date"] = types.SingleDependantEntityFieldType
	res["datetime"] = types.SingleDependantEntityFieldType
	res["time"] = types.RawJSONFieldType
	res["slug"] = types.SingleDependantEntityFieldType
	return res
}

func (e FieldTypeConfig) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e FieldTypeConfig) IsDependant() bool {
	return true
}

func FieldTypeConfigFromJSON(data json.RawMessage) FieldTypeConfig {
	entity := FieldTypeConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func FieldTypeConfigSliceFromJSON(data json.RawMessage) []FieldTypeConfig {
	entity := []FieldTypeConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e FieldTypeConfig) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldTypeConfigToJSON", err)
	}
	return res
}

func FieldTypeConfigToJSON(e FieldTypeConfig) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldTypeConfigToJSON", err)
	}
	return res
}

func FieldTypeConfigSliceToJSON(e []FieldTypeConfig) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldTypeConfigSliceToJSON", err)
	}
	return res
}

func NewFieldTypeConfigWithRandomValues() FieldTypeConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return FieldTypeConfig{
		UUID:      randomvalues.GetRandomRawJSONValue(),
		Integer:   field_type_integer_config.NewFieldTypeIntegerConfigWithRandomValues(),
		Float:     field_type_float_config.NewFieldTypeFloatConfigWithRandomValues(),
		Decimal:   field_type_decimal_config.NewFieldTypeDecimalConfigWithRandomValues(),
		Boolean:   randomvalues.GetRandomRawJSONValue(),
		Char:      field_type_char_config.NewFieldTypeCharConfigWithRandomValues(),
		Varchar:   field_type_varchar_config.NewFieldTypeVarcharConfigWithRandomValues(),
		Text:      field_type_text_config.NewFieldTypeTextConfigWithRandomValues(),
		Encrypted: field_type_encrypted_config.NewFieldTypeEncryptedConfigWithRandomValues(),
		Email:     field_type_email_config.NewFieldTypeEmailConfigWithRandomValues(),
		Phone:     field_type_phone_config.NewFieldTypePhoneConfigWithRandomValues(),
		URL:       field_type_url_config.NewFieldTypeURLConfigWithRandomValues(),
		Location:  randomvalues.GetRandomRawJSONValue(),
		Color:     randomvalues.GetRandomRawJSONValue(),
		RichText:  randomvalues.GetRandomRawJSONValue(),
		Code:      randomvalues.GetRandomRawJSONValue(),
		Markdown:  randomvalues.GetRandomRawJSONValue(),
		File:      field_type_file_config.NewFieldTypeFileConfigWithRandomValues(),
		Image:     field_type_file_config.NewFieldTypeFileConfigWithRandomValues(),
		Video:     field_type_file_config.NewFieldTypeFileConfigWithRandomValues(),
		Audio:     field_type_file_config.NewFieldTypeFileConfigWithRandomValues(),
		Enum:      field_type_enum_config.NewFieldTypeEnumConfigWithRandomValues(),
		JSON:      field_type_json_config.NewFieldTypeJSONConfigWithRandomValues(),
		Array:     field_type_array_config.NewFieldTypeArrayConfigWithRandomValues(),
		Date:      field_type_date_config.NewFieldTypeDateConfigWithRandomValues(),
		Datetime:  field_type_datetime_config.NewFieldTypeDatetimeConfigWithRandomValues(),
		Time:      randomvalues.GetRandomRawJSONValue(),
		Slug:      field_type_slug_config.NewFieldTypeSlugConfigWithRandomValues(),
	}
}

func NewFieldTypeConfigSliceWithRandomValues(n int) []FieldTypeConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []FieldTypeConfig{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewFieldTypeConfigWithRandomValues())
	}
	return res
}
