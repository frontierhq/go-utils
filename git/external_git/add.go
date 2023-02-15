package external_git

// Add stages untracked files
func (g *ExternalGit) Add(path string) error {
	_, err := g.Exec("add", path)

	return err
}

// AddAll stages all untracked files
func (g *ExternalGit) AddAll() error {
	_, err := g.Exec("add", "--all")

	return err
}
