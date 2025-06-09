package create_comment

import (
	"errors"
)

var ErrIncorrectUserID = errors.New("userID is incorrect")
var ErrUserServiceUnavailable = errors.New("user service unavailable")
var ErrProductOwnerNotFound = errors.New("product owner not found")
var ErrProductServiceUnavailable = errors.New("product service unavailable")

type Comment struct {
	ProductID      int64
	ProductOwnerID int64
	UserID         int64
	Text           string
}
