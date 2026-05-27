package cloudinaryUploader

import (
	"context"
	"net/url"

	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/pkg/utils"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type Cloudinary struct {
	Client *cloudinary.Cloudinary
}

func NewCloudinaryUploader(api_key, api_secret, cloud_name string) (*Cloudinary, error) {
	cld, err := cloudinary.NewFromParams(
		cloud_name,
		api_key,
		api_secret,
	)
	if err != nil {
		return nil, err
	}

	return &Cloudinary{
		Client: cld,
	}, nil
}

func (c *Cloudinary) SignParameters(params url.Values, secret string) (string, error) {
	// Cloudinary SDK's built-in parameter signing utility
	return api.SignParameters(params, secret)
}

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

func (c *Cloudinary) DeleteByPublicID(publicID string) error {
	_, err := c.Client.Upload.Destroy(
		context.Background(),
		uploader.DestroyParams{
			PublicID: publicID,
		},
	)

	return err
}
