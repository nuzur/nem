package list

import (
	"go.einride.tech/aip/filtering"
	"go.einride.tech/aip/ordering"

	entitytypes "github.com/nuzur/nem/core/entity/types"
	expr "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
)

type JoinParams struct {
	JoinStatement    string
	JoinWhereClauses string
}

type ListEntity interface {
	FieldIdentfierToTypeMap() map[string]entitytypes.FieldType
	DependantFieldIdentifierToTypeMap() map[string]map[string]entitytypes.FieldType
	EntityIdentifier() string
	PrimaryKeyIdentifier() string
	ArrayFieldIdentifierToType() map[string]entitytypes.FieldType
}

type ListRequest interface {
	GetFilter() filtering.Filter
	GetFilteringDeclarations() *filtering.Declarations
	GetOrderBy() ordering.OrderBy
	GetPageSize() int32
	GetOffset() int64
	GetJoinParams() *JoinParams
}

type SingleClauseRequest struct {
	cex                     *expr.Expr_Call
	declarations            *filtering.Declarations
	fieldIdentifier         string
	isDependant             bool
	isDependantMulti        bool
	dependantFieldIdentifer string
	isArray                 bool
	arrayType               entitytypes.FieldType
	entity                  ListEntity
}

type SingleClauseResponse struct {
	ResolvedClause string
	JSONTable      string
	JSONTableName  string
}

type OrderByField struct {
	path                    string
	isDependant             bool
	isDependantMulti        bool
	fieldIdentifier         string
	dependantFieldIdentifer string
	desc                    bool
	fieldType               entitytypes.FieldType
}
