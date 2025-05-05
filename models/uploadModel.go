package models

import "time"

type Upload struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Size      float64   `json:"size"`
	Type      string    `json:"type"`
	URL       string    `json:"url"`
	Ext       string    `json:"ext"`
	CreatedBy int       `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
}

func (b *Upload) TableName() string {
	return "upload"
}
