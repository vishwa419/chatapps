const socket = io('http://localhost:4000',{
    auth: {
        secret: "This is a secret"
    },
    query:{
        meaningOfLife: 42,
    }
})

socket.on('welcome', data=>{
	console.log("welcome ", data)
	socket.emit("hello all")
	}
)

socket.on('msgStoC', newMessage=>{
    document.getElementById('messages').innerHTML += `<li>${newMessage}</li>`
})

document.getElementById('messages-form').addEventListener('submit',e=>{
    e.preventDefault()
    const newMessage = document.getElementById('user-message').value
    document.getElementById('user-message').value = ""
    // this socket is sending an event to the server...
    socket.emit('messageFromClientToServer',newMessage)
})
