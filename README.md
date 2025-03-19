# Mood API - Text Sentiment Analysis Service

Mood API is a service for analyzing the emotional tone of text. The service uses machine learning to determine text sentiment and returns an assessment with confidence level.

## Quick Start

```bash
# Clone repository
git clone https://github.com/bkzntsv/mood-api.git
cd mood-api

# Run with Docker
docker compose up --build
```

The service will be available at `http://localhost:8080`

## Live Demo

The service is currently deployed at:
```
http://185.212.148.199:8080
```

### Example Request

```bash
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{"text":"I am very happy today!"}' \
  http://185.212.148.199:8080/api/v1/analyze
```

Expected response:
```json
{
    "sentiment": "Very Positive",
    "score": 0.8534
}
```

## API Endpoints

### 1. Text Sentiment Analysis

**Endpoint:** `POST /api/v1/analyze`

**Request:**
```json
{
    "text": "Your text to analyze"
}
```

**Response:**
```json
{
    "sentiment": "Very Positive",
    "score": 0.8534
}
```

### 2. Health Check

**Endpoint:** `GET /api/v1/health`

**Response:**
```json
{
    "status": "OK"
}
```

## Sentiment Categories

- `Very Positive` - very positive (score > 0.8)
- `Positive` - positive (score 0.6-0.8)
- `Neutral` - neutral (score 0.4-0.6)
- `Negative` - negative (score 0.2-0.4)
- `Very Negative` - very negative (score < 0.2)

## Usage Examples

### Using curl

```bash
# Health check
curl http://localhost:8080/api/v1/health

# Text analysis
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{"text":"I am very happy today!"}' \
  http://localhost:8080/api/v1/analyze
```

### Using Postman

1. Create new request
2. Set method (GET/POST)
3. Enter URL: `http://localhost:8080/api/v1/analyze`
4. For POST requests:
   - Set header: `Content-Type: application/json`
   - Set body (raw JSON):
   ```json
   {
       "text": "Your text to analyze"
   }
   ```

## Technical Stack

- Go 1.21
- Gin (web framework)
- Docker & Docker Compose
- Python ML service with transformers

## Development Setup

1. Install Go 1.21 or higher
2. Install dependencies:
```bash
go mod download
```

3. Create `.env` file:
```bash
cp app.env .env
```

4. Run locally:
```bash
go run .
```

## Project Structure

```
.
├── main.go          # Application entry point and router setup
├── service.go       # Business logic and ML service integration
├── handler.go       # HTTP request handlers
├── model.go         # Data structures and request/response models
├── config.go        # Configuration management
├── app.env          # Environment variables template
├── docker-compose.yaml  # Docker services configuration
└── ml_mood/         # ML model service
    ├── model_server.py  # FastAPI server for ML model
    ├── requirements.txt # Python dependencies
    └── Dockerfile      # ML service container configuration
```

### Key Files Description

- `main.go`: Sets up the HTTP server, routes, and middleware
- `service.go`: Contains the core business logic and integration with the ML service
- `handler.go`: Implements HTTP request handlers and response formatting
- `model.go`: Defines data structures for requests, responses, and internal models
- `config.go`: Manages application configuration and environment variables
- `docker-compose.yaml`: Orchestrates both Go and Python services
- `ml_mood/model_server.py`: Implements the ML model inference service
- `ml_mood/requirements.txt`: Lists Python dependencies for the ML service

## Contributing

1. Fork the repository
2. Create feature branch: `git checkout -b feature/name`
3. Commit changes: `git commit -m "Description"`
4. Push to branch: `git push origin feature/name`
5. Create Pull Request

## License

MIT

## Author

[bkzntsv](https://github.com/bkzntsv)

## Support

Create an issue in the repository or contact via GitHub. 