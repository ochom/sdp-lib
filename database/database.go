package database

import (
	"context"
	"fmt"

	"github.com/ochom/sdp-lib/models"
	"github.com/ochom/sdp-lib/utils"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//go:generate mockgen -source=database.go -destination=mocks/mock_database.go -package=mocks

// Platform ...
type Platform string

const (
	// MySQL ...
	MySQL Platform = "mysql"

	// Postgres ...
	Postgres Platform = "postgres"

	// SQLite ...
	SQLite Platform = "sqlite"
)

// Repo ...
type Repo interface {
	CreateOrganization(ctx context.Context, data *models.Organization) error
	UpdateOrganization(ctx context.Context, data *models.Organization) error
	DeleteOrganization(ctx context.Context, query *models.Organization) error
	GetOrganization(ctx context.Context, query *models.Organization) (*models.Organization, error)
	GetOrganizations(ctx context.Context, query *models.Organization) ([]*models.Organization, error)

	CreateOffer(ctx context.Context, data *models.Offer) error
	UpdateOffer(ctx context.Context, data *models.Offer) error
	DeleteOffer(ctx context.Context, query *models.Offer) error
	GetOffer(ctx context.Context, query *models.Offer) (*models.Offer, error)
	GetOffers(ctx context.Context, query *models.Offer) ([]*models.Offer, error)

	CreateUser(ctx context.Context, data *models.User) error
	UpdateUser(ctx context.Context, data *models.User) error
	DeleteUser(ctx context.Context, query *models.User) error
	GetUser(ctx context.Context, query *models.User) (*models.User, error)
	GetUsers(ctx context.Context, query *models.User) ([]*models.User, error)

	CreateSubscriber(ctx context.Context, data *models.Subscriber) error
	DeleteSubscriber(ctx context.Context, query *models.Subscriber) error
	GetSubscriber(ctx context.Context, query *models.Subscriber) (*models.Subscriber, error)
	GetSubscribers(ctx context.Context, query *models.Subscriber) ([]*models.Subscriber, error)

	CreateContactGroup(ctx context.Context, data *models.ContactGroup) error
	UpdateContactGroup(ctx context.Context, data *models.ContactGroup) error
	DeleteContactGroup(ctx context.Context, query *models.ContactGroup) error
	GetContactGroup(ctx context.Context, query *models.ContactGroup) (*models.ContactGroup, error)
	GetContactGroups(ctx context.Context, query *models.ContactGroup) ([]*models.ContactGroup, error)

	CreateContact(ctx context.Context, data *models.Contact) error
	UpdateContact(ctx context.Context, data *models.Contact) error
	DeleteContact(ctx context.Context, query *models.Contact) error
	GetContact(ctx context.Context, query *models.Contact) (*models.Contact, error)
	GetContacts(ctx context.Context, query *models.Contact) ([]*models.Contact, error)
}

type impl struct {
	DB *gorm.DB
}

func (i *impl) init(pl Platform) error {
	dns := utils.MustGetEnv("DATABASE_DNS")
	var db *gorm.DB
	var err error

	switch pl {
	case MySQL:
		db, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	case Postgres:
		db, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	case SQLite:
		db, err = gorm.Open(sqlite.Open(dns), &gorm.Config{})
	}

	if err != nil {
		return fmt.Errorf("failed to connect database: %w", err)
	}

	i.DB = db

	err = i.DB.AutoMigrate(models.AllModels()...)

	if err != nil {
		return fmt.Errorf("database migration failed: %s", err.Error())
	}

	return nil
}

// New ...
func New(dbPlatform Platform) (Repo, error) {
	i := &impl{}
	if err := i.init(dbPlatform); err != nil {
		return nil, err
	}

	return i, nil
}
