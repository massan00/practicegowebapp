package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		io.WriteString(w, "Hello, world\n")
	} else {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
	}
}

func PostArticleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		io.WriteString(w, "Posting Article...\n")
	} else {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
	}
}

func ArticleListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		io.WriteString(w, "Article List\n")
	} else {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
	}
}

func ArticleDetailHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		articleId := 1
		resString := fmt.Sprintf("Artilce NO.%d\n", articleId)
		io.WriteString(w, resString)
	} else {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
	}
}

func PostNiceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		io.WriteString(w, "Posting Nice...\n")
	} else {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
	}
}

func PostCommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		io.WriteString(w, "Posting comment...\n")
	} else {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
	}
}
