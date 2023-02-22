package external_git

import (
	"bytes"
	"os/exec"
)

// Exec executes git commands in the context of the repository
func (g *ExternalGit) Exec(arg ...string) (string, error) {
	cmd := exec.Command("git", arg...)
	cmd.Dir = g.repositoryPath

	stdoutBuffer := new(bytes.Buffer)
	cmd.Stdout = stdoutBuffer
	// cmd.Stdout = os.Stdout

	stderrBuffer := new(bytes.Buffer)
	cmd.Stderr = stderrBuffer
	// cmd.Stderr = os.Stderr

	err := cmd.Run()

	return stdoutBuffer.String(), err
}
