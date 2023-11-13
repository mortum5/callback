package server

import (
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
	_ "net/http/pprof"

	"github.com/mortum5/callback/internal/service"
)

type RequestObjectIds struct {
	Ids []int `json:"object_ids"`
}

type Server struct {
	service *service.ObjectStatusService
}

func New(service *service.ObjectStatusService) *Server {
	return &Server{
		service: service,
	}
}

func (s *Server) Start() {
	http.HandleFunc("/callback", s.handlerCallback())
	log.Fatal(http.ListenAndServe(":9090", nil))
}

func (s *Server) handlerCallback() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		var requestObjectIds RequestObjectIds

		err := json.NewDecoder(req.Body).Decode(&requestObjectIds)
		if err != nil {
			slog.Error("callback decode error", err)
		}
		defer req.Body.Close()

		slog.Info("recieve", "count", len(requestObjectIds.Ids))

		for _, id := range requestObjectIds.Ids {
			s.service.HandleId(id)
		}

		w.Write([]byte("ok"))
	}

}
