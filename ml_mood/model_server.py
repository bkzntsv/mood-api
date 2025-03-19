from datetime import datetime
from fastapi import FastAPI, Request, HTTPException
from pydantic import BaseModel, Field
from transformers import pipeline
import uvicorn
import logging

# Configure logging
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

app = FastAPI(
    title="Sentiment Analysis API",
    description="API for multilingual sentiment analysis",
    version="1.0.0"
)

# Initialize model with error handling
try:
    classifier = pipeline("text-classification", model="tabularisai/multilingual-sentiment-analysis")
except Exception as e:
    logger.error(f"Failed to load model: {str(e)}")
    raise

class SentimentRequest(BaseModel):
    text: str = Field(..., min_length=1, description="Text to analyze")

class SentimentResponse(BaseModel):
    sentiment: str
    score: float

@app.post("/predict", response_model=SentimentResponse)
async def predict_sentiment(request: SentimentRequest):
    try:
        result = classifier(request.text)[0]
        return SentimentResponse(
            sentiment=result["label"],
            score=float(result["score"])
        )
    except Exception as e:
        logger.error(f"Prediction error: {str(e)}")
        raise HTTPException(status_code=500, detail="Failed to analyze sentiment")

@app.get("/health")
async def health_check():
    return {"status": "healthy"}

if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8000)