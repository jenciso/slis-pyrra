package slos

import (
	"fmt"
	"os"
	"regexp"

	"github.com/jenciso/slis-pyrra/pkg/apps"
	"github.com/jenciso/slis-pyrra/pkg/slis"
	"github.com/jenciso/slis-pyrra/pkg/slogen"
	"gopkg.in/yaml.v2"
)

type SliSchemaApp struct {
	Name string `yaml:"name"`
	Spec struct {
		App  apps.AppYamlSpec `yaml:"app"`
		Slis any              `yaml:"slis"`
	} `yaml:"spec,omitempty"`
}

func MergeSlosResult(d []SloSchemaYaml) string {
	var output string
	listSize := len(d)
	for i, file := range d {
		s, _ := yaml.Marshal(&file)
		output += fmt.Sprint(string(s))
		if i != listSize-1 {
			output += fmt.Sprintln("---")
		}
	}
	return output
}

func SliLoadSpec(filename string) (sliList []slis.SliSchemaYaml) {
	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading YAML file: %s\n", err)
		os.Exit(1)
	}
	err = yaml.Unmarshal(yamlFile, &sliList)
	if err != nil {
		panic(err)
	}
	return sliList
}

func sanityzeEndpoint(e string) string {
	e = regexp.MustCompile(`[^a-zA-Z0-9]+`).ReplaceAllString(e, ".")
	e = slogen.TruncateString(e, 48)
	e = regexp.MustCompile(`^[^a-zA-Z0-9]`).ReplaceAllString(e, "")
	e = regexp.MustCompile(`[^a-zA-Z0-9]$`).ReplaceAllString(e, "")
	if e == "" {
		e = "rootpath"
	}
	return e
}

func (s *SloSchemaYaml) loadDefaultSliValues(sli slis.SliSchemaYaml) {
	s.APIVersion = "pyrra.dev/v1alpha1"
	s.Kind = "ServiceLevelObjective"
	s.Metadata.Namespace = sli.Spec.App.Namespace
	s.Metadata.Labels.PyrraDevApp = sli.Name
	s.Metadata.Labels.PyrraDevOwner = sli.Spec.App.Owner
	s.Metadata.Labels.PyrraDevSystem = sli.Spec.App.System
	s.Metadata.Labels.PyrraDevDomain = sli.Spec.App.Domain
	s.Metadata.Labels.PyrraDevType = sli.Spec.App.Type
	// Here CUJ
	s.Metadata.Labels.Role = "prometheus-rule"
	s.Spec.Alerting.Absent = false
	s.Spec.Alerting.Burnrates = true
	s.Spec.Window = "4w"
}
