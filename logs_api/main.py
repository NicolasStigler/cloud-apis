from fastapi import FastAPI
from pymongo import MongoClient
from pydantic import BaseModel
from datetime import datetime
import os

app = FastAPI()

# Conexión a MongoDB
client = MongoClient("mongodb://user:password@localhost:8002/") # cambiar a ip de base
db = client["casino"]
logs_collection = db["logs"]

# Modelo para el registro de logs
class LogEntry(BaseModel):
    level: str
    message: str
    timestamp: datetime = None

# Registrar un nuevo log
@app.post("/logs/")
async def create_log(log_entry: LogEntry):
    if log_entry.timestamp is None:
        log_entry.timestamp = datetime.utcnow()  # Establece la fecha y hora actual si no se proporciona
    logs_collection.insert_one(log_entry.dict())
    return {"status": "Log entry created", "log": log_entry}

# Obtener todos los logs
@app.get("/logs/")
async def get_logs():
    logs = list(logs_collection.find())
    for log in logs:
        log["_id"] = str(log["_id"])  # Convierte el ObjectId a string
    return logs

