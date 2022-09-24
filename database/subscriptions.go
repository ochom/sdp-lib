package database

import (
	"context"

	"github.com/ochom/sdp-lib/models"
)

func (i *impl) CreateSubscriber(ctx context.Context, data *models.Subscriber) error {
	return i.DB.Create(data).Error
}

func (i *impl) DeleteSubscriber(ctx context.Context, query *models.Subscriber) error {
	return i.DB.Where(query).Delete(&models.Subscriber{}).Error
}

func (i *impl) GetSubscriber(ctx context.Context, query *models.Subscriber) (*models.Subscriber, error) {
	var data models.Subscriber
	err := i.DB.Where(query).First(&data).Error
	return &data, err
}

func (i *impl) GetSubscribers(ctx context.Context, query *models.Subscriber) ([]*models.Subscriber, error) {
	var data []*models.Subscriber
	err := i.DB.Where(query).Find(&data).Error
	return data, err
}
