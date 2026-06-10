// Package mediauploader defines tools for file storage.
package mediauploader

import (
	"net/url"

	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/internal/utils"
)

// MediaUploader defines the methods required to upload files.
type MediaUploader interface {
	UploadProfilePic(filePath string) utils.ResponseFromUpload
	UploadThumbnail(filePath string) utils.ResponseFromUpload
	DeleteByPublicID(publicID string) error
	SignParameters(params url.Values, secret string) (string, error)
}
