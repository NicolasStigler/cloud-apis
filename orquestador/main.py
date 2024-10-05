from fastapi import FastAPI, HTTPException
import httpx

app = FastAPI()

# URLs de las APIs
USER_API_URL = "http://localhost:8080"
BET_API_URL = "http://localhost:8081"
LOGS_API_URL = "http://localhost:8082"

@app.get("/")
def read_root():
    return {"message": "Orchestrator API is running"}

# Endpoint para obtener la información de un usuario y su balance
@app.get("/user/{user_id}")
async def get_user_balance(user_id: int):
    async with httpx.AsyncClient() as client:
        # Llamar a la user_api para obtener información del usuario
        user_response = await client.get(f"{USER_API_URL}/usuarios/{user_id}")
        
        if user_response.status_code != 200:
            raise HTTPException(status_code=user_response.status_code, detail="Error retrieving user info")
        
        user_data = user_response.json()
        
        # Obtener el balance del usuario desde user_api o bet_api
        balance_response = await client.get(f"{USER_API_URL}/usuarios/{user_id}/balance")
        
        if balance_response.status_code != 200:
            raise HTTPException(status_code=balance_response.status_code, detail="Error retrieving balance")
        
        balance_data = balance_response.json()
        
        # Retornar datos combinados
        return {"user_info": user_data, "balance": balance_data}

# Endpoint para registrar una nueva apuesta y hacer log de la acción
@app.post("/bet/{user_id}")
async def place_bet(user_id: int, amount: float, game_type: str):
    async with httpx.AsyncClient() as client:
        # Llamar a bet_api para registrar la apuesta
        bet_response = await client.post(f"{BET_API_URL}/apuestas", json={"user_id": user_id, "amount": amount, "game_type": game_type})
        
        if bet_response.status_code != 201:
            raise HTTPException(status_code=bet_response.status_code, detail="Error placing bet")
        
        bet_data = bet_response.json()
        
        # Llamar a logs_api para registrar la transacción
        log_response = await client.post(f"{LOGS_API_URL}/logs", json={"user_id": user_id, "action": "bet placed", "details": bet_data})
        
        if log_response.status_code != 201:
            raise HTTPException(status_code=log_response.status_code, detail="Error logging action")
        
        return {"bet": bet_data, "log": log_response.json()}

# Endpoint para obtener logs de usuario
@app.get("/logs/{user_id}")
async def get_logs(user_id: int):
    async with httpx.AsyncClient() as client:
        log_response = await client.get(f"{LOGS_API_URL}/logs/{user_id}")
        
        if log_response.status_code != 200:
            raise HTTPException(status_code=log_response.status_code, detail="Error retrieving logs")
        
        return log_response.json()

