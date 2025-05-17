package main

import (
	"github.com/spf13/cobra"
	"log/slog"
	"os"
	"path/filepath"
	"project-generator/internal/config"
	"project-generator/internal/service"
)

func main() {
	var cfg config.Config
	var targetDir string

	rootCmd := &cobra.Command{
		Use:   "projgen",
		Short: "Project structure generator",
		Run: func(cmd *cobra.Command, args []string) {
			changedH := cmd.Flags().Changed("http")
			changedG := cmd.Flags().Changed("grpc")
			if !changedH && !changedG {
				cfg.EnableGRPC = true
				cfg.EnableHTTP = true
			} else {
				if !changedH {
					cfg.EnableHTTP = false
				}
				if !changedG {
					cfg.EnableGRPC = false
				}
			}

			cfg.EnableGrafana = cmd.Flags().Changed("grafana")
			cfg.EnableRedis = cmd.Flags().Changed("redis")
			cfg.EnableMetrics = cmd.Flags().Changed("metrics")
			cfg.EnableLoad = cmd.Flags().Changed("load")

			if cfg.ProjectName == "" || cfg.CmdName == "" {
				slog.Error("--name and --cmd are required")
				os.Exit(1)
			}

			projectRoot := filepath.Join(targetDir, cfg.ProjectName)
			service.CreateProjectStructure(&cfg, projectRoot)
			service.InitGoMod(cfg.ProjectName, projectRoot)
			service.CopyStaticDir("templates/pkg", filepath.Join(projectRoot, "pkg"))
			service.CopyTemplates("templates", projectRoot, &cfg)
			slog.Info("✅ Project generated successfully!")
		},
	}

	rootCmd.Flags().StringVarP(&cfg.ProjectName, "name", "n", "", "Project name (required)")
	rootCmd.Flags().StringVarP(&cfg.CmdName, "cmd", "c", "", "Subdirectory name under cmd/ (required)")
	rootCmd.Flags().StringVarP(&targetDir, "dir", "d", ".", "Target directory to create project in")
	rootCmd.Flags().BoolVarP(&cfg.EnableGRPC, "grpc", "g", false, "Enable gRPC")
	rootCmd.Flags().BoolVarP(&cfg.EnableHTTP, "http", "t", false, "Enable HTTP")
	rootCmd.Flags().BoolVarP(&cfg.EnableGrafana, "grafana", "f", false, "Enable Grafana")
	rootCmd.Flags().BoolVarP(&cfg.EnableRedis, "redis", "r", false, "Enable Redis")
	rootCmd.Flags().BoolVarP(&cfg.EnableMetrics, "metrics", "m", false, "Enable Prometheus metrics")
	rootCmd.Flags().BoolVarP(&cfg.EnableLoad, "load", "l", false, "Enable load testing")

	if err := rootCmd.Execute(); err != nil {
		slog.Error("command execution failed", "err", err)
		os.Exit(1)
	}
}
