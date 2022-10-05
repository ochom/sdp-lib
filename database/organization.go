package database

import (
	"context"

	"github.com/ochom/sdp-lib/models"
)

// CreateOrganization ...
func (i *impl) CreateOrganization(ctx context.Context, data *models.Organization) error {
	return i.DB.Create(data).Error
}

// UpdateOrganization ...
func (i *impl) UpdateOrganization(ctx context.Context, data *models.Organization) error {
	return i.DB.Save(data).Error
}

// DeleteOrganization ...
func (i *impl) DeleteOrganization(ctx context.Context, query *models.Organization) error {
	return i.DB.Where(query).Delete(&models.Organization{}).Error
}

// GetOrganization ...
func (i *impl) GetOrganization(ctx context.Context, query *models.Organization) (*models.Organization, error) {
	var data models.Organization
	err := i.DB.Where(query).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// GetOrganizations ...
func (i *impl) GetOrganizations(ctx context.Context, query *models.Organization) ([]*models.Organization, error) {
	var data []*models.Organization
	err := i.DB.Where(query).Find(&data).Error
	return data, err
}

// CreateOffer ...
func (i *impl) CreateOffer(ctx context.Context, data *models.Offer) error {
	return i.DB.Create(data).Error
}

// UpdateOffer ...
func (i *impl) UpdateOffer(ctx context.Context, data *models.Offer) error {
	return i.DB.Save(data).Error
}

// DeleteOffer ...
func (i *impl) DeleteOffer(ctx context.Context, query *models.Offer) error {
	return i.DB.Where(query).Delete(&models.Offer{}).Error
}

// GetOffer ...
func (i *impl) GetOffer(ctx context.Context, query *models.Offer) (*models.Offer, error) {
	var data models.Offer
	err := i.DB.Where(query).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// GetOffers ...
func (i *impl) GetOffers(ctx context.Context, query *models.Offer) ([]*models.Offer, error) {
	var data []*models.Offer
	err := i.DB.Where(query).Find(&data).Error
	return data, err
}
