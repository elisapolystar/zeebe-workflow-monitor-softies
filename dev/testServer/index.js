const express = require('express');
const app = express();
const PORT = 8000;
const CONNECTIONS = new Map();

const input_app = express();
const input_port = 8080;

const http = require('http').Server(app);
const cors = require('cors');

app.use(cors());

const socketIO = require('socket.io')(http, {
  cors: {
    origin: "http://localhost:3000"
  }
});

socketIO.on('connection', (socket) => {
  console.log(`Socket ${socket.id} connected`);
  CONNECTIONS.set(socket.id, socket);
  socket.on('disconnect', () => {
    CONNECTIONS.delete(socket.id);
    console.log(`Socket ${socket.id} disconnected`);
  });
  socket.emit('message', 'got connection');
});

http.listen(PORT, () => {
  console.log(`Server listening on ${PORT}`);
});

const input_http = require('http').Server(input_app);
input_app.post('/', function (req, res) {
  let input = [];
  req.on('data', (chunk) => {
    input.push(chunk);
  }).on('end', () => {
    input = Buffer.concat(input).toString();

    console.log('got input:', input);
    console.log(`emitting to ${CONNECTIONS.size} sockets`);
    CONNECTIONS.forEach((socket, id) => {
      socket.emit('message', input);
      console.log(`emitted to ${id}`);
    })
    res.writeHead(200);
    res.end('got input\n');
  });

});

input_http.listen(input_port);
