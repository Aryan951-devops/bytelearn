package mediauploader

import (
	"net/url"

	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/internal/utils"
)

type MediaUploader interface {
	UploadProfilePic(filePath string) utils.ResponseFromUpload
	UploadThumbnail(filePath string) utils.ResponseFromUpload
	DeleteByPublicID(publicID string) error
	SignParameters(params url.Values, secret string) (string, error)
}
