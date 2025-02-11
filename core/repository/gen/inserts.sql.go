// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: inserts.sql

package nemdb

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"
)

const insertChangeRequest = `-- name: InsertChangeRequest :execresult
INSERT INTO change_request
(uuid,version,title,description,review_type,ref_uuid,old_data,old_data_ref,new_data,new_data_ref,reviews,owner_uuid,status,created_at,updated_at,created_by_uuid,updated_by_uuid)
VALUES
(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)
`

type InsertChangeRequestParams struct {
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

func (q *Queries) InsertChangeRequest(ctx context.Context, arg InsertChangeRequestParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, insertChangeRequest,
		arg.UUID,
		arg.Version,
		arg.Title,
		arg.Description,
		arg.ReviewType,
		arg.RefUUID,
		arg.OldData,
		arg.OldDataRef,
		arg.NewData,
		arg.NewDataRef,
		arg.Reviews,
		arg.OwnerUUID,
		arg.Status,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.CreatedByUUID,
		arg.UpdatedByUUID,
	)
}

const insertExtension = `-- name: InsertExtension :execresult
INSERT INTO extension
(uuid,version,identifier,display_name,display_author_name,description,url,verfied,repository,extension_type,tags,public,visibility,status,owner_uuid,created_at,updated_at,created_by_uuid,updated_by_uuid)
VALUES
(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)
`

type InsertExtensionParams struct {
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

func (q *Queries) InsertExtension(ctx context.Context, arg InsertExtensionParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, insertExtension,
		arg.UUID,
		arg.Version,
		arg.Identifier,
		arg.DisplayName,
		arg.DisplayAuthorName,
		arg.Description,
		arg.URL,
		arg.Verfied,
		arg.Repository,
		arg.ExtensionType,
		arg.Tags,
		arg.Public,
		arg.Visibility,
		arg.Status,
		arg.OwnerUUID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.CreatedByUUID,
		arg.UpdatedByUUID,
	)
}

const insertExtensionExecution = `-- name: InsertExtensionExecution :execresult
INSERT INTO extension_execution
(uuid,extension_uuid,extension_version_uuid,project_extension_uuid,project_uuid,project_version_uuid,executed_by_uuid,metadata,status,status_msg,created_at,updated_at)
VALUES
(?,?,?,?,?,?,?,?,?,?,?,?)
`

type InsertExtensionExecutionParams struct {
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

func (q *Queries) InsertExtensionExecution(ctx context.Context, arg InsertExtensionExecutionParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, insertExtensionExecution,
		arg.UUID,
		arg.ExtensionUUID,
		arg.ExtensionVersionUUID,
		arg.ProjectExtensionUUID,
		arg.ProjectUUID,
		arg.ProjectVersionUUID,
		arg.ExecutedByUUID,
		arg.Metadata,
		arg.Status,
		arg.StatusMsg,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
}

const insertExtensionVersion = `-- name: InsertExtensionVersion :execresult
INSERT INTO extension_version
(uuid,version,extension_uuid,display_version,description,repository_tag,configuration_entity,execution_mode,review_status,status,created_at,updated_at,created_by_uuid,updated_by_uuid)
VALUES
(?,?,?,?,?,?,?,?,?,?,?,?,?,?)
`

type InsertExtensionVersionParams struct {
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

func (q *Queries) InsertExtensionVersion(ctx context.Context, arg InsertExtensionVersionParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, insertExtensionVersion,
		arg.UUID,
		arg.Version,
		arg.ExtensionUUID,
		arg.DisplayVersion,
		arg.Description,
		arg.RepositoryTag,
		arg.ConfigurationEntity,
		arg.ExecutionMode,
		arg.ReviewStatus,
		arg.Status,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.CreatedByUUID,
		arg.UpdatedByUUID,
	)
}

const insertOrganization = `-- name: InsertOrganization :execresult
INSERT INTO organization
(uuid,version,name,domains,admin_uuids,memberships,status,created_at,updated_at,created_by_uuid,updated_by_uuid)
VALUES
(?,?,?,?,?,?,?,?,?,?,?)
`

type InsertOrganizationParams struct {
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

func (q *Queries) InsertOrganization(ctx context.Context, arg InsertOrganizationParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, insertOrganization,
		arg.UUID,
		arg.Version,
		arg.Name,
		arg.Domains,
		arg.AdminUUIDs,
		arg.Memberships,
		arg.Status,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.CreatedByUUID,
		arg.UpdatedByUUID,
	)
}

const insertProject = `-- name: InsertProject :execresult
INSERT INTO project
(uuid,version,name,description,tags,url,owner_uuid,team_uuid,project_extensions,status,created_at,updated_at,created_by_uuid,updated_by_uuid)
VALUES
(?,?,?,?,?,?,?,?,?,?,?,?,?,?)
`

type InsertProjectParams struct {
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

func (q *Queries) InsertProject(ctx context.Context, arg InsertProjectParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, insertProject,
		arg.UUID,
		arg.Version,
		arg.Name,
		arg.Description,
		arg.Tags,
		arg.URL,
		arg.OwnerUUID,
		arg.TeamUUID,
		arg.ProjectExtensions,
		arg.Status,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.CreatedByUUID,
		arg.UpdatedByUUID,
	)
}

const insertProjectVersion = `-- name: InsertProjectVersion :execresult
INSERT INTO project_version
(uuid,version,identifier,description,project_uuid,entities,relationships,enums,services,base_version_uuid,review_status,deployments,status,created_at,updated_at,created_by_uuid,updated_by_uuid)
VALUES
(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)
`

type InsertProjectVersionParams struct {
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

func (q *Queries) InsertProjectVersion(ctx context.Context, arg InsertProjectVersionParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, insertProjectVersion,
		arg.UUID,
		arg.Version,
		arg.Identifier,
		arg.Description,
		arg.ProjectUUID,
		arg.Entities,
		arg.Relationships,
		arg.Enums,
		arg.Services,
		arg.BaseVersionUUID,
		arg.ReviewStatus,
		arg.Deployments,
		arg.Status,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.CreatedByUUID,
		arg.UpdatedByUUID,
	)
}

const insertTeam = `-- name: InsertTeam :execresult
INSERT INTO team
(uuid,version,name,enviorments,review_configs,memberships,stores,connections,object_stores,organization_uuid,default_entity,status,created_at,updated_at,created_by_uuid,updated_by_uuid)
VALUES
(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)
`

type InsertTeamParams struct {
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

func (q *Queries) InsertTeam(ctx context.Context, arg InsertTeamParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, insertTeam,
		arg.UUID,
		arg.Version,
		arg.Name,
		arg.Enviorments,
		arg.ReviewConfigs,
		arg.Memberships,
		arg.Stores,
		arg.Connections,
		arg.ObjectStores,
		arg.OrganizationUUID,
		arg.DefaultEntity,
		arg.Status,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.CreatedByUUID,
		arg.UpdatedByUUID,
	)
}

const insertUser = `-- name: InsertUser :execresult
INSERT INTO user
(uuid,identifier,name,last_name,email,user_type,country_ios2,locale,metadata,status,created_at,updated_at)
VALUES
(?,?,?,?,?,?,?,?,?,?,?,?)
`

type InsertUserParams struct {
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

func (q *Queries) InsertUser(ctx context.Context, arg InsertUserParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, insertUser,
		arg.UUID,
		arg.Identifier,
		arg.Name,
		arg.LastName,
		arg.Email,
		arg.UserType,
		arg.CountryIos2,
		arg.Locale,
		arg.Metadata,
		arg.Status,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
}

const insertUserConnection = `-- name: InsertUserConnection :execresult
INSERT INTO user_connection
(uuid,user_uuid,project_uuid,project_version_uuid,type,type_config,db_schema,executions,status,created_at,updated_at)
VALUES
(?,?,?,?,?,?,?,?,?,?,?)
`

type InsertUserConnectionParams struct {
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

func (q *Queries) InsertUserConnection(ctx context.Context, arg InsertUserConnectionParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, insertUserConnection,
		arg.UUID,
		arg.UserUUID,
		arg.ProjectUUID,
		arg.ProjectVersionUUID,
		arg.Type,
		arg.TypeConfig,
		arg.DbSchema,
		arg.Executions,
		arg.Status,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
}

const insertUserTeam = `-- name: InsertUserTeam :execresult
INSERT INTO user_team
(uuid,user_uuid,user_email,team_uuid,roles,status,created_at,updated_at,created_by_uuid,updated_by_uuid)
VALUES
(?,?,?,?,?,?,?,?,?,?)
`

type InsertUserTeamParams struct {
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

func (q *Queries) InsertUserTeam(ctx context.Context, arg InsertUserTeamParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, insertUserTeam,
		arg.UUID,
		arg.UserUUID,
		arg.UserEmail,
		arg.TeamUUID,
		arg.Roles,
		arg.Status,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.CreatedByUUID,
		arg.UpdatedByUUID,
	)
}
