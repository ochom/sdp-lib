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

// CreateShortcode ...
func (i *impl) CreateShortcode(ctx context.Context, data *models.Shortcode) error {
	return i.DB.Create(data).Error
}

// UpdateShortcode ...
func (i *impl) UpdateShortcode(ctx context.Context, data *models.Shortcode) error {
	return i.DB.Save(data).Error
}

// DeleteShortcode ...
func (i *impl) DeleteShortcode(ctx context.Context, query *models.Shortcode) error {
	return i.DB.Where(query).Delete(&models.Shortcode{}).Error
}

// GetShortcode ...
func (i *impl) GetShortcode(ctx context.Context, query *models.Shortcode) (*models.Shortcode, error) {
	var data models.Shortcode
	err := i.DB.Where(query).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// GetShortcodes ...
func (i *impl) GetShortcodes(ctx context.Context, query *models.Shortcode) ([]*models.Shortcode, error) {
	var data []*models.Shortcode
	err := i.DB.Where(query).Find(&data).Error
	return data, err
}
