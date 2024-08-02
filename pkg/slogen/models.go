package slogen

type CmdOpts struct {
	AppName        string
	AppType        string
	AppListFile    string
	SlisFile       string
	OutputFile     string
	BackstageURL   string
	BackstageToken string
	PrometheusURL  string
}

func NewOpts() CmdOpts {
	return CmdOpts{}
}
