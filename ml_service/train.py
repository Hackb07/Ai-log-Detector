# ml_service/train.py
import sys
sys.path.append("C:\\Users\\Hackb07\\AppData\\Roaming\\Python\\Python314\\site-packages")
import pandas as pd
from sklearn.svm import OneClassSVM
from sklearn.feature_extraction.text import TfidfVectorizer
from sklearn.pipeline import Pipeline
import joblib
import numpy as np

# Sample training data (normal logs) - Augmented
data = [
    "User login successful",
    "Database connection established",
    "Request processed successfully",
    "File uploaded",
    "System started",
    "Health check passed",
    "Service ready",
    "Cache cleared",
    "User logout",
    "Dashboard loaded",
    "HTTP 200 OK",
    "API request received",
    "Data synced",
    "Backup completed",
    "Job scheduled",
    "Email sent",
    "Payment processed",
    "User profile updated",
    "Session validated",
    "Token refreshed"
] * 10 # Repeat to have more data points

def train_model():
    print("Training model...")
    # High gamma forces tight clusters around training data
    model = Pipeline([
        ('vectorizer', TfidfVectorizer(max_features=500, stop_words=None)),
        ('classifier', OneClassSVM(nu=0.05, kernel='rbf', gamma=10))
    ])
    
    model.fit(data)
    
    # Save model
    joblib.dump(model, 'model.joblib')
    print("Model saved to model.joblib")

if __name__ == "__main__":
    train_model()
