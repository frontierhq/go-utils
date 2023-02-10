package external_git

import (
	netUrl "net/url"
)

// CloneOverHttp clones a Git repository using HTTP(S)
func (g *ExternalGit) CloneOverHttp(url string, username string, password string) error {
	parsedUrl, err := netUrl.Parse(url)
	if err != nil {
		return err
	}

	parsedUrl.User = netUrl.UserPassword(username, password)

	_, err = g.Exec("clone", parsedUrl.String(), g.repositoryPath)
	if err != nil {
		return err
	}

	return nil
}
