package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/renoinn/kawaguchi-child-data/datasource"
	"github.com/renoinn/kawaguchi-child-data/handler"
	"github.com/renoinn/kawaguchi-child-data/repository"
)

func main() {
	l := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(l)

	d := datasource.PreschoolCsv{}
	r := repository.NewPreschoolRepository(d)
	h := handler.NewPreschoolHandler(r)

	http.HandleFunc("/preschool", h.GetPreschool())

	slog.Info("start server")
	http.ListenAndServe(":8080", nil)
}
