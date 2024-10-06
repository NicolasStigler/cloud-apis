const express = require('express');
const { Pool } = require('pg');
const swaggerJsdoc = require('swagger-jsdoc');
const swaggerUi = require('swagger-ui-express');

const app = express();
app.use(express.json());

const pool = new Pool({
  user: 'postgres',
  host: 'localhost', // ip de MV BD
  database: 'casino',
  password: 'GNU+Casino',
  port: 8000,
});

// Swagger definition
const swaggerOptions = {
  swaggerDefinition: {
    openapi: '3.0.0',
    info: {
      title: 'User API',
      version: '1.0.0',
      description: 'API for User and Account Management',
    },
    servers: [{ url: 'http://localhost:8001' }],
  },
  apis: ['./user_api.js'],
};
const swaggerDocs = swaggerJsdoc(swaggerOptions);
app.use('/api-docs', swaggerUi.serve, swaggerUi.setup(swaggerDocs));

/**
 * @swagger
 * /register:
 *   post:
 *     description: Register a new user
 *     parameters:
 *       - in: body
 *         name: user
 *         schema:
 *           type: object
 *           required:
 *             - name
 *             - email
 *           properties:
 *             name:
 *               type: string
 *             email:
 *               type: string
 *     responses:
 *       200:
 *         description: User registered
 */
app.post('/register', async (req, res) => {
  const { name, email } = req.body;
  const result = await pool.query(
    'INSERT INTO usuarios (nombre, email, balance, fecha_registro) VALUES ($1, $2, $3, NOW()) RETURNING *',
    [name, email, 0]
  );
  res.json(result.rows[0]);
});

// Endpoint para buscar usuarios
/**
 * @swagger
 * /usuarios:
 *   get:
 *     description: Obtener en orden descendente los usuarios con mayor balance
 *     responses:
 *       200:
 *         description: Leaderboard
 *         content:
 *           application/json:
 *             schema:
 *               type: array
 *               items:
 *                 type: object
 *                 properties:
 *                   id:
 *                     type: integer
 *                   nombre:
 *                     type: string
 *                   email:
 *                     type: string
 *                   balance:
 *                     type: number
 *                   fecha_registro:
 *                     type: string
 *                     format: date-time
 */
app.get('/leaderboard', async (req, res) => {
  try {
    const result = await pool.query('SELECT * FROM usuarios ORDER BY balance DESC');
    res.json(result.rows);
  } catch (err) {
    console.error(err);
    res.status(500).send('Error retrieving users');
  }
});

// Endpoint para buscar un usuario por nombre
/**
 * @swagger
 * /usuarios/{nombre}:
 *   get:
 *     description: Buscar un usuario por su nombre
 *     parameters:
 *       - in: path
 *         name: nombre
 *         schema:
 *           type: string
 *         required: true
 *         description: Nombre del usuario
 *     responses:
 *       200:
 *         description: Usuario encontrado
 *         content:
 *           application/json:
 *             schema:
 *               type: object
 *               properties:
 *                 id:
 *                   type: integer
 *                 nombre:
 *                   type: string
 *                 email:
 *                   type: string
 *                 balance:
 *                   type: number
 *                 fecha_registro:
 *                   type: string
 *                   format: date-time
 *       404:
 *         description: Usuario no encontrado
 */
app.get('/usuarios/:nombre', async (req, res) => {
  const { nombre } = req.params; // Obtenemos el nombre desde la URL
  try {
    const result = await pool.query(
      'SELECT * FROM usuarios WHERE nombre = $1',
      [nombre]
    );
    if (result.rows.length > 0) {
      res.json(result.rows[0]); // Si se encuentra el usuario, lo enviamos en formato JSON
    } else {
      res.status(404).send('Usuario no encontrado');
    }
  } catch (err) {
    console.error(err);
    res.status(500).send('Error al buscar el usuario');
  }
});

app.listen(8001, () => {
  console.log('User API listening on port 8001');
});
