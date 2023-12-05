import React, { useState, useEffect} from 'react';
import NavBar from './components/NavBar.tsx';

const App: React.FC = () => {

  const [socket, setSocket] = useState<WebSocket | null>(null);
  const [content, setContent] = useState<JSX.Element | null>(null);

  useEffect(() => {
    const newSocket = new WebSocket("ws://localhost:8001/ws");

    newSocket.onopen = () => {
      console.log("Successfully Connected");
      setSocket(newSocket);
    };

    newSocket.onclose = (event) => {
      console.log("Socket closed connection: ", event);
    };

    newSocket.onerror = (error) => {
      console.log("Socket Error: ", error);
    };

    return () => {
      newSocket.close();
    };
  }, []);

  return (
    <div>
      <NavBar socket={socket} setContent={setContent} />
      <div className="content">{content}</div>
    </div>
  );
};

export default App;