package user_team

import (
	main_entity "github.com/nuzur/nem/core/entity/user_team"
	nemdb "github.com/nuzur/nem/core/repository/gen"

	"github.com/gofrs/uuid"
)

func mapModelsToEntities(models []nemdb.UserTeam) []main_entity.UserTeam {
	result := []main_entity.UserTeam{}
	for _, p := range models {
		result = append(result, mapModelToEntity(p))
	}
	return result
}

func mapModelToEntity(model nemdb.UserTeam) main_entity.UserTeam {
	return main_entity.UserTeam{
		UUID:          uuid.FromStringOrNil(model.UUID),
		UserUUID:      uuid.FromStringOrNil(model.UserUUID),
		UserEmail:     model.UserEmail,
		TeamUUID:      uuid.FromStringOrNil(model.TeamUUID),
		Roles:         main_entity.JSONToRoleSlice(model.Roles),
		Status:        main_entity.Status(model.Status),
		CreatedAt:     model.CreatedAt,
		UpdatedAt:     model.UpdatedAt,
		CreatedByUUID: uuid.FromStringOrNil(model.CreatedByUUID),
		UpdatedByUUID: uuid.FromStringOrNil(model.UpdatedByUUID),
	}
}
