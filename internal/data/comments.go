package data

import (
	"time"

	"github.com/tchenbz/comments/internal/validator"
)

type Comment struct {
	ID int64				`json:"id"`
	Content string			`json:"content"`
	Author string			`json:"author"`
	CreatedAt time.Time		`json:"-"`
	Version int32			`json:"version"`
}

func ValidateComment(v *validator.Validator, comment *Comment) {
// check if the Content field is empty
    v.Check(comment.Content != "", "content", "must be provided")
// check if the Author field is empty
    v.Check(comment.Author != "", "author", "must be provided")
// check if the Content field is empty
    v.Check(len(comment.Content) <= 100, "content", "must not be more than 100 bytes long")
// check if the Author field is empty
     v.Check(len(comment.Author) <= 25, "author", "must not be more than 25 bytes long")

}