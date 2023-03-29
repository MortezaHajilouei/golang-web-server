package config

import (
	"github.com/harranali/authority"
	"gorm.io/gorm"
)

func Auth(db *gorm.DB) {
	authority.New(authority.Options{
		TablesPrefix: "authority_",
		DB:           db,
	})
}
