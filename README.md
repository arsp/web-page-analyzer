# Web Page Analyzer

## 1. Project Overview
A Go-based web application that analyzes a given web page URL and extracts
structural and accessibility information.

- HTML version
- Page title
- Head Tags counts (h1–h6)
- Internal vs External links counts
- Inaccessible links counts
- Login form detection

## 2. Prerequisites
- Go 1.25.6+ version

## 3. Build
- Clone the repository:
```
git clone <repository-url>
cd web-page-analyzer
```

- Install dependencies:
```
go mod tidy
```

- Build the application:
```
go build ./cmd/webanalyzer
```

## 4. Run
```
go run ./cmd/webanalyzer
```
Access application:
```
http://localhost:8080
```

## 5.  Deployment (Docker)
Build Docker image:
```
docker build -t web-page-analyzer .
```

Run container:
```
docker run -p 8080:8080 web-page-analyzer
```

## 6. Testing
Run unit tests:
```
go test ./... -cover
```

## 7. Architecture Overview
```
web-page-analyzer/
─ cmd/webanalyzer        # Application entry point
─ internal/server        # HTTP server and routing
─ internal/analyzer      # Core analysis logic
─ internal/validator     # URL validation
─ internal/model         # Data models
─ web/templates          # HTML templates
─ web/static             # CSS
```

## 8. Usage
1. Enter a valid URL - start with http:// or https://
2. Click Analyze
3. View analysis results:
   - HTML version
   - Title
   - Headings
   - Link statistics
   - Login form detection

4. If the page is not reachable, an error page displays the HTTP status code and message.

## 9. Technologies Used
Backend
- Golang

Frontend
- HTML
- CSS

## 10. Assumptions
- Internal links are defined as same host or relative paths.
- External links are pointing to different hosts.
