package create_comment

import (
	"context"
	"errors"
)

type SaveCommentRepository interface {
	SaveComment(_ context.Context, comment Comment) (int64, error)
}

type UserService interface {
	CheckUserID(_ context.Context, userID int64) (bool, error)
}

type ProductsService interface {
	GetProductOwner(_ context.Context, productID int64) (int64, error)
}

type CreateCommentService struct {
	rep            SaveCommentRepository
	userService    UserService
	productService ProductsService
}

func NewCreateCommentService(rep SaveCommentRepository, products ProductsService, users UserService) *CreateCommentService {
	return &CreateCommentService{
		rep:            rep,
		productService: products,
		userService:    users,
	}
}

func (s *CreateCommentService) CreateComment(ctx context.Context, userID int64, productID int64, tx string) (int64, error) {
	isCorrectUserID, err := s.userService.CheckUserID(ctx, userID)
	if err != nil {
		return 0, errors.Join(ErrUserServiceUnavailable, err)
	}
	if !isCorrectUserID {
		return 0, ErrIncorrectUserID
	}
	productOwnerID, err := s.productService.GetProductOwner(ctx, productID)
	if err != nil {
		return 0, errors.Join(ErrProductServiceUnavailable, err)
	}
	if productOwnerID == 0 {
		return 0, ErrProductOwnerNotFound
	}
	comment := Comment{
		UserID:         userID,
		ProductID:      productID,
		ProductOwnerID: productOwnerID,
		Text:           tx,
	}
	commentID, err := s.rep.SaveComment(ctx, comment)
	return commentID, err
}
