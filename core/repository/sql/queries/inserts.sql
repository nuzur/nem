
-- name: InsertTeam :execresult
INSERT INTO team
(uuid,version,name,enviorments,review_configs,memberships,stores,connections,object_stores,organization_uuid,default_entity,status,created_at,updated_at,created_by_uuid,updated_by_uuid)
VALUES
(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);

-- name: InsertOrganization :execresult
INSERT INTO organization
(uuid,version,name,domains,admin_uuids,memberships,status,created_at,updated_at,created_by_uuid,updated_by_uuid)
VALUES
(?,?,?,?,?,?,?,?,?,?,?);

-- name: InsertProject :execresult
INSERT INTO project
(uuid,version,name,description,tags,url,owner_uuid,team_uuid,access_type,project_extensions,status,created_at,updated_at,created_by_uuid,updated_by_uuid)
VALUES
(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);

-- name: InsertExtension :execresult
INSERT INTO extension
(uuid,version,identifier,display_name,display_author_name,description,url,verfied,repository,extension_type,tags,public,visibility,status,owner_uuid,created_at,updated_at,created_by_uuid,updated_by_uuid)
VALUES
(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);

-- name: InsertExtensionVersion :execresult
INSERT INTO extension_version
(uuid,version,extension_uuid,display_version,description,repository_tag,configuration_entity,execution_mode,review_status,status,created_at,updated_at,created_by_uuid,updated_by_uuid)
VALUES
(?,?,?,?,?,?,?,?,?,?,?,?,?,?);

-- name: InsertUser :execresult
INSERT INTO user
(uuid,identifier,name,last_name,email,user_type,country_ios2,locale,metadata,status,created_at,updated_at)
VALUES
(?,?,?,?,?,?,?,?,?,?,?,?);

-- name: InsertChangeRequest :execresult
INSERT INTO change_request
(uuid,version,title,description,project_version_uuid,change_type,data_changes,version_changes,reviews,review_status,owner_uuid,status,created_at,updated_at,created_by_uuid,updated_by_uuid)
VALUES
(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);

-- name: InsertProjectVersion :execresult
INSERT INTO project_version
(uuid,version,identifier,description,project_uuid,entities,relationships,enums,services,base_version_uuid,review_status,reviews,deployments,status,created_at,updated_at,created_by_uuid,updated_by_uuid)
VALUES
(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);

-- name: InsertUserTeam :execresult
INSERT INTO user_team
(uuid,user_uuid,user_email,team_uuid,role,review_required_structure,review_required_data,status,created_at,updated_at,created_by_uuid,updated_by_uuid)
VALUES
(?,?,?,?,?,?,?,?,?,?,?,?);

-- name: InsertExtensionExecution :execresult
INSERT INTO extension_execution
(uuid,extension_uuid,extension_version_uuid,project_extension_uuid,project_uuid,project_version_uuid,executed_by_uuid,metadata,status,status_msg,created_at,updated_at)
VALUES
(?,?,?,?,?,?,?,?,?,?,?,?);

-- name: InsertUserConnection :execresult
INSERT INTO user_connection
(uuid,user_uuid,project_uuid,project_version_uuid,type,type_config,db_schema,executions,status,created_at,updated_at)
VALUES
(?,?,?,?,?,?,?,?,?,?,?);

-- name: InsertUserProjectVersion :execresult
INSERT INTO user_project_version
(uuid,version,project_version_uuid,user_uuid,data,status,created_at,updated_at,created_by_uuid,updated_by_uuid)
VALUES
(?,?,?,?,?,?,?,?,?,?);

-- name: InsertUserProject :execresult
INSERT INTO user_project
(uuid,user_uuid,user_email,project_uuid,role,review_required_structure,review_required_data,status,created_at,updated_at,created_by_uuid,updated_by_uuid)
VALUES
(?,?,?,?,?,?,?,?,?,?,?,?);
