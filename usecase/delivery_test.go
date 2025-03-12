package usecase

import (
	"context"
	"errors"
	"io"
	"reflect"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"github.com/javiertlopez/awesome/errorcodes"
	"github.com/javiertlopez/awesome/model"
	"github.com/javiertlopez/awesome/usecase/mocks"
)

// Generate mocks
// mockery --keeptree --name=Videos --dir=usecase --output=usecase/mocks

func Test_delivery_GetByID(t *testing.T) {
	logger := logrus.New()
	logger.Out = io.Discard

	uuid := "4e5bf8f2-9c50-4576-b9d4-1d1fd0705885"
	asset := model.Asset{
		ID:        uuid,
		Poster:    "https://image.mux.com/5iNFJg9dIww2AgUryhgghbP00Dc4ogoxn00gzitOdjICg/thumbnail.png?width=1920\u0026height=1080\u0026smart_crop=true\u0026time=7",
		Thumbnail: "https://image.mux.com/5iNFJg9dIww2AgUryhgghbP00Dc4ogoxn00gzitOdjICg/thumbnail.png?width=640\u0026height=360\u0026smart_crop=true\u0026time=7",
		Sources: []model.Source{
			{
				Source: "https://stream.mux.com/5iNFJg9dIww2AgUryhgghbP00Dc4ogoxn00gzitOdjICg.m3u8",
				Type:   "application/x-mpegURL",
			},
		},
	}
	err := errors.New("failed")
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		args    args
		want    model.Video
		wantErr bool
		err     error
	}{
		{
			"Only video",
			args{
				ctx: context.Background(),
				id:  uuid,
			},
			model.Video{},
			false,
			nil,
		},
		{
			"Error",
			args{
				ctx: context.Background(),
				id:  uuid,
			},
			model.Video{},
			true,
			err,
		},
		{
			"With asset",
			args{
				ctx: context.Background(),
				id:  uuid,
			},
			model.Video{
				ID:        uuid,
				Poster:    "https://image.mux.com/5iNFJg9dIww2AgUryhgghbP00Dc4ogoxn00gzitOdjICg/thumbnail.png?width=1920\u0026height=1080\u0026smart_crop=true\u0026time=7",
				Thumbnail: "https://image.mux.com/5iNFJg9dIww2AgUryhgghbP00Dc4ogoxn00gzitOdjICg/thumbnail.png?width=640\u0026height=360\u0026smart_crop=true\u0026time=7",
				Sources: []model.Source{
					{
						Source: "https://stream.mux.com/5iNFJg9dIww2AgUryhgghbP00Dc4ogoxn00gzitOdjICg.m3u8",
						Type:   "application/x-mpegURL",
					},
				},
			},
			false,
			nil,
		},
		{
			"With invalid ID",
			args{
				ctx: context.Background(),
				id:  "invalid",
			},
			model.Video{},
			true,
			errorcodes.ErrInvalidID,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assets := &mocks.Assets{}
			videos := &mocks.Videos{}
			usecase := &delivery{
				assets,
				videos,
				logger,
			}

			if tt.wantErr {
				assets.On("GetByID", tt.args.ctx, tt.args.id).Return(model.Asset{}, errors.New("failed"))
				videos.On("GetByID", tt.args.ctx, tt.args.id).Return(model.Video{}, errors.New("failed"))
			} else {
				assets.On("GetByID", tt.args.ctx, tt.args.id).Return(asset, nil)
				videos.On("GetByID", tt.args.ctx, tt.args.id).Return(tt.want, nil)
			}

			got, err := usecase.GetByID(tt.args.ctx, tt.args.id)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.err, err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("delivery.GetByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
