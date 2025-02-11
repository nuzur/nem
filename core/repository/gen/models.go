// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package nemdb

import (
	"encoding/json"
	"time"
)

type ChangeRequest struct {
	UUID          string          `json:"uuid"`
	Version       int64           `json:"version"`
	Title         string          `json:"title"`
	Description   string          `json:"description"`
	ReviewType    int64           `json:"review_type"`
	RefUUID       string          `json:"ref_uuid"`
	OldData       json.RawMessage `json:"old_data"`
	OldDataRef    string          `json:"old_data_ref"`
	NewData       json.RawMessage `json:"new_data"`
	NewDataRef    string          `json:"new_data_ref"`
	Reviews       json.RawMessage `json:"reviews"`
	OwnerUUID     string          `json:"owner_uuid"`
	Status        int64           `json:"status"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
	CreatedByUUID string          `json:"created_by_uuid"`
	UpdatedByUUID string          `json:"updated_by_uuid"`
}

type Extension struct {
	UUID              string          `json:"uuid"`
	Version           int64           `json:"version"`
	Identifier        string          `json:"identifier"`
	DisplayName       string          `json:"display_name"`
	DisplayAuthorName string          `json:"display_author_name"`
	Description       string          `json:"description"`
	URL               string          `json:"url"`
	Verfied           bool            `json:"verfied"`
	Repository        string          `json:"repository"`
	ExtensionType     int64           `json:"extension_type"`
	Tags              json.RawMessage `json:"tags"`
	Public            bool            `json:"public"`
	Visibility        json.RawMessage `json:"visibility"`
	Status            int64           `json:"status"`
	OwnerUUID         string          `json:"owner_uuid"`
	CreatedAt         time.Time       `json:"created_at"`
	UpdatedAt         time.Time       `json:"updated_at"`
	CreatedByUUID     string          `json:"created_by_uuid"`
	UpdatedByUUID     string          `json:"updated_by_uuid"`
}

type ExtensionExecution struct {
	UUID                 string          `json:"uuid"`
	ExtensionUUID        string          `json:"extension_uuid"`
	ExtensionVersionUUID string          `json:"extension_version_uuid"`
	ProjectExtensionUUID string          `json:"project_extension_uuid"`
	ProjectUUID          string          `json:"project_uuid"`
	ProjectVersionUUID   string          `json:"project_version_uuid"`
	ExecutedByUUID       string          `json:"executed_by_uuid"`
	Metadata             json.RawMessage `json:"metadata"`
	Status               int64           `json:"status"`
	StatusMsg            string          `json:"status_msg"`
	CreatedAt            time.Time       `json:"created_at"`
	UpdatedAt            time.Time       `json:"updated_at"`
}

type ExtensionVersion struct {
	UUID                string          `json:"uuid"`
	Version             int64           `json:"version"`
	ExtensionUUID       string          `json:"extension_uuid"`
	DisplayVersion      string          `json:"display_version"`
	Description         string          `json:"description"`
	RepositoryTag       string          `json:"repository_tag"`
	ConfigurationEntity json.RawMessage `json:"configuration_entity"`
	ExecutionMode       json.RawMessage `json:"execution_mode"`
	ReviewStatus        int64           `json:"review_status"`
	Status              int64           `json:"status"`
	CreatedAt           time.Time       `json:"created_at"`
	UpdatedAt           time.Time       `json:"updated_at"`
	CreatedByUUID       string          `json:"created_by_uuid"`
	UpdatedByUUID       string          `json:"updated_by_uuid"`
}

type Organization struct {
	UUID          string          `json:"uuid"`
	Version       int64           `json:"version"`
	Name          string          `json:"name"`
	Domains       json.RawMessage `json:"domains"`
	AdminUUIDs    json.RawMessage `json:"admin_uuids"`
	Memberships   json.RawMessage `json:"memberships"`
	Status        int64           `json:"status"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
	CreatedByUUID string          `json:"created_by_uuid"`
	UpdatedByUUID string          `json:"updated_by_uuid"`
}

type Project struct {
	UUID              string          `json:"uuid"`
	Version           int64           `json:"version"`
	Name              string          `json:"name"`
	Description       string          `json:"description"`
	Tags              json.RawMessage `json:"tags"`
	URL               string          `json:"url"`
	OwnerUUID         string          `json:"owner_uuid"`
	TeamUUID          string          `json:"team_uuid"`
	ProjectExtensions json.RawMessage `json:"project_extensions"`
	Status            int64           `json:"status"`
	CreatedAt         time.Time       `json:"created_at"`
	UpdatedAt         time.Time       `json:"updated_at"`
	CreatedByUUID     string          `json:"created_by_uuid"`
	UpdatedByUUID     string          `json:"updated_by_uuid"`
}

type ProjectVersion struct {
	UUID            string          `json:"uuid"`
	Version         int64           `json:"version"`
	Identifier      string          `json:"identifier"`
	Description     string          `json:"description"`
	ProjectUUID     string          `json:"project_uuid"`
	Entities        json.RawMessage `json:"entities"`
	Relationships   json.RawMessage `json:"relationships"`
	Enums           json.RawMessage `json:"enums"`
	Services        json.RawMessage `json:"services"`
	BaseVersionUUID string          `json:"base_version_uuid"`
	ReviewStatus    int64           `json:"review_status"`
	Deployments     json.RawMessage `json:"deployments"`
	Status          int64           `json:"status"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
	CreatedByUUID   string          `json:"created_by_uuid"`
	UpdatedByUUID   string          `json:"updated_by_uuid"`
}

type Team struct {
	UUID             string          `json:"uuid"`
	Version          int64           `json:"version"`
	Name             string          `json:"name"`
	Enviorments      json.RawMessage `json:"enviorments"`
	ReviewConfigs    json.RawMessage `json:"review_configs"`
	Memberships      json.RawMessage `json:"memberships"`
	Stores           json.RawMessage `json:"stores"`
	Connections      json.RawMessage `json:"connections"`
	ObjectStores     json.RawMessage `json:"object_stores"`
	OrganizationUUID string          `json:"organization_uuid"`
	DefaultEntity    json.RawMessage `json:"default_entity"`
	Status           int64           `json:"status"`
	CreatedAt        time.Time       `json:"created_at"`
	UpdatedAt        time.Time       `json:"updated_at"`
	CreatedByUUID    string          `json:"created_by_uuid"`
	UpdatedByUUID    string          `json:"updated_by_uuid"`
}

type User struct {
	UUID        string    `json:"uuid"`
	Identifier  string    `json:"identifier"`
	Name        string    `json:"name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	UserType    int64     `json:"user_type"`
	CountryIos2 string    `json:"country_ios2"`
	Locale      string    `json:"locale"`
	Metadata    string    `json:"metadata"`
	Status      int64     `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserConnection struct {
	UUID               string          `json:"uuid"`
	UserUUID           string          `json:"user_uuid"`
	ProjectUUID        string          `json:"project_uuid"`
	ProjectVersionUUID string          `json:"project_version_uuid"`
	Type               int64           `json:"type"`
	TypeConfig         json.RawMessage `json:"type_config"`
	DbSchema           string          `json:"db_schema"`
	Executions         json.RawMessage `json:"executions"`
	Status             int64           `json:"status"`
	CreatedAt          time.Time       `json:"created_at"`
	UpdatedAt          time.Time       `json:"updated_at"`
}

type UserTeam struct {
	UUID          string          `json:"uuid"`
	UserUUID      string          `json:"user_uuid"`
	UserEmail     string          `json:"user_email"`
	TeamUUID      string          `json:"team_uuid"`
	Roles         json.RawMessage `json:"roles"`
	Status        int64           `json:"status"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
	CreatedByUUID string          `json:"created_by_uuid"`
	UpdatedByUUID string          `json:"updated_by_uuid"`
}
