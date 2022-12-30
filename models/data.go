package models

import "time"

var (
	Comment1 = Comment{
		CommentID: 1,
		ArticleID: 1,
		Message:   "test Comment1",
		CreatedAt: time.Now(),
	}
	Comment2 = Comment{
		CommentID: 2,
		ArticleID: 1,
		Message:   "test Comment2",
		CreatedAt: time.Now(),
	}
	Article1 = Article{
		ID:          1,
		Title:       "First Article",
		Contents:    "This is the test article",
		UserName:    "saki",
		NiceNum:     1,
		CommentList: []Comment{Comment1, Comment2},
		CreatedAt:   time.Now(),
	}

	Article2 = Article{
		ID:        2,
		Title:     "Second Article",
		Contents:  "This is the test article.",
		UserName:  "saki",
		NiceNum:   2,
		CreatedAt: time.Now(),
	}
)
