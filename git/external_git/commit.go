package external_git

import "strings"

// Commit commits all staged changes
func (g *ExternalGit) Commit(message string) (string, error) {
	_, err := g.Exec("commit", "-m", message)
	if err != nil {
		return "", err
	}

	commitSha, err := g.Exec("rev-parse", "--short", "HEAD")
	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(commitSha, "\n"), nil
}
