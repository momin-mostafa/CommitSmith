package git

import (
	"fmt"
	"os/exec"
	"strings"
)

func IsRepo() bool {
	cmd := exec.Command("git", "rev-parse", "--git-dir")
	return cmd.Run() == nil
}

func HasStagedChanges() bool {
	cmd := exec.Command("git", "diff", "--cached", "--quiet")
	cmd.Run()
	return cmd.ProcessState.ExitCode() != 0
}

func GetStagedDiff() (string, error) {
	cmd := exec.Command("git", "diff", "--cached")
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get staged diff: %w", err)
	}
	return string(out), nil
}

func Commit(message string) (string, error) {
	cmd := exec.Command("git", "commit", "-m", message)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return string(out), fmt.Errorf("commit failed: %w", err)
	}
	return string(out), nil
}

func Push() (string, error) {
	cmd := exec.Command("git", "push")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return string(out), fmt.Errorf("push failed: %w", err)
	}
	return string(out), nil
}

func GetDiffStats(diff string) string {
	var added, removed, files []string
	lines := strings.Split(diff, "\n")
	
	for _, line := range lines {
		if strings.HasPrefix(line, "diff --git") {
			parts := strings.Split(line, " b/")
			if len(parts) > 1 {
				files = append(files, parts[1])
			}
		} else if strings.HasPrefix(line, "+") && !strings.HasPrefix(line, "+++") {
			added = append(added, line[1:])
		} else if strings.HasPrefix(line, "-") && !strings.HasPrefix(line, "---") {
			removed = append(removed, line[1:])
		}
	}
	
	return fmt.Sprintf("Files changed: %d\nLines added: %d\nLines removed: %d", 
		len(files), len(added), len(removed))
}

func TruncateDiff(diff string, maxChars int) string {
	if len(diff) <= maxChars {
		return diff
	}
	return diff[:maxChars] + "\n\n[Diff truncated due to size]"
}
