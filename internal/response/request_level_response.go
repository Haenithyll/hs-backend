package response

type RequestLevelResponse struct {
	ID    uint8  `json:"id"`
	Label string `json:"label"`
}

type RequestLevelResponses []RequestLevelResponse
