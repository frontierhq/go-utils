package external_git

import (
	"os"
	"path/filepath"
)

type ExternalGit struct {
	repositoryPath string
}

// GetFilePath returns the absolute path to a file in the repository
func (g *ExternalGit) GetFilePath(filePath string) string {
	return filepath.Join(g.repositoryPath, filePath)
}

// GetRepositoryPath returns the absolute path to the repository
func (g *ExternalGit) GetRepositoryPath() string {
	return g.repositoryPath
}

// NewGit creates a new ExternalGit
func NewGit(repositoryPath string) *ExternalGit {
	return &ExternalGit{
		repositoryPath: repositoryPath,
	}
}

// NewClonedGit creates a new ExternalGit from a clone
func NewClonedGit(remoteUrl string, remoteUsername string, remotePassword string, gitEmail string, gitUsername string) (*ExternalGit, error) {
	repositoryPath, err := os.MkdirTemp("", "")
	if err != nil {
		return nil, err
	}

	git := NewGit(repositoryPath)

	err = git.CloneOverHttp(remoteUrl, remoteUsername, remotePassword)
	if err != nil {
		return nil, err
	}

	err = git.SetConfig("user.email", gitEmail)
	if err != nil {
		return nil, err
	}

	err = git.SetConfig("user.name", gitUsername)
	if err != nil {
		return nil, err
	}

	return &ExternalGit{
		repositoryPath: repositoryPath,
	}, nil
}
