package user_connection

import (
	"context"
	//"fmt"
	main_entity "github.com/nuzur/nem/core/entity/user_connection"
	"github.com/nuzur/nem/core/module/user_connection/types"
	repogen "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
	"slices"
)

func (m *module) List(ctx context.Context,
	request types.ListRequest,
	opts ...Option) (types.ListResponse, error) {

	optConfig := applyAllOptions(opts)
	query, err := m.repository.BuildListEntityQuery(
		ctx,
		request,
		main_entity.UserConnection{},
		false,
		optConfig.ListIncludeColumns,
		optConfig.ListExcludeColumns)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "list_user_connection_query",
			Message:          "error in building query for ListUserConnection",
			EntityIdentifier: "user_connection",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             request,
			Error:            err,
		})
		return types.ListResponse{}, err
	}

	// increase sort buffer size
	// TODO make this configurable
	txn, _ := m.repository.DB.Begin()
	defer txn.Commit()
	bufferRows, err := txn.QueryContext(ctx, "SET sort_buffer_size=5000000")
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "list_user_connection_sort_buffer",
			Message:          "error in setting sort buffer for ListUserConnection",
			EntityIdentifier: "user_connection",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             request,
			Error:            err,
		})
		return types.ListResponse{}, err
	}
	bufferRows.Close()
	rows, err := txn.QueryContext(ctx, query)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "list_user_connection_db",
			Message:          "error in executing query for ListUserConnection",
			EntityIdentifier: "user_connection",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             request,
			ExtraData:        map[string]string{"query": query},
			Error:            err,
		})
		return types.ListResponse{}, err
	}
	defer rows.Close()
	var items []repogen.UserConnection
	for rows.Next() {
		var i repogen.UserConnection
		fields := []any{}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "uuid") {
				fields = append(fields, &i.UUID)
			} else {
				var skip interface{}
				fields = append(fields, &skip)
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "uuid") {
				fields = append(fields, &i.UUID)
			} else {
				var skip interface{}
				fields = append(fields, &skip)
			}
		} else {
			fields = append(fields, &i.UUID)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "user_uuid") {
				fields = append(fields, &i.UserUUID)
			} else {
				var skip interface{}
				fields = append(fields, &skip)
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "user_uuid") {
				fields = append(fields, &i.UserUUID)
			} else {
				var skip interface{}
				fields = append(fields, &skip)
			}
		} else {
			fields = append(fields, &i.UserUUID)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "project_uuid") {
				fields = append(fields, &i.ProjectUUID)
			} else {
				var skip interface{}
				fields = append(fields, &skip)
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "project_uuid") {
				fields = append(fields, &i.ProjectUUID)
			} else {
				var skip interface{}
				fields = append(fields, &skip)
			}
		} else {
			fields = append(fields, &i.ProjectUUID)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "project_version_uuid") {
				fields = append(fields, &i.ProjectVersionUUID)
			} else {
				var skip interface{}
				fields = append(fields, &skip)
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "project_version_uuid") {
				fields = append(fields, &i.ProjectVersionUUID)
			} else {
				var skip interface{}
				fields = append(fields, &skip)
			}
		} else {
			fields = append(fields, &i.ProjectVersionUUID)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "type") {
				fields = append(fields, &i.Type)
			} else {
				var skip interface{}
				fields = append(fields, &skip)
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "type") {
				fields = append(fields, &i.Type)
			} else {
				var skip interface{}
				fields = append(fields, &skip)
			}
		} else {
			fields = append(fields, &i.Type)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "type_config") {
				fields = append(fields, &i.TypeConfig)
			} else {
				var skip interface{}
				fields = append(fields, &skip)
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "type_config") {
				fields = append(fields, &i.TypeConfig)
			} else {
				var skip interface{}
				fields = append(fields, &skip)
			}
		} else {
			fields = append(fields, &i.TypeConfig)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "db_schema") {
				fields = append(fields, &i.DbSchema)
			} else {
				var skip interface{}
				fields = append(fields, &skip)
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "db_schema") {
				fields = append(fields, &i.DbSchema)
			} else {
				var skip interface{}
				fields = append(fields, &skip)
			}
		} else {
			fields = append(fields, &i.DbSchema)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "executions") {
				fields = append(fields, &i.Executions)
			} else {
				var skip interface{}
				fields = append(fields, &skip)
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "executions") {
				fields = append(fields, &i.Executions)
			} else {
				var skip interface{}
				fields = append(fields, &skip)
			}
		} else {
			fields = append(fields, &i.Executions)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "status") {
				fields = append(fields, &i.Status)
			} else {
				var skip interface{}
				fields = append(fields, &skip)
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "status") {
				fields = append(fields, &i.Status)
			} else {
				var skip interface{}
				fields = append(fields, &skip)
			}
		} else {
			fields = append(fields, &i.Status)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "created_at") {
				fields = append(fields, &i.CreatedAt)
			} else {
				var skip interface{}
				fields = append(fields, &skip)
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "created_at") {
				fields = append(fields, &i.CreatedAt)
			} else {
				var skip interface{}
				fields = append(fields, &skip)
			}
		} else {
			fields = append(fields, &i.CreatedAt)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "updated_at") {
				fields = append(fields, &i.UpdatedAt)
			} else {
				var skip interface{}
				fields = append(fields, &skip)
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "updated_at") {
				fields = append(fields, &i.UpdatedAt)
			} else {
				var skip interface{}
				fields = append(fields, &skip)
			}
		} else {
			fields = append(fields, &i.UpdatedAt)
		}

		if err := rows.Scan(fields...); err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "list_user_connection_scan",
				Message:          "error in scanning rows for ListUserConnection",
				EntityIdentifier: "user_connection",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             request,
				ExtraData:        map[string]string{"query": query},
				Error:            err,
			})
			return types.ListResponse{}, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "list_user_connection_rows",
			Message:          "error closing rows for ListUserConnection",
			EntityIdentifier: "user_connection",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             request,
			ExtraData:        map[string]string{"query": query},
			Error:            err,
		})
		return types.ListResponse{}, err
	}
	if err := rows.Err(); err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "list_user_connection_rows",
			Message:          "error in rows for ListUserConnection",
			EntityIdentifier: "user_connection",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             request,
			ExtraData:        map[string]string{"query": query},
			Error:            err,
		})
		return types.ListResponse{}, err
	}

	hasNextPage, err := m.listHasNextPage(ctx, request, opts...)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "list_user_connection_next",
			Message:          "error determining if list has next page for ListUserConnection",
			EntityIdentifier: "user_connection",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             request,
			ExtraData:        map[string]string{"query": query},
			Error:            err,
		})
		return types.ListResponse{}, err
	}

	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "list_user_connection",
		Message:          "successfully handled ListUserConnection",
		EntityIdentifier: "user_connection",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeSuccess,
		Data:             request,
		ExtraData:        map[string]string{"query": query},
	})
	return types.ListResponse{
		UserConnections: mapModelsToEntities(items),
		HasNextPage:     hasNextPage,
	}, nil
}

func (m *module) ListCount(ctx context.Context,
	request types.ListRequest,
	opts ...Option) (int64, error) {
	query, err := m.repository.BuildListEntityQuery(ctx, request, main_entity.UserConnection{}, true, nil, nil)
	if err != nil {
		return -1, err
	}

	//fmt.Printf("query count: %s \n", query)

	var count int64
	err = m.repository.DB.QueryRow(query).Scan(&count)
	if err != nil {
		return -1, err
	}

	return count, nil
}

func (m *module) listHasNextPage(ctx context.Context,
	request types.ListRequest,
	opts ...Option) (bool, error) {
	count, err := m.ListCount(ctx, request, opts...)
	if err != nil {
		return false, err
	}

	if request.GetOffset()+int64(request.GetPageSize()) < count {
		return true, nil
	}
	return false, nil
}
