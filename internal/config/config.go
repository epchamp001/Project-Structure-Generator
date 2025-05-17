package config

type Config struct {
	ProjectName   string
	CmdName       string
	EnableGRPC    bool
	EnableHTTP    bool
	EnableRedis   bool
	EnableGrafana bool
	EnableMetrics bool
	EnableLoad    bool
}
