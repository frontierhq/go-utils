package external_git

type ExternalGit struct {
	repositoryPath string
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
