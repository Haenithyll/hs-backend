package request

type CreateRequestRequest struct {
	Topic      string `json:"topic" binding:"required"`
	Level      uint8  `json:"level" binding:"required"`
	ReceiverID string `json:"receiverId" binding:"required"`
}

type MarkRequestAsReadRequest struct {
	RequestID uint8 `uri:"requestId" json:"-" binding:"required"`
}
