# WebRTC Video Calling System: Go + HTMX vs. Node.js

[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Maintenance](https://img.shields.io/badge/Maintained-yes-green.svg)](https://github.com/vish419)

This project demonstrates the implementation of a WebRTC-based video calling system using two distinct technology stacks: **Go with HTMX** and **Node.js with Socket.IO**. It compares developer experience, performance, and architectural differences in building real-time communication features using these stacks.

---

## 🚀 Overview

This project focuses on core peer-to-peer WebRTC communication. The server handles signaling, while the browser handles media acquisition and peer connection logic.

### Core WebRTC Components Implemented

- **Signaling Setup:** Server enables peer discovery and communication setup.
- **SDP Exchange:** Peers negotiate media session parameters.
- **ICE Candidate Exchange:** Helps peers discover optimal network paths.
- **Media Streaming:** Video/audio flows directly between peers.
- **Call Termination:** Clean handling of disconnection and call end events.

For deeper insights and technical breakdowns, read the blog series:  
👉 [What Happens When You Make the Same Software on Go & HTMX vs JS & Node.js](https://vish419.hashnode.dev)

---

## 🛠 Implementation Strategies

### Go + HTMX

- Backend: Go
- Frontend: HTMX + minimal Vanilla JavaScript
- Signaling: Raw WebSockets
- Focus: Performance, simplicity, and control

### Node.js

- Backend: Node.js with Express.js
- Frontend: Vanilla JS + Socket.IO
- Signaling: Socket.IO abstraction
- Focus: Developer ergonomics and rapid development

---

## ✅ Advantages of This WebRTC Setup

- **Lightweight Servers:** No media relay; server only handles signaling.
- **Bandwidth Efficiency:** Peer-to-peer streaming reduces server load.
- **Low Latency:** Direct peer connections improve responsiveness.
- **Highly Scalable:** Servers can support many concurrent sessions.

---

## 🔧 Technology Decisions

### Signaling Mechanisms

| Stack        | Signaling Approach       | Notes                                                                 |
|--------------|--------------------------|-----------------------------------------------------------------------|
| Go + HTMX    | Raw WebSockets           | Fine-grained control; great with goroutines and Go's concurrency model |
| Node.js      | Socket.IO                | Built-in reconnection, rooms, events; developer-friendly abstraction  |

### Frontend Asynchrony

- Both implementations use `async/await` for WebRTC events and API calls.
- JavaScript handles media acquisition, SDP/ICE negotiation, and peer connection.

---

## 🧪 Getting Started

### Go + HTMX

1. **Clone the repo:**
    ```bash
    git clone https://github.com/vish419/webrtc-go-node-comparison.git
    cd webrtc-go-node-comparison/go-htmx-webrtc
    ```

2. **Navigate to server directory:**
    ```bash
    cd server
    ```

3. **Install dependencies:**
    ```bash
    go get -u github.com/gorilla/websocket
    ```

4. **Run the server:**
    ```bash
    go run main.go
    ```

5. **Open the client:**
    Open the `client/index.html` in a browser (WebRTC-compatible like Chrome or Firefox).

---

### Node.js

1. **Clone the repo:**
    ```bash
    git clone https://github.com/vish419/webrtc-go-node-comparison.git
    cd webrtc-go-node-comparison/nodejs-webrtc
    ```

2. **Navigate to server directory:**
    ```bash
    cd server
    ```

3. **Install dependencies:**
    ```bash
    npm install
    ```

4. **Run the server:**
    ```bash
    npm start
    ```

5. **Open the client:**
    Open the `public/index.html` file in a browser (WebRTC-compatible).

---

## 📁 Project Structure

