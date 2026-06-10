// Package cloudinaryuploader provides file uploading tools using Cloudinary.
package cloudinaryuploader

import (
	"context"
	"net/url"

	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/internal/utils"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

// Cloudinary handles uploading files to Cloudinary storage.
type Cloudinary struct {
	Client *cloudinary.Cloudinary
}

// NewCloudinaryUploader creates a new uploader instance.
func NewCloudinaryUploader(apiKey, apiSecret, cloudName string) (*Cloudinary, error) {
	cld, err := cloudinary.NewFromParams(
		cloudName,
		apiKey,
		apiSecret,
	)
	if err != nil {
		return nil, err
	}

	return &Cloudinary{
		Client: cld,
	}, nil
}

// SignParameters signs the request parameters for security.
func (c *Cloudinary) SignParameters(params url.Values, secret string) (string, error) {
	// Cloudinary SDK's built-in parameter signing utility
	return api.SignParameters(params, secret)
}

// UploadProfilePic uploads a user profile image.
func (c *Cloudinary) UploadProfilePic(filePath string) utils.ResponseFromUpload {

	resp, err := c.Client.Upload.Upload(
		context.Background(),
		filePath,
		uploader.UploadParams{
			Folder: "bytelearn/profile_pics",
		},
	)

	if err != nil {
		return utils.ResponseFromUpload{
			PublicURL: "",
			PublicID:  "",
			Err:       err,
		}
	}

	return utils.ResponseFromUpload{
		PublicURL: resp.SecureURL,
		PublicID:  resp.PublicID,
		Err:       nil,
	}
}

// UploadThumbnail uploads a video thumbnail image.
func (c *Cloudinary) UploadThumbnail(filepath string) utils.ResponseFromUpload {
	resp, err := c.Client.Upload.Upload(
		context.Background(),
		filepath,
		uploader.UploadParams{
			Folder: "bytelearn/thumbnails",
		},
	)

	if err != nil {
		return utils.ResponseFromUpload{
			PublicURL: "",
			PublicID:  "",
			Err:       err,
		}
	}

	return utils.ResponseFromUpload{
		PublicURL: resp.SecureURL,
		PublicID:  resp.PublicID,
		Err:       nil,
	}
}

// DeleteByPublicID deletes a file using its ID.
func (c *Cloudinary) DeleteByPublicID(publicID string) error {
	_, err := c.Client.Upload.Destroy(
		context.Background(),
		uploader.DestroyParams{
			PublicID: publicID,
		},
	)

	return err
}
