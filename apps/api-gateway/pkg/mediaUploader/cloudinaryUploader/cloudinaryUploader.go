package cloudinaryUploader

import (
	"context"

	"github.com/cloudinary/cloudinary-go/v2"
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

func (c *Cloudinary) Upload(filePath string) (string, error) {

	resp, err := c.Client.Upload.Upload(
		context.Background(),
		filePath,
		uploader.UploadParams{
			Folder: "bytelearn/profile_pics",
		},
	)

	if err != nil {
		return "", err
	}

	return resp.SecureURL, nil
}
