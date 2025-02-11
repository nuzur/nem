package field

import (
	"encoding/json"
	"fmt"
)

//go:generate go run github.com/dmarkham/enumer -type=Type -json
type Type int64

const (
	TYPE_INVALID Type = iota
	TYPE_UUID
	TYPE_INTEGER
	TYPE_FLOAT
	TYPE_DECIMAL
	TYPE_BOOLEAN
	TYPE_CHAR
	TYPE_VARCHAR
	TYPE_TEXT
	TYPE_ENCRYPTED
	TYPE_EMAIL
	TYPE_PHONE
	TYPE_URL
	TYPE_LOCATION
	TYPE_COLOR
	TYPE_RICHTEXT
	TYPE_CODE
	TYPE_MARKDOWN
	TYPE_FILE
	TYPE_IMAGE
	TYPE_AUDIO
	TYPE_VIDEO
	TYPE_ENUM
	TYPE_JSON
	TYPE_ARRAY
	TYPE_DATE
	TYPE_DATETIME
	TYPE_TIME
	TYPE_SLUG
)

func (e Type) ToInt64() int64 {
	return int64(e)
}

func TypeFromString(in string) Type {
	switch in {
	case "uuid":
		return TYPE_UUID
	case "integer":
		return TYPE_INTEGER
	case "float":
		return TYPE_FLOAT
	case "decimal":
		return TYPE_DECIMAL
	case "boolean":
		return TYPE_BOOLEAN
	case "char":
		return TYPE_CHAR
	case "varchar":
		return TYPE_VARCHAR
	case "text":
		return TYPE_TEXT
	case "encrypted":
		return TYPE_ENCRYPTED
	case "email":
		return TYPE_EMAIL
	case "phone":
		return TYPE_PHONE
	case "url":
		return TYPE_URL
	case "location":
		return TYPE_LOCATION
	case "color":
		return TYPE_COLOR
	case "richtext":
		return TYPE_RICHTEXT
	case "code":
		return TYPE_CODE
	case "markdown":
		return TYPE_MARKDOWN
	case "file":
		return TYPE_FILE
	case "image":
		return TYPE_IMAGE
	case "audio":
		return TYPE_AUDIO
	case "video":
		return TYPE_VIDEO
	case "enum":
		return TYPE_ENUM
	case "json":
		return TYPE_JSON
	case "array":
		return TYPE_ARRAY
	case "date":
		return TYPE_DATE
	case "datetime":
		return TYPE_DATETIME
	case "time":
		return TYPE_TIME
	case "slug":
		return TYPE_SLUG
	}
	return TYPE_INVALID
}

func TypeFromPointerString(in *string) Type {
	if in == nil {
		return TYPE_INVALID
	}
	return TypeFromString(*in)
}

func (e Type) String() string {
	switch e {
	case TYPE_UUID:
		return "uuid"
	case TYPE_INTEGER:
		return "integer"
	case TYPE_FLOAT:
		return "float"
	case TYPE_DECIMAL:
		return "decimal"
	case TYPE_BOOLEAN:
		return "boolean"
	case TYPE_CHAR:
		return "char"
	case TYPE_VARCHAR:
		return "varchar"
	case TYPE_TEXT:
		return "text"
	case TYPE_ENCRYPTED:
		return "encrypted"
	case TYPE_EMAIL:
		return "email"
	case TYPE_PHONE:
		return "phone"
	case TYPE_URL:
		return "url"
	case TYPE_LOCATION:
		return "location"
	case TYPE_COLOR:
		return "color"
	case TYPE_RICHTEXT:
		return "richtext"
	case TYPE_CODE:
		return "code"
	case TYPE_MARKDOWN:
		return "markdown"
	case TYPE_FILE:
		return "file"
	case TYPE_IMAGE:
		return "image"
	case TYPE_AUDIO:
		return "audio"
	case TYPE_VIDEO:
		return "video"
	case TYPE_ENUM:
		return "enum"
	case TYPE_JSON:
		return "json"
	case TYPE_ARRAY:
		return "array"
	case TYPE_DATE:
		return "date"
	case TYPE_DATETIME:
		return "datetime"
	case TYPE_TIME:
		return "time"
	case TYPE_SLUG:
		return "slug"
	}

	return "invalid"
}

func (e Type) StringPtr() *string {
	val := e.String()
	return &val
}

func TypeSliceToJSON(in []Type) json.RawMessage {
	res := make([]int64, len(in))
	for i, e := range in {
		res[i] = int64(e)
	}
	jr, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("error marshalling Type slice to json: %v", err)
		return json.RawMessage{}
	}
	return jr
}

func JSONToTypeSlice(in json.RawMessage) []Type {
	res := []int64{}
	err := json.Unmarshal(in, &res)
	if err != nil {
		fmt.Printf("error unmarshing Type slice to int slice: %v", err)
		return nil
	}
	if len(res) == 0 {
		return nil
	}
	finalRes := []Type{}
	for _, r := range res {
		finalRes = append(finalRes, Type(r))
	}
	return finalRes
}
