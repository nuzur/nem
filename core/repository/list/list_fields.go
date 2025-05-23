package list

import (
	"errors"
	"fmt"

	entitytypes "github.com/nuzur/nem/core/entity/types"
	"go.einride.tech/aip/filtering"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func handleClauseByType(req SingleClauseRequest, entity ListEntity, fieldType entitytypes.FieldType) (SingleClauseResponse, error) {
	switch fieldType {
	case entitytypes.StringFieldType, entitytypes.UUIDFieldType, entitytypes.RawJSONFieldType:
		return buildStringClause(req)
	case entitytypes.SingleEnumFieldType:
		return buildEnumClause(req)
	case entitytypes.MultiEnumFieldType:
		req.isDependantMulti = true
		return buildEnumClause(req)
	case entitytypes.ArrayFieldType:
		req.isArray = true
		arrayFieldTypes := entity.ArrayFieldIdentifierToType()
		req.arrayType = arrayFieldTypes[req.fieldIdentifier]
		return buildArrayClause(req)
	case entitytypes.SingleDependantEntityFieldType:
		typeMap := entity.DependantFieldIdentifierToTypeMap()
		dependantTypeMap := typeMap[req.fieldIdentifier]
		dependantFieldType := dependantTypeMap[req.dependantFieldIdentifer]
		return handleClauseByType(req, entity, dependantFieldType)
	case entitytypes.MultiDependantEntityFieldType:
		typeMap := entity.DependantFieldIdentifierToTypeMap()
		dependantTypeMap := typeMap[req.fieldIdentifier]
		dependantFieldType := dependantTypeMap[req.dependantFieldIdentifer]
		req.isDependantMulti = true
		return handleClauseByType(req, entity, dependantFieldType)
	case entitytypes.IntFieldType:
		return buildIntClause(req)
	case entitytypes.FloatFieldType:
		return buildFloatClause(req)
	case entitytypes.BooleanFieldType:
		return buildBooleanClause(req)
	case entitytypes.TimestampFieldType:
		return buildTimestampClause(req)
	}
	return SingleClauseResponse{}, errors.New("unsupported field type")
}

func buildStringClause(req SingleClauseRequest) (SingleClauseResponse, error) {
	cex := req.cex
	left := buildLeftOperator(req)
	switch cex.Function {
	case filtering.FunctionEquals, filtering.FunctionNotEquals:
		return SingleClauseResponse{
			ResolvedClause: fmt.Sprintf("%s %s '%s'", left, cex.Function, cex.Args[1].GetConstExpr().GetStringValue()),
		}, nil
	case filtering.FunctionHas:
		return SingleClauseResponse{
			ResolvedClause: fmt.Sprintf("%s like '%%%s%%'", left, cex.Args[1].GetConstExpr().GetStringValue()),
		}, nil
	}
	return SingleClauseResponse{}, errors.New("unsupported function for type string")
}

func buildEnumClause(req SingleClauseRequest) (SingleClauseResponse, error) {
	cex := req.cex
	declarations := req.declarations
	left := buildLeftOperator(req)
	name := req.fieldIdentifier
	if req.isDependant {
		name = fmt.Sprintf("%s.%s", req.fieldIdentifier, req.dependantFieldIdentifer)
	}

	enumType, found := declarations.LookupEnumIdent(name)
	if found {
		value := cex.Args[1].GetIdentExpr().Name
		numValue := enumType.Descriptor().Values().ByName(protoreflect.Name(value)).Number()
		switch cex.Function {
		case filtering.FunctionEquals, filtering.FunctionNotEquals:
			if req.isDependant && req.isDependantMulti {
				return SingleClauseResponse{
					ResolvedClause: fmt.Sprintf("JSON_CONTAINS(%s, '%d', '$')", left, numValue),
				}, nil
			}
			return SingleClauseResponse{
				ResolvedClause: fmt.Sprintf("%s %s %d", left, cex.Function, numValue),
			}, nil
		}
		return SingleClauseResponse{}, errors.New("unsupported function for type single enum")
	}

	return SingleClauseResponse{}, errors.New("enum declaration not found")
}

func buildIntClause(req SingleClauseRequest) (SingleClauseResponse, error) {
	cex := req.cex
	left := buildLeftOperator(req)
	right := cex.Args[1].GetConstExpr().GetInt64Value()
	sqlType := "INT"
	if req.isDependant {
		if req.isDependantMulti {
			json_table_name, json_table := resolveJSONTable(req, sqlType)
			return SingleClauseResponse{
				ResolvedClause: fmt.Sprintf("%s.%s %s %d", json_table_name, req.dependantFieldIdentifer, cex.Function, right),
				JSONTable:      json_table,
				JSONTableName:  json_table_name,
			}, nil
		}
		return SingleClauseResponse{
			ResolvedClause: fmt.Sprintf("convert(%s, %s) %s %d", left, sqlType, cex.Function, right),
		}, nil
	}
	return SingleClauseResponse{
		ResolvedClause: fmt.Sprintf("%s %s %d", left, cex.Function, right),
	}, nil
}

func buildFloatClause(req SingleClauseRequest) (SingleClauseResponse, error) {
	cex := req.cex
	left := buildLeftOperator(req)
	right := cex.Args[1].GetConstExpr().GetDoubleValue()
	sqlType := "DOUBLE"
	if req.isDependant {
		if req.isDependantMulti {
			json_table_name, json_table := resolveJSONTable(req, sqlType)
			return SingleClauseResponse{
				ResolvedClause: fmt.Sprintf("%s.%s %s %f", json_table_name, req.dependantFieldIdentifer, cex.Function, right),
				JSONTable:      json_table,
				JSONTableName:  json_table_name,
			}, nil
		}
		return SingleClauseResponse{
			ResolvedClause: fmt.Sprintf("convert(%s, %s) %s %f", left, sqlType, cex.Function, right),
		}, nil
	}
	return SingleClauseResponse{
		ResolvedClause: fmt.Sprintf("%s %s %f", left, cex.Function, right),
	}, nil
}

func buildBooleanClause(req SingleClauseRequest) (SingleClauseResponse, error) {
	cex := req.cex
	if cex.Function != filtering.FunctionEquals {
		return SingleClauseResponse{}, errors.New("unsupported function in dependant entity")
	}
	left := buildLeftOperator(req)
	booleanIdent := cex.Args[1].GetIdentExpr().GetName()
	numValue := 0
	if booleanIdent == "true" {
		numValue = 1
	}
	if req.isDependant {
		if req.isDependantMulti {
			return SingleClauseResponse{
				ResolvedClause: fmt.Sprintf("JSON_CONTAINS(%s, '%s', '$')", left, booleanIdent),
			}, nil
		}
		return SingleClauseResponse{
			ResolvedClause: fmt.Sprintf("%s %s '%s'", left, cex.Function, booleanIdent),
		}, nil
	}
	return SingleClauseResponse{
		ResolvedClause: fmt.Sprintf("%s %s %d", left, cex.Function, numValue),
	}, nil
}

func buildTimestampClause(req SingleClauseRequest) (SingleClauseResponse, error) {
	cex := req.cex
	left := buildLeftOperator(req)
	right := cex.Args[1].GetConstExpr().GetStringValue()
	sqlType := "DATETIME"
	if req.isDependant {
		if req.isDependantMulti {
			json_table_name, json_table := resolveJSONTable(req, sqlType)
			return SingleClauseResponse{
				ResolvedClause: fmt.Sprintf("%s.%s %s '%s'", json_table_name, req.dependantFieldIdentifer, cex.Function, right),
				JSONTable:      json_table,
				JSONTableName:  json_table_name,
			}, nil
		}
		return SingleClauseResponse{
			ResolvedClause: fmt.Sprintf("convert(%s, %s) %s '%s'", left, sqlType, cex.Function, right),
		}, nil
	}

	return SingleClauseResponse{
		ResolvedClause: fmt.Sprintf("%s %s '%s'", left, cex.Function, right),
	}, nil
}

func buildArrayClause(req SingleClauseRequest) (SingleClauseResponse, error) {
	cex := req.cex
	switch cex.Function {
	case filtering.FunctionEquals, filtering.FunctionNotEquals:
		switch req.arrayType {
		case entitytypes.StringFieldType, entitytypes.UUIDFieldType, entitytypes.RawJSONFieldType:
			return SingleClauseResponse{
				ResolvedClause: fmt.Sprintf("JSON_CONTAINS(%s,'\"%s\"','$')", req.fieldIdentifier, cex.Args[1].GetConstExpr().GetStringValue()),
			}, nil
		case entitytypes.IntFieldType:
			return SingleClauseResponse{
				ResolvedClause: fmt.Sprintf("JSON_CONTAINS(%s,'%d','$')", req.fieldIdentifier, cex.Args[1].GetConstExpr().GetInt64Value()),
			}, nil
		case entitytypes.FloatFieldType:
			return SingleClauseResponse{
				ResolvedClause: fmt.Sprintf("JSON_CONTAINS(%s,'%f','$')", req.fieldIdentifier, cex.Args[1].GetConstExpr().GetDoubleValue()),
			}, nil
		case entitytypes.BooleanFieldType:
			return SingleClauseResponse{
				ResolvedClause: fmt.Sprintf("JSON_CONTAINS(%s,'%s','$')", req.fieldIdentifier, cex.Args[1].GetConstExpr().GetStringValue()),
			}, nil
		case entitytypes.TimestampFieldType:
			return SingleClauseResponse{
				ResolvedClause: fmt.Sprintf("JSON_CONTAINS(%s,'\"%s\"','$')", req.fieldIdentifier, cex.Args[1].GetConstExpr().GetStringValue()),
			}, nil
		}

	}
	return SingleClauseResponse{}, errors.New("unsupported function for type array")
}

// utilities
func buildLeftOperator(req SingleClauseRequest) string {
	if req.isDependant {
		return buildJSONExtract(req)
	}
	return fmt.Sprintf("%s.%s", req.entity.EntityIdentifier(), req.fieldIdentifier)
}

func buildJSONExtract(req SingleClauseRequest) string {
	if req.isDependantMulti {
		return fmt.Sprintf("JSON_EXTRACT(%s.%s, '$[*].%s')", req.entity.EntityIdentifier(), req.fieldIdentifier, req.dependantFieldIdentifer)
	}
	return fmt.Sprintf("JSON_EXTRACT(%s.%s, '$.%s')", req.entity.EntityIdentifier(), req.fieldIdentifier, req.dependantFieldIdentifer)
}

func resolveJSONTable(req SingleClauseRequest, sqlType string) (json_table_name, json_table string) {
	json_table_name = fmt.Sprintf("%s_%s", req.fieldIdentifier, req.dependantFieldIdentifer)
	json_table = fmt.Sprintf(`
			INNER JOIN JSON_TABLE(
				%s,
				'$[*]'
				COLUMNS(
					%s %s PATH '$.%s'            
				)
			) AS %s
			`, req.fieldIdentifier, req.dependantFieldIdentifer, sqlType, req.dependantFieldIdentifer, json_table_name)
	return
}
