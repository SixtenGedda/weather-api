# Weather API service written in go
A lightweight weather service utilizing Redis for 12h caching and Gin for HTTP routing.

## How to run

### 1. Add your own [API key](https://www.visualcrossing.com/weather-api/) to compose.yaml:
```yaml
environment:
    REDIS_ADDR: "redis:6379"
    API_KEY: "YOUR-API_KEY-HERE"
```


### 2. Start the service:
```bash
docker compose up --build
```

### 3. Then to use the API, input whatever location you want:
```bash 
curl http://localhost:8080/london
``` 
