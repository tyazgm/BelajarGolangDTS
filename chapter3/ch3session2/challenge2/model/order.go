package model

type Order struct {
	ID     string `gorm:"primaryKey;type:varchar(255)"`
	Price  int    `gorm:"not null"`
	UserID string
}

type OrderCreateRequest struct {
	Price int `json:"price"`
}

type OrderCreateResponse struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	Price  int    `json:"price"`
}

type OrderGetResponse struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	Price  int    `json:"price"`
}
