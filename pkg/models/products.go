package models

const (
	ProductsTableName = "products"
)

type Product struct {
	Id          int    `gorm:"column:id; primary_key;auto_increment"`
	SKU         string `gorm:"column:sku;unique;not null"`
	CountryCode string `gorm:"column:country_code;not null"`
	Name        string `gorm:"column:name;not null"`
	Amount      int    `gorm:"column:amount; not null"`
}

func (p *Product) TableName() string {
	return ProductsTableName
}
