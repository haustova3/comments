package create_comment

import (
	"context"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateCommentService_CreateComment(t *testing.T) {
	type testCase struct {
		name      string
		ctx       context.Context
		userID    int64
		productID int64
		text      string
		mockSetup func(ctx context.Context, rep *SaveCommentRepositoryMock, users *UserServiceMock, products *ProductsServiceMock)
		expErr    error
	}

	testCases := []testCase{
		{
			name:      "create_comment_success",
			ctx:       context.Background(),
			userID:    123,
			productID: 456,
			text:      "super",
			mockSetup: func(ctx context.Context, rep *SaveCommentRepositoryMock, users *UserServiceMock, products *ProductsServiceMock) {
				rep.SaveCommentMock.Expect(ctx, Comment{
					UserID:         123,
					ProductID:      456,
					ProductOwnerID: 789,
					Text:           "super",
				}).Return(1, nil)
				users.CheckUserIDMock.Expect(ctx, 123).Return(true, nil)
				products.GetProductOwnerMock.Expect(ctx, 456).Return(789, nil)
			},
			expErr: nil,
		},
		{
			name:      "create_comment_wrong_user_id",
			ctx:       context.Background(),
			userID:    123,
			productID: 456,
			text:      "super",
			mockSetup: func(ctx context.Context, rep *SaveCommentRepositoryMock, users *UserServiceMock, products *ProductsServiceMock) {
				users.CheckUserIDMock.Expect(ctx, 123).Return(false, nil)
			},
			expErr: ErrIncorrectUserID,
		},
		{
			name:      "create_comment_product_owner_not_found",
			ctx:       context.Background(),
			userID:    123,
			productID: 456,
			text:      "super",
			mockSetup: func(ctx context.Context, rep *SaveCommentRepositoryMock, users *UserServiceMock, products *ProductsServiceMock) {
				users.CheckUserIDMock.Expect(ctx, 123).Return(true, nil)
				products.GetProductOwnerMock.Expect(ctx, 456).Return(0, nil)
			},
			expErr: ErrProductOwnerNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			controller := minimock.NewController(t)
			mockRep := NewSaveCommentRepositoryMock(controller)
			mockUsers := NewUserServiceMock(controller)
			mockProducts := NewProductsServiceMock(controller)
			tc.mockSetup(tc.ctx, mockRep, mockUsers, mockProducts)
			service := NewCreateCommentService(mockRep, mockProducts, mockUsers)
			id, err := service.CreateComment(tc.ctx, tc.userID, tc.productID, tc.text)
			if tc.expErr != nil {
				assert.ErrorIs(t, err, tc.expErr)
			} else {
				assert.NoError(t, err)
				if id <= 0 {
					t.Errorf("Wrong id value: %d", id)
				}
			}
		})
	}
}
