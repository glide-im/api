package cs

import (
	comm2 "github.com/glide-im/api/internal/api/comm"
	route "github.com/glide-im/api/internal/api/router"
	"github.com/glide-im/api/internal/config"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
)

var (
	ErrFailToken = comm2.NewApiBizError(2001, "qiniu not configured")
)

type ToolApi struct {
}

func (q *ToolApi) GetQiniuToken(ctx *route.Context) error {
	uploadInfo := make(map[string]string)

	bucket := config.Qiniu.QINIU_BUKET_PATH
	accessKey := config.Qiniu.QINIU_AK
	secretKey := config.Qiniu.QINIU_SK
	if len(bucket) == 0 || len(accessKey) == 0 || len(secretKey) == 0 {
		return ErrFailToken
	}

	putPolicy := storage.PutPolicy{
		Scope:   bucket,
		Expires: 60,
		SaveKey: "$(fname)",
	}
	putPolicy.Expires = 60
	mac := qbox.NewMac(accessKey, secretKey)
	token := putPolicy.UploadToken(mac)
	uploadInfo = map[string]string{
		"token":      token,
		"bucket":     bucket,
		"url":        config.Qiniu.QINIU_UPLOAD_URL,
		"upload_dir": config.Qiniu.QINIU_UPLOAD_DIR,
		"host":       config.Qiniu.QINIU_HOST,
	}

	ctx.ReturnSuccess(uploadInfo)
	return nil
}
