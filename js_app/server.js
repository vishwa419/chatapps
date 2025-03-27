const express = require('express')
const {Server} = require('socket.io')

const app = express()
app.use(express.static('public'))
const es = app.listen(4000)

const io = new Server(es,{
	cors: ['http://localhost:4000']
})

io.on(
	'connection', socket=>{
		console.log(socket.handshake)
		console.log(socket.id, " has joined")
		socket.on("msgCtoS", data=>{
			io.emit('msgStoAllC', data)
			}
		)
	}
)
