
-- name: UpdateTeam :exec
UPDATE team
SET
version = ?, name = ?, enviorments = ?, review_configs = ?, memberships = ?, stores = ?, connections = ?, object_stores = ?, organization_uuid = ?, default_entity = ?, status = ?, created_at = ?, updated_at = ?, created_by_uuid = ?, updated_by_uuid = ?
WHERE uuid = ?;

-- name: UpdateOrganization :exec
UPDATE organization
SET
version = ?, name = ?, domains = ?, admin_uuids = ?, memberships = ?, status = ?, created_at = ?, updated_at = ?, created_by_uuid = ?, updated_by_uuid = ?
WHERE uuid = ?;

-- name: UpdateProject :exec
UPDATE project
SET
version = ?, name = ?, description = ?, tags = ?, url = ?, owner_uuid = ?, team_uuid = ?, project_extensions = ?, status = ?, created_at = ?, updated_at = ?, created_by_uuid = ?, updated_by_uuid = ?
WHERE uuid = ?;

-- name: UpdateExtension :exec
UPDATE extension
SET
version = ?, identifier = ?, display_name = ?, display_author_name = ?, description = ?, url = ?, verfied = ?, repository = ?, extension_type = ?, tags = ?, public = ?, visibility = ?, status = ?, owner_uuid = ?, created_at = ?, updated_at = ?, created_by_uuid = ?, updated_by_uuid = ?
WHERE uuid = ?;

-- name: UpdateExtensionVersion :exec
UPDATE extension_version
SET
version = ?, extension_uuid = ?, display_version = ?, description = ?, repository_tag = ?, configuration_entity = ?, execution_mode = ?, review_status = ?, status = ?, created_at = ?, updated_at = ?, created_by_uuid = ?, updated_by_uuid = ?
WHERE uuid = ?;

-- name: UpdateUser :exec
UPDATE user
SET
identifier = ?, name = ?, last_name = ?, email = ?, user_type = ?, country_ios2 = ?, locale = ?, metadata = ?, status = ?, created_at = ?, updated_at = ?
WHERE uuid = ?;

-- name: UpdateChangeRequest :exec
UPDATE change_request
SET
version = ?, title = ?, description = ?, review_type = ?, ref_uuid = ?, old_data = ?, old_data_ref = ?, new_data = ?, new_data_ref = ?, reviews = ?, owner_uuid = ?, status = ?, created_at = ?, updated_at = ?, created_by_uuid = ?, updated_by_uuid = ?
WHERE uuid = ?;

-- name: UpdateProjectVersion :exec
UPDATE project_version
SET
version = ?, identifier = ?, description = ?, project_uuid = ?, entities = ?, relationships = ?, enums = ?, services = ?, base_version_uuid = ?, review_status = ?, deployments = ?, status = ?, created_at = ?, updated_at = ?, created_by_uuid = ?, updated_by_uuid = ?
WHERE uuid = ?;

-- name: UpdateUserTeam :exec
UPDATE user_team
SET
user_uuid = ?, user_email = ?, team_uuid = ?, roles = ?, status = ?, created_at = ?, updated_at = ?, created_by_uuid = ?, updated_by_uuid = ?
WHERE uuid = ?;

-- name: UpdateExtensionExecution :exec
UPDATE extension_execution
SET
extension_uuid = ?, extension_version_uuid = ?, project_extension_uuid = ?, project_uuid = ?, project_version_uuid = ?, executed_by_uuid = ?, metadata = ?, status = ?, status_msg = ?, created_at = ?, updated_at = ?
WHERE uuid = ?;

-- name: UpdateUserConnection :exec
UPDATE user_connection
SET
user_uuid = ?, project_uuid = ?, project_version_uuid = ?, type = ?, type_config = ?, db_schema = ?, executions = ?, status = ?, created_at = ?, updated_at = ?
WHERE uuid = ?;

