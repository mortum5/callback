package service

import (
	"log/slog"

	"github.com/mortum5/callback/internal/db"
)

type Worker struct {
	n        int
	idCh     chan int
	statusCh chan db.ObjectModel
	client   *Client
}

func NewWorker(n int, idCh chan int, statusCh chan db.ObjectModel, client *Client) *Worker {
	return &Worker{
		n:        n,
		idCh:     idCh,
		statusCh: statusCh,
		client:   client,
	}
}

func (w *Worker) Start() {
	for i := 0; i < w.n; i++ {
		// slog.Info("create worked with id ", "id", i)
		go func() {
			for {
				objectId := <-w.idCh
				slog.Info("worker received id ", "id", objectId)
				status := w.client.getStatus(objectId)
				slog.Info("worket get ", "id", objectId, "status", status)
				w.statusCh <- db.ObjectModel{
					Id:     objectId,
					Status: status,
				}
			}
		}()
	}
}
