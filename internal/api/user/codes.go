package user

import (
	"github.com/glide-im/api/internal/api/comm"
)

var (
	errAddSelf         = comm.NewApiBizError(2001, "unable to add yourself as a contact")
	errUserNotExist    = comm.NewApiBizError(2002, "user does not exist")
	errAlreadyContacts = comm.NewApiBizError(2003, "already in the contact list")
)
