package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/renoinn/kawaguchi-child-data/datasource"
	"github.com/renoinn/kawaguchi-child-data/handler"
	"github.com/renoinn/kawaguchi-child-data/repository"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

var tracer = otel.Tracer("github.com/renoinn/kawaguchi-child-data")

func main() {
	l := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(l)

	_, span := tracer.Start(context.Background(), "Server start", trace.WithAttributes(
		attribute.String("foo", "foo"),
		attribute.String("hoge", "hoge"),
	))
	defer span.End()

	d := datasource.PreschoolCsv{}
	r := repository.NewPreschoolRepository(d)
	h := handler.NewPreschoolHandler(r)

	http.HandleFunc("/preschool", h.GetPreschool())

	slog.Info("start server")
	http.ListenAndServe(":8080", nil)
}
