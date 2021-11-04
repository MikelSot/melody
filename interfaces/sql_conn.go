package interfaces

import "gorm.io/gorm"

type Pooler interface{
	Pool() *gorm.DB
}

type Connectioner interface{
	Connection()
}