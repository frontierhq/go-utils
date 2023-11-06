package external_git

func (g *ExternalGit) SetConfig(key string, value string) error {
	_, err := g.Exec("config", "--local", key, value)

	return err
}
