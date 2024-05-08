package server

import (
	"net/http"

	imgai "githup.com/Therocking/go_gemini/internal/api/http/imgAi"
)

func Server() {
	http.HandleFunc("/img", imgai.ImgAi)

	http.ListenAndServe(":8080", nil)
}
