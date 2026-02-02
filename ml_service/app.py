# ml_service/app.py
import sys
sys.path.append("C:\\Users\\Hackb07\\AppData\\Roaming\\Python\\Python314\\site-packages")
from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
from model import LogAnomalyDetector
import os

app = FastAPI()
detector = LogAnomalyDetector()

class LogRequest(BaseModel):
    message: str

@app.post("/predict")
async def predict(log: LogRequest):
    if not detector.model:
        raise HTTPException(status_code=503, detail="Model not initialized")
    
    result = detector.predict(log.message)
    return result

@app.get("/health")
def health_check():
    return {"status": "ok", "model_loaded": detector.model is not None}
