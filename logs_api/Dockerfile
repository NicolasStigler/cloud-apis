# Usa una imagen oficial de Python como base
FROM python:3.9-slim

# Establece el directorio de trabajo
WORKDIR /app

# Copia el archivo de requerimientos y los instala
COPY requirements.txt requirements.txt
RUN pip install --no-cache-dir -r requirements.txt

# Copia el código fuente de la API
COPY . .

# Exponer el puerto en el que correrá la aplicación
EXPOSE 8000

# Comando para ejecutar la API usando Uvicorn
CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8000"]

