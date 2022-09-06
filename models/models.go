package models

import (
	"encoding/json"
	"time"

	"github.com/ochom/sdp-lib/dto"
	"github.com/ochom/sdp-lib/utils"
	"gorm.io/gorm"
)

// Subscription ...
type Subscription struct {
	ID                  uint           `json:"id" gorm:"primary_key"`
	RequestID           string         `json:"requestID" gorm:"column:requestID"`
	RequestTimeStamp    string         `json:"requestTimeStamp" gorm:"column:requestTimeStamp"`
	Operation           string         `json:"operation"`
	OfferCode           string         `json:"offerCode" gorm:"column:offerCode"`
	TransactionID       string         `json:"transactionID" gorm:"column:transactionID"`
	ClientTransactionID string         `json:"clientTransactionID" gorm:"column:clientTransactionID"`
	Language            string         `json:"language"`
	SubscriberLifeCycle string         `json:"subscriberLifeCycle" gorm:"column:subscriberLifeCycle"`
	SubscriptionStatus  string         `json:"subscriptionStatus" gorm:"column:subscriptionStatus"`
	Channel             string         `json:"channel"`
	Reason              string         `json:"reason"`
	Type                string         `json:"type"`
	OfferName           string         `json:"offerName" gorm:"column:offerName"`
	Mobile              string         `json:"mobile"`
	CreatedAt           time.Time      `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt           time.Time      `json:"updatedAt" gorm:"column:updatedAt"`
	DeletedAt           gorm.DeletedAt `json:"deletedAt" gorm:"column:deletedAt"`
}

// SetValuesFromInput sets values in payload param
func (s *Subscription) SetValuesFromInput(data []dto.ParamValue) {
	for _, v := range data {
		if v.Name == "OfferCode" {
			s.OfferCode = v.Value
		}

		if v.Name == "TransactionId" {
			s.TransactionID = v.Value
		}

		if v.Name == "ClientTransactionId" {
			s.ClientTransactionID = v.Value
		}

		if v.Name == "Language" {
			s.Language = v.Value
		}

		if v.Name == "SubscriberLifeCycle" {
			s.SubscriberLifeCycle = v.Value
		}

		if v.Name == "SubscriptionStatus" {
			s.SubscriptionStatus = v.Value
		}

		if v.Name == "Channel" {
			s.Channel = v.Value
		}

		if v.Name == "Reason" {
			s.Reason = v.Value
		}

		if v.Name == "Type" {
			s.Type = v.Value
		}

		if v.Name == "OfferName" {
			s.OfferName = v.Value
		}

		if v.Name == "Msisdn" {
			s.Mobile = utils.ParseMobile(v.Value)
		}
	}
}

// ToJSON ...
func (s *Subscription) ToJSON() []byte {
	b, err := json.Marshal(s)
	if err != nil {
		return nil
	}

	return b
}
