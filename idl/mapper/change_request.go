package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/change_request"
	pb "github.com/nuzur/nem/idl/gen"

	"github.com/gofrs/uuid"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ChangeRequestToProto(e main_entity.ChangeRequest) *pb.ChangeRequest {
	return &pb.ChangeRequest{
		Uuid:          e.UUID.String(),
		Version:       int64(e.Version),
		Title:         e.Title,
		Description:   e.Description,
		ReviewType:    pb.ChangeRequestReviewType(e.ReviewType),
		RefUuid:       e.RefUUID.String(),
		OldData:       e.OldData,
		OldDataRef:    e.OldDataRef,
		NewData:       e.NewData,
		NewDataRef:    e.NewDataRef,
		Reviews:       ChangeRequestReviewSliceToProto(e.Reviews),
		OwnerUuid:     e.OwnerUUID.String(),
		Status:        pb.ChangeRequestStatus(e.Status),
		CreatedAt:     timestamppb.New(e.CreatedAt),
		UpdatedAt:     timestamppb.New(e.UpdatedAt),
		CreatedByUuid: e.CreatedByUUID.String(),
		UpdatedByUuid: e.UpdatedByUUID.String(),
	}
}

func ChangeRequestSliceToProto(es []main_entity.ChangeRequest) []*pb.ChangeRequest {
	res := []*pb.ChangeRequest{}
	for _, e := range es {
		res = append(res, ChangeRequestToProto(e))
	}
	return res
}

func ChangeRequestFromProto(m *pb.ChangeRequest) main_entity.ChangeRequest {
	if m == nil {
		return main_entity.ChangeRequest{}
	}
	return main_entity.ChangeRequest{
		UUID:          uuid.FromStringOrNil(m.GetUuid()),
		Version:       int64(m.GetVersion()),
		Title:         m.GetTitle(),
		Description:   m.GetDescription(),
		ReviewType:    main_entity.ReviewType(m.GetReviewType()),
		RefUUID:       uuid.FromStringOrNil(m.GetRefUuid()),
		OldData:       m.GetOldData(),
		OldDataRef:    m.GetOldDataRef(),
		NewData:       m.GetNewData(),
		NewDataRef:    m.GetNewDataRef(),
		Reviews:       ChangeRequestReviewSliceFromProto(m.GetReviews()),
		OwnerUUID:     uuid.FromStringOrNil(m.GetOwnerUuid()),
		Status:        main_entity.Status(m.GetStatus()),
		CreatedAt:     m.GetCreatedAt().AsTime(),
		UpdatedAt:     m.GetUpdatedAt().AsTime(),
		CreatedByUUID: uuid.FromStringOrNil(m.GetCreatedByUuid()),
		UpdatedByUUID: uuid.FromStringOrNil(m.GetUpdatedByUuid()),
	}
}

func ChangeRequestSliceFromProto(es []*pb.ChangeRequest) []main_entity.ChangeRequest {
	if es == nil {
		return []main_entity.ChangeRequest{}
	}
	res := []main_entity.ChangeRequest{}
	for _, e := range es {
		res = append(res, ChangeRequestFromProto(e))
	}
	return res
}
