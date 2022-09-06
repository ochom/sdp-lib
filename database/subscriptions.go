package database

import (
	"context"

	"github.com/ochom/sdp-lib/models"
)

func (i *impl) CreateSubscription(ctx context.Context, data *models.Subscription) error {
	return i.DB.Create(data).Error
}

func (i *impl) UpdateSubscription(ctx context.Context, query, data *models.Subscription) error {
	return i.DB.Model(&models.Subscription{}).Where(query).Updates(data).Error
}

func (i *impl) GetSubscription(ctx context.Context, query *models.Subscription) (*models.Subscription, error) {
	var data models.Subscription
	err := i.DB.Where(query).First(&data).Error
	return &data, err
}

func (i *impl) GetSubscriptions(ctx context.Context, query *models.Subscription) ([]*models.Subscription, error) {
	var data []*models.Subscription
	err := i.DB.Where(query).Find(&data).Error
	return data, err
}
