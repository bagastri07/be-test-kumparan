package models

type MessageResponse struct {
	Message string `json:"message"`
}

type DataResponse struct {
	Data interface{} `json:"data"`
}

type ItemsResponse struct {
	Data struct {
		Items interface{} `json:"items,omitempty"`
	} `json:"data"`
}
