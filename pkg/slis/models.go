package slis

import "github.com/jenciso/slis-pyrra/pkg/apps"

type SliSchemaYaml struct {
	Name string `yaml:"name"`
	Spec struct {
		App  apps.AppYamlSpec `yaml:"app"`
		Conf Conf             `yaml:"conf,omitempty"`
		Slis Slis             `yaml:"slis"`
	} `yaml:"spec"`
}

type Conf struct {
	MetricName             string `yaml:"metricName,omitempty"`
	EndpointLabel          string `yaml:"endpointLabel,omitempty"`
	TargetDefault          string `yaml:"targetDefault,omitempty"`
	BucketDefaultThreshold string `yaml:"bucketDefaultThreshold,omitempty"`
}
type Slis struct {
	Ratio   Ratio   `yaml:"ratio,omitempty"`
	Latency Latency `yaml:"latency,omitempty"`
}
type Ratio struct {
	Endpoints []Endpoints `yaml:"endpoints,omitempty"`
}
type Latency struct {
	Endpoints []Endpoints `yaml:"endpoints,omitempty"`
}
type Endpoints struct {
	Name   string `yaml:"name,omitempty"`
	Target string `yaml:"target,omitempty"`
	Bucket string `yaml:"bucket,omitempty"`
}

type SliSchemaProduct struct {
	SliYaml []SliSchemaYaml
}

type SliSchemaService struct {
	SliYaml []SliSchemaYaml
}

type SliSchemaEntity struct {
	SliYaml []SliSchemaYaml
}

type SliSchemaIntegration struct {
	SliYaml []SliSchemaYaml
}

type SliSchemaFrontend struct {
	SliYaml []SliSchemaYaml
}

type SliSchemaBackend struct {
	SliYaml []SliSchemaYaml
}

type BucketHistogram struct {
	BucketString string
	BucketNumber float64
}

type SliConf struct {
	App     string
	Conf    Conf
	Ratio   Ratio
	Latency Latency
}

func NewConfProductBackend(app string) SliConf {
	return SliConf{
		App: app + `-backend`,
		Conf: Conf{
			MetricName:             "mycompany_bff",
			EndpointLabel:          "endpoint",
			TargetDefault:          "99",
			BucketDefaultThreshold: "0.25",
		},
	}
}

func NewConfProductFrontend(app string) SliConf {
	return SliConf{
		App: app + `-frontend`,
		Conf: Conf{
			MetricName:             "mycompany_web_server",
			EndpointLabel:          "path",
			TargetDefault:          "99",
			BucketDefaultThreshold: "0.25",
		},
	}
}

func NewConfEntity(app string) SliConf {
	return SliConf{
		App: app,
		Conf: Conf{
			MetricName:             "mycompany_rpc",
			EndpointLabel:          "endpoint",
			TargetDefault:          "99",
			BucketDefaultThreshold: "0.1",
		},
	}
}

func NewConfService(app string) SliConf {
	return SliConf{
		App: app,
		Conf: Conf{
			MetricName:             "mycompany_rpc",
			EndpointLabel:          "endpoint",
			TargetDefault:          "99",
			BucketDefaultThreshold: "0.25",
		},
	}
}

func NewConfIntegration(app string) SliConf {
	return SliConf{
		App: app,
		Conf: Conf{
			MetricName:             "mycompany_bff",
			EndpointLabel:          "endpoint",
			TargetDefault:          "99",
			BucketDefaultThreshold: "0.1",
		},
	}
}

func NewConfFrontend(app string) SliConf {
	return SliConf{
		App: app,
		Conf: Conf{
			MetricName:             "mycompany_web_server",
			EndpointLabel:          "path",
			TargetDefault:          "99",
			BucketDefaultThreshold: "0.25",
		},
	}
}

func NewConfBackend(app string) SliConf {
	return SliConf{
		App: app,
		Conf: Conf{
			MetricName:             "mycompany_bff",
			EndpointLabel:          "endpoint",
			TargetDefault:          "99",
			BucketDefaultThreshold: "0.25",
		},
	}
}

func NewProduct() *SliSchemaProduct {
	return &SliSchemaProduct{}
}

func NewService() *SliSchemaService {
	return &SliSchemaService{}
}

func NewEntity() *SliSchemaEntity {
	return &SliSchemaEntity{}
}

func NewIntegration() *SliSchemaIntegration {
	return &SliSchemaIntegration{}
}

func NewFrontend() *SliSchemaFrontend {
	return &SliSchemaFrontend{}
}

func NewBackend() *SliSchemaBackend {
	return &SliSchemaBackend{}
}
