package apps

import (
	"cmp"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/jenciso/slis-pyrra/pkg/slogen"
	"go.einride.tech/backstage/catalog"
	"gopkg.in/yaml.v2"
)

func FetchAppsFromBackstage(opt slogen.CmdOpts) []AppSchemaYaml {
	var appList []AppSchemaYaml
	var query string

	query = "kind=Component,spec.type=" + opt.AppType

	if opt.AppType == "all" {
		query = "kind=Component"
	}

	//Connect to backstage
	client := catalog.NewClient(
		catalog.WithBaseURL(opt.BackstageURL),
		catalog.WithToken(opt.BackstageToken),
	)

	response := BackstageListEntities(client, query)
	for _, entity := range response.Entities {
		spec, err := entity.ComponentSpec()
		if err != nil {
			panic(err)
		}
		var app AppSchemaYaml
		app.Name = entity.Metadata.Name
		app.Spec.Autoconf = true
		app.Spec.Type = spec.Type
		app.Spec.Owner = strings.ReplaceAll(spec.Owner, "group:", "")
		if spec.System == "TBD" {
			app.Spec.System = ""
		} else {
			app.Spec.System = spec.System
		}
		app.Spec.Domain = GetDomain(client, app.Spec.System)
		app.Spec.Namespace = entity.Metadata.Name

		if app.Spec.Type != "template" {
			appList = append(appList, app)
		}
	}
	return appList
}

func GetDomain(client *catalog.Client, system string) string {
	defer func() {
		recover()
	}()
	response, _ := client.GetEntityByName(context.Background(), &catalog.GetEntityByNameRequest{
		Kind:      "System",
		Name:      system,
		Namespace: "default",
	})
	d, _ := response.SystemSpec()
	return d.Domain
}

func UpdateAppList(baseAppList []AppSchemaYaml, updatedAppList []AppSchemaYaml) (appList []AppSchemaYaml) {
	for _, uApp := range updatedAppList {
		for _, bApp := range baseAppList {

			if bApp.Name == uApp.Name {
				uApp.Spec.Autoconf = bApp.Spec.Autoconf // Preserve autoconf spec to not generate SLIs
				uApp.Spec.Cuj = bApp.Spec.Cuj           // Preserve CUJ spec since it is not in backstage
			}
		}
		appList = append(appList, uApp)
	}
	return
}

func BackstageListEntities(client *catalog.Client, query string) *catalog.ListEntitiesResponse {
	response, err := client.ListEntities(context.Background(), &catalog.ListEntitiesRequest{
		Filters: []string{query},
	})
	if err != nil {
		panic(err)
	}
	return response
}

func ReadAppsList(filename string) []AppSchemaYaml {
	var appList []AppSchemaYaml
	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading YAML file: %s\n", err)
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &appList)
	if err != nil {
		panic(err)
	}
	return appList
}

func CmpName(a, b AppSchemaYaml) int {
	return cmp.Compare(a.Name, b.Name)
}
