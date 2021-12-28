package schema

import (
	"github.com/jinzhu/gorm"
)

var (
	Database *gorm.DB // Scope is only within schema package
)
