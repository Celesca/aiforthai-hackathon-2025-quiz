from flask import Flask, jsonify
import logging
import os

logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

app = Flask(__name__)

@app.route('/')
def hello():
    logger.info("API2: Received request")
    response = {"message": "Hello from API2"}
    logger.info(f"API2: Sending response: {response['message']}")
    return jsonify(response)

port = int(os.environ.get('PORT', 8081))
logger.info(f"API2 starting on port {port}")
app.run(host='0.0.0.0', port=port)