package apps

type AppSchemaYaml struct {
	Name string      `yaml:"name"`
	Spec AppYamlSpec `yaml:"spec"`
}

type AppYamlSpec struct {
	Autoconf  bool     `yaml:"autoconf,omitempty"`
	Owner     string   `yaml:"owner,omitempty"`
	System    string   `yaml:"system,omitempty"`
	Domain    string   `yaml:"domain,omitempty"`
	Namespace string   `yaml:"namespace,omitempty"`
	Type      string   `yaml:"type,omitempty"`
	Cuj       []string `yaml:"cuj,omitempty"`
}

type AppSchema struct {
	AppYaml []AppSchemaYaml
}

func New() *AppSchema {
	return &AppSchema{}
}
