package user_team

import (
	main_entity "github.com/nuzur/nem/core/entity/user_team"
	nemdb "github.com/nuzur/nem/core/repository/gen"

	"github.com/nuzur/nem/core/entity/mapper"

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
		UUID:                    uuid.FromStringOrNil(model.UUID),
		UserUUID:                uuid.FromStringOrNil(mapper.SqlNullStringToString(model.UserUUID)),
		UserEmail:               mapper.SqlNullStringToStringPtr(model.UserEmail),
		TeamUUID:                uuid.FromStringOrNil(model.TeamUUID),
		Role:                    main_entity.Role(model.Role),
		ReviewRequiredStructure: model.ReviewRequiredStructure,
		ReviewRequiredData:      model.ReviewRequiredData,
		Status:                  main_entity.Status(model.Status),
		CreatedAt:               model.CreatedAt,
		UpdatedAt:               model.UpdatedAt,
		CreatedByUUID:           uuid.FromStringOrNil(model.CreatedByUUID),
		UpdatedByUUID:           uuid.FromStringOrNil(model.UpdatedByUUID),
	}
}
