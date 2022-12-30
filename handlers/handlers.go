package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/massan00/practicegowebapp/models"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, world\n")
}

func PostArticleHandler(w http.ResponseWriter, r *http.Request) {
	var reqArticle models.Article

	if err := json.NewDecoder(r.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to get request body\n", http.StatusBadRequest)
		return
	}

	// article := reqArticle

	json.NewEncoder(w).Encode(reqArticle)

}

func ArticleListHandler(w http.ResponseWriter, r *http.Request) {
	queryMap := r.URL.Query()
	var page int
	if q, ok := queryMap["page"]; ok && len(q) > 0 {
		var err error
		page, err = strconv.Atoi(q[0])
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}

	articleList := []models.Article{models.Article1, models.Article2}
	json.NewEncoder(w).Encode(articleList)
	log.Println(page)
}

func ArticleDetailHandler(w http.ResponseWriter, r *http.Request) {
	articleId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	article := models.Article1
	json.NewEncoder(w).Encode(article)
	log.Println(articleId)

}

func PostNiceHandler(w http.ResponseWriter, r *http.Request) {
	var reqArticle models.Article
	decode := json.NewDecoder(r.Body)

	if err := decode.Decode(&reqArticle); err != nil {
		http.Error(w, "failed decode request body\n", http.StatusBadRequest)
		return
	}
	article := reqArticle
	json.NewEncoder(w).Encode(article)
}

func PostCommentHandler(w http.ResponseWriter, r *http.Request) {
	var reqComment models.Comment

	req := json.NewDecoder(r.Body)
	if err := req.Decode(&reqComment); err != nil {
		http.Error(w, "failed docode request body\n", http.StatusBadRequest)
		return
	}

	comment := reqComment
	json.NewEncoder(w).Encode(comment)
}
