package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/renoinn/kawaguchi-child-data/repository"
)

type PreschoolHandler interface {
	GetPreschool() func(http.ResponseWriter, *http.Request)
}

type preschoolHandler struct {
	repository repository.PreschoolRepository
}

func (p preschoolHandler) GetPreschool() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			slog.Info("invalid http method")
			w.WriteHeader(405)
			return
		}

		data, err := p.repository.GetData()
		if err != nil {
			slog.Error("faild load preschool data")
			// internal error
		}
		slog.Info("get preschool data")

		output, err := json.MarshalIndent(data, "", "\t\t")
		if err != nil {
			slog.Error("faild marshal json")
			// internal error
		}

		w.Header().Set("Contet-Type", "application/json")
		w.Write(output)
	}
}

func NewPreschoolHandler(pr repository.PreschoolRepository) PreschoolHandler {
	return preschoolHandler{pr}
}
