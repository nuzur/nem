package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/user"
	pb "github.com/nuzur/nem/idl/gen"

	"github.com/gofrs/uuid"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func UserToProto(e main_entity.User) *pb.User {
	return &pb.User{
		Uuid:        e.UUID.String(),
		Identifier:  e.Identifier,
		Name:        StringPtrToString(e.Name),
		LastName:    StringPtrToString(e.LastName),
		Email:       e.Email,
		UserType:    pb.UserUserType(e.UserType),
		CountryIos2: StringPtrToString(e.CountryIos2),
		Locale:      StringPtrToString(e.Locale),
		Metadata:    StringPtrToString(e.Metadata),
		Status:      pb.UserStatus(e.Status),
		CreatedAt:   timestamppb.New(e.CreatedAt),
		UpdatedAt:   timestamppb.New(e.UpdatedAt),
	}
}

func UserSliceToProto(es []main_entity.User) []*pb.User {
	res := []*pb.User{}
	for _, e := range es {
		res = append(res, UserToProto(e))
	}
	return res
}

func UserFromProto(m *pb.User) main_entity.User {
	if m == nil {
		return main_entity.User{}
	}
	return main_entity.User{
		UUID:        uuid.FromStringOrNil(m.GetUuid()),
		Identifier:  m.GetIdentifier(),
		Name:        &m.Name,
		LastName:    &m.LastName,
		Email:       m.GetEmail(),
		UserType:    main_entity.UserType(m.GetUserType()),
		CountryIos2: &m.CountryIos2,
		Locale:      &m.Locale,
		Metadata:    &m.Metadata,
		Status:      main_entity.Status(m.GetStatus()),
		CreatedAt:   m.GetCreatedAt().AsTime(),
		UpdatedAt:   m.GetUpdatedAt().AsTime(),
	}
}

func UserSliceFromProto(es []*pb.User) []main_entity.User {
	if es == nil {
		return []main_entity.User{}
	}
	res := []main_entity.User{}
	for _, e := range es {
		res = append(res, UserFromProto(e))
	}
	return res
}
