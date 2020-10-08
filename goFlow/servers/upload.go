package servers

import (
	"context"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"goFlow/utils"
	"goFlow/utils/errmsg"
	"mime/multipart"
)

var (
	accessKey = utils.AccessKey
	secretKey = utils.SecretKey
	bucket    = utils.Bucket
	imgUrl    = utils.Qiniuserver
)

func UploadFile(file multipart.File, fileSize int64) (string, int) {
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan, //华南地区
		UseHTTPS:      false,
		UseCdnDomains: false,
	}

	putExtra := storage.PutExtra{}

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	if err != nil {
		return "", errmsg.ERROR
	}
	url := imgUrl + ret.Key
	return url, errmsg.SUCCESS
}
