package get_comment

import (
	"context"
)

type GetCommentsRepository interface {
	GetComments(_ context.Context, productID int64) ([]Comment, error)
}

type GetCommentsService struct {
	rep GetCommentsRepository
}

func NewGetCommentsService(rep GetCommentsRepository) *GetCommentsService {
	return &GetCommentsService{
		rep: rep,
	}
}

func (service *GetCommentsService) GetComments(ctx context.Context, productID int64) ([]Comment, error) {
	commentID, err := service.rep.GetComments(ctx, productID)
	return commentID, err
}
