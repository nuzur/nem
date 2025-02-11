package change_request

import (
	main_entity "nem/core/entity/change_request"
	nemdb "nem/core/repository/gen"

	"github.com/gofrs/uuid"
	"nem/core/entity/change_request_review"
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
		Description:   model.Description,
		ReviewType:    main_entity.ReviewType(model.ReviewType),
		RefUUID:       uuid.FromStringOrNil(model.RefUUID),
		OldData:       string(model.OldData),
		OldDataRef:    model.OldDataRef,
		NewData:       string(model.NewData),
		NewDataRef:    model.NewDataRef,
		Reviews:       change_request_review.ChangeRequestReviewSliceFromJSON(model.Reviews),
		OwnerUUID:     uuid.FromStringOrNil(model.OwnerUUID),
		Status:        main_entity.Status(model.Status),
		CreatedAt:     model.CreatedAt,
		UpdatedAt:     model.UpdatedAt,
		CreatedByUUID: uuid.FromStringOrNil(model.CreatedByUUID),
		UpdatedByUUID: uuid.FromStringOrNil(model.UpdatedByUUID),
	}
}
