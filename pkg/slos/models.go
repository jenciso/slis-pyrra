package slos

import (
	"github.com/jenciso/slis-pyrra/pkg/slis"
)

type SloSchemaYaml struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Annotations struct {
		} `yaml:"annotations,omitempty"`
		Labels struct {
			PyrraDevCuj      string `yaml:"pyrra.dev/cuj,omitempty"`
			PyrraDevSli      string `yaml:"pyrra.dev/sli,omitempty"`
			PyrraDevType     string `yaml:"pyrra.dev/type,omitempty"`
			PyrraDevApp      string `yaml:"pyrra.dev/app,omitempty"`
			PyrraDevEndpoint string `yaml:"pyrra.dev/endpoint,omitempty"`
			PyrraDevOwner    string `yaml:"pyrra.dev/owner,omitempty"`
			PyrraDevSystem   string `yaml:"pyrra.dev/system,omitempty"`
			PyrraDevDomain   string `yaml:"pyrra.dev/domain,omitempty"`
			Role             string `yaml:"role"`
		} `yaml:"labels"`
		Name      string `yaml:"name"`
		Namespace string `yaml:"namespace"`
	} `yaml:"metadata"`
	Spec struct {
		Alerting struct {
			Absent    bool `yaml:"absent"`
			Burnrates bool `yaml:"burnrates"`
		} `yaml:"alerting"`
		Description string `yaml:"description"`
		Indicator   struct {
			Latency struct {
				Success struct {
					Metric string `yaml:"metric"`
				} `yaml:"success"`
				Total struct {
					Metric string `yaml:"metric"`
				} `yaml:"total"`
			} `yaml:"latency,omitempty"`
			Ratio struct {
				Errors struct {
					Metric string `yaml:"metric"`
				} `yaml:"errors"`
				Total struct {
					Metric string `yaml:"metric"`
				} `yaml:"total"`
			} `yaml:"ratio,omitempty"`
		} `yaml:"indicator"`
		Target string `yaml:"target"`
		Window string `yaml:"window"`
	} `yaml:"spec"`
}

type SloSchema struct {
	SloYaml []SloSchemaYaml
	SliConf []slis.SliSchemaYaml
}

func New() *SloSchema {
	return &SloSchema{}
}
