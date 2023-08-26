package main

import (
	"github.com/ujunglangit-id/makaryo/pkg"
	_ "go.uber.org/automaxprocs"
)

type OrganizationData struct {
	Index            string
	OrganizationID   string
	Name             string
	Website          string
	Country          string
	Description      string
	Founded          string
	Industry         string
	NumberOfEmployee string
}

func main() {
	inputFileName := "files/organizations-2000000.csv"
	maxPool := 10
	pool := pkg.WorkerPool{
		MaxPool: maxPool,
		Handler: func(payload interface{}) {
			ConstructSummary(payload.(OrganizationData))
		},
		WorkerQueue: make(chan *interface{}, maxPool),
	}
	pool.InitWorkerPool()

	pool.AddPayloadIntoQueue()
}

func ConstructSummary(data OrganizationData) {

}
