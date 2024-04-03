package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi/v5"
)

func main() {
	go standardLib()
	go frameworkGin()
	frameworkChi()
}

func standardLib() {
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello")
	})
	http.ListenAndServe(":8080", nil)
}

func frameworkGin() {
	r := gin.New()
	r.GET("/", func(ctx *gin.Context) {
		fmt.Fprintf(ctx.Writer, "Hello")
	})
	r.Run(":8081")
}

func frameworkChi() {
	c := chi.NewRouter()
	c.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})
	http.ListenAndServe(":8082", c)
}
