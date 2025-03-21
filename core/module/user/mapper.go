package user

import (
	main_entity "github.com/nuzur/nem/core/entity/user"
	nemdb "github.com/nuzur/nem/core/repository/gen"

	"github.com/nuzur/nem/core/entity/mapper"

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
		Name:        mapper.SqlNullStringToStringPtr(model.Name),
		LastName:    mapper.SqlNullStringToStringPtr(model.LastName),
		Email:       model.Email,
		UserType:    main_entity.UserType(model.UserType),
		CountryIos2: mapper.SqlNullStringToStringPtr(model.CountryIos2),
		Locale:      mapper.SqlNullStringToStringPtr(model.Locale),
		Metadata:    mapper.SqlNullStringToStringPtr(model.Metadata),
		Status:      main_entity.Status(model.Status),
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}
}
