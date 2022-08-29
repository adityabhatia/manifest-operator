package types

type ManifestResolver interface {
	Get(object BaseCustomObject) (InstallationSpec, error)
}

type InstallationSpec struct {
	ChartPath   string
	ReleaseName string

	// configFlags only string, bool and int types are supported
	// check: https://github.com/helm/helm/blob/d7b4c38c42cb0b77f1bcebf9bb4ae7695a10da0b/pkg/action/install.go#L67
	ConfigFlags map[string]interface{}
	SetFlags    map[string]interface{}
}
