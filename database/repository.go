package database

import (
	"context"

	"github.com/ochom/sdp-lib/models"
	"github.com/ochom/sdp-lib/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var log = utils.NewLogger()

// Repo ...
type Repo interface {
	CreateSubscription(ctx context.Context, data *models.Subscription) error
	UpdateSubscription(ctx context.Context, query, data *models.Subscription) error
	DeleteSubscription(ctx context.Context, query *models.Subscription) error
	GetSubscription(ctx context.Context, query *models.Subscription) (*models.Subscription, error)
	GetSubscriptions(ctx context.Context, query *models.Subscription) ([]*models.Subscription, error)
}

type impl struct {
	DB *gorm.DB
}

func (i *impl) init() error {
	dns := utils.MustGetEnv("DATABASE_DNS")
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Error("failed to connect database", err.Error())
		return err
	}

	i.DB = db
	return nil
}

func (i *impl) migrate() error {
	return i.DB.AutoMigrate(
		&models.Subscription{},
	)
}

// New ...
func New() Repo {
	i := &impl{}
	if err := i.init(); err != nil {
		log.Error("failed to initialize database", err.Error())
		return &impl{}
	}

	if err := i.migrate(); err != nil {
		log.Error("failed to migrate database", err.Error())
		return &impl{}
	}

	return i
}
