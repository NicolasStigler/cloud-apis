#!/bin/bash

cd /BDs/

docker-compose up -d

sleep 10

PGPASSWORD=GNUpluscasino psql -h localhost -p 8000 -U postgres -d casino_db -c "
CREATE TABLE IF NOT EXISTS usuarios (
  usuario_id INT PRIMARY KEY,
  nombre VARCHAR(50),
  email VARCHAR(50),
  balance DECIMAL(10, 2) DEFAULT 0,
  fecha_registro TIMESTAMP DEFAULT NOW()
);
"

PGPASSWORD=GNUpluscasino psql -h localhost -p 8000 -U postgres -d casino_db -c "
CREATE TABLE IF NOT EXISTS transacciones (
  transaccion_id INT PRIMARY KEY,
  usuario_id INT NOT NULL,
  tipo VARCHAR(50),
  monto DECIMAL(10, 2) NOT NULL,
  fecha_transaccion TIMESTAMP DEFAULT NOW(),
  FOREIGN KEY (usuario_id) REFERENCES usuarios(usuario_id)
);
"

mysql -h 127.0.0.1 -P 8001 -u root GNUpluscasino -D casino_db -e "
CREATE TABLE IF NOT EXISTS Juego (
    juego_id INT AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    tipo VARCHAR(50) NOT NULL
);
"

mysql -h 127.0.0.1 -P 8001 -u root GNUpluscasino -D casino_db -e "
CREATE TABLE IF NOT EXISTS Apuesta (
    apuesta_id INT AUTO_INCREMENT PRIMARY KEY,
    usuario_id INT NOT NULL,
    juego_id INT NOT NULL,
    monto DECIMAL(10, 2) NOT NULL,
    resultado VARCHAR(50) NOT NULL,
    fecha TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (juego_id) REFERENCES Juego(juego_id)
);
"

echo "Containers are up."

