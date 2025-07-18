# 🚀 Microservices Architecture - API1 (Golang) + API2 (Python)

## 📋 Overview

This project implements a simple microservices architecture where:
- **User** sends requests to **API1** (Golang)
- **API1** forwards requests to **API2** (Python) 
- **API2** processes the request and returns response
- **API1** sends the final response back to the **User**

```
User → API1 (Golang:8080) → API2 (Python:8081) → Response → User
```

## 🏗️ Architecture

```
┌─────────┐    ┌──────────────┐    ┌──────────────┐
│  User   │───▶│ API1 (Go)    │───▶│ API2 (Python)│
│         │    │ Port: 8080   │    │ Port: 8081   │
│         │◀───│ Gateway      │◀───│ Backend      │
└─────────┘    └──────────────┘    └──────────────┘
```

## 🛠️ Services

### API1 - Golang Service (Gateway)
- **Language**: Go
- **Port**: 8080
- **Role**: Receives user requests and forwards to API2
- **Endpoints**:
  - `GET /` - Hello World endpoint
  - `GET/POST /api/hello` - API endpoint
- **Logging**: Yes ✅

### API2 - Python Service (Backend)  
- **Language**: Python (Flask)
- **Port**: 8081
- **Role**: Processes requests and returns responses
- **Endpoints**:
  - `GET /` - Hello World endpoint
  - `GET/POST /api/hello` - API endpoint
  - `GET /health` - Health check
- **Logging**: Yes ✅

## 🚀 Quick Start

### Prerequisites
- Docker
- Docker Compose

### 1. Build and Run
```bash
# Build and start all services
docker-compose up --build

# Run in background
docker-compose up --build -d
```

### 2. Test the APIs

#### Test Hello World
```bash
# Test via API1 (which forwards to API2)
curl http://localhost:8080/

# Direct test API2
curl http://localhost:8081/
```

#### Test API Endpoint
```bash
# GET request
curl http://localhost:8080/api/hello

# POST request with data
curl -X POST http://localhost:8080/api/hello \
  -H "Content-Type: application/json" \
  -d '{"message": "Hello from user!"}'
```

#### Health Check
```bash
# Check API2 health
curl http://localhost:8081/health
```

### 3. View Logs
```bash
# View all logs
docker-compose logs

# View specific service logs
docker-compose logs api1
docker-compose logs api2

# Follow logs in real-time
docker-compose logs -f
```

### 4. Stop Services
```bash
# Stop all services
docker-compose down

# Stop and remove volumes
docker-compose down -v
```

## 📊 Expected Response

### Hello World Response
```json
{
  "message": "Hello from API1 (Golang)!",
  "status": "success",
  "timestamp": "2025-07-18T10:30:00Z",
  "processed_by": "API1-Golang",
  "from_api2": {
    "message": "Hello from API2 (Python)! 🐍",
    "status": "success",
    "timestamp": "2025-07-18T10:30:00Z",
    "processed_by": "API2-Python",
    "language": "Python",
    "framework": "Flask"
  }
}
```

## 🔧 Development

### Local Development (without Docker)

#### API1 (Golang)
```bash
cd service-1
go run main.go
# Runs on http://localhost:8080
```

#### API2 (Python)
```bash
cd service-2
pip install -r requirements.txt
python app.py
# Runs on http://localhost:8081
```

## 📁 Project Structure
```
.
├── docker-compose.yml          # Docker orchestration
├── README.md                   # This file
├── service-1/                  # API1 - Golang Service
│   ├── main.go                # Go application
│   ├── go.mod                 # Go modules
│   └── Dockerfile             # Docker config
└── service-2/                  # API2 - Python Service
    ├── app.py                 # Python Flask app
    ├── requirements.txt       # Python dependencies
    └── Dockerfile             # Docker config
```

## 🌟 Features

- ✅ **Simple Hello World** implementation
- ✅ **Request forwarding** from API1 to API2
- ✅ **Comprehensive logging** in both services
- ✅ **Docker containerization**
- ✅ **Docker Compose** orchestration
- ✅ **Health checks**
- ✅ **Error handling**
- ✅ **JSON responses**
- ✅ **Cross-service communication**

## 📝 API Endpoints Summary

| Service | Endpoint | Method | Description |
|---------|----------|--------|-------------|
| API1 | `/` | GET | Hello World (forwards to API2) |
| API1 | `/api/hello` | GET/POST | API endpoint (forwards to API2) |
| API2 | `/` | GET/POST | Hello World |
| API2 | `/api/hello` | GET/POST | API endpoint |
| API2 | `/health` | GET | Health check |

## 🔍 Troubleshooting

### Common Issues
1. **Port conflicts**: Make sure ports 8080 and 8081 are available
2. **Docker issues**: Ensure Docker is running
3. **Network issues**: Check if services can communicate within Docker network

### Debug Commands
```bash
# Check running containers
docker ps

# Check logs
docker-compose logs api1
docker-compose logs api2

# Restart specific service
docker-compose restart api1

# Rebuild specific service
docker-compose up --build api1
```

## 🎯 Next Steps

To extend this project, you could add:
- Database integration
- Authentication
- More complex business logic
- API documentation (Swagger)
- Monitoring and metrics
- Load balancing
- Message queues

---

**สร้างโดย**: AI Thailand Project 🇹🇭  
**เทคโนโลยี**: Golang + Python + Docker  
**วัตถุประสงค์**: แสดงการทำงานของ Microservices Architecture
