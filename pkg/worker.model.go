package pkg

type WorkerPool struct {
	WorkerQueue chan *interface{}
	Handler     func(payload interface{})
	MaxPool     int
}
