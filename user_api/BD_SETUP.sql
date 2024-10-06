CREATE TABLE usuarios (
    usuario_id UUID PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    balance DECIMAL(10, 2) NOT NULL DEFAULT 0,
    fecha_registro TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE transacciones (
    transaccion_id UUID PRIMARY KEY,
    usuario_id UUID REFERENCES usuarios(usuario_id),
    tipo VARCHAR(10) CHECK (tipo IN ('ingreso', 'retiro')),
    monto DECIMAL(10, 2) NOT NULL,
    fecha_transaccion TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
