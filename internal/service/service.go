package service

import (
	"log/slog"
	"time"

	"github.com/mortum5/callback/internal/db"
)

type ObjectStatusService struct {
	db *db.Repository
	w  *Worker
}

func New(repo *db.Repository) *ObjectStatusService {
	idCh := make(chan int)
	statusCh := make(chan db.ObjectModel)
	client := &Client{}

	worker := NewWorker(200, idCh, statusCh, client)

	return &ObjectStatusService{
		db: repo,
		w:  worker,
	}
}

func (service *ObjectStatusService) Start() {

	service.w.Start()

	go func() {
		for {
			object := <-service.w.statusCh
			if object.Status {
				service.db.UpdateUserLastSeen(object.Id)
				slog.Info("save new object to db", "object", object)
			}
		}
	}()

	go func() {
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()

		for {
			<-ticker.C
			service.db.DeleteObjectsThatOffline()
		}
	}()
}

func (service *ObjectStatusService) HandleId(id int) {
	slog.Info("new id received in service", "id", id)
	service.w.idCh <- id
}
