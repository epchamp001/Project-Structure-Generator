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

			service.CreateProjectStructure(&cfg, cfg.ProjectName)
			service.InitGoMod(cfg.ProjectName, cfg.ProjectName)
			service.CopyStaticDir("templates/pkg", filepath.Join(cfg.ProjectName, "pkg"))
			service.CopyTemplates("templates", cfg.ProjectName, &cfg)
			slog.Info("✅ Project generated successfully!")
		},
	}

	rootCmd.Flags().StringVarP(&cfg.ProjectName, "name", "n", "", "Project name (required)")
	rootCmd.Flags().StringVarP(&cfg.CmdName, "cmd", "c", "", "Subdirectory name under cmd/ (required)")
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
