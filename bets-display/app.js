const express = require('express');
const { Pool } = require('pg');
const dotenv = require('dotenv');

// Load environment variables from .env file
dotenv.config();

const app = express();
const port = 3000;

// Set up the PostgreSQL connection
const pool = new Pool({
    user: process.env.POSTGRES_USER,
    host: process.env.POSTGRES_HOST,
    database: process.env.POSTGRES_DB,
    password: process.env.POSTGRES_PASSWORD,
    port: process.env.POSTGRES_PORT,
});

// Set EJS as the templating engine
app.set('view engine', 'ejs');

// Route to display the bets
app.get('/', (req, res) => {
    pool.query('SELECT * FROM bets ORDER BY posted_at DESC', (error, results) => {
        if (error) {
            throw error;
        }
        res.render('index', { bets: results.rows });
    });
});

// Start the server
app.listen(port, () => {
    console.log(`App running on port ${port}.`);
});