package git

type Git interface {
	Add(path string) error
	AddAll() error
	Checkout(branchName string, create bool) error
	CloneOverHttp(url string, username string, password string) error
	Commit(message string) (string, error)
	Configure(gitEmail string, gitUsername string) error
	GetFilePath(filePath string) string
	GetRepositoryPath() string
	Push(force bool) error
	SetConfig(key string, value string) error
}
