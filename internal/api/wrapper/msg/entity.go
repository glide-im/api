package msg

type MessageQueryRequest struct {
	PageSize int   `json:"pageSize"`
	Page     int   `json:"page"`
	To       int64 `json:"to"`
	EndMid   int64 `json:"end_mid"`
	StartMid int64 `json:"start_mid"`
}

type MessageReadRequest struct {
	SessionId string `json:"session_id"`
}
