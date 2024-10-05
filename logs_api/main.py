from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
from datetime import datetime
import json

app = FastAPI()

# Modelo para el log
class Log(BaseModel):
    usuario_id: int
    accion: str
    descripcion: str
    fecha: datetime = datetime.now()

# Lista en memoria para almacenar los logs (puedes reemplazar esto con una base de datos)
logs = []

# Endpoint para obtener todos los logs
@app.get("/logs", response_model=list[Log])
async def get_logs():
    return logs

# Endpoint para crear un nuevo log
@app.post("/logs", response_model=Log)
async def create_log(log: Log):
    logs.append(log)
    return log

# Endpoint para obtener los logs por usuario
@app.get("/logs/{usuario_id}", response_model=list[Log])
async def get_logs_by_usuario(usuario_id: int):
    usuario_logs = [log for log in logs if log.usuario_id == usuario_id]
    if not usuario_logs:
        raise HTTPException(status_code=404, detail="No logs found for this user")
    return usuario_logs

