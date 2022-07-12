package main

type Service struct {
	CommentRepo CommentRepo
}

type CommentRepo interface {
	GetComments() ([]Comment, error)
}

type Comment struct {
	Author string
	Body   string
	Slug   string
}

func NewService() {

}
