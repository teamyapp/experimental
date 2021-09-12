package service

import (
	"github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/entity"
	"github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/hosting"
)

type Repository struct {
	teamService Team
	hostingProviderPlugInRegistry hosting.ProviderPlugInRegistry
	teamIdToHostingProviders map[int]map[entity.HostingProviderType]hosting.ProviderClient
}

func (r Repository) SyncRepos(teamId int) {
	team := r.teamService.GetTeam(teamId)

	for providerType := range team.HostingProviderCredential {
		client := r.getHostingProviderClient(team, providerType)
		providerRepos := hosting.ProviderClient.ListRepos()
		// Update existing repos
		// Record new repos
		// How to handle repos whose name or namespace has changed?
		// How to handle repos that are deleted?
	}
}

func (r Repository) getHostingProviderClient(
	team entity.Team,
	providerType entity.HostingProviderType) hosting.ProviderClient {

	hostingProviderClient := r.hostingProviderPlugInRegistry.GetHostingProviderClient(providerType, team.HostingProviderCredential[providerType])

}

func (r Repository) ListRepos(teamId int) []entity.Repository {
}

func (r Repository) ActivateRepo(teamId int, repositoryId int) {
	repo := r.GetRepo(repositoryId)
	team := r.teamService.GetTeam(teamId)
	providerClient := r.getHostingProviderClient(team, repo.ProviderType)
	providerClient.SubscribeEvents(repo)
}

func (r Repository) GetRepo(repoId int) entity.Repository{

}

func (r Repository) DeactivateRepo(repositoryId int) {
}

func (r Repository) UpdateRepo(repo entity.Repository) {
}