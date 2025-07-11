package types

import (
	"encoding/json"
	main_entity "github.com/nuzur/nem/core/entity/user_project"

	"github.com/nuzur/nem/core/repository/list"
	"go.einride.tech/aip/filtering"
	"go.einride.tech/aip/ordering"
	"go.uber.org/zap/zapcore"
)

type ListRequest struct {
	PageSize              int32
	Offset                int64
	Filter                filtering.Filter
	FilteringDeclarations *filtering.Declarations
	OrderBy               ordering.OrderBy
	JoinParams            *list.JoinParams
}

func (r ListRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {
	e.AddInt32("page-size", r.PageSize)
	e.AddInt64("offset", r.Offset)
	filterJson, err := json.Marshal(r.Filter)
	if err != nil {
		return err
	}
	e.AddString("filter", string(filterJson))
	orderJson, err := json.Marshal(r.OrderBy)
	if err != nil {
		return err
	}
	e.AddString("OrderBy", string(orderJson))
	return nil
}

func (r ListRequest) GetPageSize() int32 {
	return r.PageSize
}

func (r ListRequest) GetOffset() int64 {
	return r.Offset
}

func (r ListRequest) GetFilter() filtering.Filter {
	return r.Filter
}

func (r ListRequest) GetFilteringDeclarations() *filtering.Declarations {
	return r.FilteringDeclarations
}

func (r ListRequest) GetOrderBy() ordering.OrderBy {
	return r.OrderBy
}

func (r ListRequest) GetJoinParams() *list.JoinParams {
	return r.JoinParams
}

type ListResponse struct {
	UserProjects []main_entity.UserProject
	HasNextPage  bool
}
