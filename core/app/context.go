package app

import (
	"testapi/core/config"

	"gorm.io/gorm"
)

type Context struct {
	Config *config.Config
	Db     *gorm.DB
}
