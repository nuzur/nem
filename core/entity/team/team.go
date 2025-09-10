package team

import (
	"encoding/json"

	"github.com/gofrs/uuid"
	"github.com/nuzur/nem/core/entity/connection"
	"github.com/nuzur/nem/core/entity/entity"
	"github.com/nuzur/nem/core/entity/enviorment"
	"github.com/nuzur/nem/core/entity/object_store"
	"github.com/nuzur/nem/core/entity/review_config"
	"github.com/nuzur/nem/core/entity/store"
	"time"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"
)

type Team struct {
	UUID             uuid.UUID                    `json:"uuid"`
	Version          int64                        `json:"version"`
	Name             string                       `json:"name"`
	Enviorments      []enviorment.Enviorment      `json:"enviorments"`
	ReviewConfigs    []review_config.ReviewConfig `json:"review_configs"`
	Stores           []store.Store                `json:"stores"`
	Connections      []connection.Connection      `json:"connections"`
	ObjectStores     []object_store.ObjectStore   `json:"object_stores"`
	OrganizationUUID uuid.UUID                    `json:"organization_uuid"`
	DefaultEntity    entity.Entity                `json:"default_entity"`
	Status           Status                       `json:"status"`
	CreatedAt        time.Time                    `json:"created_at"`
	UpdatedAt        time.Time                    `json:"updated_at"`
	CreatedByUUID    uuid.UUID                    `json:"created_by_uuid"`
	UpdatedByUUID    uuid.UUID                    `json:"updated_by_uuid"`
}

func (e Team) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e Team) EntityIdentifier() string {
	return "team"
}

func (e Team) PrimaryKeyIdentifier() string {
	return "uuid"
}

func (e Team) PrimaryKeyValue() string {
	return e.UUID.String()
}

func (e Team) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["uuid"] = types.UUIDFieldType
	res["version"] = types.IntFieldType
	res["name"] = types.StringFieldType
	res["enviorments"] = types.MultiDependantEntityFieldType
	res["review_configs"] = types.MultiDependantEntityFieldType
	res["stores"] = types.MultiDependantEntityFieldType
	res["connections"] = types.MultiDependantEntityFieldType
	res["object_stores"] = types.MultiDependantEntityFieldType
	res["organization_uuid"] = types.UUIDFieldType
	res["default_entity"] = types.SingleDependantEntityFieldType
	res["status"] = types.SingleEnumFieldType
	res["created_at"] = types.TimestampFieldType
	res["updated_at"] = types.TimestampFieldType
	res["created_by_uuid"] = types.UUIDFieldType
	res["updated_by_uuid"] = types.UUIDFieldType
	return res
}

func (e Team) DependantFieldIdentifierToTypeMap() map[string]map[string]types.FieldType {
	res := make(map[string]map[string]types.FieldType)

	res["enviorments"] = enviorment.Enviorment{}.FieldIdentfierToTypeMap()
	res["review_configs"] = review_config.ReviewConfig{}.FieldIdentfierToTypeMap()
	res["stores"] = store.Store{}.FieldIdentfierToTypeMap()
	res["connections"] = connection.Connection{}.FieldIdentfierToTypeMap()
	res["object_stores"] = object_store.ObjectStore{}.FieldIdentfierToTypeMap()
	res["default_entity"] = entity.Entity{}.FieldIdentfierToTypeMap()
	return res
}

func (e Team) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e Team) IsDependant() bool {
	return false
}

func TeamFromJSON(data json.RawMessage) Team {
	entity := Team{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func TeamSliceFromJSON(data json.RawMessage) []Team {
	entity := []Team{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e Team) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error TeamToJSON", err)
	}
	return res
}

func TeamToJSON(e Team) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error TeamToJSON", err)
	}
	return res
}

func TeamSliceToJSON(e []Team) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error TeamSliceToJSON", err)
	}
	return res
}

func NewTeamWithRandomValues() Team {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return Team{
		UUID:             randomvalues.GetRandomUUIDValue(),
		Version:          randomvalues.GetRandomIntValue(),
		Name:             randomvalues.GetRandomStringValue(),
		Enviorments:      enviorment.NewEnviormentSliceWithRandomValues(rand.Intn(10)),
		ReviewConfigs:    review_config.NewReviewConfigSliceWithRandomValues(rand.Intn(10)),
		Stores:           store.NewStoreSliceWithRandomValues(rand.Intn(10)),
		Connections:      connection.NewConnectionSliceWithRandomValues(rand.Intn(10)),
		ObjectStores:     object_store.NewObjectStoreSliceWithRandomValues(rand.Intn(10)),
		OrganizationUUID: randomvalues.GetRandomUUIDValue(),
		DefaultEntity:    entity.NewEntityWithRandomValues(),
		Status:           randomvalues.GetRandomOptionValue[Status](2),
		CreatedAt:        randomvalues.GetRandomTimeValue(),
		UpdatedAt:        randomvalues.GetRandomTimeValue(),
		CreatedByUUID:    randomvalues.GetRandomUUIDValue(),
		UpdatedByUUID:    randomvalues.GetRandomUUIDValue(),
	}
}

func NewTeamSliceWithRandomValues(n int) []Team {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []Team{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewTeamWithRandomValues())
	}
	return res
}
