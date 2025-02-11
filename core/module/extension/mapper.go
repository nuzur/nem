package extension

import (
	main_entity "nem/core/entity/extension"
	nemdb "nem/core/repository/gen"

	"nem/core/entity/mapper"

	"github.com/gofrs/uuid"
	"nem/core/entity/visibility"
)

func mapModelsToEntities(models []nemdb.Extension) []main_entity.Extension {
	result := []main_entity.Extension{}
	for _, p := range models {
		result = append(result, mapModelToEntity(p))
	}
	return result
}

func mapModelToEntity(model nemdb.Extension) main_entity.Extension {
	return main_entity.Extension{
		UUID:              uuid.FromStringOrNil(model.UUID),
		Version:           model.Version,
		Identifier:        model.Identifier,
		DisplayName:       model.DisplayName,
		DisplayAuthorName: model.DisplayAuthorName,
		Description:       model.Description,
		URL:               model.URL,
		Verfied:           model.Verfied,
		Repository:        model.Repository,
		ExtensionType:     main_entity.ExtensionType(model.ExtensionType),
		Tags:              mapper.JSONToStringSlice(model.Tags),
		Public:            model.Public,
		Visibility:        visibility.VisibilitySliceFromJSON(model.Visibility),
		Status:            main_entity.Status(model.Status),
		OwnerUUID:         uuid.FromStringOrNil(model.OwnerUUID),
		CreatedAt:         model.CreatedAt,
		UpdatedAt:         model.UpdatedAt,
		CreatedByUUID:     uuid.FromStringOrNil(model.CreatedByUUID),
		UpdatedByUUID:     uuid.FromStringOrNil(model.UpdatedByUUID),
	}
}
