// App.tsx or other container component
import React, { useState, useEffect } from 'react';
import NavBar from './components/NavBar.tsx';
import MessageListener from './messageListener.tsx';

const App: React.FC = () => {
  const [socket, setSocket] = useState<WebSocket | null>(null);

  const handleSocketOpen = (newSocket: WebSocket) => {
    setSocket(newSocket);
  };

  return (
    <div>
      <MessageListener onSocketOpen={handleSocketOpen} />
      <NavBar socket={socket} />
    </div>
  );
};

export default App;