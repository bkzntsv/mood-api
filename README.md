# Mood API - Text Sentiment Analysis Service

Mood API is a service for analyzing the emotional tone of text. The service uses machine learning to determine text sentiment and returns an assessment with confidence level.

## Endpoints

The service provides the following endpoints:

### 1. Text Sentiment Analysis

**Endpoint:** `POST /api/v1/analyze`

**Content-Type:** `application/json`

**Request body:**
```json
{
    "text": "Your text to analyze"
}
```

**Example response:**
```json
{
    "sentiment": "Very Positive",
    "score": 0.8534
}
```

Where:
- `sentiment` - text sentiment assessment:
  - `Very Positive` - very positive
  - `Positive` - positive
  - `Neutral` - neutral
  - `Negative` - negative
  - `Very Negative` - very negative
- `score` - model confidence in the assessment (from 0 to 1):
  - Values closer to 1.0 indicate high confidence
  - Values around 0.5 indicate medium confidence
  - Values closer to 0 indicate low confidence

### 2. Health Check

**Endpoint:** `GET /api/v1/health`

**Example response:**
```json
{
    "status": "OK"
}
```

## Usage Examples

### Using curl

1. Health check:
```bash
curl http://localhost:8080/api/v1/health
```

2. Text analysis:
```bash
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{"text":"I am very happy today!"}' \
  http://localhost:8080/api/v1/analyze
```

### Using Postman

1. Health Check:
   - Method: `GET`
   - URL: `http://localhost:8080/api/v1/health`
   - No additional headers or request body required

2. Text Analysis:
   - Method: `POST`
   - URL: `http://localhost:8080/api/v1/analyze`
   - Headers:
     - Key: `Content-Type`
     - Value: `application/json`
   - Body:
     - Select type `raw`
     - Select format `JSON`
     - Enter:
     ```json
     {
         "text": "Your text to analyze"
     }
     ```

### Remote Access via Postman

1. Health Check:
   - Method: `GET`
   - URL: `http://<server-ip>:8080/api/v1/health`
   - No additional headers or request body required

2. Text Analysis:
   - Method: `POST`
   - URL: `http://<server-ip>:8080/api/v1/analyze`
   - Headers:
     - Key: `Content-Type`
     - Value: `application/json`
   - Body:
     - Select type `raw`
     - Select format `JSON`
     - Enter:
     ```json
     {
         "text": "Your text to analyze"
     }
     ```

Replace `<server-ip>` with your server's IP address.

Example for current server:
```
http://185.212.148.199:8080/api/v1/analyze
```

#### Setting up Postman for Remote Access:

1. Open Postman
2. Create a new collection (New Collection)
3. Add a new request (New Request)
4. Enter the full URL with server IP
5. Select the appropriate method (GET or POST)
6. For POST requests:
   - Go to the "Body" tab
   - Select "raw"
   - Select "JSON" from the dropdown on the right
   - Enter the request body in JSON format

#### Tips for Working with Remote API:

1. Ensure port 8080 is open on the server
2. Check server availability via ping
3. First check the /health endpoint to confirm service availability
4. If errors occur, check:
   - Correct IP address
   - Correct port
   - Valid JSON format in request body
   - All required headers are present

## Response Codes

- `200 OK` - request successfully processed
- `400 Bad Request` - invalid request format
- `500 Internal Server Error` - internal server error

## Limitations

1. The `text` field cannot be empty
2. Russian or English text is recommended
3. For more accurate results, use texts between 3 and 500 words

## Request and Response Examples

### Positive Text

**Request:**
```json
{
    "text": "I am very happy today!"
}
```

**Response:**
```json
{
    "sentiment": "Very Positive",
    "score": 0.8534
}
```

### Negative Text

**Request:**
```json
{
    "text": "This was a terrible day."
}
```

**Response:**
```json
{
    "sentiment": "Very Negative",
    "score": 0.7645
}
```

### Neutral Text

**Request:**
```json
{
    "text": "Today is an ordinary day."
}
```

**Response:**
```json
{
    "sentiment": "Neutral",
    "score": 0.6123
}
```

## Technical Stack

- Go
- Gin (web framework)
- Docker
- ML model for sentiment analysis

## Installation and Setup

### Requirements

- Go 1.16 or higher
- Docker and Docker Compose (optional)

### Local Setup

1. Clone the repository:
```bash
git clone [repository-url]
cd mood
```

2. Install dependencies:
```bash
go mod download
```

3. Create `.env` file based on `app.env`:
```bash
cp app.env .env
```

4. Run the service:
```bash
go run .
```

### Docker Setup

```bash
docker-compose up --build
```

## Project Structure

```
.
├── main.go          # Application entry point
├── service.go       # Business logic
├── handler.go       # HTTP handlers
├── model.go         # Data models
├── config.go        # Configuration
├── docker-compose.yaml
└── ml_mood/         # ML model
```

## Development

### Adding New Features

1. Add new models in `model.go`
2. Implement business logic in `service.go`
3. Create handlers in `handler.go`
4. Add new routes in `setupRouter` in `main.go`

### Testing

```bash
go test ./...
```

## Contributing

### How to Contribute

1. Fork the repository
2. Create a branch for your changes:
   ```bash
   git checkout -b feature/feature-name
   ```
3. Make changes and commit them:
   ```bash
   git add .
   git commit -m "Description of changes"
   ```
4. Push changes to your fork:
   ```bash
   git push origin feature/feature-name
   ```
5. Create a Pull Request to the main repository

### Code Style Guidelines

1. Use `gofmt` for Go code formatting
2. Add comments to public functions and structures
3. Update documentation when making changes
4. Add tests for new functionality

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test ./... -cover
```

### Building the Project

```bash
# Local build
go build

# Build Docker image
docker build -t mood-api .
```

### Running with Docker

```bash
# Run all services
docker compose up --build

# Run in background
docker compose up -d --build

# Stop services
docker compose down
```

## License

MIT

## Authors

- [bkzntsv](https://github.com/bkzntsv) - Project Creator

## Support

If you have questions or suggestions, create an issue in the repository or contact the author via GitHub. 