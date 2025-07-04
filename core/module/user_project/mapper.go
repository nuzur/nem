package user_project

import (
	main_entity "github.com/nuzur/nem/core/entity/user_project"
	nemdb "github.com/nuzur/nem/core/repository/gen"

	"github.com/nuzur/nem/core/entity/mapper"

	"github.com/gofrs/uuid"
)

func mapModelsToEntities(models []nemdb.UserProject) []main_entity.UserProject {
	result := []main_entity.UserProject{}
	for _, p := range models {
		result = append(result, mapModelToEntity(p))
	}
	return result
}

func mapModelToEntity(model nemdb.UserProject) main_entity.UserProject {
	return main_entity.UserProject{
		UUID:                    uuid.FromStringOrNil(model.UUID),
		UserUUID:                uuid.FromStringOrNil(mapper.SqlNullStringToString(model.UserUUID)),
		UserEmail:               mapper.SqlNullStringToStringPtr(model.UserEmail),
		ProjectUUID:             uuid.FromStringOrNil(model.ProjectUUID),
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
