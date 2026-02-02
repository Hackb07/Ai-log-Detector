package storage

import (
	"encoding/json"
	"os"
	"sync"
	"time"
)

// LogEntry represents the structure of a log message
type LogEntry struct {
	Timestamp    time.Time `json:"timestamp"`
	Level        string    `json:"level"`
	Message      string    `json:"message"`
	Source       string    `json:"source"`
	IsAnomaly    bool      `json:"is_anomaly"`
	AnomalyScore float64   `json:"anomaly_score"`
}

// Store interface for log storage directly
type Store interface {
	SaveLog(entry LogEntry) error
}

// FileStore implements Store using JSON files
type FileStore struct {
	mu          sync.Mutex
	NormalPath  string
	AnomalyPath string
}

func NewFileStore(normalPath, anomalyPath string) *FileStore {
	return &FileStore{
		NormalPath:  normalPath,
		AnomalyPath: anomalyPath,
	}
}

func (fs *FileStore) SaveLog(entry LogEntry) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	path := fs.NormalPath
	if entry.IsAnomaly {
		path = fs.AnomalyPath
	}

	// Read existing logs (inefficient for large files, but okay for demo)
	// In production, use append-only file or DB
	var logs []LogEntry

	fileBytes, err := os.ReadFile(path)
	if err == nil {
		_ = json.Unmarshal(fileBytes, &logs)
	}

	logs = append(logs, entry)

	data, err := json.MarshalIndent(logs, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
