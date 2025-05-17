package service

import (
	"bytes"
	"os"
	"path/filepath"
	"project-generator/internal/config"
	"strings"
	"text/template"
)

func CopyTemplates(srcDir, targetRoot string, cfg *config.Config) {
	_ = filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || !strings.HasSuffix(path, ".tmpl") {
			return nil
		}

		rel, err := filepath.Rel(srcDir, path)
		if err != nil {
			return err
		}

		switch {
		case strings.Contains(path, "internal/handler/grpc") && !cfg.EnableGRPC:
			return nil
		case strings.Contains(path, "internal/handler/http") && !cfg.EnableHTTP:
			return nil
		case strings.Contains(path, "grafana") && !cfg.EnableGrafana:
			return nil
		case strings.Contains(path, "prometheus") && !cfg.EnableMetrics:
			return nil
		case strings.Contains(path, "scripts/k6") && !cfg.EnableLoad:
			return nil
		}

		var dstPath string
		if rel == "main.go.tmpl" {
			dstPath = filepath.Join(targetRoot, "cmd", cfg.CmdName, "main.go")
		} else {
			dstPath = filepath.Join(targetRoot, strings.TrimSuffix(rel, ".tmpl"))
		}

		_ = os.MkdirAll(filepath.Dir(dstPath), 0755)

		tmpl, err := template.ParseFiles(path)
		if err != nil {
			return err
		}

		var buf bytes.Buffer
		err = tmpl.Execute(&buf, cfg)
		if err != nil {
			return err
		}

		return os.WriteFile(dstPath, buf.Bytes(), 0644)
	})
}
