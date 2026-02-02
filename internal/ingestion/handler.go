package ingestion

import (
	"ai-log-detector/internal/processor"
	"ai-log-detector/internal/storage"
	"encoding/json"
	"net/http"
	"time"
)

type LogHandler struct {
	Processor *processor.LogProcessor
}

func NewLogHandler(proc *processor.LogProcessor) *LogHandler {
	return &LogHandler{Processor: proc}
}

type IngestRequest struct {
	Message string `json:"message"`
	Source  string `json:"source"`
	Level   string `json:"level"`
}

func (h *LogHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req IngestRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	entry := storage.LogEntry{
		Timestamp: time.Now(),
		Message:   req.Message,
		Source:    req.Source,
		Level:     req.Level,
	}

	// Non-blocking send to channel (unless buffer full)
	select {
	case h.Processor.IngestChan <- entry:
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(map[string]string{"status": "queued"})
	default:
		http.Error(w, "Service overloaded", http.StatusServiceUnavailable)
	}
}
