import React, { useState, useEffect } from 'react';
import io from 'socket.io-client';
import MessageDisplayer from './messageDisplayer'; 

const MessageListener = () => {
    const [messages, setMessages] = useState([]);

    useEffect(() => {
        // Update the URL once known
        const socket = io("ws://localhost:8001/ws");

        // Define event handlers for messages received from the server.
        socket.on('message', (message) => {
            // Handle incoming messages here and update your component state.
            setMessages((prevMessages) => [...prevMessages, message]);
        });

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