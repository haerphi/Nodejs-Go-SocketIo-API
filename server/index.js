const express = require("express");
const http = require("http");
const socketIo = require("socket.io");
//port for the app
const { APP_PORT, PORT } = process.env;
const port = APP_PORT || PORT || 4001;

const app = express();

const server = http.createServer(app);
const io = socketIo(server); // < Interesting!

app.get("/", (req, res) => {
  res.send(`Welcome !`);
});

app.get("/:anything", (req, res) => {
  console.log(`ℹ️  (${req.method.toUpperCase()}) ${req.url}`);
  res.send(`404 : /${req.params.anything} doesn't exit`);
});

//start server
server.listen(port, () => console.log(`🚀 Listening on port ${port}`));

//SocketIO event
io.on("connection", socket => {
  console.log(`New client connected ${socket.id}`);

  //disconnect
  socket.on("disconnect", () => {
    console.log(`Client disconnected ${socket.id}`);
  });

  //receved a new message
  socket.on("newMessage", receve => {
    console.log(receve);
    socket.emit("newMessage", `Did you said ${receve} ?`);
  });
});
