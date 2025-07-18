# AI Thailand Hackathon 2025 Quiz by Ex_Machina

## API1 (Golang) + API2 (Python)

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
- **Endpoints**:
  - `GET /` - Forwards request to API2 and returns response

### API2 - Python Service (Backend)  
- **Language**: Python (Flask)
- **Port**: 8081
- **Endpoints**:
  - `GET /` - Returns hello message

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

#### Test the Service Chain
```bash
# Test via API1 (which forwards to API2)
curl http://localhost:8080/

# Direct test API2
curl http://localhost:8081/
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

### API Response Format
```json
{
  "message": "Hello from API2"
}
```

### Example Request Flow
1. **User** → `GET http://localhost:8080/`
2. **API1** → Logs: "API1: Received request, calling API2"
3. **API1** → Calls `GET http://api2:8081/`
4. **API2** → Logs: "API2: Received request"
5. **API2** → Returns `{"message": "Hello from API2"}`
6. **API1** → Logs: "API1: Got response from API2: Hello from API2"
7. **API1** → Returns same response to user

## 🔧 Development

### Local Development (without Docker)

#### API1 (Golang)
```bash
cd service-1
export SERVICE_2_URL="http://localhost:8081"
go run main.go
# Runs on http://localhost:8080
```

#### API2 (Python)
```bash
cd service-2
pip install flask
python app.py
# Runs on http://localhost:8081
```

## 📁 Project Structure
```
.
├── docker-compose.yml          # Docker orchestration
├── README.md                   # This file
├── service-1/                  # API1 - Golang Service
│   ├── main.go                # Go application (35 lines)
│   ├── go.mod                 # Go modules
│   └── Dockerfile             # Docker config
└── service-2/                  # API2 - Python Service
    ├── app.py                 # Python Flask app (18 lines)
    ├── requirements.txt       # Python dependencies
    └── Dockerfile             # Docker config
```

## 📝 API Endpoints Summary

| Service | Endpoint | Method | Description |
|---------|----------|--------|-------------|
| API1 | `/` | GET | Forwards to API2 and returns response |
| API2 | `/` | GET | Returns hello message |

### Debug / Monitor Commands

```bash
# Check running containers
docker ps

# Check container logs
docker-compose logs api1
docker-compose logs api2

# Test connectivity between services
docker-compose exec api1 wget -qO- http://api2:8081/

# Restart services
docker-compose restart
```

### Expected Log Output
```
api1-golang  | 2025/07/18 13:42:35 API1 starting on :8080
api2-python  | INFO:__main__:API2 starting on port 8081
api1-golang  | 2025/07/18 13:42:35 API1: Received request, calling API2
api2-python  | INFO:__main__:API2: Received request
api2-python  | INFO:__main__:API2: Sending response: Hello from API2
api1-golang  | 2025/07/18 13:42:35 API1: Got response from API2: Hello from API2
```