CREATE TABLE usuarios (
  usuario_id SERIAL PRIMARY KEY,
  nombre VARCHAR(50),
  email VARCHAR(50),
  balance DECIMAL(10, 2) DEFAULT 0,
  fecha_registro TIMESTAMP DEFAULT NOW()
);
