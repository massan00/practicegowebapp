package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, world\n")
}

func PostArticleHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Posting Article...\n")
}

func ArticleListHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Article List\n")
}

func ArticleDetailHandler(w http.ResponseWriter, r *http.Request) {
	articleId := 1
	resString := fmt.Sprintf("Artilce NO.%d\n", articleId)
	io.WriteString(w, resString)
}

func PostNiceHandler(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, "Posting Nice...\n")
}

func PostCommentHandler(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, "Posting comment...\n")
}
