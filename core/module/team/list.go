package team

import (
	"context"
	//"fmt"
	main_entity "github.com/nuzur/nem/core/entity/team"
	"github.com/nuzur/nem/core/module/team/types"
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
		main_entity.Team{},
		false,
		optConfig.ListIncludeColumns,
		optConfig.ListExcludeColumns)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "list_team_query",
			Message:          "error in building query for ListTeam",
			EntityIdentifier: "team",
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
			ActionIdentifier: "list_team_sort_buffer",
			Message:          "error in setting sort buffer for ListTeam",
			EntityIdentifier: "team",
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
			ActionIdentifier: "list_team_db",
			Message:          "error in executing query for ListTeam",
			EntityIdentifier: "team",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             request,
			ExtraData:        map[string]string{"query": query},
			Error:            err,
		})
		return types.ListResponse{}, err
	}
	defer rows.Close()
	var items []repogen.Team
	for rows.Next() {
		var i repogen.Team
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
			if slices.Contains(optConfig.ListIncludeColumns, "version") {
				fields = append(fields, &i.Version)
				continue
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "version") {
				fields = append(fields, &i.Version)
				continue
			}
		} else {
			fields = append(fields, &i.Version)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "name") {
				fields = append(fields, &i.Name)
				continue
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "name") {
				fields = append(fields, &i.Name)
				continue
			}
		} else {
			fields = append(fields, &i.Name)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "enviorments") {
				fields = append(fields, &i.Enviorments)
				continue
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "enviorments") {
				fields = append(fields, &i.Enviorments)
				continue
			}
		} else {
			fields = append(fields, &i.Enviorments)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "review_configs") {
				fields = append(fields, &i.ReviewConfigs)
				continue
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "review_configs") {
				fields = append(fields, &i.ReviewConfigs)
				continue
			}
		} else {
			fields = append(fields, &i.ReviewConfigs)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "stores") {
				fields = append(fields, &i.Stores)
				continue
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "stores") {
				fields = append(fields, &i.Stores)
				continue
			}
		} else {
			fields = append(fields, &i.Stores)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "connections") {
				fields = append(fields, &i.Connections)
				continue
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "connections") {
				fields = append(fields, &i.Connections)
				continue
			}
		} else {
			fields = append(fields, &i.Connections)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "object_stores") {
				fields = append(fields, &i.ObjectStores)
				continue
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "object_stores") {
				fields = append(fields, &i.ObjectStores)
				continue
			}
		} else {
			fields = append(fields, &i.ObjectStores)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "default_entity") {
				fields = append(fields, &i.DefaultEntity)
				continue
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "default_entity") {
				fields = append(fields, &i.DefaultEntity)
				continue
			}
		} else {
			fields = append(fields, &i.DefaultEntity)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "owner_uuid") {
				fields = append(fields, &i.OwnerUUID)
				continue
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "owner_uuid") {
				fields = append(fields, &i.OwnerUUID)
				continue
			}
		} else {
			fields = append(fields, &i.OwnerUUID)
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

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "created_by_uuid") {
				fields = append(fields, &i.CreatedByUUID)
				continue
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "created_by_uuid") {
				fields = append(fields, &i.CreatedByUUID)
				continue
			}
		} else {
			fields = append(fields, &i.CreatedByUUID)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "updated_by_uuid") {
				fields = append(fields, &i.UpdatedByUUID)
				continue
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "updated_by_uuid") {
				fields = append(fields, &i.UpdatedByUUID)
				continue
			}
		} else {
			fields = append(fields, &i.UpdatedByUUID)
		}

		if err := rows.Scan(fields...); err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "list_team_scan",
				Message:          "error in scanning rows for ListTeam",
				EntityIdentifier: "team",
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
			ActionIdentifier: "list_team_rows",
			Message:          "error closing rows for ListTeam",
			EntityIdentifier: "team",
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
			ActionIdentifier: "list_team_rows",
			Message:          "error in rows for ListTeam",
			EntityIdentifier: "team",
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
			ActionIdentifier: "list_team_next",
			Message:          "error determining if list has next page for ListTeam",
			EntityIdentifier: "team",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             request,
			ExtraData:        map[string]string{"query": query},
			Error:            err,
		})
		return types.ListResponse{}, err
	}

	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "list_team",
		Message:          "successfully handled ListTeam",
		EntityIdentifier: "team",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeSuccess,
		Data:             request,
		ExtraData:        map[string]string{"query": query},
	})
	return types.ListResponse{
		Teams:       mapModelsToEntities(items),
		HasNextPage: hasNextPage,
	}, nil
}

func (m *module) ListCount(ctx context.Context,
	request types.ListRequest,
	opts ...Option) (int64, error) {
	query, err := m.repository.BuildListEntityQuery(ctx, request, main_entity.Team{}, true, nil, nil)
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
