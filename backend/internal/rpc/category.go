package rpc

import (
	"context"

	apiv1 "github.com/amlcx/tablero/backend/gen/api/v1"
	"github.com/amlcx/tablero/backend/internal/auth"
	"github.com/amlcx/tablero/backend/sentinel"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CategoryHandler interface {
	Create(ctx context.Context, req *apiv1.CreateRequest) (*apiv1.CreateResponse, error)
	List(ctx context.Context, req *apiv1.ListRequest) (*apiv1.ListResponse, error)
	SelectByID(ctx context.Context, req *apiv1.SelectByIDRequest) (*apiv1.SelectByIDResponse, error)
	SelectByShortName(ctx context.Context, req *apiv1.SelectByShortNameRequest) (*apiv1.SelectByShortNameResponse, error)
}

type categoryHandler struct {
	BaseHandler
}

var _ CategoryHandler = (*categoryHandler)(nil)

func NewCategoryHandler(
	bh BaseHandler,
) CategoryHandler {
	sentinel.Assert(bh != nil, "failed to initialize category handler: nil base repo")

	return &categoryHandler{
		BaseHandler: bh,
	}
}

func (h *categoryHandler) Create(ctx context.Context, req *apiv1.CreateRequest) (*apiv1.CreateResponse, error) {
	h.log().Debug("category handler create method received request")

	cid := uuid.New()
	mid := uuid.New()

	user, err := auth.UserFromContext(ctx)
	if err != nil {
		return nil, h.normalize(err)
	}

	response := &apiv1.CreateResponse{
		Category: &apiv1.Category{
			Id:        cid.String(),
			CreatedBy: user.ID.String(),
			UpdatedBy: user.ID.String(),
			Name:      "name 1",
			ShortName: "sn1",
			Blurb:     "blurb 1",
			Nsfw:      true,
			Media: &apiv1.Media{
				Id:                  mid.String(),
				Checksum:            "checksum1",
				Mime:                "mime1",
				Type:                "type1",
				Extension:           "ext1",
				Size:                1,
				Url:                 "url1",
				SquareThumbnailUrl:  "sq1",
				RegularThumbnailUrl: "rg1",
				CreatedAt:           timestamppb.Now(),
			},
			CreatedAt: timestamppb.Now(),
			UpdatedAt: timestamppb.Now(),
			DeletedAt: nil,
		},
	}

	return response, nil
}

func (h *categoryHandler) List(ctx context.Context, req *apiv1.ListRequest) (*apiv1.ListResponse, error) {
	return nil, nil
}

func (h *categoryHandler) SelectByID(ctx context.Context, req *apiv1.SelectByIDRequest) (*apiv1.SelectByIDResponse, error) {
	return nil, nil
}

func (h *categoryHandler) SelectByShortName(ctx context.Context, req *apiv1.SelectByShortNameRequest) (*apiv1.SelectByShortNameResponse, error) {
	return nil, nil
}
