package processor

import (
	"ai-log-detector/internal/mlclient"
	"ai-log-detector/internal/storage"
	"fmt"
	"log"
	"sync"
)

type LogProcessor struct {
	IngestChan chan storage.LogEntry
	MLClient   *mlclient.MLClient
	Store      storage.Store
	Workers    int
	wg         sync.WaitGroup
}

func NewLogProcessor(workers int, mlClient *mlclient.MLClient, store storage.Store) *LogProcessor {
	return &LogProcessor{
		IngestChan: make(chan storage.LogEntry, 1000), // Buffered channel
		MLClient:   mlClient,
		Store:      store,
		Workers:    workers,
	}
}

func (p *LogProcessor) Start() {
	for i := 0; i < p.Workers; i++ {
		p.wg.Add(1)
		go p.worker(i)
	}
	log.Printf("Started %d log processor workers", p.Workers)
}

func (p *LogProcessor) Stop() {
	close(p.IngestChan)
	p.wg.Wait()
	log.Println("All workers stopped")
}

func (p *LogProcessor) worker(id int) {
	defer p.wg.Done()
	for entry := range p.IngestChan {
		// Call ML Service for anomaly detection
		prediction, err := p.MLClient.Predict(entry.Message)
		if err != nil {
			log.Printf("Worker %d: Error predicting log: %v", id, err)
			continue
		}

		// Update entry with ML results
		entry.IsAnomaly = prediction.IsAnomaly
		entry.AnomalyScore = prediction.AnomalyScore

		// Store result
		if err := p.Store.SaveLog(entry); err != nil {
			log.Printf("Worker %d: Error saving log: %v", id, err)
		}

		if entry.IsAnomaly {
			fmt.Printf("ALARM: Anomaly detected! Score: %.4f | Message: %s\n", entry.AnomalyScore, entry.Message)
		} else {
			fmt.Printf("NORMAL: Score: %.4f | Message: %s\n", entry.AnomalyScore, entry.Message)
		}
	}
}
