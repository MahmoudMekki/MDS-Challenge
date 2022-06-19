package models

import "time"

const (
	OrdersTableName = "orders"
	OrdersMQTopic   = "orders"
)

type Order struct {
	Id        int       `gorm:"column:id; primary_key;auto_increment" json:"-"`
	ProductId int       `gorm:"column:product_id;not null" json:"-"`
	Amount    int       `gorm:"column:amount;not null" json:"amount"`
	CreatedAt time.Time `gorm:"column:created_at;not null" json:"created_at"`
}

type OrderCSV struct {
	Country     string `csv:"country" json:"country"`
	SKU         string `csv:"sku" json:"sku"`
	Name        string `csv:"name" json:"name"`
	StockChange int    `csv:"stock_change" json:"stock_change"`
}

func (o *Order) TableName() string {
	return OrdersTableName
}
