package service

import (
	"bytes"
	"io/fs"
	"os"
	"path/filepath"
	"project-generator/internal/config"
	"strings"
	"text/template"
)

func CopyTemplates(srcFS fs.FS, targetRoot string, cfg *config.Config) error {
	// «вырезаем» поддиректорию templates, чтобы WALKDIR шёл оттуда
	tplFS, err := fs.Sub(srcFS, "templates")
	if err != nil {
		return err
	}

	return fs.WalkDir(tplFS, ".", func(relPath string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		// пропускаем директории
		if d.IsDir() {
			return nil
		}
		// работаем только с файлами *.tmpl
		if !strings.HasSuffix(relPath, ".tmpl") {
			return nil
		}

		// флаги включения/отключения шаблонов
		switch {
		case strings.HasPrefix(relPath, "internal/handler/grpc") && !cfg.EnableGRPC:
			return nil
		case strings.HasPrefix(relPath, "internal/handler/http") && !cfg.EnableHTTP:
			return nil
		case strings.HasPrefix(relPath, "grafana") && !cfg.EnableGrafana:
			return nil
		case strings.HasPrefix(relPath, "prometheus") && !cfg.EnableMetrics:
			return nil
		case strings.HasPrefix(relPath, "scripts/k6") && !cfg.EnableLoad:
			return nil
		}

		// определяем куда писать
		var dstPath string
		if relPath == "main.go.tmpl" {
			dstPath = filepath.Join(targetRoot, "cmd", cfg.CmdName, "main.go")
		} else {
			// убираем суффикс .tmpl
			cleanRel := strings.TrimSuffix(relPath, ".tmpl")
			dstPath = filepath.Join(targetRoot, cleanRel)
		}

		// создаём директорию, если нужно
		if err := os.MkdirAll(filepath.Dir(dstPath), 0o755); err != nil {
			return err
		}

		// считываем шаблон из fs
		data, err := fs.ReadFile(tplFS, relPath)
		if err != nil {
			return err
		}
		// парсим и выполняем шаблон
		tmpl, err := template.New(relPath).Parse(string(data))
		if err != nil {
			return err
		}
		var buf bytes.Buffer
		if err := tmpl.Execute(&buf, cfg); err != nil {
			return err
		}

		// пишем файл
		return os.WriteFile(dstPath, buf.Bytes(), 0o644)
	})
}
