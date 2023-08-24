package pkg

type WorkerPool struct {
	WorkerQueue chan *interface{}
}
