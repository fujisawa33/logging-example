package main

import (
	"net/http"

	"go.uber.org/zap"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("ほげほげ",
		zap.String("ふがふが", "fugafuga"),
		zap.String("ぴよぴよ", "piyopiyo"),
	)
}
