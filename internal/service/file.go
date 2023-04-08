package service

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/url"
	"os"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"penahub.gitlab.yandexcloud.net/pena-services/accruals-service/internal/models"
)

type FileService struct {
	logger   *logrus.Logger
	distPath string
}

func NewFileService(logger *logrus.Logger, distPath string) *FileService {
	return &FileService{
		logger:   logger,
		distPath: distPath,
	}
}

func (receiver *FileService) FormFilePath(filename string) (string, error) {
	path, err := url.JoinPath(receiver.distPath, fmt.Sprintf("%s_%s", uuid.New().String(), filename))
	if err != nil {
		return "", err
	}

	return path, nil
}

func (receiver *FileService) SaveFile(ctx context.Context, request *models.SaveFile) error {
	if request == nil {
		return ErrEmptyArgs
	}

	select {
	case <-ctx.Done():
		return errors.New("context closed on <SaveFile>")
	default:
		openedFile, err := request.Fileheader.Open()
		if err != nil {
			return err
		}

		defer openedFile.Close()

		createdFile, err := os.Create(request.Path)
		if err != nil {
			return err
		}

		defer createdFile.Close()

		if _, err = io.Copy(createdFile, openedFile); err != nil {
			return err
		}

		return nil
	}
}
