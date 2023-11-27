import React, { useState, useEffect} from 'react';
import NavBar from './components/NavBar.tsx';

const App: React.FC = () => {

  const [socket, setSocket] = useState<WebSocket | null>(null);
  const [messages, setMessages] = useState<string[]>([]);
  let connection = false;

  useEffect(() => {
    const newSocket = new WebSocket("ws://localhost:8001/ws");

    if(!connection){
    newSocket.onopen = () => {
      console.log("Successfully Connected");
      
      //newSocket.send('{ "process": "" }');
      setSocket(newSocket);
      connection = true;
    };

    newSocket.onclose = (event) => {
      console.log("Socket closed connection: ", event);
      connection = false;
    };

    newSocket.onerror = (error) => {
      console.log("Socket Error: ", error);
    };

    newSocket.addEventListener('message', (event) => {
      const message = event.data;
      //setMessages(message);
    });

    return () => {
      newSocket.close();
    };
  };
  }, []);


  return (
    <div>
      <NavBar socket={socket}/>
    </div>
  );
};

export default App;