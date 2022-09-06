package dto

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
