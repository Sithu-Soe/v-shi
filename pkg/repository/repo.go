package repository

import "v-shi/pkg/ds"

type Repository struct {
	DS       *ds.DataSource
	Category *categoryRepository
	Shop     *shopRepository
}
type RepoConfig struct {
	DS *ds.DataSource
}

func NewRepository(rConfig *RepoConfig) *Repository {
	return &Repository{
		DS:       rConfig.DS,
		Category: newCategoryRepository(rConfig),
		Shop:     newShopRepository(rConfig),
	}
}