package apps

import (
	"slices"

	"github.com/jenciso/slis-pyrra/pkg/slogen"
)

func (s AppSchema) GenerateAPP(opts slogen.CmdOpts) {
	if opts.AppListFile != "" {
		s.AppYaml = UpdateAppList(ReadAppsList(opts.AppListFile), FetchAppsFromBackstage(opts))
	} else {
		s.AppYaml = FetchAppsFromBackstage(opts)
	}

	slices.SortFunc(s.AppYaml, CmpName)
	slogen.PrintYamlResult(opts.OutputFile, slogen.Yaml2String(s.AppYaml))
}
