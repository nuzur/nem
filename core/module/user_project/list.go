package user_project

import (
	"context"
	//"fmt"
	main_entity "github.com/nuzur/nem/core/entity/user_project"
	"github.com/nuzur/nem/core/module/user_project/types"
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
		main_entity.UserProject{},
		false,
		optConfig.ListIncludeColumns,
		optConfig.ListExcludeColumns)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "list_user_project_query",
			Message:          "error in building query for ListUserProject",
			EntityIdentifier: "user_project",
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
			ActionIdentifier: "list_user_project_sort_buffer",
			Message:          "error in setting sort buffer for ListUserProject",
			EntityIdentifier: "user_project",
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
			ActionIdentifier: "list_user_project_db",
			Message:          "error in executing query for ListUserProject",
			EntityIdentifier: "user_project",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             request,
			ExtraData:        map[string]string{"query": query},
			Error:            err,
		})
		return types.ListResponse{}, err
	}
	defer rows.Close()
	var items []repogen.UserProject
	for rows.Next() {
		var i repogen.UserProject
		fields := []any{}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "uuid") {
				fields = append(fields, &i.UUID)
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "uuid") {
				fields = append(fields, &i.UUID)
			}
		} else {
			fields = append(fields, &i.UUID)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "user_uuid") {
				fields = append(fields, &i.UserUUID)
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "user_uuid") {
				fields = append(fields, &i.UserUUID)
			}
		} else {
			fields = append(fields, &i.UserUUID)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "user_email") {
				fields = append(fields, &i.UserEmail)
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "user_email") {
				fields = append(fields, &i.UserEmail)
			}
		} else {
			fields = append(fields, &i.UserEmail)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "project_uuid") {
				fields = append(fields, &i.ProjectUUID)
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "project_uuid") {
				fields = append(fields, &i.ProjectUUID)
			}
		} else {
			fields = append(fields, &i.ProjectUUID)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "role") {
				fields = append(fields, &i.Role)
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "role") {
				fields = append(fields, &i.Role)
			}
		} else {
			fields = append(fields, &i.Role)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "review_required_structure") {
				fields = append(fields, &i.ReviewRequiredStructure)
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "review_required_structure") {
				fields = append(fields, &i.ReviewRequiredStructure)
			}
		} else {
			fields = append(fields, &i.ReviewRequiredStructure)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "review_required_data") {
				fields = append(fields, &i.ReviewRequiredData)
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "review_required_data") {
				fields = append(fields, &i.ReviewRequiredData)
			}
		} else {
			fields = append(fields, &i.ReviewRequiredData)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "status") {
				fields = append(fields, &i.Status)
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "status") {
				fields = append(fields, &i.Status)
			}
		} else {
			fields = append(fields, &i.Status)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "created_at") {
				fields = append(fields, &i.CreatedAt)
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "created_at") {
				fields = append(fields, &i.CreatedAt)
			}
		} else {
			fields = append(fields, &i.CreatedAt)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "updated_at") {
				fields = append(fields, &i.UpdatedAt)
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "updated_at") {
				fields = append(fields, &i.UpdatedAt)
			}
		} else {
			fields = append(fields, &i.UpdatedAt)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "created_by_uuid") {
				fields = append(fields, &i.CreatedByUUID)
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "created_by_uuid") {
				fields = append(fields, &i.CreatedByUUID)
			}
		} else {
			fields = append(fields, &i.CreatedByUUID)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "updated_by_uuid") {
				fields = append(fields, &i.UpdatedByUUID)
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "updated_by_uuid") {
				fields = append(fields, &i.UpdatedByUUID)
			}
		} else {
			fields = append(fields, &i.UpdatedByUUID)
		}

		if err := rows.Scan(fields...); err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "list_user_project_scan",
				Message:          "error in scanning rows for ListUserProject",
				EntityIdentifier: "user_project",
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
			ActionIdentifier: "list_user_project_rows",
			Message:          "error closing rows for ListUserProject",
			EntityIdentifier: "user_project",
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
			ActionIdentifier: "list_user_project_rows",
			Message:          "error in rows for ListUserProject",
			EntityIdentifier: "user_project",
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
			ActionIdentifier: "list_user_project_next",
			Message:          "error determining if list has next page for ListUserProject",
			EntityIdentifier: "user_project",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             request,
			ExtraData:        map[string]string{"query": query},
			Error:            err,
		})
		return types.ListResponse{}, err
	}

	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "list_user_project",
		Message:          "successfully handled ListUserProject",
		EntityIdentifier: "user_project",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeSuccess,
		Data:             request,
		ExtraData:        map[string]string{"query": query},
	})
	return types.ListResponse{
		UserProjects: mapModelsToEntities(items),
		HasNextPage:  hasNextPage,
	}, nil
}

func (m *module) ListCount(ctx context.Context,
	request types.ListRequest,
	opts ...Option) (int64, error) {
	query, err := m.repository.BuildListEntityQuery(ctx, request, main_entity.UserProject{}, true, nil, nil)
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
