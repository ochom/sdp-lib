package dto

import (
	"encoding/json"

	"github.com/ochom/sdp-lib/utils"
)

// ParamValue ...
type ParamValue struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// RequestParam ...
type RequestParam struct {
	Data []ParamValue `json:"data"`
}

// ResponseParam ...
type ResponseParam struct {
	Status      string `json:"status"`
	StatusCode  string `json:"statusCode"`
	Description string `json:"description"`
}

// InputSDP ..
type InputSDP struct {
	RequestID        string       `json:"requestId"`
	RequestTimeStamp string       `json:"requestTimeStamp"`
	Operation        string       `json:"operation"`
	RequestParam     RequestParam `json:"requestParam"`
}

// SubscriptionRequest ..
type SubscriptionRequest struct {
	RequestID        string       `json:"requestId"`
	Channel          string       `json:"channel"`
	RequestTimeStamp string       `json:"requestTimeStamp"`
	Operation        string       `json:"operation"`
	RequestParam     RequestParam `json:"requestParam"`
}

// SubscriptionResponse ...
type SubscriptionResponse struct {
	RequestID         string        `json:"requestId"`
	ResponseID        string        `json:"responseId"`
	ResponseTimeStamp string        `json:"responseTimeStamp"`
	Channel           string        `json:"channel"`
	Operation         string        `json:"operation"`
	RequestParam      RequestParam  `json:"requestParam"`
	ResponseParam     ResponseParam `json:"responseParam"`
}

// NewSubscription ...
type NewSubscription struct {
	RequestID           string `json:"requestID"`
	RequestTimeStamp    string `json:"requestTimeStamp"`
	Operation           string `json:"operation"`
	OfferCode           string `json:"offerCode"`
	TransactionID       string `json:"transactionID"`
	ClientTransactionID string `json:"clientTransactionID"`
	Language            string `json:"language"`
	SubscriberLifeCycle string `json:"subscriberLifeCycle"`
	SubscriptionStatus  string `json:"subscriptionStatus"`
	Channel             string `json:"channel"`
	Reason              string `json:"reason"`
	Type                string `json:"type"`
	OfferName           string `json:"offerName"`
	Mobile              string `json:"mobile"`
}

// SetValuesFromInput sets values in payload param
func (s *NewSubscription) SetValuesFromInput(data []ParamValue) {
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
func (s *NewSubscription) ToJSON() []byte {
	b, err := json.Marshal(s)
	if err != nil {
		return nil
	}

	return b
}

// NewUserAccount ...
type NewUserAccount struct {
	FirstName        string `json:"firstName" binding:"required"`
	LastName         string `json:"lastName" binding:"required"`
	Email            string `json:"email" binding:"required"`
	Mobile           string `json:"mobile" binding:"required"`
	Password         string `json:"password" binding:"required"`
	OrganizationName string `json:"organizationName" binding:"required"`
}

// Validate ...
func (a *NewUserAccount) Validate() bool {
	if a.FirstName == "" || a.LastName == "" || a.Email == "" || a.Mobile == "" || a.Password == "" || a.OrganizationName == "" {
		return false
	}

	return true
}
