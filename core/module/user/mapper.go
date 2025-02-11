package user

import (
	main_entity "nem/core/entity/user"
	nemdb "nem/core/repository/gen"

	"github.com/gofrs/uuid"
)

func mapModelsToEntities(models []nemdb.User) []main_entity.User {
	result := []main_entity.User{}
	for _, p := range models {
		result = append(result, mapModelToEntity(p))
	}
	return result
}

func mapModelToEntity(model nemdb.User) main_entity.User {
	return main_entity.User{
		UUID:        uuid.FromStringOrNil(model.UUID),
		Identifier:  model.Identifier,
		Name:        model.Name,
		LastName:    model.LastName,
		Email:       model.Email,
		UserType:    main_entity.UserType(model.UserType),
		CountryIos2: model.CountryIos2,
		Locale:      model.Locale,
		Metadata:    model.Metadata,
		Status:      main_entity.Status(model.Status),
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}
}
