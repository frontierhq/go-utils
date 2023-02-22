package external_git

// SetConfig sets the value for the given key
func (g *ExternalGit) SetConfig(key string, value string) error {
	_, err := g.Exec("config", "--local", key, value)

	return err
}
