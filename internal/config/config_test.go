package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadDefaults(t *testing.T) {
	homeDir, _ := os.UserHomeDir()
	configPath := filepath.Join(homeDir, ".git_comment.yaml")
	_, err := os.Stat(configPath)
	if err == nil {
		t.Skip("Skipping test: config file exists")
	}

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	if cfg.Model != defaultModel {
		t.Errorf("Model = %q, want %q", cfg.Model, defaultModel)
	}
	if cfg.Host != defaultHost {
		t.Errorf("Host = %q, want %q", cfg.Host, defaultHost)
	}
	if cfg.Temperature != defaultTemperature {
		t.Errorf("Temperature = %f, want %f", cfg.Temperature, defaultTemperature)
	}
	if cfg.MaxOptions != defaultMaxOptions {
		t.Errorf("MaxOptions = %d, want %d", cfg.MaxOptions, defaultMaxOptions)
	}
}

func TestLoadFromFile(t *testing.T) {
	dir := t.TempDir()
	configPath := filepath.Join(dir, ".git_comment.yaml")

	content := `model: deepseek-coder
host: http://localhost:11435
temperature: 0.5
max_options: 5
use_conventional_commits: false
`
	os.WriteFile(configPath, []byte(content), 0644)

	origHome := os.Getenv("HOME")
	os.Setenv("HOME", dir)
	defer os.Setenv("HOME", origHome)

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	if cfg.Model != "deepseek-coder" {
		t.Errorf("Model = %q, want %q", cfg.Model, "deepseek-coder")
	}
	if cfg.Host != "http://localhost:11435" {
		t.Errorf("Host = %q, want %q", cfg.Host, "http://localhost:11435")
	}
	if cfg.Temperature != 0.5 {
		t.Errorf("Temperature = %f, want %f", cfg.Temperature, 0.5)
	}
	if cfg.MaxOptions != 5 {
		t.Errorf("MaxOptions = %d, want %d", cfg.MaxOptions, 5)
	}
	if cfg.UseConventionalCommits {
		t.Error("UseConventionalCommits = true, want false")
	}
}
