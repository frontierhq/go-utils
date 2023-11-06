package external_git

// Configure configures the repository
func (g *ExternalGit) Configure(gitEmail string, gitUsername string) error {
	err := g.SetConfig("user.email", gitEmail)
	if err != nil {
		return err
	}

	err = g.SetConfig("user.name", gitUsername)
	if err != nil {
		return err
	}

	return nil
}

// SetConfig sets the value for the given key
func (g *ExternalGit) SetConfig(key string, value string) error {
	_, err := g.Exec("config", "--local", key, value)

	return err
}
