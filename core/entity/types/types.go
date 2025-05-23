package types

type FieldType string

const (
	StringFieldType                FieldType = "string"
	UUIDFieldType                  FieldType = "uuid"
	SingleDependantEntityFieldType FieldType = "dependant_single"
	MultiDependantEntityFieldType  FieldType = "dependant_multi"
	RawJSONFieldType               FieldType = "raw_json"
	SingleEnumFieldType            FieldType = "enum_single"
	MultiEnumFieldType             FieldType = "enum_multi"
	TimestampFieldType             FieldType = "timestamp"
	BooleanFieldType               FieldType = "bool"
	IntFieldType                   FieldType = "int"
	FloatFieldType                 FieldType = "float"
	ArrayFieldType                 FieldType = "array"
)

func FieldTypeToSQL(ft FieldType) string {
	switch ft {
	case StringFieldType, UUIDFieldType:
		return "VARCHAR"
	case SingleDependantEntityFieldType, MultiDependantEntityFieldType, RawJSONFieldType, MultiEnumFieldType, ArrayFieldType:
		return "JSON"
	case SingleEnumFieldType, BooleanFieldType, IntFieldType:
		return "INT"
	case TimestampFieldType:
		return "DATETIME"
	case FloatFieldType:
		return "DOUBLE"
	}
	return ""
}
