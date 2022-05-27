package msg

import (
	"github.com/glide-im/api/internal/api/comm"
)

var (
	errRecentMsgLoadFailed = comm.NewApiBizError(3001, "message load failed")
)
