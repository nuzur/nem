package user

import (
	"context"
	//"fmt"
	main_entity "github.com/nuzur/nem/core/entity/user"
	"github.com/nuzur/nem/core/module/user/types"
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
		main_entity.User{},
		false,
		optConfig.ListIncludeColumns,
		optConfig.ListExcludeColumns)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "list_user_query",
			Message:          "error in building query for ListUser",
			EntityIdentifier: "user",
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
			ActionIdentifier: "list_user_sort_buffer",
			Message:          "error in setting sort buffer for ListUser",
			EntityIdentifier: "user",
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
			ActionIdentifier: "list_user_db",
			Message:          "error in executing query for ListUser",
			EntityIdentifier: "user",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             request,
			ExtraData:        map[string]string{"query": query},
			Error:            err,
		})
		return types.ListResponse{}, err
	}
	defer rows.Close()
	var items []repogen.User
	for rows.Next() {
		var i repogen.User
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
			if slices.Contains(optConfig.ListIncludeColumns, "identifier") {
				fields = append(fields, &i.Identifier)
			} else {
				var skip interface{}
				fields = append(fields, &skip)
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "identifier") {
				fields = append(fields, &i.Identifier)
			} else {
				var skip interface{}
				fields = append(fields, &skip)
			}
		} else {
			fields = append(fields, &i.Identifier)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "name") {
				fields = append(fields, &i.Name)
			} else {
				var skip interface{}
				fields = append(fields, &skip)
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "name") {
				fields = append(fields, &i.Name)
			} else {
				var skip interface{}
				fields = append(fields, &skip)
			}
		} else {
			fields = append(fields, &i.Name)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "last_name") {
				fields = append(fields, &i.LastName)
			} else {
				var skip interface{}
				fields = append(fields, &skip)
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "last_name") {
				fields = append(fields, &i.LastName)
			} else {
				var skip interface{}
				fields = append(fields, &skip)
			}
		} else {
			fields = append(fields, &i.LastName)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "email") {
				fields = append(fields, &i.Email)
			} else {
				var skip interface{}
				fields = append(fields, &skip)
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "email") {
				fields = append(fields, &i.Email)
			} else {
				var skip interface{}
				fields = append(fields, &skip)
			}
		} else {
			fields = append(fields, &i.Email)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "user_type") {
				fields = append(fields, &i.UserType)
			} else {
				var skip interface{}
				fields = append(fields, &skip)
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "user_type") {
				fields = append(fields, &i.UserType)
			} else {
				var skip interface{}
				fields = append(fields, &skip)
			}
		} else {
			fields = append(fields, &i.UserType)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "country_ios2") {
				fields = append(fields, &i.CountryIos2)
			} else {
				var skip interface{}
				fields = append(fields, &skip)
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "country_ios2") {
				fields = append(fields, &i.CountryIos2)
			} else {
				var skip interface{}
				fields = append(fields, &skip)
			}
		} else {
			fields = append(fields, &i.CountryIos2)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "locale") {
				fields = append(fields, &i.Locale)
			} else {
				var skip interface{}
				fields = append(fields, &skip)
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "locale") {
				fields = append(fields, &i.Locale)
			} else {
				var skip interface{}
				fields = append(fields, &skip)
			}
		} else {
			fields = append(fields, &i.Locale)
		}

		if len(optConfig.ListIncludeColumns) > 0 {
			if slices.Contains(optConfig.ListIncludeColumns, "metadata") {
				fields = append(fields, &i.Metadata)
			} else {
				var skip interface{}
				fields = append(fields, &skip)
			}
		} else if len(optConfig.ListExcludeColumns) > 0 {
			if !slices.Contains(optConfig.ListExcludeColumns, "metadata") {
				fields = append(fields, &i.Metadata)
			} else {
				var skip interface{}
				fields = append(fields, &skip)
			}
		} else {
			fields = append(fields, &i.Metadata)
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
				ActionIdentifier: "list_user_scan",
				Message:          "error in scanning rows for ListUser",
				EntityIdentifier: "user",
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
			ActionIdentifier: "list_user_rows",
			Message:          "error closing rows for ListUser",
			EntityIdentifier: "user",
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
			ActionIdentifier: "list_user_rows",
			Message:          "error in rows for ListUser",
			EntityIdentifier: "user",
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
			ActionIdentifier: "list_user_next",
			Message:          "error determining if list has next page for ListUser",
			EntityIdentifier: "user",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             request,
			ExtraData:        map[string]string{"query": query},
			Error:            err,
		})
		return types.ListResponse{}, err
	}

	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "list_user",
		Message:          "successfully handled ListUser",
		EntityIdentifier: "user",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeSuccess,
		Data:             request,
		ExtraData:        map[string]string{"query": query},
	})
	return types.ListResponse{
		Users:       mapModelsToEntities(items),
		HasNextPage: hasNextPage,
	}, nil
}

func (m *module) ListCount(ctx context.Context,
	request types.ListRequest,
	opts ...Option) (int64, error) {
	query, err := m.repository.BuildListEntityQuery(ctx, request, main_entity.User{}, true, nil, nil)
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
