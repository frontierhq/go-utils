package external_git

// HasChanges checks for changes in the current branch
func (g *ExternalGit) HasChanges() (bool, error) {
	var args []string
	args = []string{"status", "-s"}

	content, err := g.Exec(args...)
	if err != nil {
		return false, err
	}

	if content == "" {
		return false, nil
	}

	return true, nil
}
