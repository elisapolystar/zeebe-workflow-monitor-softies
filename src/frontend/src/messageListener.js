import React, { useState, useEffect } from 'react';
import io from 'socket.io-client';
import MessageDisplayer from './messageDisplayer'; 

const MessageListener = () => {
    const [messages, setMessages] = useState([]);

    useEffect(() => {
        let socket = new WebSocket("ws://localhost:8001/ws")
        console.log("Attempting Websocket Connection")

        socket.onopen = () => {
            console.log("Successfully Connected");
            socket.send("Hi from the client!")
        }

        socket.onclose = (event) => {
            console.log("Socket closed connection: ", event)

        }

        socket.onerror = (error) => {
            console.log("SOcket Error: ", error)
        }

        socket.addEventListener('message', (event) => {
                setMessages((prevMessages) => [...prevMessages, event.data]);
        })

        // Clean up the WebSocket connection when the component unmounts.
        return () => {
            socket.disconnect();
        };
    }, []);

    return (
        <div>
            <MessageDisplayer messages={messages} /> 
        </div>
    );
};

export default MessageListener;