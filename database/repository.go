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
	CreateOrganization(ctx context.Context, data *models.Organization) error
	UpdateOrganization(ctx context.Context, data *models.Organization) error
	DeleteOrganization(ctx context.Context, query *models.Organization) error
	GetOrganization(ctx context.Context, query *models.Organization) (*models.Organization, error)
	GetOrganizations(ctx context.Context, query *models.Organization) ([]*models.Organization, error)

	CreateUser(ctx context.Context, data *models.User) error
	UpdateUser(ctx context.Context, data *models.User) error
	DeleteUser(ctx context.Context, query *models.User) error
	GetUser(ctx context.Context, query *models.User) (*models.User, error)
	GetUsers(ctx context.Context, query *models.User) ([]*models.User, error)

	CreateSubscriber(ctx context.Context, data *models.Subscriber) error
	DeleteSubscriber(ctx context.Context, query *models.Subscriber) error
	GetSubscriber(ctx context.Context, query *models.Subscriber) (*models.Subscriber, error)
	GetSubscribers(ctx context.Context, query *models.Subscriber) ([]*models.Subscriber, error)
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
		&models.Subscriber{},
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
