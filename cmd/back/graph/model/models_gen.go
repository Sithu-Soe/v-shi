// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"

	"github.com/99designs/gqlgen/graphql"
)

type CategoriesResp struct {
	List  []*Category `json:"list"`
	Total int         `json:"total"`
}

type Category struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type FilterCatrgory struct {
	ID        *int    `json:"id"`
	Name      *string `json:"name"`
	StartTime *string `json:"start_time"`
	EndTime   *string `json:"end_time"`
}

type FilterShop struct {
	ID        *int    `json:"id"`
	Name      *string `json:"name"`
	StartTime *string `json:"start_time"`
	EndTime   *string `json:"end_time"`
}

type Shop struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	LogoFilename string    `json:"logo_filename"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type ShopsResp struct {
	List  []*Shop `json:"list"`
	Total int     `json:"total"`
}

type UpdateCategory struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CreateShop struct {
	Name string         `json:"name"`
	File graphql.Upload `json:"file"`
}
