package usecase

import (
	"context"

	"github.com/javiertlopez/awesome/errorcodes"
	"github.com/javiertlopez/awesome/model"
	"github.com/javiertlopez/awesome/repository"
)

// Ingestion usecase
type Ingestion interface {
	Create(ctx context.Context, anyVideo model.Video) (model.Video, error)
}

type ingestion struct {
	assets repository.AssetRepo
	videos repository.VideoRepo
}

// NewIngestion returns the usecase implementation
func NewIngestion(
	a repository.AssetRepo,
	v repository.VideoRepo,
) Ingestion {
	return &ingestion{
		assets: a,
		videos: v,
	}
}

// Create method
func (u *ingestion) Create(ctx context.Context, anyVideo model.Video) (model.Video, error) {
	// Title and Description are mandatory fields
	if len(anyVideo.Title) == 0 || len(anyVideo.Description) == 0 {
		return model.Video{}, errorcodes.ErrVideoUnprocessable
	}

	// If body contains a Source File URL, send it to Ingestion
	if len(anyVideo.SourceURL) > 0 {
		var isPublic bool
		switch anyVideo.Policy {
		case "public":
			isPublic = true
		case "signed":
			isPublic = false
		default:
			return model.Video{}, errorcodes.ErrIngestionFailed
		}

		assetID, err := u.assets.Create(ctx, anyVideo.SourceURL, isPublic)
		if err == nil {
			anyVideo.Asset = &model.Asset{
				ID: assetID,
			}
		}
	}

	response, err := u.videos.Create(ctx, anyVideo)

	if err != nil {
		return model.Video{}, err
	}

	return response, nil
}
