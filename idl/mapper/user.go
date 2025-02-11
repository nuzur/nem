package mapper

import (
	main_entity "nem/core/entity/user"
	pb "nem/idl/gen"

	"github.com/gofrs/uuid"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func UserToProto(e main_entity.User) *pb.User {
	return &pb.User{
		Uuid:        e.UUID.String(),
		Identifier:  e.Identifier,
		Name:        e.Name,
		LastName:    e.LastName,
		Email:       e.Email,
		UserType:    pb.UserUserType(e.UserType),
		CountryIos2: e.CountryIos2,
		Locale:      e.Locale,
		Metadata:    e.Metadata,
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
		Name:        m.GetName(),
		LastName:    m.GetLastName(),
		Email:       m.GetEmail(),
		UserType:    main_entity.UserType(m.GetUserType()),
		CountryIos2: m.GetCountryIos2(),
		Locale:      m.GetLocale(),
		Metadata:    m.GetMetadata(),
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
