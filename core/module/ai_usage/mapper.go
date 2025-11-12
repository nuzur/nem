package ai_usage

import (
	main_entity "github.com/nuzur/nem/core/entity/ai_usage"
	nemdb "github.com/nuzur/nem/core/repository/gen"

	"github.com/gofrs/uuid"
)

func mapModelsToEntities(models []nemdb.AiUsage) []main_entity.AiUsage {
	result := []main_entity.AiUsage{}
	for _, p := range models {
		result = append(result, mapModelToEntity(p))
	}
	return result
}

func mapModelToEntity(model nemdb.AiUsage) main_entity.AiUsage {
	return main_entity.AiUsage{
		UUID:               uuid.FromStringOrNil(model.UUID),
		UserUUID:           uuid.FromStringOrNil(model.UserUUID),
		ProjectUUID:        uuid.FromStringOrNil(model.ProjectUUID),
		ProjectVersionUUID: uuid.FromStringOrNil(model.ProjectVersionUUID),
		UserPrompt:         model.UserPrompt,
		Step:               model.Step,
		Context:            main_entity.Context(model.Context),
		Provider:           main_entity.Provider(model.Provider),
		Tokens:             model.Tokens,
		Status:             main_entity.Status(model.Status),
		CreatedAt:          model.CreatedAt,
		UpdatedAt:          model.UpdatedAt,
		CreatedByUUID:      uuid.FromStringOrNil(model.CreatedByUUID),
		UpdatedByUUID:      uuid.FromStringOrNil(model.UpdatedByUUID),
	}
}
