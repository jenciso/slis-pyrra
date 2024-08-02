package main

import (
	"fmt"
	"os"

	"github.com/jenciso/slis-pyrra/pkg/apps"
	"github.com/jenciso/slis-pyrra/pkg/slis"
	"github.com/jenciso/slis-pyrra/pkg/slogen"
	"github.com/jenciso/slis-pyrra/pkg/slos"
)

type Application struct {
	app appGenerator
	sli sliGenerator
	slo sloGenerator
}

type appGenerator interface {
	GenerateAPP(opts slogen.CmdOpts)
}

type sliGenerator interface {
	GenerateSLI(opts slogen.CmdOpts)
}

type sloGenerator interface {
	GenerateSLO(opts slogen.CmdOpts)
}

func (a Application) RunAPP(opts slogen.CmdOpts) {
	a.app.GenerateAPP(opts)
}

func (a Application) RunSLI(opts slogen.CmdOpts) {
	a.sli.GenerateSLI(opts)
}

func (a Application) RunSLO(opts slogen.CmdOpts) {
	a.slo.GenerateSLO(opts)
}

func newApps(g appGenerator) *Application {
	return &Application{app: g}
}

func newSlis(g sliGenerator) *Application {
	return &Application{sli: g}
}

func newSlos(g sloGenerator) *Application {
	return &Application{slo: g}
}

func main() {

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("No command provided")
		os.Exit(1)
	}

	switch args[0] {
	case "apps":
		opt := slogen.AppsCmd(args[1:])
		newApps(apps.New()).RunAPP(opt)

	case "slis":
		opt := slogen.SlisCmd(args[1:])
		appType := slis.GetAppTypeFromAppList(opt.AppListFile, opt.AppName)
		switch appType {
		case "product":
			newSlis(slis.NewProduct()).RunSLI(opt)
		case "service":
			newSlis(slis.NewService()).RunSLI(opt)
		case "entity":
			newSlis(slis.NewEntity()).RunSLI(opt)
		case "integration":
			newSlis(slis.NewIntegration()).RunSLI(opt)
		case "frontend":
			newSlis(slis.NewFrontend()).RunSLI(opt)
		case "backend":
			newSlis(slis.NewBackend()).RunSLI(opt)
		default:
			fmt.Fprintf(os.Stderr, "SLIS generation: Unrecognized application type: %s. "+
				"Argument must be one of: product, service, entity, integration, frontend, backend\n", appType)
			os.Exit(1)
		}

	case "slos":
		opt := slogen.SlosCmd(args[1:])
		newSlos(slos.New()).RunSLO(opt)

	default:
		fmt.Fprintf(os.Stderr, "Unrecognized command %q. "+
			"Command must be one of: 'apps', 'slis', 'slos'\n", args[0])
		os.Exit(1)
	}

}
