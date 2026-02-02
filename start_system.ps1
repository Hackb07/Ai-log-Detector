# Start ML Service
$env:PYTHONPATH = "C:\Users\Hackb07\AppData\Roaming\Python\Python314\site-packages;" + $env:PYTHONPATH
Start-Process -FilePath "python" -ArgumentList "-m uvicorn app:app --host 0.0.0.0 --port 8000" -WorkingDirectory "ml_service" -NoNewWindow

# Wait for ML service to start
Start-Sleep -Seconds 5

# Start Go Backend
go run cmd/server/main.go
