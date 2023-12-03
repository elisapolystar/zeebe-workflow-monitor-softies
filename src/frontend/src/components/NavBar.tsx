import React, { useState, useEffect } from 'react';
import ReactDOM from 'react-dom/client';
import Processes from './Processes.tsx';
import Instances from './Instances.tsx';
import Incidents from './Incidents.tsx';
import './NavBar.css';

interface NavBarProps {
  socket: WebSocket | null;
}

const NavBar: React.FC<NavBarProps> = ({ socket }) => {
  const [processesData, setProcesses] = useState<string | null>(null);

  const fetchProcesses = (id: string | undefined) => {
    if (socket && socket.readyState === WebSocket.OPEN) {
      const messageObject = !id ? '{ "process": "" }' : `{ "process": "${id}" }`;
      socket.send(messageObject);
      console.log(`Process request ${messageObject} sent from frontend`);
    }
  };

  const fetchInstances = (id: string | undefined) => {
    if (socket && socket.readyState === WebSocket.OPEN) {
      const messageObject = `{ "instance": "${id}" }`;
      socket.send(messageObject);
      console.log(`Instance request ${messageObject} sent from frontend`);
    }
  };

  const navigate = (path: string) => {
    window.history.pushState({}, '', path);
    ReactDOM.createRoot(document.getElementById('content') as HTMLElement).render(getComponentForPath(path));
  };

  const getComponentForPath = (path: string) => {
    switch (path) {
      case '/processes':
        fetchProcesses(undefined);
        return processesData ? <Processes socket={socket} processes={processesData} /> : <div>Loading...</div>;
      case '/instances':
        fetchInstances(undefined);
        return <Instances /*instances={instances}*/ />;
      case '/incidents':
        return <Incidents />;
      default:
        if(!processesData) fetchProcesses(undefined);
        return processesData ? <Processes socket={socket} processes={processesData} /> : <div>Loading...</div>;
    }
  };

  useEffect(() => {
    if (socket && socket.readyState === WebSocket.OPEN) {
      socket.addEventListener('message', (event) => {
        console.log('Recieved a message from backend!');

        const message = JSON.parse(event.data);
        console.log(`Process ${message}`)
        if(message.type === 'process'){
          console.log(`Process recieved: ${message.data}`)
          setProcesses(message.data);
        }
      });
    }
  }, [socket]);

  return (
    <div>
      <nav>
        <ul id="NavBarComponent">
          <li onClick={() => navigate('/processes')}>Processes</li>
          <li onClick={() => navigate('/instances')}>Instances</li>
          <li onClick={() => navigate('/incidents')}>Incidents</li>
        </ul>
      </nav>
      <div id="content">
        {getComponentForPath(window.location.pathname)}
      </div>
    </div>
  );
};

export default NavBar;


