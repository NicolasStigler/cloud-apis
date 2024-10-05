const express = require('express');
const { Pool } = require('pg');
const swaggerJsdoc = require('swagger-jsdoc');
const swaggerUi = require('swagger-ui-express');

const app = express();
app.use(express.json());

const pool = new Pool({
  user: 'postgres',
  host: 'localhost',
  database: 'casino',
  //password: 'GNU+Casino',
  port: 5432,
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
    servers: [{ url: 'http://localhost:3000' }],
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

app.listen(3000, () => {
  console.log('User API listening on port 3000');
});

