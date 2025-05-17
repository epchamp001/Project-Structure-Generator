package service

import (
	"os"
	"path/filepath"
	"project-generator/internal/config"
)

func CreateProjectStructure(cfg *config.Config, root string) {
	dirs := []string{
		"cmd/" + cfg.CmdName,
		"internal/repository/storage/postgres",
		"configs",
		"internal/app",
		"internal/handler",
		"internal/handler/dto",
		"internal/mapper",
		"internal/models",
		"internal/service",
		"internal/config",
		"migrations",
		"docs",
		"tests/e2e",
	}

	if cfg.EnableRedis {
		dirs = append(dirs, "internal/repository/cache/redis")
	}
	if cfg.EnableGRPC {
		dirs = append(dirs,
			"internal/service/grpc",
			"internal/handler/grpc/middleware",
			"api/proto",
			"api/pb",
		)
	}
	if cfg.EnableHTTP {
		dirs = append(dirs,
			"internal/service/http",
			"internal/handler/http",
			"internal/handler/http/middleware",
		)
	}
	if cfg.EnableGrafana {
		dirs = append(dirs,
			"grafana/provisioning/dashboards/dashboards",
			"grafana/provisioning/datasourses",
		)
	}
	if cfg.EnableLoad {
		dirs = append(dirs, "scripts/k6")
	}
	if cfg.EnableMetrics {
		dirs = append(dirs, "internal/metrics")
	}

	for _, dir := range dirs {
		path := filepath.Join(root, dir)
		_ = os.MkdirAll(path, 0755)
	}
}
