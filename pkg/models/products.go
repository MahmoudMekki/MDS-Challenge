package models

const (
	ProductsTableName = "products"
)

type Product struct {
	Id          int     `gorm:"column:id; primary_key;auto_increment" json:"-"`
	SKU         string  `gorm:"column:sku;unique;not null" json:"sku"`
	CountryCode string  `gorm:"column:country_code;not null" json:"country_code"`
	Name        string  `gorm:"column:name;not null" json:"name"`
	Amount      int     `gorm:"column:amount; not null" json:"amount"`
	Orders      []Order `json:"orders,omitempty"`
}

func (p *Product) TableName() string {
	return ProductsTableName
}
