package models

import "time"

const (
	OrdersTableName = "orders"
)

type Order struct {
	Id        int       `gorm:"column:id; primary_key;auto_increment"`
	ProductId int       `gorm:"column:product_id;not null"`
	Amount    uint      `gorm:"column:amount;not null"`
	CreatedAt time.Time `gorm:"column:created_at;not null"`
	Product   Product   `gorm:"foreign_key:product_id"`
}

func (o *Order) TableName() string {
	return OrdersTableName
}
