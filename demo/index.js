const express = require('express');
const app = express();
const port = process.env.PORT || 8080;

app.use(express.json());

app.get('/healthz', (req, res) => {
  res.status(200).json({ status: 'ok', service: 'secure-demo' });
});

app.get('/', (req, res) => {
  res.send('Hello from secure-demo!');
});

app.listen(port, () => {
  console.log(`secure-demo listening on port ${port}`);
});
