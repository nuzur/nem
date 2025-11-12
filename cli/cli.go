package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"

	"github.com/nuzur/nem/config"
	"github.com/nuzur/nem/core"
	"github.com/nuzur/nem/monitoring"

	"github.com/nuzur/nem/core/entity/team"
	teamtypes "github.com/nuzur/nem/core/module/team/types"

	"github.com/nuzur/nem/core/entity/project"
	projecttypes "github.com/nuzur/nem/core/module/project/types"

	"github.com/nuzur/nem/core/entity/extension"
	extensiontypes "github.com/nuzur/nem/core/module/extension/types"

	"github.com/nuzur/nem/core/entity/extension_version"
	extension_versiontypes "github.com/nuzur/nem/core/module/extension_version/types"

	"github.com/nuzur/nem/core/entity/user"
	usertypes "github.com/nuzur/nem/core/module/user/types"

	"github.com/nuzur/nem/core/entity/change_request"
	change_requesttypes "github.com/nuzur/nem/core/module/change_request/types"

	"github.com/nuzur/nem/core/entity/project_version"
	project_versiontypes "github.com/nuzur/nem/core/module/project_version/types"

	"github.com/nuzur/nem/core/entity/user_team"
	user_teamtypes "github.com/nuzur/nem/core/module/user_team/types"

	"github.com/nuzur/nem/core/entity/extension_execution"
	extension_executiontypes "github.com/nuzur/nem/core/module/extension_execution/types"

	"github.com/nuzur/nem/core/entity/user_connection"
	user_connectiontypes "github.com/nuzur/nem/core/module/user_connection/types"

	"github.com/nuzur/nem/core/entity/user_project_version"
	user_project_versiontypes "github.com/nuzur/nem/core/module/user_project_version/types"

	"github.com/nuzur/nem/core/entity/user_project"
	user_projecttypes "github.com/nuzur/nem/core/module/user_project/types"

	"github.com/nuzur/nem/core/entity/membership"
	membershiptypes "github.com/nuzur/nem/core/module/membership/types"

	"github.com/nuzur/nem/core/entity/ai_usage"
	ai_usagetypes "github.com/nuzur/nem/core/module/ai_usage/types"

	"github.com/gofrs/uuid"
	"github.com/manifoldco/promptui"
	"github.com/urfave/cli"
	"go.uber.org/zap"
)

func main() {
	app := cli.NewApp()
	app.Name = "GPG CLI"

	app.Flags = []cli.Flag{}

	config, err := config.NewWithPathAndEnviorment("../config", "cli")
	if err != nil {
		panic("error creating config")
	}
	m, err := monitoring.New(monitoring.Params{
		Logger:   zap.NewNop(),
		Provider: config,
	})
	if err != nil {
		panic("error creating monitoring")
	}
	c, err := core.New(core.Params{
		Provider:   config,
		Monitoring: m,
	})
	if err != nil {
		panic("error creating core")
	}

	app.Commands = []cli.Command{
		{
			Name:  "create-user",
			Usage: "Creates a user",
			Action: func(cliCtx *cli.Context) error {

				ctx := context.Background()
				fields := map[string]promptContent{

					"uuid": {
						label:     "UUID",
						fieldType: "uuid",
					},

					"identifier": {
						label:     "Identifier",
						fieldType: "string",
					},

					"email": {
						label:     "Email",
						fieldType: "string",
					},

					"user_type": {
						label:     "UserType",
						fieldType: "options_single",
					},
				}

				for k, f := range fields {
					res, err := promptGetInput(f)
					if err != nil {
						return err
					}
					f.value = res
					fields[k] = f
				}

				res, err := c.User().Upsert(ctx, usertypes.UpsertRequest{
					User: user.User{

						Identifier: fields["identifier"].value,

						Email: fields["email"].value,

						Status: user.Status(1),
					},
				}, false)

				if err != nil {
					return err
				}

				fmt.Printf("%v", res)

				return nil
			},
		},
		{
			Name:  "create-token",
			Usage: "Creates a token for a given user",
			Action: func(c *cli.Context) error {
				url := "https://nem.nuzur.com/signin"
				fmt.Println("URL: ", url)

				fields := map[string]promptContent{
					"email": {
						label:     "Email",
						fieldType: "string",
					},

					"password": {
						label:     "Password",
						fieldType: "string",
					},
				}

				for k, f := range fields {
					res, err := promptGetInput(f)
					if err != nil {
						return err
					}
					f.value = res
					fields[k] = f
				}

				var jsonStr = []byte(fmt.Sprintf(`{"email":"%s", "password":"%s"}`,
					fields["email"].value,
					fields["password"].value,
				),
				)
				req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
				req.Header.Set("Content-Type", "application/json")

				client := &http.Client{}
				resp, err := client.Do(req)
				if err != nil {
					return err
				}
				defer resp.Body.Close()

				fmt.Println("Status:", resp.Status)
				body, _ := io.ReadAll(resp.Body)
				resData := map[string]string{}
				err = json.Unmarshal(body, &resData)
				if err != nil {
					return err
				}

				fmt.Println("Token:", resData["Token"])
				fmt.Println("Expires:", resData["Expires"])
				return nil

			},
		},
		{
			Name:  "fill-db",
			Usage: "Fills the entity with dummy values for testing",
			Action: func(ctx *cli.Context) error {
				entityIdentifier, err := promptGetInput(promptContent{
					label:     "Entity Identifier",
					fieldType: "string",
				})

				if err != nil {
					return err
				}

				switch entityIdentifier {

				case "team":
					entities := team.NewTeamSliceWithRandomValues(100)
					for _, e := range entities {
						e.UUID = uuid.Nil
						res, err := c.Team().Upsert(context.Background(), teamtypes.UpsertRequest{
							Team: e,
						}, false)
						if err != nil {
							fmt.Printf("error writing entity: %s, %v \n", "team", err)
						} else {
							fmt.Printf("wrote entity: %s, %v \n", "team", res)
						}
					}
					c.Destroy()

				case "project":
					entities := project.NewProjectSliceWithRandomValues(100)
					for _, e := range entities {
						e.UUID = uuid.Nil
						res, err := c.Project().Upsert(context.Background(), projecttypes.UpsertRequest{
							Project: e,
						}, false)
						if err != nil {
							fmt.Printf("error writing entity: %s, %v \n", "project", err)
						} else {
							fmt.Printf("wrote entity: %s, %v \n", "project", res)
						}
					}
					c.Destroy()

				case "extension":
					entities := extension.NewExtensionSliceWithRandomValues(100)
					for _, e := range entities {
						e.UUID = uuid.Nil
						res, err := c.Extension().Upsert(context.Background(), extensiontypes.UpsertRequest{
							Extension: e,
						}, false)
						if err != nil {
							fmt.Printf("error writing entity: %s, %v \n", "extension", err)
						} else {
							fmt.Printf("wrote entity: %s, %v \n", "extension", res)
						}
					}
					c.Destroy()

				case "extension_version":
					entities := extension_version.NewExtensionVersionSliceWithRandomValues(100)
					for _, e := range entities {
						e.UUID = uuid.Nil
						res, err := c.ExtensionVersion().Upsert(context.Background(), extension_versiontypes.UpsertRequest{
							ExtensionVersion: e,
						}, false)
						if err != nil {
							fmt.Printf("error writing entity: %s, %v \n", "extension_version", err)
						} else {
							fmt.Printf("wrote entity: %s, %v \n", "extension_version", res)
						}
					}
					c.Destroy()

				case "user":
					entities := user.NewUserSliceWithRandomValues(100)
					for _, e := range entities {
						e.UUID = uuid.Nil
						res, err := c.User().Upsert(context.Background(), usertypes.UpsertRequest{
							User: e,
						}, false)
						if err != nil {
							fmt.Printf("error writing entity: %s, %v \n", "user", err)
						} else {
							fmt.Printf("wrote entity: %s, %v \n", "user", res)
						}
					}
					c.Destroy()

				case "change_request":
					entities := change_request.NewChangeRequestSliceWithRandomValues(100)
					for _, e := range entities {
						e.UUID = uuid.Nil
						res, err := c.ChangeRequest().Upsert(context.Background(), change_requesttypes.UpsertRequest{
							ChangeRequest: e,
						}, false)
						if err != nil {
							fmt.Printf("error writing entity: %s, %v \n", "change_request", err)
						} else {
							fmt.Printf("wrote entity: %s, %v \n", "change_request", res)
						}
					}
					c.Destroy()

				case "project_version":
					entities := project_version.NewProjectVersionSliceWithRandomValues(100)
					for _, e := range entities {
						e.UUID = uuid.Nil
						res, err := c.ProjectVersion().Upsert(context.Background(), project_versiontypes.UpsertRequest{
							ProjectVersion: e,
						}, false)
						if err != nil {
							fmt.Printf("error writing entity: %s, %v \n", "project_version", err)
						} else {
							fmt.Printf("wrote entity: %s, %v \n", "project_version", res)
						}
					}
					c.Destroy()

				case "user_team":
					entities := user_team.NewUserTeamSliceWithRandomValues(100)
					for _, e := range entities {
						e.UUID = uuid.Nil
						res, err := c.UserTeam().Upsert(context.Background(), user_teamtypes.UpsertRequest{
							UserTeam: e,
						}, false)
						if err != nil {
							fmt.Printf("error writing entity: %s, %v \n", "user_team", err)
						} else {
							fmt.Printf("wrote entity: %s, %v \n", "user_team", res)
						}
					}
					c.Destroy()

				case "extension_execution":
					entities := extension_execution.NewExtensionExecutionSliceWithRandomValues(100)
					for _, e := range entities {
						e.UUID = uuid.Nil
						res, err := c.ExtensionExecution().Upsert(context.Background(), extension_executiontypes.UpsertRequest{
							ExtensionExecution: e,
						}, false)
						if err != nil {
							fmt.Printf("error writing entity: %s, %v \n", "extension_execution", err)
						} else {
							fmt.Printf("wrote entity: %s, %v \n", "extension_execution", res)
						}
					}
					c.Destroy()

				case "user_connection":
					entities := user_connection.NewUserConnectionSliceWithRandomValues(100)
					for _, e := range entities {
						e.UUID = uuid.Nil
						res, err := c.UserConnection().Upsert(context.Background(), user_connectiontypes.UpsertRequest{
							UserConnection: e,
						}, false)
						if err != nil {
							fmt.Printf("error writing entity: %s, %v \n", "user_connection", err)
						} else {
							fmt.Printf("wrote entity: %s, %v \n", "user_connection", res)
						}
					}
					c.Destroy()

				case "user_project_version":
					entities := user_project_version.NewUserProjectVersionSliceWithRandomValues(100)
					for _, e := range entities {
						e.UUID = uuid.Nil
						res, err := c.UserProjectVersion().Upsert(context.Background(), user_project_versiontypes.UpsertRequest{
							UserProjectVersion: e,
						}, false)
						if err != nil {
							fmt.Printf("error writing entity: %s, %v \n", "user_project_version", err)
						} else {
							fmt.Printf("wrote entity: %s, %v \n", "user_project_version", res)
						}
					}
					c.Destroy()

				case "user_project":
					entities := user_project.NewUserProjectSliceWithRandomValues(100)
					for _, e := range entities {
						e.UUID = uuid.Nil
						res, err := c.UserProject().Upsert(context.Background(), user_projecttypes.UpsertRequest{
							UserProject: e,
						}, false)
						if err != nil {
							fmt.Printf("error writing entity: %s, %v \n", "user_project", err)
						} else {
							fmt.Printf("wrote entity: %s, %v \n", "user_project", res)
						}
					}
					c.Destroy()

				case "membership":
					entities := membership.NewMembershipSliceWithRandomValues(100)
					for _, e := range entities {
						e.UUID = uuid.Nil
						res, err := c.Membership().Upsert(context.Background(), membershiptypes.UpsertRequest{
							Membership: e,
						}, false)
						if err != nil {
							fmt.Printf("error writing entity: %s, %v \n", "membership", err)
						} else {
							fmt.Printf("wrote entity: %s, %v \n", "membership", res)
						}
					}
					c.Destroy()

				case "ai_usage":
					entities := ai_usage.NewAiUsageSliceWithRandomValues(100)
					for _, e := range entities {
						e.UUID = uuid.Nil
						res, err := c.AiUsage().Upsert(context.Background(), ai_usagetypes.UpsertRequest{
							AiUsage: e,
						}, false)
						if err != nil {
							fmt.Printf("error writing entity: %s, %v \n", "ai_usage", err)
						} else {
							fmt.Printf("wrote entity: %s, %v \n", "ai_usage", res)
						}
					}
					c.Destroy()

				}
				return nil
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

type promptContent struct {
	label     string
	fieldType string
	value     string
}

func promptGetInput(pc promptContent) (string, error) {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New("required field")
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     pc.label,
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return result, nil
}
