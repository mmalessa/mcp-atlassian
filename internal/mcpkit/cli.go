package mcpkit

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// claudeName is the name used when registering with the `claude` CLI.
// It strips the conventional "mcp-" prefix from the app name:
//   "mcp-atlassian" -> "atlassian".
func (a App) claudeName() string {
	return strings.TrimPrefix(a.Name, "mcp-")
}

// binaryPath returns the absolute, symlink-resolved path to the running binary.
func binaryPath() (string, error) {
	exe, err := os.Executable()
	if err != nil {
		return "", err
	}
	if resolved, err := filepath.EvalSymlinks(exe); err == nil {
		exe = resolved
	}
	return filepath.Abs(exe)
}

func (a App) registerInClaude() error {
	exe, err := binaryPath()
	if err != nil {
		return fmt.Errorf("resolving binary path: %w", err)
	}

	name := a.claudeName()

	// Remove any existing registration first so `add` is idempotent.
	// Ignore error — may simply not be registered yet.
	_ = exec.Command("claude", "mcp", "remove", name).Run()

	cmd := exec.Command("claude", "mcp", "add", name, "--", exe)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("claude mcp add %s: %w", name, err)
	}
	fmt.Fprintf(os.Stderr, "Registered %q -> %s\n", name, exe)
	return nil
}

func (a App) unregisterFromClaude() error {
	name := a.claudeName()
	cmd := exec.Command("claude", "mcp", "remove", name)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("claude mcp remove %s: %w", name, err)
	}
	return nil
}
