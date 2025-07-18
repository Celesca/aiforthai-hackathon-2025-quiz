from flask import Flask, jsonify, request
import logging
from datetime import datetime
import os

# Configure logging
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

app = Flask(__name__)

@app.route('/', methods=['GET', 'POST'])
def hello():
    logger.info(f"API2: Received {request.method} request to /")
    
    response = {
        "message": "Hello from API2 (Python)! 🐍",
        "status": "success",
        "timestamp": datetime.now().isoformat(),
        "processed_by": "API2-Python",
        "language": "Python",
        "framework": "Flask"
    }
    
    logger.info("API2: Sending Hello World response")
    return jsonify(response)

@app.route('/api/hello', methods=['GET', 'POST'])
def api_hello():
    logger.info(f"API2: Received {request.method} request to /api/hello")
    
    # Get data from request if POST
    data = None
    if request.method == 'POST' and request.is_json:
        data = request.get_json()
        logger.info(f"API2: Received data: {data}")
    
    response = {
        "message": "API2 processed your request successfully! ✅",
        "status": "success",
        "timestamp": datetime.now().isoformat(),
        "processed_by": "API2-Python",
        "request_data": data,
        "python_version": "3.11+",
        "response_from": "Flask API"
    }
    
    logger.info("API2: Successfully processed request")
    return jsonify(response)

@app.route('/health', methods=['GET'])
def health():
    logger.info("API2: Health check requested")
    return jsonify({
        "message": "API2 (Python) is healthy and running! 💚",
        "status": "healthy",
        "timestamp": datetime.now().isoformat(),
        "service": "API2-Python"
    })

@app.errorhandler(404)
def not_found(error):
    logger.warning(f"API2: 404 error for path: {request.path}")
    return jsonify({
        "error": "Endpoint not found",
        "status": "error",
        "timestamp": datetime.now().isoformat(),
        "processed_by": "API2-Python"
    }), 404

@app.errorhandler(500)
def internal_error(error):
    logger.error(f"API2: Internal error: {error}")
    return jsonify({
        "error": "Internal server error",
        "status": "error",
        "timestamp": datetime.now().isoformat(),
        "processed_by": "API2-Python"
    }), 500

if __name__ == '__main__':
    port = int(os.environ.get('PORT', 8081))
    logger.info(f"🚀 API2 (Python) starting on port :{port}")
    logger.info("API2: Ready to process requests from API1")
    app.run(host='0.0.0.0', port=port, debug=False)
