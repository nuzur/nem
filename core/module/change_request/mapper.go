package change_request

import (
	main_entity "github.com/nuzur/nem/core/entity/change_request"
	nemdb "github.com/nuzur/nem/core/repository/gen"

	"github.com/nuzur/nem/core/entity/mapper"

	"github.com/gofrs/uuid"
	"github.com/nuzur/nem/core/entity/change_request_review"
)

func mapModelsToEntities(models []nemdb.ChangeRequest) []main_entity.ChangeRequest {
	result := []main_entity.ChangeRequest{}
	for _, p := range models {
		result = append(result, mapModelToEntity(p))
	}
	return result
}

func mapModelToEntity(model nemdb.ChangeRequest) main_entity.ChangeRequest {
	return main_entity.ChangeRequest{
		UUID:          uuid.FromStringOrNil(model.UUID),
		Version:       model.Version,
		Title:         model.Title,
		Description:   mapper.SqlNullStringToStringPtr(model.Description),
		ReviewType:    main_entity.ReviewType(model.ReviewType),
		RefUUID:       uuid.FromStringOrNil(model.RefUUID),
		OldData:       mapper.StringToStringPtr(string(model.OldData)),
		OldDataRef:    mapper.SqlNullStringToStringPtr(model.OldDataRef),
		NewData:       mapper.StringToStringPtr(string(model.NewData)),
		NewDataRef:    mapper.SqlNullStringToStringPtr(model.NewDataRef),
		Reviews:       change_request_review.ChangeRequestReviewSliceFromJSON(model.Reviews),
		OwnerUUID:     uuid.FromStringOrNil(model.OwnerUUID),
		Status:        main_entity.Status(model.Status),
		CreatedAt:     model.CreatedAt,
		UpdatedAt:     model.UpdatedAt,
		CreatedByUUID: uuid.FromStringOrNil(model.CreatedByUUID),
		UpdatedByUUID: uuid.FromStringOrNil(model.UpdatedByUUID),
	}
}
