package pkg

import "github.com/rs/zerolog/log"

func (w *WorkerPool) InitWorkerPool() {
	w.WorkerQueue = make(chan *interface{}, w.MaxPool)
	for p := 0; p < w.MaxPool; p++ {
		w.runWorkerPool()
		log.Info().Msgf("start worker %d\n", p)
	}
}

func (w *WorkerPool) runWorkerPool() {
	go func() {
		for {
			select {
			case msg := <-w.WorkerQueue:
				w.Handler(msg)
			}
		}
	}()
}

func (w *WorkerPool) AddPayloadIntoQueue(data *interface{}) {
	if w.WorkerQueue == nil {
		w.WorkerQueue = make(chan *interface{}, w.MaxPool)
	}
	w.WorkerQueue <- data
}
