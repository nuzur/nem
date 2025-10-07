package extension_execution

import (
	"context"
	//"fmt"
	main_entity "github.com/nuzur/nem/core/entity/extension_execution"
	"github.com/nuzur/nem/core/module/extension_execution/types"
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
		main_entity.ExtensionExecution{},
		false,
		optConfig.ListIncludeColumns,
		optConfig.ListExcludeColumns)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "list_extension_execution_query",
			Message:          "error in building query for ListExtensionExecution",
			EntityIdentifier: "extension_execution",
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
			ActionIdentifier: "list_extension_execution_sort_buffer",
			Message:          "error in setting sort buffer for ListExtensionExecution",
			EntityIdentifier: "extension_execution",
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
			ActionIdentifier: "list_extension_execution_db",
			Message:          "error in executing query for ListExtensionExecution",
			EntityIdentifier: "extension_execution",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             request,
			ExtraData:        map[string]string{"query": query},
			Error:            err,
		})
		return types.ListResponse{}, err
	}
	defer rows.Close()
	var items []repogen.ExtensionExecution
	for rows.Next() {
		var i repogen.ExtensionExecution
		fields := []any{}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "uuid") {
				fields = append(fields, &i.UUID)
				continue
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "uuid") {
				fields = append(fields, &i.UUID)
				continue
			}
		} else {
			fields = append(fields, &i.UUID)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "extension_uuid") {
				fields = append(fields, &i.ExtensionUUID)
				continue
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "extension_uuid") {
				fields = append(fields, &i.ExtensionUUID)
				continue
			}
		} else {
			fields = append(fields, &i.ExtensionUUID)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "extension_version_uuid") {
				fields = append(fields, &i.ExtensionVersionUUID)
				continue
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "extension_version_uuid") {
				fields = append(fields, &i.ExtensionVersionUUID)
				continue
			}
		} else {
			fields = append(fields, &i.ExtensionVersionUUID)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "project_extension_uuid") {
				fields = append(fields, &i.ProjectExtensionUUID)
				continue
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "project_extension_uuid") {
				fields = append(fields, &i.ProjectExtensionUUID)
				continue
			}
		} else {
			fields = append(fields, &i.ProjectExtensionUUID)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "project_uuid") {
				fields = append(fields, &i.ProjectUUID)
				continue
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "project_uuid") {
				fields = append(fields, &i.ProjectUUID)
				continue
			}
		} else {
			fields = append(fields, &i.ProjectUUID)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "project_version_uuid") {
				fields = append(fields, &i.ProjectVersionUUID)
				continue
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "project_version_uuid") {
				fields = append(fields, &i.ProjectVersionUUID)
				continue
			}
		} else {
			fields = append(fields, &i.ProjectVersionUUID)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "executed_by_uuid") {
				fields = append(fields, &i.ExecutedByUUID)
				continue
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "executed_by_uuid") {
				fields = append(fields, &i.ExecutedByUUID)
				continue
			}
		} else {
			fields = append(fields, &i.ExecutedByUUID)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "metadata") {
				fields = append(fields, &i.Metadata)
				continue
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "metadata") {
				fields = append(fields, &i.Metadata)
				continue
			}
		} else {
			fields = append(fields, &i.Metadata)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "status") {
				fields = append(fields, &i.Status)
				continue
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "status") {
				fields = append(fields, &i.Status)
				continue
			}
		} else {
			fields = append(fields, &i.Status)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "status_msg") {
				fields = append(fields, &i.StatusMsg)
				continue
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "status_msg") {
				fields = append(fields, &i.StatusMsg)
				continue
			}
		} else {
			fields = append(fields, &i.StatusMsg)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "created_at") {
				fields = append(fields, &i.CreatedAt)
				continue
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "created_at") {
				fields = append(fields, &i.CreatedAt)
				continue
			}
		} else {
			fields = append(fields, &i.CreatedAt)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "updated_at") {
				fields = append(fields, &i.UpdatedAt)
				continue
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "updated_at") {
				fields = append(fields, &i.UpdatedAt)
				continue
			}
		} else {
			fields = append(fields, &i.UpdatedAt)
		}

		if err := rows.Scan(fields...); err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "list_extension_execution_scan",
				Message:          "error in scanning rows for ListExtensionExecution",
				EntityIdentifier: "extension_execution",
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
			ActionIdentifier: "list_extension_execution_rows",
			Message:          "error closing rows for ListExtensionExecution",
			EntityIdentifier: "extension_execution",
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
			ActionIdentifier: "list_extension_execution_rows",
			Message:          "error in rows for ListExtensionExecution",
			EntityIdentifier: "extension_execution",
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
			ActionIdentifier: "list_extension_execution_next",
			Message:          "error determining if list has next page for ListExtensionExecution",
			EntityIdentifier: "extension_execution",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             request,
			ExtraData:        map[string]string{"query": query},
			Error:            err,
		})
		return types.ListResponse{}, err
	}

	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "list_extension_execution",
		Message:          "successfully handled ListExtensionExecution",
		EntityIdentifier: "extension_execution",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeSuccess,
		Data:             request,
		ExtraData:        map[string]string{"query": query},
	})
	return types.ListResponse{
		ExtensionExecutions: mapModelsToEntities(items),
		HasNextPage:         hasNextPage,
	}, nil
}

func (m *module) ListCount(ctx context.Context,
	request types.ListRequest,
	opts ...Option) (int64, error) {
	query, err := m.repository.BuildListEntityQuery(ctx, request, main_entity.ExtensionExecution{}, true, nil, nil)
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
