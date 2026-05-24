package mediauploader

type MediaUploader interface {
	Upload(tempPath string) (string, error)
}
