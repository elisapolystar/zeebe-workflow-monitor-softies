import React, { useState, useEffect} from 'react';
import NavBar from './components/NavBar.tsx';

const App: React.FC = () => {

  const [socket, setSocket] = useState<WebSocket | null>(null);

  useEffect(() => {
    const newSocket = new WebSocket("ws://localhost:8001/ws");

    newSocket.onopen = () => {
      console.log("Successfully Connected");
      //newSocket.send('Hi from the client!');
      setSocket(newSocket);
    };

    newSocket.onclose = (event) => {
      console.log("Socket closed connection: ", event);
    };

    newSocket.onerror = (error) => {
      console.log("Socket Error: ", error);
    };

    newSocket.addEventListener('message', (event) => {
      //const message = event.data;
      //console.log(message);
    });

    return () => {
      newSocket.close();
    };
  }, []);


  return (
    <div>
      <NavBar socket={socket}/>
    </div>
  );
};

export default App;