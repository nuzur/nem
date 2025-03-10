package list

import (
	"context"
	"errors"
	"fmt"
	"strings"

	entitytypes "github.com/nuzur/nem/core/entity/types"
	"go.einride.tech/aip/filtering"
	expr "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
)

type Implementation struct {
}

func New() *Implementation {
	return &Implementation{}
}

func (i *Implementation) BuildListEntityQuery(
	ctx context.Context,
	request ListRequest,
	entity ListEntity,
	onlyCount bool,
) (string, error) {

	if request.GetFilter().CheckedExpr == nil || request.GetFilter().CheckedExpr.Expr == nil {
		if onlyCount {
			return fmt.Sprintf(
				"SELECT DISTINCT count(%s.%s) FROM %s",
				entity.EntityIdentifier(),
				entity.PrimaryKeyIdentifier(),
				entity.EntityIdentifier(),
			), nil
		} else {
			return fmt.Sprintf(
				"SELECT DISTINCT %s.* FROM %s limit 10",
				entity.EntityIdentifier(),
				entity.EntityIdentifier(),
			), nil
		}
	}

	jsonTables := make(map[string]string)
	whereClause, err := buildWhereClauses(
		request.GetFilter().CheckedExpr.Expr,
		entity,
		request.GetFilteringDeclarations(),
		jsonTables)
	if err != nil {
		return "", err
	}

	orderByClause, err := buildOrderByClause(request, entity, jsonTables)
	if err != nil {
		return "", err
	}

	jsonTablesSlice := []string{}
	for _, jt := range jsonTables {
		jsonTablesSlice = append(jsonTablesSlice, jt)
	}

	jsonTablesFinal := strings.Join(jsonTablesSlice, " ")

	limit := buildLimit(request)

	if request.GetJoinParams() != nil {
		finalWhereClause := fmt.Sprintf("%s AND (%s)", whereClause, request.GetJoinParams().JoinWhereClauses)
		if onlyCount {
			return fmt.Sprintf(
				"SELECT DISTINCT count(%s.%s) FROM %s %s %s WHERE %s",
				entity.EntityIdentifier(),
				entity.PrimaryKeyIdentifier(),
				entity.EntityIdentifier(),
				request.GetJoinParams().JoinStatement,
				jsonTablesFinal,
				finalWhereClause,
			), nil
		}
		return fmt.Sprintf(
			"SELECT DISTINCT %s.* FROM %s %s %s WHERE %s %s %s",
			entity.EntityIdentifier(),
			entity.EntityIdentifier(),
			request.GetJoinParams().JoinStatement,
			jsonTablesFinal,
			finalWhereClause,
			orderByClause,
			limit,
		), nil
	}
	if onlyCount {
		return fmt.Sprintf(
			"SELECT DISTINCT count(%s.%s) FROM %s %s WHERE %s",
			entity.EntityIdentifier(),
			entity.PrimaryKeyIdentifier(),
			entity.EntityIdentifier(),
			jsonTablesFinal,
			whereClause,
		), nil
	}
	return fmt.Sprintf(
		"SELECT DISTINCT %s.* FROM %s %s WHERE %s %s %s",
		entity.EntityIdentifier(),
		entity.EntityIdentifier(),
		jsonTablesFinal,
		whereClause,
		orderByClause,
		limit,
	), nil
}

func buildWhereClauses(ex *expr.Expr, entity ListEntity, declarations *filtering.Declarations, jsonTables map[string]string) (string, error) {
	cex := ex.GetCallExpr()
	if cex != nil {
		if isBaseFunction(cex.Function) {
			res, err := buildSingleClause(cex, entity, declarations)
			if err != nil {
				return "", nil
			}
			if res.JSONTable != "" {
				if _, found := jsonTables[res.JSONTableName]; !found {
					jsonTables[res.JSONTableName] = res.JSONTable
				}
			}
			return res.ResolvedClause, nil
		}

		operator := ""
		switch cex.Function {
		case filtering.FunctionAnd:
			operator = "AND"
		case filtering.FunctionOr:
			operator = "OR"
		}

		left, err := buildWhereClauses(cex.Args[0], entity, declarations, jsonTables)
		if err != nil {
			return "", err
		}
		right, err := buildWhereClauses(cex.Args[1], entity, declarations, jsonTables)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("(%s %s %s)", left, operator, right), nil
	}
	return "", errors.New("invalid call expression")
}

func buildSingleClause(cex *expr.Expr_Call, entity ListEntity, declarations *filtering.Declarations) (SingleClauseResponse, error) {
	if len(cex.Args) != 2 {
		return SingleClauseResponse{}, errors.New("invalid argument count")
	}
	fieldIdentifier := ""
	isDependant := false
	dependantField := ""
	if cex.Args[0].GetIdentExpr() == nil {
		if cex.Args[0].GetSelectExpr() != nil {
			isDependant = true
			fieldIdentifier = cex.Args[0].GetSelectExpr().GetOperand().GetIdentExpr().GetName()
			dependantField = cex.Args[0].GetSelectExpr().GetField()
		}
	} else {
		fieldIdentifier = cex.Args[0].GetIdentExpr().Name
	}

	request := SingleClauseRequest{
		cex:                     cex,
		declarations:            declarations,
		fieldIdentifier:         fieldIdentifier,
		isDependant:             isDependant,
		dependantFieldIdentifer: dependantField,
		entity:                  entity,
	}

	fieldMap := entity.FieldIdentfierToTypeMap()
	fieldType := fieldMap[fieldIdentifier]
	return handleClauseByType(request, entity, fieldType)
}

func isBaseFunction(f string) bool {
	if f == filtering.FunctionEquals ||
		f == filtering.FunctionGreaterEquals ||
		f == filtering.FunctionGreaterThan ||
		f == filtering.FunctionLessEquals ||
		f == filtering.FunctionLessThan ||
		f == filtering.FunctionNot ||
		f == filtering.FunctionNotEquals ||
		f == filtering.FunctionHas ||
		f == filtering.FunctionDuration ||
		f == filtering.FunctionTimestamp {
		return true
	}
	return false
}

func buildOrderByClause(request ListRequest, entity ListEntity, jsonTables map[string]string) (string, error) {
	ofs, err := isValidOrderby(request, entity)
	if err != nil {
		return "", err
	}

	res := ""
	if len(ofs) > 0 {
		res = "ORDER BY "
		orderByPairs := []string{}
		for _, f := range ofs {
			order := "ASC"
			if f.desc {
				order = "DESC"
			}
			if !f.isDependant {
				orderByPairs = append(orderByPairs, fmt.Sprintf("%s %s", f.fieldIdentifier, order))
			} else if f.isDependantMulti {
				/*
					json_table_name, json_table := resolveJSONTable(SingleClauseRequest{
						fieldIdentifier:         f.fieldIdentifier,
						dependantFieldIdentifer: f.dependantFieldIdentifer,
						isDependant:             f.isDependant,
						isDependantMulti:        f.isDependantMulti,
					}, entitytypes.FieldTypeToSQL(f.fieldType))
					jsonTables[json_table_name] = json_table
					orderByPairs = append(orderByPairs, fmt.Sprintf("%s.%s %s", json_table_name, f.dependantFieldIdentifer, order))
				*/
				return "", errors.New("dependant list entities are not supported in sorting")
			} else {
				left := buildJSONExtract(SingleClauseRequest{
					fieldIdentifier:         f.fieldIdentifier,
					dependantFieldIdentifer: f.dependantFieldIdentifer,
					isDependant:             f.isDependant,
					isDependantMulti:        f.isDependantMulti,
				})
				orderByPairs = append(orderByPairs, fmt.Sprintf("%s %s", left, order))
			}
		}
		res += strings.Join(orderByPairs, ", ")
	}
	return res, nil
}

func isValidOrderby(request ListRequest, entity ListEntity) ([]OrderByField, error) {
	res := []OrderByField{}
	fields := request.GetOrderBy().Fields
	fieldsMap := entity.FieldIdentfierToTypeMap()
	dependantFieldMap := entity.DependantFieldIdentifierToTypeMap()
	if len(fields) > 0 {
		for _, f := range fields {

			if strings.Contains(f.Path, ".") {
				parts := strings.Split(f.Path, ".")
				if len(parts) != 2 {
					return res, errors.New("invalid dependant sort by field")
				}

				subFieldsMap, found := dependantFieldMap[parts[0]]
				if !found {
					return res, errors.New("invalid sort by field (2)")
				}
				subFieldType, found := subFieldsMap[parts[1]]
				if !found {
					return res, errors.New("invalid sort by field (3)")
				}

				isDependantMulti := false
				if fieldsMap[parts[0]] == entitytypes.MultiDependantEntityFieldType {
					isDependantMulti = true
				}
				res = append(res, OrderByField{
					path:                    f.Path,
					isDependant:             true,
					isDependantMulti:        isDependantMulti,
					fieldIdentifier:         parts[0],
					dependantFieldIdentifer: parts[1],
					desc:                    f.Desc,
					fieldType:               subFieldType,
				})
			} else {
				fieldType, found := fieldsMap[f.Path]
				if !found {
					return res, errors.New("invalid sort by field (1)")
				}
				res = append(res, OrderByField{
					path:            f.Path,
					fieldIdentifier: f.Path,
					desc:            f.Desc,
					fieldType:       fieldType,
				})

			}
		}
	}
	return res, nil
}

func buildLimit(request ListRequest) string {
	return fmt.Sprintf("LIMIT %d, %d ", request.GetOffset(), request.GetPageSize())
}
