package hosting

import "github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/entity"

//type Provider interface {
//	ListRepos() entity.Repository
//}



type ProviderPlugInRegistry struct {
}

func (p ProviderPlugInRegistry) RegisterProviderPlugIn(providerPlugIn ProviderPlugIn) {
}

func (p ProviderPlugInRegistry) GetHostingProviderClient(
	providerType entity.HostingProviderType,
	credential entity.Credential) ProviderClient {
}

func (p ProviderPlugInRegistry) ListProviderPlugIns() []ProviderPlugIn {
}

type ProviderPlugIn interface {
	NewProviderClient(credential entity.Credential) ProviderClient
}

type ProviderClient interface {
	ListRepos() entity.Repository
	SubscribeEvents(repository entity.Repository)
}

//// Store Credentials (maybe DB)
//// Find the associated credential for that repo
//type CredentialRegistry struct {
//}
//
//func (c CredentialRegistry) AddCredential(credential Credential) {
//}
//
//func (c CredentialRegistry) RemoveCredential(credentialId int) {
//}
//
//func (c CredentialRegistry) GetCredential(credentialId int) Credential {
//}
//
//type FindCredentialsQuery struct {
//	CredentialName string
//	HostingProviderName string
//	CredentialType CredentialType
//}
//
//func (c CredentialRegistry) SearchCredentials(query FindCredentialsQuery) []Credential {
//}
//
//func (c CredentialRegistry) ListCredentials() []Credential {
//}
//
//func (c CredentialRegistry) UpdateCredential(credential Credential) {
//}
//
//type Credential struct {
//	id int
//	name string
//	providerName string
//	credentialType CredentialType
//	value string
//}
//
//type Randomizer struct {
//
//}
//
//func (r Randomizer) GenerateString(length int) string {
//}
//
//type CredentialType int
//
//type RepositoryRegistry struct {
//	hostingProviderRegistry ProviderRegistry
//	randomizer Randomizer
//	secretLength int
//}
//
//func (r RepositoryRegistry) CreateRepo(repository entity.Repository) {
//}
//
//func (r RepositoryRegistry) ActivateRepo(repository entity.Repository) {
//	provider := r.hostingProviderRegistry.GetProvider(repository.ProviderName)
//	secret := r.randomizer.GenerateString(r.secretLength)
//	provider.ActivateRepo(repository, secret)
//	repository.HostingSecret = secret
//	r.UpdateRepo(repository)
//}
//
//func (r RepositoryRegistry) UpdateRepo(repository entity.Repository) {
//}
//
//func (r RepositoryRegistry) ListRepos() []entity.Repository {
//}
//
//
//type Provider interface {
//	Init(credential Credential)
//	GetName() string
//	ActivateRepo(repo entity.Repository, secret string)
//	ListRepos(credential Credential)
//}
//
//type ProviderRegistry struct {
//	hostingProviders map[string]Provider
//}
//
//func (p ProviderRegistry) RegisterProvider(provider Provider) {
//	name := provider.GetName()
//	p.hostingProviders[name] = provider
//}
//
//func (p ProviderRegistry) GetProvider(name string) Provider{
//	return p.hostingProviders[name]
//}


