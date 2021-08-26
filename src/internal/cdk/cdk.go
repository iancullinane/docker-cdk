package cdk

import (
	"fmt"
	"os/exec"
	"strings"
)

type cdk struct{}

func New() *cdk {
	return &cdk{}
}

// List stack
func (c *cdk) List(repoPath string, contexts map[string]string) ([]string, error) {
	args := []string{"ls", "-la"}
	cmd := exec.Command("npm", args...)
	out, err := cmd.CombinedOutput()
	if err != nil || cmd.ProcessState.ExitCode() != 0 {
		return nil, fmt.Errorf("cdk list failed: %s %v", string(out), err)
	}
	lists := strings.Split(strings.Trim(string(out), "\n"), "\n")[3:]
	return lists, nil
}
