
-- name: UpdateTeam :exec
UPDATE team
SET
version = ?, name = ?, enviorments = ?, review_configs = ?, stores = ?, connections = ?, object_stores = ?, default_entity = ?, owner_uuid = ?, status = ?, created_at = ?, updated_at = ?, created_by_uuid = ?, updated_by_uuid = ?
WHERE uuid = ?;

-- name: UpdateProject :exec
UPDATE project
SET
version = ?, name = ?, description = ?, tags = ?, url = ?, owner_uuid = ?, team_uuid = ?, access_type = ?, project_extensions = ?, status = ?, created_at = ?, updated_at = ?, created_by_uuid = ?, updated_by_uuid = ?
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
version = ?, title = ?, description = ?, project_uuid = ?, project_version_uuid = ?, change_type = ?, data_changes = ?, metadata = ?, reviews = ?, review_status = ?, owner_uuid = ?, status = ?, created_at = ?, updated_at = ?, created_by_uuid = ?, updated_by_uuid = ?
WHERE uuid = ?;

-- name: UpdateProjectVersion :exec
UPDATE project_version
SET
version = ?, identifier = ?, description = ?, project_uuid = ?, entities = ?, relationships = ?, enums = ?, services = ?, base_version_uuid = ?, review_status = ?, deployments = ?, status = ?, created_at = ?, updated_at = ?, created_by_uuid = ?, updated_by_uuid = ?
WHERE uuid = ?;

-- name: UpdateUserTeam :exec
UPDATE user_team
SET
user_uuid = ?, user_email = ?, team_uuid = ?, role = ?, review_required_structure = ?, review_required_data = ?, status = ?, created_at = ?, updated_at = ?, created_by_uuid = ?, updated_by_uuid = ?
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

-- name: UpdateUserProjectVersion :exec
UPDATE user_project_version
SET
version = ?, project_version_uuid = ?, user_uuid = ?, data = ?, status = ?, created_at = ?, updated_at = ?, created_by_uuid = ?, updated_by_uuid = ?
WHERE uuid = ?;

-- name: UpdateUserProject :exec
UPDATE user_project
SET
user_uuid = ?, user_email = ?, project_uuid = ?, role = ?, review_required_structure = ?, review_required_data = ?, status = ?, created_at = ?, updated_at = ?, created_by_uuid = ?, updated_by_uuid = ?
WHERE uuid = ?;

-- name: UpdateMembership :exec
UPDATE membership
SET
owner_uuid = ?, type = ?, start_date = ?, billing_metadata = ?, status = ?, created_at = ?, updated_at = ?, created_by_uuid = ?, updated_by_uuid = ?
WHERE uuid = ?;

-- name: UpdateAiUsage :exec
UPDATE ai_usage
SET
user_uuid = ?, project_uuid = ?, project_version_uuid = ?, user_prompt = ?, step = ?, context = ?, provider = ?, tokens = ?, status = ?, created_at = ?, updated_at = ?, created_by_uuid = ?, updated_by_uuid = ?
WHERE uuid = ?;

