package slis

import (
	"fmt"

	"github.com/jenciso/slis-pyrra/pkg/slogen"
)

func (s SliSchemaProduct) GenerateSLI(opts slogen.CmdOpts) {

	var subtypes = []string{"backend", "frontend"}
	printSli := true

	for _, t := range subtypes {
		var sli SliSchemaYaml
		var conf SliConf
		sli.Spec.App = AppLoadSpec(opts.AppListFile, opts.AppName)

		if sli.Spec.App.Autoconf {
			if t == "backend" {
				conf = NewConfProductBackend(opts.AppName)
				sli.Name = conf.App
				sli.Spec.Conf.MetricName = conf.Conf.MetricName
				sli.Spec.Conf.EndpointLabel = conf.Conf.EndpointLabel
				sli.Spec.Conf.BucketDefaultThreshold = conf.Conf.BucketDefaultThreshold
				sli.Spec.Conf.TargetDefault = conf.Conf.TargetDefault
				sli.setValuesEndpointsRatio(conf, opts)
				sli.setValuesEndpointsLatency(conf, opts)
			}
			if t == "frontend" {
				conf = NewConfProductFrontend(opts.AppName)
				sli.Name = conf.App
				sli.Spec.Conf.MetricName = conf.Conf.MetricName
				sli.Spec.Conf.EndpointLabel = conf.Conf.EndpointLabel
				sli.Spec.Conf.BucketDefaultThreshold = conf.Conf.BucketDefaultThreshold
				sli.Spec.Conf.TargetDefault = conf.Conf.TargetDefault
				sli.setValuesEndpointsRatio(conf, opts)
				sli.setValuesEndpointsLatency(conf, opts)
			}
			s.SliYaml = append(s.SliYaml, sli)
		} else {
			printSli = false
		}
	}

	if printSli {
		slogen.PrintYamlResult(opts.OutputFile, slogen.Yaml2String(s.SliYaml))
	} else {
		fmt.Printf("Nothing to generate for %s. Autoconf is disabled in %s file\n", opts.AppName, opts.AppListFile)
	}

}

func (s SliSchemaService) GenerateSLI(opts slogen.CmdOpts) {
	var sli SliSchemaYaml
	sli.Spec.App = AppLoadSpec(opts.AppListFile, opts.AppName)
	if sli.Spec.App.Autoconf {
		conf := NewConfService(opts.AppName)
		sli.Name = conf.App
		sli.Spec.Conf.MetricName = conf.Conf.MetricName
		sli.Spec.Conf.EndpointLabel = conf.Conf.EndpointLabel
		sli.Spec.Conf.BucketDefaultThreshold = conf.Conf.BucketDefaultThreshold
		sli.Spec.Conf.TargetDefault = conf.Conf.TargetDefault
		sli.setValuesEndpointsRatio(conf, opts)
		sli.setValuesEndpointsLatency(conf, opts)
		s.SliYaml = append(s.SliYaml, sli)
		slogen.PrintYamlResult(opts.OutputFile, slogen.Yaml2String(s.SliYaml))
	} else {
		fmt.Printf("Nothing to generate for %s. Autoconf is disabled in %s file\n", opts.AppName, opts.AppListFile)
	}
}

func (s SliSchemaEntity) GenerateSLI(opts slogen.CmdOpts) {
	var sli SliSchemaYaml
	sli.Spec.App = AppLoadSpec(opts.AppListFile, opts.AppName)
	if sli.Spec.App.Autoconf {
		conf := NewConfEntity(opts.AppName)
		sli.Name = conf.App
		sli.Spec.Conf.MetricName = conf.Conf.MetricName
		sli.Spec.Conf.EndpointLabel = conf.Conf.EndpointLabel
		sli.Spec.Conf.BucketDefaultThreshold = conf.Conf.BucketDefaultThreshold
		sli.Spec.Conf.TargetDefault = conf.Conf.TargetDefault
		sli.setValuesEndpointsRatio(conf, opts)
		sli.setValuesEndpointsLatency(conf, opts)
		s.SliYaml = append(s.SliYaml, sli)
		slogen.PrintYamlResult(opts.OutputFile, slogen.Yaml2String(s.SliYaml))
	} else {
		fmt.Printf("Nothing to generate for %s. Autoconf is disabled in %s file\n", opts.AppName, opts.AppListFile)
	}

}

func (s SliSchemaIntegration) GenerateSLI(opts slogen.CmdOpts) {
	var sli SliSchemaYaml
	sli.Spec.App = AppLoadSpec(opts.AppListFile, opts.AppName)
	if sli.Spec.App.Autoconf {
		conf := NewConfIntegration(opts.AppName)
		sli.Name = conf.App
		sli.Spec.Conf.MetricName = conf.Conf.MetricName
		sli.Spec.Conf.EndpointLabel = conf.Conf.EndpointLabel
		sli.Spec.Conf.BucketDefaultThreshold = conf.Conf.BucketDefaultThreshold
		sli.Spec.Conf.TargetDefault = conf.Conf.TargetDefault
		sli.setValuesEndpointsRatio(conf, opts)
		sli.setValuesEndpointsLatency(conf, opts)
		s.SliYaml = append(s.SliYaml, sli)
		slogen.PrintYamlResult(opts.OutputFile, slogen.Yaml2String(s.SliYaml))
	} else {
		fmt.Printf("Nothing to generate for %s. Autoconf is disabled in %s file\n", opts.AppName, opts.AppListFile)
	}
}

func (s SliSchemaFrontend) GenerateSLI(opts slogen.CmdOpts) {
	var sli SliSchemaYaml
	sli.Spec.App = AppLoadSpec(opts.AppListFile, opts.AppName)
	if sli.Spec.App.Autoconf {
		conf := NewConfFrontend(opts.AppName)
		sli.Name = conf.App
		sli.Spec.Conf.MetricName = conf.Conf.MetricName
		sli.Spec.Conf.EndpointLabel = conf.Conf.EndpointLabel
		sli.Spec.Conf.BucketDefaultThreshold = conf.Conf.BucketDefaultThreshold
		sli.Spec.Conf.TargetDefault = conf.Conf.TargetDefault
		sli.setValuesEndpointsRatio(conf, opts)
		sli.setValuesEndpointsLatency(conf, opts)
		s.SliYaml = append(s.SliYaml, sli)
		slogen.PrintYamlResult(opts.OutputFile, slogen.Yaml2String(s.SliYaml))
	} else {
		fmt.Printf("Nothing to generate for %s. Autoconf is disabled in %s file\n", opts.AppName, opts.AppListFile)
	}
}

func (s SliSchemaBackend) GenerateSLI(opts slogen.CmdOpts) {
	var sli SliSchemaYaml
	sli.Spec.App = AppLoadSpec(opts.AppListFile, opts.AppName)
	if sli.Spec.App.Autoconf {
		conf := NewConfBackend(opts.AppName)
		sli.Name = conf.App
		sli.Spec.Conf.MetricName = conf.Conf.MetricName
		sli.Spec.Conf.EndpointLabel = conf.Conf.EndpointLabel
		sli.Spec.Conf.BucketDefaultThreshold = conf.Conf.BucketDefaultThreshold
		sli.Spec.Conf.TargetDefault = conf.Conf.TargetDefault
		sli.setValuesEndpointsRatio(conf, opts)
		sli.setValuesEndpointsLatency(conf, opts)
		s.SliYaml = append(s.SliYaml, sli)
		slogen.PrintYamlResult(opts.OutputFile, slogen.Yaml2String(s.SliYaml))
	} else {
		fmt.Printf("Nothing to generate for %s. Autoconf is disabled in %s file\n", opts.AppName, opts.AppListFile)
	}
}
