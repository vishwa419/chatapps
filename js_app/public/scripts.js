const socket = io('http://localhost:4000',{
    auth: {
        secret: "This is a secret"
    },
    query:{
        meaningOfLife: 42,
    }
})

socket.on('welcome', data=>{
	console.log( data)
	socket.emit("hello all")
	}
)

socket.on('newClient',data=>{
    console.log('Message to all clients: A new socket has joined', data)
})

socket.on('msgStoAllC', newMessage=>{
    console.log("here")
    document.getElementById('messages').innerHTML += `<li>${newMessage}</li>`
})

document.getElementById('messages-form').addEventListener('submit',e=>{
    e.preventDefault()
    const newMessage = document.getElementById('user-message').value
    document.getElementById('user-message').value = ""
    // this socket is sending an event to the server...
    socket.emit('msgCtoS',newMessage)
})
