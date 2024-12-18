package services

import (
	"genggam-makna-api/dto"
	"genggam-makna-api/repositories"
)

type CompService interface{
	RegisterUserCredential(data dto.User) (*string, error)
	LoginUserCredentials(email string, password string) (*string, error)
	LoginUserGoogle(data dto.User) (*string, error)

	ImagePredict(image_data []byte) (*dto.MLResponse, error)
	VideoPredict(video_data []byte) (*dto.MLResponse, error)
}

type compServices struct {
	repo repositories.CompRepository
}

func NewService(r repositories.CompRepository) *compServices {
	return &compServices{
		repo: r,
	}
}
