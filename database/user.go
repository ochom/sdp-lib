package database

import (
	"context"

	"github.com/ochom/sdp-lib/models"
)

// CreateUser ...
func (i *impl) CreateUser(ctx context.Context, data *models.User) error {
	return i.DB.Create(data).Error
}

// UpdateUser ...
func (i *impl) UpdateUser(ctx context.Context, data *models.User) error {
	return i.DB.Save(data).Error
}

// DeleteUser ...
func (i *impl) DeleteUser(ctx context.Context, query *models.User) error {
	return i.DB.Where(query).Delete(&models.User{}).Error
}

// GetUser ...
func (i *impl) GetUser(ctx context.Context, query *models.User) (*models.User, error) {
	var data models.User
	err := i.DB.Where(query).First(&data).Error
	return &data, err
}

// GetUsers ...
func (i *impl) GetUsers(ctx context.Context, query *models.User) ([]*models.User, error) {
	var data []*models.User
	err := i.DB.Where(query).Find(&data).Error
	return data, err
}
