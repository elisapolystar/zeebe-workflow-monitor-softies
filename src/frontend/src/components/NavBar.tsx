import React, { useEffect } from 'react';
import Processes from './Processes.tsx';
import Instances from './Instances.tsx';
import Incidents from './Incidents.tsx';
import './component styles/NavBar.css';


interface NavBarProps {
  socket: WebSocket | null;
  setContent: React.Dispatch<React.SetStateAction<JSX.Element | null>>;
}

const NavBar: React.FC<NavBarProps> = ({ socket, setContent }) => {

  // Send a request to backend for all processes.
  const fetchProcesses = () => {
    if (socket && socket.readyState === WebSocket.OPEN) {
      const messageObject = '{ "process": "" }';
      socket.send(messageObject);
    }
  };

  // Send a request to backend for all instances.
  const fetchInstances = () => {
    if (socket && socket.readyState === WebSocket.OPEN) {
      const messageObject = '{ "instance": "" }';
      socket.send(messageObject);
    }
  };

  // Send a request to backend for all incidents.
  const fetchIncidents = () => {
    if (socket && socket.readyState === WebSocket.OPEN) {
      const messageObject = '{ "incident": "" }';
      socket.send(messageObject);
    }
  };


  /**
   * Catch ws messages and set render with setContent accordingly.
   */
  useEffect(() => {
    
    if (socket && socket.readyState === WebSocket.OPEN) {
      let path;
      let data;
    
      socket.addEventListener('message', (event) => {

        const message = JSON.parse(event.data);
        const type = message.type;
        switch(type) {

          case 'all-processes':
            path = '/processes';
            data = <Processes socket={socket} processes={message.data} setContent={setContent} />
            break;
          
          case 'all-instances':
            path = '/instances';
            data = <Instances socket={socket} instances={message.data} setContent={setContent} />
            break;

          case 'all-incidents':
            path = '/incidents';
            data = <Incidents socket={socket} incidents={message.data} setContent={setContent} />
            break;
          
          default: return;
        }

        window.history.pushState({}, '', path);
        setContent(data);
      });
    }
  }, [socket]);

  return (
    <div>
      <nav>
        <ul id="NavBarComponent">
          <li onClick={() => fetchProcesses()}>Processes</li>
          <li onClick={() => fetchInstances()}>Instances</li>
          <li onClick={() => fetchIncidents()}>Incidents</li>
        </ul>
      </nav>
    </div>
  );
};

export default NavBar;


