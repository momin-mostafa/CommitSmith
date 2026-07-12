package git

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestIsRepo(t *testing.T) {
	dir := t.TempDir()
	cmd := exec.Command("git", "init")
	cmd.Dir = dir
	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed to init git repo: %v", err)
	}

	origDir, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(origDir)

	if !IsRepo() {
		t.Error("IsRepo() = false, want true")
	}
}

func TestIsRepoFalse(t *testing.T) {
	dir := t.TempDir()
	origDir, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(origDir)

	if IsRepo() {
		t.Error("IsRepo() = true, want false")
	}
}

func TestHasStagedChanges(t *testing.T) {
	dir := t.TempDir()
	cmd := exec.Command("git", "init")
	cmd.Dir = dir
	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed to init git repo: %v", err)
	}

	cmd = exec.Command("git", "config", "user.email", "test@test.com")
	cmd.Dir = dir
	cmd.Run()

	cmd = exec.Command("git", "config", "user.name", "Test")
	cmd.Dir = dir
	cmd.Run()

	os.WriteFile(filepath.Join(dir, "test.txt"), []byte("hello"), 0644)

	origDir, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(origDir)

	if HasStagedChanges() {
		t.Error("HasStagedChanges() = true before staging, want false")
	}

	cmd = exec.Command("git", "add", "test.txt")
	cmd.Dir = dir
	cmd.Run()

	if !HasStagedChanges() {
		t.Error("HasStagedChanges() = false after staging, want true")
	}
}
