package external_git

import "path/filepath"

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
