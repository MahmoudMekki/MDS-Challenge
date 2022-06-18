package migration

import (
	"github.com/go-gormigrate/gormigrate/v2"
)

func listProduct() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		{},
	}
}
