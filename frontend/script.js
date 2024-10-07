const API_URL = "http://localhost:8004";  // URL de la API orquestadora

// Obtener información del usuario y balance
async function getUser() {
    const userId = document.getElementById("user-id").value;
    const userInfoDiv = document.getElementById("user-info");

    try {
        const response = await fetch(`${API_URL}/user/${userId}`);
        const data = await response.json();
        
        if (response.ok) {
            userInfoDiv.innerHTML = `<p>Nombre: ${data.user_info.nombre}</p>
                                     <p>Email: ${data.user_info.email}</p>
                                     <p>Balance: ${data.balance.balance}</p>`;
        } else {
            userInfoDiv.innerHTML = `<p>Error: ${data.detail}</p>`;
        }
    } catch (error) {
        userInfoDiv.innerHTML = `<p>Error: ${error.message}</p>`;
    }
}

// Registrar una apuesta
async function placeBet() {
    const userId = document.getElementById("bet-user-id").value;
    const amount = document.getElementById("amount").value;
    const gameType = document.getElementById("game-type").value;
    const betResultDiv = document.getElementById("bet-result");

    try {
        const response = await fetch(`${API_URL}/bet/${userId}`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ amount: parseFloat(amount), game_type: gameType })
        });
        const data = await response.json();
        
        if (response.ok) {
            betResultDiv.innerHTML = `<p>Apuesta registrada con éxito!</p>
                                      <p>Detalles de la apuesta: ${JSON.stringify(data.bet)}</p>`;
        } else {
            betResultDiv.innerHTML = `<p>Error: ${data.detail}</p>`;
        }
    } catch (error) {
        betResultDiv.innerHTML = `<p>Error: ${error.message}</p>`;
    }
}

// Obtener logs de usuario
async function getLogs() {
    const userId = document.getElementById("log-user-id").value;
    const logsDiv = document.getElementById("logs");

    try {
        const response = await fetch(`${API_URL}/logs/${userId}`);
        const data = await response.json();
        
        if (response.ok) {
            logsDiv.innerHTML = `<p>Logs:</p><pre>${JSON.stringify(data, null, 2)}</pre>`;
        } else {
            logsDiv.innerHTML = `<p>Error: ${data.detail}</p>`;
        }
    } catch (error) {
        logsDiv.innerHTML = `<p>Error: ${error.message}</p>`;
    }
}

