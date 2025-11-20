package project

import (
	"context"
	"fmt"
	main_entity "github.com/nuzur/nem/core/entity/project"
	"github.com/nuzur/nem/core/module/project/types"
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
		main_entity.Project{},
		false,
		optConfig.ListIncludeColumns,
		optConfig.ListExcludeColumns)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "list_project_query",
			Message:          "error in building query for ListProject",
			EntityIdentifier: "project",
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
			ActionIdentifier: "list_project_sort_buffer",
			Message:          "error in setting sort buffer for ListProject",
			EntityIdentifier: "project",
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
			ActionIdentifier: "list_project_db",
			Message:          "error in executing query for ListProject",
			EntityIdentifier: "project",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             request,
			ExtraData:        map[string]string{"query": query},
			Error:            err,
		})
		fmt.Printf("error in executing query for ListProject: %v \n %v\n", query, err)
		return types.ListResponse{}, err
	}
	defer rows.Close()
	var items []repogen.Project
	for rows.Next() {
		var i repogen.Project
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
			if slices.Contains(optConfig.ListIncludeColumns, "version") {
				fields = append(fields, &i.Version)
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "version") {
				fields = append(fields, &i.Version)
			}
		} else {
			fields = append(fields, &i.Version)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "name") {
				fields = append(fields, &i.Name)
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "name") {
				fields = append(fields, &i.Name)
			}
		} else {
			fields = append(fields, &i.Name)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "description") {
				fields = append(fields, &i.Description)
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "description") {
				fields = append(fields, &i.Description)
			}
		} else {
			fields = append(fields, &i.Description)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "tags") {
				fields = append(fields, &i.Tags)
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "tags") {
				fields = append(fields, &i.Tags)
			}
		} else {
			fields = append(fields, &i.Tags)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "url") {
				fields = append(fields, &i.URL)
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "url") {
				fields = append(fields, &i.URL)
			}
		} else {
			fields = append(fields, &i.URL)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "owner_uuid") {
				fields = append(fields, &i.OwnerUUID)
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "owner_uuid") {
				fields = append(fields, &i.OwnerUUID)
			}
		} else {
			fields = append(fields, &i.OwnerUUID)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "team_uuid") {
				fields = append(fields, &i.TeamUUID)
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "team_uuid") {
				fields = append(fields, &i.TeamUUID)
			}
		} else {
			fields = append(fields, &i.TeamUUID)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "access_type") {
				fields = append(fields, &i.AccessType)
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "access_type") {
				fields = append(fields, &i.AccessType)
			}
		} else {
			fields = append(fields, &i.AccessType)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "project_extensions") {
				fields = append(fields, &i.ProjectExtensions)
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "project_extensions") {
				fields = append(fields, &i.ProjectExtensions)
			}
		} else {
			fields = append(fields, &i.ProjectExtensions)
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
				ActionIdentifier: "list_project_scan",
				Message:          "error in scanning rows for ListProject",
				EntityIdentifier: "project",
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
			ActionIdentifier: "list_project_rows",
			Message:          "error closing rows for ListProject",
			EntityIdentifier: "project",
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
			ActionIdentifier: "list_project_rows",
			Message:          "error in rows for ListProject",
			EntityIdentifier: "project",
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
			ActionIdentifier: "list_project_next",
			Message:          "error determining if list has next page for ListProject",
			EntityIdentifier: "project",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             request,
			ExtraData:        map[string]string{"query": query},
			Error:            err,
		})
		return types.ListResponse{}, err
	}

	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "list_project",
		Message:          "successfully handled ListProject",
		EntityIdentifier: "project",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeSuccess,
		Data:             request,
		ExtraData:        map[string]string{"query": query},
	})
	return types.ListResponse{
		Projects:    mapModelsToEntities(items),
		HasNextPage: hasNextPage,
	}, nil
}

func (m *module) ListCount(ctx context.Context,
	request types.ListRequest,
	opts ...Option) (int64, error) {
	query, err := m.repository.BuildListEntityQuery(ctx, request, main_entity.Project{}, true, nil, nil)
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
