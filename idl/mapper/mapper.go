package mapper

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func JSONRawToString(raw json.RawMessage) string {
	j, err := json.Marshal(&raw)
	if err != nil {
		return ""
	}
	return string(j)
}

func StringToJSONRaw(s string) json.RawMessage {
	var r json.RawMessage
	err := json.Unmarshal([]byte(s), &r)
	if err != nil {
		return json.RawMessage{}
	}
	return r
}

func UUIDSliceToStringSlice(uuids []uuid.UUID) []string {
	res := []string{}
	for _, u := range uuids {
		res = append(res, u.String())
	}
	return res
}

func TimeSliceToProtoTimeSlice(ts []time.Time) []*timestamppb.Timestamp {
	res := []*timestamppb.Timestamp{}
	for _, t := range ts {
		res = append(res, timestamppb.New(t))
	}
	return res
}

func StringSliceToUUIDSlice(ss []string) []uuid.UUID {
	res := []uuid.UUID{}
	for _, s := range ss {
		res = append(res, uuid.FromStringOrNil(s))
	}
	return res
}

func ProtoTimeSliceToTimeSlice(ts []*timestamppb.Timestamp) []time.Time {
	res := []time.Time{}
	for _, t := range ts {
		res = append(res, t.AsTime())
	}
	return res
}

func StringPtrToString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
