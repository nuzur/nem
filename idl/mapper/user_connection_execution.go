package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/user_connection_execution"
	pb "github.com/nuzur/nem/idl/gen"

	"github.com/gofrs/uuid"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func UserConnectionExecutionToProto(e main_entity.UserConnectionExecution) *pb.UserConnectionExecution {
	return &pb.UserConnectionExecution{
		Uuid:          e.UUID.String(),
		Status:        pb.UserConnectionExecutionStatus(e.Status),
		ResultsPath:   StringPtrToString(e.ResultsPath),
		NumResults:    int64(e.NumResults),
		StartedAt:     timestamppb.New(e.StartedAt),
		FinishedAt:    timestamppb.New(e.FinishedAt),
		EstimatedTime: StringPtrToString(e.EstimatedTime),
		UserMsg:       StringPtrToString(e.UserMsg),
	}
}

func UserConnectionExecutionSliceToProto(es []main_entity.UserConnectionExecution) []*pb.UserConnectionExecution {
	res := []*pb.UserConnectionExecution{}
	for _, e := range es {
		res = append(res, UserConnectionExecutionToProto(e))
	}
	return res
}

func UserConnectionExecutionFromProto(m *pb.UserConnectionExecution) main_entity.UserConnectionExecution {
	if m == nil {
		return main_entity.UserConnectionExecution{}
	}
	return main_entity.UserConnectionExecution{
		UUID:          uuid.FromStringOrNil(m.GetUuid()),
		Status:        main_entity.Status(m.GetStatus()),
		ResultsPath:   &m.ResultsPath,
		NumResults:    int64(m.GetNumResults()),
		StartedAt:     m.GetStartedAt().AsTime(),
		FinishedAt:    m.GetFinishedAt().AsTime(),
		EstimatedTime: &m.EstimatedTime,
		UserMsg:       &m.UserMsg,
	}
}

func UserConnectionExecutionSliceFromProto(es []*pb.UserConnectionExecution) []main_entity.UserConnectionExecution {
	if es == nil {
		return []main_entity.UserConnectionExecution{}
	}
	res := []main_entity.UserConnectionExecution{}
	for _, e := range es {
		res = append(res, UserConnectionExecutionFromProto(e))
	}
	return res
}
