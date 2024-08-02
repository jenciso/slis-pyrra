package slogen

import "flag"

func AppsCmd(args []string) CmdOpts {
	fs := flag.NewFlagSet("apps", flag.ExitOnError)

	t := fs.String("t", "", "The apps type allowed: product, service, entity, backend, frontend, integration, all")
	a := fs.String("a", "", "The application file list already created")
	b := fs.String("b", "https://backstage.mycompany.com", "The Backstage URL")
	o := fs.String("o", "", "Output file to be written")

	fs.Parse(args)

	return CmdOpts{
		AppType:      *t,
		AppListFile:  *a,
		BackstageURL: *b,
		OutputFile:   *o,
	}
}

func SlisCmd(args []string) CmdOpts {
	fs := flag.NewFlagSet("apps", flag.ExitOnError)

	n := fs.String("n", "", "The application name")
	a := fs.String("a", "", "The application file list already created")
	p := fs.String("p", "https://thanos-query.mycompany.com", "The Backstage URL")
	o := fs.String("o", "", "Output file to be written")

	fs.Parse(args)

	return CmdOpts{
		AppName:       *n,
		AppListFile:   *a,
		PrometheusURL: *p,
		OutputFile:    *o,
	}
}

func SlosCmd(args []string) CmdOpts {
	fs := flag.NewFlagSet("apps", flag.ExitOnError)

	s := fs.String("s", "", "The SLIs file generated")
	o := fs.String("o", "", "Output file to be written")

	fs.Parse(args)

	return CmdOpts{
		SlisFile:   *s,
		OutputFile: *o,
	}
}
