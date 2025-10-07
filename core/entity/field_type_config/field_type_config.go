package field_type_config

import (
	"encoding/json"

	"github.com/nuzur/nem/core/entity/field_type_array_config"
	"github.com/nuzur/nem/core/entity/field_type_char_config"
	"github.com/nuzur/nem/core/entity/field_type_date_config"
	"github.com/nuzur/nem/core/entity/field_type_datetime_config"
	"github.com/nuzur/nem/core/entity/field_type_decimal_config"
	"github.com/nuzur/nem/core/entity/field_type_email_config"
	"github.com/nuzur/nem/core/entity/field_type_encrypted_config"
	"github.com/nuzur/nem/core/entity/field_type_enum_config"
	"github.com/nuzur/nem/core/entity/field_type_file_config"
	"github.com/nuzur/nem/core/entity/field_type_float_config"
	"github.com/nuzur/nem/core/entity/field_type_integer_config"
	"github.com/nuzur/nem/core/entity/field_type_json_config"
	"github.com/nuzur/nem/core/entity/field_type_phone_config"
	"github.com/nuzur/nem/core/entity/field_type_slug_config"
	"github.com/nuzur/nem/core/entity/field_type_text_config"
	"github.com/nuzur/nem/core/entity/field_type_url_config"
	"github.com/nuzur/nem/core/entity/field_type_varchar_config"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"

	"time"
)

type FieldTypeConfig struct {
	UUID      *string                                              `json:"uuid"`
	Integer   field_type_integer_config.FieldTypeIntegerConfig     `json:"integer"`
	Float     field_type_float_config.FieldTypeFloatConfig         `json:"float"`
	Decimal   field_type_decimal_config.FieldTypeDecimalConfig     `json:"decimal"`
	Boolean   *string                                              `json:"boolean"`
	Char      field_type_char_config.FieldTypeCharConfig           `json:"char"`
	Varchar   field_type_varchar_config.FieldTypeVarcharConfig     `json:"varchar"`
	Text      field_type_text_config.FieldTypeTextConfig           `json:"text"`
	Encrypted field_type_encrypted_config.FieldTypeEncryptedConfig `json:"encrypted"`
	Email     field_type_email_config.FieldTypeEmailConfig         `json:"email"`
	Phone     field_type_phone_config.FieldTypePhoneConfig         `json:"phone"`
	URL       field_type_url_config.FieldTypeURLConfig             `json:"url"`
	Location  *string                                              `json:"location"`
	Color     *string                                              `json:"color"`
	RichText  *string                                              `json:"rich_text"`
	Code      *string                                              `json:"code"`
	Markdown  *string                                              `json:"markdown"`
	File      field_type_file_config.FieldTypeFileConfig           `json:"file"`
	Image     field_type_file_config.FieldTypeFileConfig           `json:"image"`
	Video     field_type_file_config.FieldTypeFileConfig           `json:"video"`
	Audio     field_type_file_config.FieldTypeFileConfig           `json:"audio"`
	Enum      field_type_enum_config.FieldTypeEnumConfig           `json:"enum"`
	JSON      field_type_json_config.FieldTypeJSONConfig           `json:"json"`
	Array     field_type_array_config.FieldTypeArrayConfig         `json:"array"`
	Date      field_type_date_config.FieldTypeDateConfig           `json:"date"`
	Datetime  field_type_datetime_config.FieldTypeDatetimeConfig   `json:"datetime"`
	Time      *string                                              `json:"time"`
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

func (e FieldTypeConfig) OrderedFieldIdentifiers() []string {
	res := []string{}
	res = append(res, "uuid")
	res = append(res, "integer")
	res = append(res, "float")
	res = append(res, "decimal")
	res = append(res, "boolean")
	res = append(res, "char")
	res = append(res, "varchar")
	res = append(res, "text")
	res = append(res, "encrypted")
	res = append(res, "email")
	res = append(res, "phone")
	res = append(res, "url")
	res = append(res, "location")
	res = append(res, "color")
	res = append(res, "rich_text")
	res = append(res, "code")
	res = append(res, "markdown")
	res = append(res, "file")
	res = append(res, "image")
	res = append(res, "video")
	res = append(res, "audio")
	res = append(res, "enum")
	res = append(res, "json")
	res = append(res, "array")
	res = append(res, "date")
	res = append(res, "datetime")
	res = append(res, "time")
	res = append(res, "slug")

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
		UUID:      randomvalues.GetRandomRawJSONValuePtr(),
		Integer:   field_type_integer_config.NewFieldTypeIntegerConfigWithRandomValues(),
		Float:     field_type_float_config.NewFieldTypeFloatConfigWithRandomValues(),
		Decimal:   field_type_decimal_config.NewFieldTypeDecimalConfigWithRandomValues(),
		Boolean:   randomvalues.GetRandomRawJSONValuePtr(),
		Char:      field_type_char_config.NewFieldTypeCharConfigWithRandomValues(),
		Varchar:   field_type_varchar_config.NewFieldTypeVarcharConfigWithRandomValues(),
		Text:      field_type_text_config.NewFieldTypeTextConfigWithRandomValues(),
		Encrypted: field_type_encrypted_config.NewFieldTypeEncryptedConfigWithRandomValues(),
		Email:     field_type_email_config.NewFieldTypeEmailConfigWithRandomValues(),
		Phone:     field_type_phone_config.NewFieldTypePhoneConfigWithRandomValues(),
		URL:       field_type_url_config.NewFieldTypeURLConfigWithRandomValues(),
		Location:  randomvalues.GetRandomRawJSONValuePtr(),
		Color:     randomvalues.GetRandomRawJSONValuePtr(),
		RichText:  randomvalues.GetRandomRawJSONValuePtr(),
		Code:      randomvalues.GetRandomRawJSONValuePtr(),
		Markdown:  randomvalues.GetRandomRawJSONValuePtr(),
		File:      field_type_file_config.NewFieldTypeFileConfigWithRandomValues(),
		Image:     field_type_file_config.NewFieldTypeFileConfigWithRandomValues(),
		Video:     field_type_file_config.NewFieldTypeFileConfigWithRandomValues(),
		Audio:     field_type_file_config.NewFieldTypeFileConfigWithRandomValues(),
		Enum:      field_type_enum_config.NewFieldTypeEnumConfigWithRandomValues(),
		JSON:      field_type_json_config.NewFieldTypeJSONConfigWithRandomValues(),
		Array:     field_type_array_config.NewFieldTypeArrayConfigWithRandomValues(),
		Date:      field_type_date_config.NewFieldTypeDateConfigWithRandomValues(),
		Datetime:  field_type_datetime_config.NewFieldTypeDatetimeConfigWithRandomValues(),
		Time:      randomvalues.GetRandomRawJSONValuePtr(),
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
