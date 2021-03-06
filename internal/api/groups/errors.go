package groups

import (
	"github.com/glide-im/api/internal/api/comm"
)

var (
	ErrGroupNotExit       = comm.NewApiBizError(3001, "ErrGroupNotExit")
	ErrMemberAlreadyExist = comm.NewApiBizError(3002, "ErrMemberAlreadyExist")
)
