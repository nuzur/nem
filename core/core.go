package core

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	nemconfig "github.com/nuzur/nem/config"
	coretypes "github.com/nuzur/nem/core/types"
	"go.uber.org/config"
	"go.uber.org/fx"

	"github.com/nuzur/nem/core/module/team"

	"github.com/nuzur/nem/core/module/organization"

	"github.com/nuzur/nem/core/module/project"

	"github.com/nuzur/nem/core/module/extension"

	"github.com/nuzur/nem/core/module/extension_version"

	"github.com/nuzur/nem/core/module/user"

	"github.com/nuzur/nem/core/module/change_request"

	"github.com/nuzur/nem/core/module/project_version"

	"github.com/nuzur/nem/core/module/user_team"

	"github.com/nuzur/nem/core/module/extension_execution"

	"github.com/nuzur/nem/core/module/user_connection"

	"github.com/nuzur/nem/core/module/user_project_version"

	"github.com/nuzur/nem/core/module/user_project"

	"github.com/nuzur/nem/core/module/membership"

	"github.com/nuzur/nem/core/repository"
	"github.com/nuzur/nem/monitoring"

	"github.com/nuzur/nem/core/events"
)

type Implementation struct {
	db         *sql.DB
	repository *repository.Implementation

	team team.Module

	organization organization.Module

	project project.Module

	extension extension.Module

	extension_version extension_version.Module

	user user.Module

	change_request change_request.Module

	project_version project_version.Module

	user_team user_team.Module

	extension_execution extension_execution.Module

	user_connection user_connection.Module

	user_project_version user_project_version.Module

	user_project user_project.Module

	membership membership.Module

	monitoring *monitoring.Implementation

	events *events.Implementation
}

type Params struct {
	fx.In
	Provider   config.Provider
	Lifecycle  fx.Lifecycle
	Monitoring *monitoring.Implementation

	Events *events.Implementation
}

func New(params Params) (*Implementation, error) {

	var dbs nemconfig.DBs
	if err := params.Provider.Get("db").Populate(&dbs); err != nil {
		return nil, err
	}

	if len(dbs) == 0 {
		return nil, errors.New("db configuration not found")
	}

	dbconfig := dbs[0]
	db, err := sql.Open(dbconfig.Driver, dbconfig.Path())
	if err != nil {
		return nil, fmt.Errorf("error connecting to DB: %v", err)
	}
	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(0)
	repository := repository.New(db)

	if params.Lifecycle != nil {
		params.Lifecycle.Append(fx.Hook{
			OnStop: func(ctx context.Context) error {
				db.Close()
				return nil
			},
		})
	}

	return &Implementation{
		db:         db,
		repository: repository,
		monitoring: params.Monitoring,

		events: params.Events,
	}, nil
}

func (i *Implementation) Destroy() {
	i.db.Close()
}

func (i *Implementation) DB() *sql.DB {
	return i.db
}

func (i Implementation) Team() team.Module {
	if i.team == nil {
		i.team = team.New(coretypes.ModuleParams{
			Repository: i.repository,
			Monitoring: i.monitoring,

			Events: i.events,
		})
	}
	return i.team
}

func (i Implementation) Organization() organization.Module {
	if i.organization == nil {
		i.organization = organization.New(coretypes.ModuleParams{
			Repository: i.repository,
			Monitoring: i.monitoring,

			Events: i.events,
		})
	}
	return i.organization
}

func (i Implementation) Project() project.Module {
	if i.project == nil {
		i.project = project.New(coretypes.ModuleParams{
			Repository: i.repository,
			Monitoring: i.monitoring,

			Events: i.events,
		})
	}
	return i.project
}

func (i Implementation) Extension() extension.Module {
	if i.extension == nil {
		i.extension = extension.New(coretypes.ModuleParams{
			Repository: i.repository,
			Monitoring: i.monitoring,

			Events: i.events,
		})
	}
	return i.extension
}

func (i Implementation) ExtensionVersion() extension_version.Module {
	if i.extension_version == nil {
		i.extension_version = extension_version.New(coretypes.ModuleParams{
			Repository: i.repository,
			Monitoring: i.monitoring,

			Events: i.events,
		})
	}
	return i.extension_version
}

func (i Implementation) User() user.Module {
	if i.user == nil {
		i.user = user.New(coretypes.ModuleParams{
			Repository: i.repository,
			Monitoring: i.monitoring,

			Events: i.events,
		})
	}
	return i.user
}

func (i Implementation) ChangeRequest() change_request.Module {
	if i.change_request == nil {
		i.change_request = change_request.New(coretypes.ModuleParams{
			Repository: i.repository,
			Monitoring: i.monitoring,

			Events: i.events,
		})
	}
	return i.change_request
}

func (i Implementation) ProjectVersion() project_version.Module {
	if i.project_version == nil {
		i.project_version = project_version.New(coretypes.ModuleParams{
			Repository: i.repository,
			Monitoring: i.monitoring,

			Events: i.events,
		})
	}
	return i.project_version
}

func (i Implementation) UserTeam() user_team.Module {
	if i.user_team == nil {
		i.user_team = user_team.New(coretypes.ModuleParams{
			Repository: i.repository,
			Monitoring: i.monitoring,

			Events: i.events,
		})
	}
	return i.user_team
}

func (i Implementation) ExtensionExecution() extension_execution.Module {
	if i.extension_execution == nil {
		i.extension_execution = extension_execution.New(coretypes.ModuleParams{
			Repository: i.repository,
			Monitoring: i.monitoring,

			Events: i.events,
		})
	}
	return i.extension_execution
}

func (i Implementation) UserConnection() user_connection.Module {
	if i.user_connection == nil {
		i.user_connection = user_connection.New(coretypes.ModuleParams{
			Repository: i.repository,
			Monitoring: i.monitoring,

			Events: i.events,
		})
	}
	return i.user_connection
}

func (i Implementation) UserProjectVersion() user_project_version.Module {
	if i.user_project_version == nil {
		i.user_project_version = user_project_version.New(coretypes.ModuleParams{
			Repository: i.repository,
			Monitoring: i.monitoring,

			Events: i.events,
		})
	}
	return i.user_project_version
}

func (i Implementation) UserProject() user_project.Module {
	if i.user_project == nil {
		i.user_project = user_project.New(coretypes.ModuleParams{
			Repository: i.repository,
			Monitoring: i.monitoring,

			Events: i.events,
		})
	}
	return i.user_project
}

func (i Implementation) Membership() membership.Module {
	if i.membership == nil {
		i.membership = membership.New(coretypes.ModuleParams{
			Repository: i.repository,
			Monitoring: i.monitoring,

			Events: i.events,
		})
	}
	return i.membership
}
