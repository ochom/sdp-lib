package database

import (
	"github.com/ochom/sdp-lib/models"
	"github.com/ochom/sdp-lib/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var log = utils.NewLogger()

// Repo ...
type Repo interface {
}

type impl struct {
	DB *gorm.DB
}

func (i *impl) init() {
	dns := utils.MustGetEnv("DATABASE_DNS")
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		log.Error("failed to connect database", err.Error())
		return
	}

	i.DB = db
}

func (i *impl) migrate() {
	i.DB.AutoMigrate(
		&models.Subscription{},
	)
}

// New ...
func New() Repo {
	i := &impl{}
	i.init()
	i.migrate()
	return i
}
