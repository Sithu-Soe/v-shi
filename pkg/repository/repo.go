package repository

import "v-shi/pkg/ds"

type Repository struct {
	DS *ds.DataSource
}
type RepoConfig struct {
	DS *ds.DataSource
}

func NewRepository(rConfig *RepoConfig) *Repository {
	return &Repository{
		DS: rConfig.DS,
	}
}
