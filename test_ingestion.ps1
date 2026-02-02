$headers = @{
    "Content-Type" = "application/json"
}

$uri = "http://localhost:8080/api/v1/logs"

# 1. Normal Log (Just one)
$normalLogs = @(
    "User login successful"
    # "Health check passed" - Commented out to have 'less normal logs'
)

Write-Host "`n[TEST] Sending Normal Logs..." -ForegroundColor Green
foreach ($msg in $normalLogs) {
    $body = @{
        message = $msg
        source  = "auth-service"
        level   = "INFO"
    } | ConvertTo-Json -Depth 2
    
    Invoke-RestMethod -Uri $uri -Method Post -Headers $headers -Body $body | Out-Null
    Write-Host "  -> Sent: $msg"
}

# 2. Anomalous Logs (Many)
$anomalousLogs = @(
    "DATABASE INTEGRITY ERROR: TABLE CORRUPTION DETECTED",
    "FATAL EXCEPTION: NullPointerException in module payment-gateway",
    "Security Alert: Repeated failed login attempts from IP 203.0.113.42",
    "CRITICAL: Memory usage exceeded 99%. OOM Killer invoked.",
    "SQL Injection attempt blocked: SELECT * FROM users WHERE '1'='1'",
    "Process deadlocked in transaction manager",
    "Unexpected system shutdown initiated by watchdog",
    "Buffer overflow detected in input stream processing"
)

Write-Host "`n[TEST] Sending Anomalous Logs..." -ForegroundColor Red
foreach ($msg in $anomalousLogs) {
    $body = @{
        message = $msg
        source  = "core-system"
        level   = "CRITICAL"
    } | ConvertTo-Json -Depth 2

    Invoke-RestMethod -Uri $uri -Method Post -Headers $headers -Body $body | Out-Null
    Write-Host "  -> Sent: $msg"
}

Write-Host "`nAll logs sent. Please check the Go backend console for 'ALARM: Anomaly detected!' messages." -ForegroundColor Yellow
