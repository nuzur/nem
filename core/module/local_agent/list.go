package local_agent

import (
	"context"
	"fmt"
	main_entity "github.com/nuzur/nem/core/entity/local_agent"
	"github.com/nuzur/nem/core/module/local_agent/types"
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
		main_entity.LocalAgent{},
		false,
		optConfig.ListIncludeColumns,
		optConfig.ListExcludeColumns)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "list_local_agent_query",
			Message:          "error in building query for ListLocalAgent",
			EntityIdentifier: "local_agent",
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
			ActionIdentifier: "list_local_agent_sort_buffer",
			Message:          "error in setting sort buffer for ListLocalAgent",
			EntityIdentifier: "local_agent",
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
			ActionIdentifier: "list_local_agent_db",
			Message:          "error in executing query for ListLocalAgent",
			EntityIdentifier: "local_agent",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             request,
			ExtraData:        map[string]string{"query": query},
			Error:            err,
		})
		fmt.Printf("error in executing query for ListLocalAgent: %v \n %v\n", query, err)
		return types.ListResponse{}, err
	}
	defer rows.Close()
	var items []repogen.LocalAgent
	for rows.Next() {
		var i repogen.LocalAgent
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
			if slices.Contains(optConfig.ListIncludeColumns, "token_hash") {
				fields = append(fields, &i.TokenHash)
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "token_hash") {
				fields = append(fields, &i.TokenHash)
			}
		} else {
			fields = append(fields, &i.TokenHash)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "machine_name") {
				fields = append(fields, &i.MachineName)
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "machine_name") {
				fields = append(fields, &i.MachineName)
			}
		} else {
			fields = append(fields, &i.MachineName)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "os") {
				fields = append(fields, &i.Os)
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "os") {
				fields = append(fields, &i.Os)
			}
		} else {
			fields = append(fields, &i.Os)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "cli_version") {
				fields = append(fields, &i.CliVersion)
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "cli_version") {
				fields = append(fields, &i.CliVersion)
			}
		} else {
			fields = append(fields, &i.CliVersion)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "connections") {
				fields = append(fields, &i.Connections)
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "connections") {
				fields = append(fields, &i.Connections)
			}
		} else {
			fields = append(fields, &i.Connections)
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
			if slices.Contains(optConfig.ListIncludeColumns, "last_seen_at") {
				fields = append(fields, &i.LastSeenAt)
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "last_seen_at") {
				fields = append(fields, &i.LastSeenAt)
			}
		} else {
			fields = append(fields, &i.LastSeenAt)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "revoked_at") {
				fields = append(fields, &i.RevokedAt)
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "revoked_at") {
				fields = append(fields, &i.RevokedAt)
			}
		} else {
			fields = append(fields, &i.RevokedAt)
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
				ActionIdentifier: "list_local_agent_scan",
				Message:          "error in scanning rows for ListLocalAgent",
				EntityIdentifier: "local_agent",
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
			ActionIdentifier: "list_local_agent_rows",
			Message:          "error closing rows for ListLocalAgent",
			EntityIdentifier: "local_agent",
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
			ActionIdentifier: "list_local_agent_rows",
			Message:          "error in rows for ListLocalAgent",
			EntityIdentifier: "local_agent",
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
			ActionIdentifier: "list_local_agent_next",
			Message:          "error determining if list has next page for ListLocalAgent",
			EntityIdentifier: "local_agent",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             request,
			ExtraData:        map[string]string{"query": query},
			Error:            err,
		})
		return types.ListResponse{}, err
	}

	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "list_local_agent",
		Message:          "successfully handled ListLocalAgent",
		EntityIdentifier: "local_agent",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeSuccess,
		Data:             request,
		ExtraData:        map[string]string{"query": query},
	})
	return types.ListResponse{
		LocalAgents: mapModelsToEntities(items),
		HasNextPage: hasNextPage,
	}, nil
}

func (m *module) ListCount(ctx context.Context,
	request types.ListRequest,
	opts ...Option) (int64, error) {
	query, err := m.repository.BuildListEntityQuery(ctx, request, main_entity.LocalAgent{}, true, nil, nil)
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
