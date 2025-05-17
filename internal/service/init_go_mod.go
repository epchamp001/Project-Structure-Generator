package service

import (
	"log/slog"
	"os"
	"os/exec"
)

func InitGoMod(projectName string, dir string) {
	cmd := exec.Command("go", "mod", "init", projectName)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		slog.Error("failed to init go module", "err", err)
		os.Exit(1)
	}
}
