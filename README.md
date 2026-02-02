# AI-Based Log Anomaly Detection System

A scalable, real-time log anomaly detection system that identifies abnormal patterns in system logs using machine learning. The system utilizes a **Golang** backend for high-performance log ingestion and processing, and a **Python** microservice for ML-based anomaly detection (Isolation Forest).

## 🚀 Features

*   **Real-time Log Ingestion**: REST API to accept logs from various sources.
*   **Concurrent Processing**: Golang worker pools for efficient log handling.
*   **ML-Based Detection**: Uses Isolation Forest (unsupervised learning) to detect anomalies in log messages without requiring labeled data.
*   **Anomaly Scoring**: Assigns an anomaly score to each log entry.
*   **Alerting**: Automatically flags high-risk logs.
*   **Storage**: Separates normal logs from anomalous logs for easier auditing.

## 🏗️ System Architecture

```
graph TD
    subgraph "Real Life / Production Environment"
        Users[User/Apps] --> |HTTP POST Logs| LB[Load Balancer]
        LB --> |Distribute| GoService[Go Backend Service]
    end

    subgraph "Go Backend Service"
        GoService --> |接收 (Ingest)| Buffer[Buffered Channel]
        Buffer --> |Consume| WorkerPool[Go Worker Pool]
        WorkerPool --> |Orchestrate| MLClient[ML Client]
    end

    subgraph "AI/ML Engine (Google Nano / Banana Dev*)"
        MLClient --> |HTTP Request| PyService[Python ML Service]
        PyService --> |Vectorize & Predict| IsoForest[Isolation Forest Model]
        IsoForest --> |Result| PyService
    end

    subgraph "Storage & Alerting"
        WorkerPool --> |Check Result| Decision{Is Anomaly?}
        Decision --> |Yes| Alert[🚨 GENERATE ALERT]
        Decision --> |Yes| StoreAnomaly[Save to anomalous_logs.json]
        Decision --> |No| StoreNormal[Save to normal_logs.json]
    end

    style Users fill:#f9f,stroke:#333
    style GoService fill:#bbf,stroke:#333
    style PyService fill:#dfd,stroke:#333
    style Alert fill:#f00,stroke:#333,color:#fff
```

*> Note: In a real-world scenario, the Python ML Service could be hosted on serverless GPU platforms (like Banana.dev) or use on-device models (like Google Gemini Nano) for edge processing.*

1.  **Log Source**: Applications send logs via HTTP POST.
2.  **Ingestion Service (Go)**: Receives logs and pushes them to a buffered channel.
3.  **Processor (Go)**: Worker pool consumes logs, acts as an orchestration layer.
4.  **Inference Engine (Python)**:
    *   Receives text-based log messages.
    *   Vectorizes text (TF-IDF).
    *   Predicts anomalies.
5.  **Storage Layer**: Saves results to `logs/normal_logs.json` and `logs/anomalous_logs.json`.

## 🛠️ Prerequisites

*   **Go**: Version 1.20+
*   **Python**: Version 3.8+
*   **PowerShell** (for running start scripts on Windows)

## 📦 Installation & Setup

1.  **Clone the repository** (if applicable) or navigate to the project directory.

2.  **Python Environment Setup**:
    It is recommended to use the provided system site-packages or a virtual environment. The project depends on:
    *   `fastapi`
    *   `uvicorn`
    *   `scikit-learn`
    *   `pandas`
    *   `numpy`

    Install dependencies:
    ```bash
    pip install -r ml_service/requirements.txt
    ```

3.  **Train the Initial Model**:
    Before running the system, train the anomaly detection model on sample data:
    ```bash
    python ml_service/train.py
    ```
    This will create a `model.joblib` file in the project root.

## 🚦 Running the System

We have provided a robust PowerShell script to start everything automatically:

```powershell
./fix_and_start.ps1
```

found any error: 
```
powershell -ExecutionPolicy Bypass -File fix_and_start.ps1
```
This script will:
1.  **Kill** any hanging processes on ports 8000 (ML Service) and 8080 (Go Server).
2.  **Train** the model if it doesn't need exist.
3.  **Start** the Python ML Service and Go Backend Server.

Alternatively, you can use `./start_system.ps1` if you are sure ports are free.

**Stopping the System:**
Press `Ctrl+C` in the terminal where the Go server is running. You may need to manually stop the uvicorn process if it persists.

## 🧪 Testing & Verification

Use the provided test script to send sample logs:

```powershell
./test_ingestion.ps1
```

This sends:
1.  A "Normal" log entry ("User login successful").
2.  An "Anomalous" log entry ("DATABASE INTEGRITY ERROR...").

Check the output in the console and the files in the `logs/` directory.

## 📂 Project Structure

```
/
├── cmd/server/         # Go Main entry point
├── internal/
│   ├── ingestion/      # HTTP Handlers
│   ├── processor/      # Logic/Goroutines
│   ├── mlclient/       # Client for Python ML Service
│   └── storage/        # Storage Interface
├── ml_service/         # Python ML Service
│   ├── app.py          # FastAPI App
│   ├── model.py        # Model Logic
│   ├── train.py        # Training Script
│   └── requirements.txt
├── logs/               # Output directory for processed logs
├── scripts/            # Helper scripts
├── start_system.ps1    # Startup script
└── test_ingestion.ps1  # Verification script
```

## 📊 API Documentation

### Ingestion API (Go)
**POST** `/api/v1/logs`

request body:
```json
{
  "message": "User login failed",
  "source": "auth-service",
  "level": "WARN"
}
```

### Inference API (Python)
**POST** `/predict`

request body:
```json
{
  "message": "User login failed"
}
```

response:
```json
{
  "is_anomaly": true,
  "anomaly_score": -0.123,
  "prediction": -1
}
```
=======
# Ai-log-Detector
