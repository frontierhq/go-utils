package external_git

import (
	"fmt"

	exec "github.com/gofrontier-com/go-utils/exec"
)

// Exec executes git commands in the context of the repository
func (g *ExternalGit) Exec(arg ...string) (string, error) {
	stdout, stderr, exitCode := exec.RunCommand("git", g.repositoryPath, arg...)
	if exitCode != 0 {
		return stdout, fmt.Errorf("(%d) %s", exitCode, stderr)
	}

	return stdout, nil
}
