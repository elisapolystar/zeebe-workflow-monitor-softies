import React, { useState, useEffect } from 'react';
import MessageDisplayer from './messageDisplayer';
/*
interface MessageListenerProps {
  onSocketOpen: (socket: WebSocket) => void;
}
*/
const MessageListener: React.FC/*<MessageListenerProps>*/ = (/*{ onSocketOpen }*/) => {
  const [messages, setMessages] = useState<string[]>([]);
  
  useEffect(() => {
    let socket = new WebSocket("ws://localhost:8001/ws");
    console.log("Attempting WebSocket Connection");

    socket.onopen = () => {
      console.log("Successfully Connected");
      //onSocketOpen(socket);
      socket.send("Hi from the client!");
    }

    socket.onclose = (event) => {
      console.log("Socket closed connection: ", event);
    }

    socket.onerror = (error) => {
      console.log("Socket Error: ", error);
    }

    socket.addEventListener('message', (event) => {
      setMessages((prevMessages) => [...prevMessages, event.data]);
    });

    // Clean up the WebSocket connection when the component unmounts.
    return () => {
      socket.close();
    };
  }, [/*onSocketOpen*/]);

  return (
    <div>
      <MessageDisplayer messages={messages} />
    </div>
  );
};

export default MessageListener;