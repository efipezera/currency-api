package models

type Transfer struct {
	ID    uint    `json:"id" gorm:"primary_key"`
	Value float64 `json:"value"`
}

type CreateTransferInput struct {
	Value float64 `json:"value" binding:"required"`
}
