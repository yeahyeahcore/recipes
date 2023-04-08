package models

import "mime/multipart"

type SaveFile struct {
	Fileheader *multipart.FileHeader
	Path       string
}
