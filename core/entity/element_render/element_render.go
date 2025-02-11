package element_render

import (
	"encoding/json"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"

	"time"
)

type ElementRender struct {
	X         float64 `json:"x"`
	Y         float64 `json:"y"`
	Width     float64 `json:"width"`
	Height    float64 `json:"height"`
	Collapsed bool    `json:"collapsed"`
}

func (e ElementRender) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e ElementRender) EntityIdentifier() string {
	return "element_render"
}

func (e ElementRender) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["x"] = types.FloatFieldType
	res["y"] = types.FloatFieldType
	res["width"] = types.FloatFieldType
	res["height"] = types.FloatFieldType
	res["collapsed"] = types.BooleanFieldType
	return res
}

func (e ElementRender) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e ElementRender) IsDependant() bool {
	return true
}

func ElementRenderFromJSON(data json.RawMessage) ElementRender {
	entity := ElementRender{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func ElementRenderSliceFromJSON(data json.RawMessage) []ElementRender {
	entity := []ElementRender{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e ElementRender) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ElementRenderToJSON", err)
	}
	return res
}

func ElementRenderToJSON(e ElementRender) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ElementRenderToJSON", err)
	}
	return res
}

func ElementRenderSliceToJSON(e []ElementRender) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ElementRenderSliceToJSON", err)
	}
	return res
}

func NewElementRenderWithRandomValues() ElementRender {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return ElementRender{
		X:         randomvalues.GetRandomFloatValue(),
		Y:         randomvalues.GetRandomFloatValue(),
		Width:     randomvalues.GetRandomFloatValue(),
		Height:    randomvalues.GetRandomFloatValue(),
		Collapsed: randomvalues.GetRandomBoolValue(),
	}
}

func NewElementRenderSliceWithRandomValues(n int) []ElementRender {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []ElementRender{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewElementRenderWithRandomValues())
	}
	return res
}
