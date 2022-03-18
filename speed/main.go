package main

import (
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

const maxSize = 26214400

func downloadHandler() http.HandlerFunc {
	src := rand.NewSource(0)
	return func(w http.ResponseWriter, r *http.Request) {
		queries := r.URL.Query()
		size := queries.Get("size")
		max, err := strconv.Atoi(size)
		if err != nil {
			max = maxSize
		}
		read := rand.New(src)
		_, err = io.CopyN(w, read, int64(max))
		if err != nil {
			log.Printf("failed to write random data: %s", err)
			return
		}
	}
}
