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
  const [processes, setProcesses] = useState([]);
  const [instances, setInstances] = useState([]);

  const fetchProcesses = (id: string | undefined) => {
    if (socket && socket.readyState === WebSocket.OPEN) {
      const messageObject = !id
        ? '{ "process": "" }'
        : `{ "process": "${id}" }`;

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

  useEffect(() => {
    const responseListener = (event: MessageEvent) => {
      const data = JSON.parse(event.data);
      if (data.process) {
        setProcesses(data.process);
      } else if (data.instance) {
        setInstances(data.instance);
      }
    };

    if (socket && socket.readyState === WebSocket.OPEN) {
      socket.addEventListener('message', responseListener);
    }

    return () => {
      if (socket && socket.readyState === WebSocket.OPEN) {
        socket.removeEventListener('message', responseListener);
      }
    };
  }, [socket]);

  const navigate = (path: string) => {
    window.history.pushState({}, '', path);
    ReactDOM.createRoot(document.getElementById('content') as HTMLElement).render(getComponentForPath(path));
    return getComponentForPath(path);
  };

  const getComponentForPath = (path: string) => {
    switch (path) {
      case '/processes':
        fetchProcesses(undefined);
        return <Processes />;
      case '/instances':
        fetchInstances(undefined);
        return <Instances />;
      case '/incidents':
        return <Incidents />;
      default:
        fetchProcesses(undefined);
        return <Processes />;
    }
  };

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
