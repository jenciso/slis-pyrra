package slos

import (
	"fmt"
	"strings"

	"github.com/jenciso/slis-pyrra/pkg/slogen"
)

func (s SloSchema) GenerateSLO(opts slogen.CmdOpts) {

	s.SliConf = SliLoadSpec(opts.SlisFile)
	for _, sli := range s.SliConf {

		for _, e := range sli.Spec.Slis.Ratio.Endpoints {
			var slo SloSchemaYaml

			slo.loadDefaultSliValues(sli)

			slo.Metadata.Name = fmt.Sprintf(`%s-%s-%s`, sli.Name, "ratio", slogen.GenerateHashID(e.Name))
			slo.Metadata.Labels.PyrraDevSli = "ratio"
			slo.Metadata.Labels.PyrraDevEndpoint = sanityzeEndpoint(e.Name)
			slo.Spec.Description = fmt.Sprintf(`SLO definition for %s`, sli.Name)

			slo.Spec.Indicator.Ratio.Errors.Metric = fmt.Sprintf(`%s_count{app="%s",%s="%s", status=~"5.*"}`,
				sli.Spec.Conf.MetricName, sli.Name, sli.Spec.Conf.EndpointLabel, strings.Replace(e.Name, `\`, `\\`, -1))
			slo.Spec.Indicator.Ratio.Total.Metric = fmt.Sprintf(`%s_count{app="%s",%s="%s"}`,
				sli.Spec.Conf.MetricName, sli.Name, sli.Spec.Conf.EndpointLabel, strings.Replace(e.Name, `\`, `\\`, -1))
			slo.Spec.Target = e.Target

			s.SloYaml = append(s.SloYaml, slo)

		}

		for _, e := range sli.Spec.Slis.Latency.Endpoints {
			var slo SloSchemaYaml

			slo.loadDefaultSliValues(sli)

			slo.Metadata.Name = fmt.Sprintf(`%s-%s-%s`, sli.Name, "latency", slogen.GenerateHashID(e.Name))
			slo.Metadata.Labels.PyrraDevSli = "latency"
			slo.Metadata.Labels.PyrraDevEndpoint = sanityzeEndpoint(e.Name)
			slo.Spec.Description = fmt.Sprintf(`SLO definition for %s`, sli.Name)

			slo.Spec.Indicator.Latency.Success.Metric = fmt.Sprintf(`%s_bucket{app="%s",%s="%s",status!~"5..",le="%s"}`,
				sli.Spec.Conf.MetricName, sli.Name, sli.Spec.Conf.EndpointLabel, strings.Replace(e.Name, `\`, `\\`, -1), e.Bucket)
			slo.Spec.Indicator.Latency.Total.Metric = fmt.Sprintf(`%s_count{app="%s",%s="%s"}`,
				sli.Spec.Conf.MetricName, sli.Name, sli.Spec.Conf.EndpointLabel, strings.Replace(e.Name, `\`, `\\`, -1))
			slo.Spec.Target = e.Target

			s.SloYaml = append(s.SloYaml, slo)

		}
	}
	slogen.PrintYamlResult(opts.OutputFile, MergeSlosResult(s.SloYaml))
}
