package database

import (
	"context"

	"github.com/ochom/sdp-lib/models"
)

// CreateContactGroup ...
func (r *impl) CreateContactGroup(ctx context.Context, data *models.ContactGroup) error {
	return r.DB.Create(data).Error
}

// UpdateContactGroup ...
func (r *impl) UpdateContactGroup(ctx context.Context, data *models.ContactGroup) error {
	return r.DB.Save(data).Error
}

// DeleteContactGroup ...
func (r *impl) DeleteContactGroup(ctx context.Context, query *models.ContactGroup) error {
	return r.DB.Where(query).Delete(&models.ContactGroup{}).Error
}

// GetContactGroup ...
func (r *impl) GetContactGroup(ctx context.Context, query *models.ContactGroup) (*models.ContactGroup, error) {
	var data models.ContactGroup
	err := r.DB.Where(query).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// GetContactGroups ...
func (r *impl) GetContactGroups(ctx context.Context, query *models.ContactGroup) ([]*models.ContactGroup, error) {
	var data []*models.ContactGroup
	err := r.DB.Where(query).Find(&data).Error
	return data, err
}

// CreateContact ...
func (r *impl) CreateContact(ctx context.Context, data *models.Contact) error {
	return r.DB.Create(data).Error
}

// UpdateContact ...
func (r *impl) UpdateContact(ctx context.Context, data *models.Contact) error {
	return r.DB.Save(data).Error
}

// DeleteContact ...
func (r *impl) DeleteContact(ctx context.Context, query *models.Contact) error {
	return r.DB.Where(query).Delete(&models.Contact{}).Error
}

// GetContact ...
func (r *impl) GetContact(ctx context.Context, query *models.Contact) (*models.Contact, error) {
	var data models.Contact
	err := r.DB.Where(query).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// GetContacts ...
func (r *impl) GetContacts(ctx context.Context, query *models.Contact) ([]*models.Contact, error) {
	var data []*models.Contact
	err := r.DB.Where(query).Find(&data).Error
	return data, err
}
