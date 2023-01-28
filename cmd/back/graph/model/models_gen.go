// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"

	"github.com/99designs/gqlgen/graphql"
)

type Category struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CategoryListResp struct {
	List  []*Category `json:"list"`
	Total int         `json:"total"`
}

type CategoryWithFoods struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Foods     []*Food   `json:"Foods"`
}

type CategoryWithFoodsListResp struct {
	List  []*CategoryWithFoods `json:"list"`
	Total int                  `json:"total"`
}

type CreateCategoryResponse struct {
	Code     int       `json:"code"`
	Success  bool      `json:"success"`
	Message  string    `json:"message"`
	Category *Category `json:"category"`
}

type FilterCategory struct {
	ID        *int    `json:"id"`
	Name      *string `json:"name"`
	StartTime *string `json:"start_time"`
	EndTime   *string `json:"end_time"`
	FoodName  *string `json:"food_name"`
}

type Food struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	ShopID      int         `json:"shop_id"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	Shop        *Shop       `json:"shop"`
	Categories  []*Category `json:"Categories"`
}

type Shop struct {
	ID           int        `json:"id"`
	Name         string     `json:"name"`
	LogoFilename string     `json:"logo_filename"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	ShopOwner    *ShopOwner `json:"shop_owner"`
}

type ShopListResp struct {
	List  []*Shop `json:"list"`
	Total int     `json:"total"`
}

type ShopOwner struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Shops     []*Shop   `json:"shops"`
}

type ShopOwnerListResponse struct {
	List  []*ShopOwner `json:"list"`
	Total int          `json:"total"`
}

type CreateShop struct {
	Name        string         `json:"name"`
	File        graphql.Upload `json:"file"`
	ShopOwnerID int            `json:"shop_owner_id"`
}

type CreateShopOwner struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type FilterShop struct {
	ID            *int    `json:"id"`
	Name          *string `json:"name"`
	StartTime     *string `json:"start_time"`
	EndTime       *string `json:"end_time"`
	ShopOwnerID   *int    `json:"shop_owner_id"`
	ShopOwnerName *string `json:"shop_owner_name"`
	Page          int     `json:"page"`
	PageSize      int     `json:"page_size"`
}

type FilterShopOwner struct {
	ID       *int    `json:"id"`
	Name     *string `json:"name"`
	Email    *string `json:"email"`
	ShopID   *int    `json:"shop_id"`
	ShopName *string `json:"shop_name"`
	Page     int     `json:"page"`
	PageSize int     `json:"page_size"`
}

type UpdateCategory struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UpdateShop struct {
	ID          int     `json:"id"`
	Name        *string `json:"name"`
	ShopOwnerID *int    `json:"shop_owner_id"`
}

type UpdateShopOwner struct {
	ID       int     `json:"id"`
	Name     *string `json:"name"`
	Email    *string `json:"email"`
	Password *string `json:"password"`
}
