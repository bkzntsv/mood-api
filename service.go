package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Service определяет интерфейс для анализа настроений текста.
// Он предоставляет методы для анализа текста и получения эндпоинта модели.
type Service interface {
	// AnalyzeSentiment анализирует настроение переданного текста.
	// Возвращает SentimentResponse с результатами анализа и возможную ошибку.
	AnalyzeSentiment(text string) (SentimentResponse, error)

	// GetModelEndpoint возвращает URL эндпоинта модели для анализа настроений.
	GetModelEndpoint() string
}

// service реализует интерфейс Service
type service struct {
	config     Config
	httpClient *http.Client
}

// NewService создает новый экземпляр сервиса с указанной конфигурацией.
// Возвращает реализацию интерфейса Service.
func NewService(config Config) Service {
	return &service{
		config: config,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// AnalyzeSentiment отправляет текст на анализ настроений в удаленный сервис.
// Параметры:
//   - text: текст для анализа
//
// Возвращает:
//   - SentimentResponse: результат анализа настроений
//   - error: ошибка, если анализ не удался
func (s *service) AnalyzeSentiment(text string) (SentimentResponse, error) {
	modelEndpoint := s.config.ModelEndpoint

	request := SentimentRequest{Text: text}
	requestBody, err := json.Marshal(request)
	if err != nil {
		return SentimentResponse{}, fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, modelEndpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		return SentimentResponse{}, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return SentimentResponse{}, fmt.Errorf("failed to call external service: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return SentimentResponse{}, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return SentimentResponse{}, fmt.Errorf("external service returned error: %s, status code: %d", string(body), resp.StatusCode)
	}

	var result SentimentResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return SentimentResponse{}, fmt.Errorf("failed to parse response: %w, body: %s", err, string(body))
	}

	return result, nil
}

// GetModelEndpoint возвращает URL эндпоинта модели, используемого для анализа настроений.
// Возвращает:
//   - string: URL эндпоинта модели
func (s *service) GetModelEndpoint() string {
	return s.config.ModelEndpoint
}
