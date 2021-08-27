package vcs

type VersionControlSystem interface {
	NewRepository(rootPath string) Repository
}
