Write-Host "========================================="
Write-Host " AI Log Anomaly Detection System Launcher "
Write-Host "========================================="

# Ports used by the system
$ports = @(8000, 8080)

foreach ($port in $ports) {
    Write-Host "`nChecking port $port..."

    $connections = netstat -ano | findstr ":$port"

    if ($connections) {
        $pids = $connections | ForEach-Object {
            ($_ -split "\s+")[-1]
        } | Sort-Object -Unique

        foreach ($procId in $pids) {
            Write-Host "Killing process on port $port (PID: $procId)"
            Stop-Process -Id $procId -Force -ErrorAction SilentlyContinue
        }
    }
    else {
        Write-Host "Port $port is free."
    }
}

Write-Host "`nAll required ports are now free."

# Optional: Train ML model if not present
if (-Not (Test-Path ".\ml_service\model.joblib")) {
    Write-Host "`nModel not found. Training model..."
    Set-Location ml_service
    $env:PYTHONPATH = "C:\Users\Hackb07\AppData\Roaming\Python\Python314\site-packages;" + $env:PYTHONPATH
    python train.py
    Set-Location ..
}
else {
    Write-Host "`nModel already exists. Skipping training."
}

Write-Host "`nStarting AI Log Anomaly Detection System..."
.\start_system.ps1
