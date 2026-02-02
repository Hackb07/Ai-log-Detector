# ml_service/model.py
import sys
sys.path.append("C:\\Users\\Hackb07\\AppData\\Roaming\\Python\\Python314\\site-packages")
import joblib
import os

class LogAnomalyDetector:
    def __init__(self, model_path='model.joblib'):
        self.model_path = model_path
        self.model = None
        self.load_model()

    def load_model(self):
        if os.path.exists(self.model_path):
            self.model = joblib.load(self.model_path)
            print("Model loaded.")
        else:
            print("Model file not found. Please train the model first.")

    def predict(self, message: str):
        if not self.model:
            return {"error": "Model not loaded"}
        
        # Check for Out-Of-Vocabulary (OOV) words
        # If the message contains NO words from the training vocabulary, it's definitely an anomaly (novelty)
        vectorizer = self.model.named_steps['vectorizer']
        vector = vectorizer.transform([message])
        
        if vector.nnz == 0:
            return {
                "is_anomaly": True,
                "anomaly_score": -1.0,
                "prediction": -1,
                "note": "Unknown log pattern (OOV)"
            }

        # Model prediction (1: Normal, -1: Anomaly)
        prediction = self.model.predict([message])[0]
        score = self.model.decision_function([message])[0]
        
        is_anomaly = True if prediction == -1 else False
        
        return {
            "is_anomaly": is_anomaly,
            "anomaly_score": float(score),
            "prediction": int(prediction)
        }
