package external_git

// Push pushes the current branch to the "origin" remote
func (g *ExternalGit) Push(force bool) error {
	var args []string
	if force {
		args = []string{"push", "-u", "origin", "HEAD", "--force"}
	} else {
		args = []string{"push", "-u", "origin", "HEAD"}
	}

	_, err := g.Exec(args...)

	return err
}
