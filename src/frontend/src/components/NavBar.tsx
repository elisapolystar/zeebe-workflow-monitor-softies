import React, { useState, useEffect } from 'react';
import ReactDOM, { Root } from 'react-dom/client';
import Processes from './Processes.tsx';
import Instances from './Instances.tsx';
import Incidents from './Incidents.tsx';
import './NavBar.css';
import App from '../App.tsx';


interface NavBarProps {
  socket: WebSocket | null;
}

const NavBar: React.FC<NavBarProps> = ({ socket }) => {
  const [processesData, setProcesses] = useState<string | null>(null);
  const [instancesData, setInstances] = useState<string | null>(null);
  const [incidentsData, setIncidents] = useState<string | null>(null);

  const fetchProcesses = () => {
    if (socket && socket.readyState === WebSocket.OPEN) {
      const messageObject = '{ "process": "" }';
      socket.send(messageObject);
      console.log(`Process request ${messageObject} sent from frontend`);
    }
  };

  const fetchInstances = () => {
    if (socket && socket.readyState === WebSocket.OPEN) {
      const messageObject = '{ "instance": "" }';
      socket.send(messageObject);
      console.log(`Instance request ${messageObject} sent from frontend`);
    }
  };

  const fetchIncidents = () => {
    if (socket && socket.readyState === WebSocket.OPEN) {
      const messageObject = '{ "incident": "" }';
      socket.send(messageObject);
      console.log(`Incident request ${messageObject} sent from frontend`);
    }
  };

  const navigate = async (path: string) => {
    getComponentForPath(path);
    return;
  };

  const getComponentForPath = (path: string) => {
    switch (path) {
      case '/processes':
        fetchProcesses();
        return 
      case '/instances':
        fetchInstances();
        return
      case '/incidents':
        fetchIncidents();
        return
        //fetchInstances(undefined);
        //return incidentsData ? <Incidents socket={socket} incidents={incidentsData} /> : <div>Loading...</div>;

      default:
        if(!processesData) fetchProcesses();
    }
  };

  useEffect(() => {
    console.log('checking connection');
    if (socket && socket.readyState === WebSocket.OPEN) {
      console.log('connection OK');
      let path;
      let data;
    
      socket.addEventListener('message', (event) => {

        const message = JSON.parse(event.data);
        const type = message.type;
        switch(type) {
          
          case 'all-processes':
            console.log(`Processes recieved: ${message.data}`)
            path = '/processes';
            data = <Processes socket={socket} processes={message.data} />
            break;
          
          case 'all-instances':
            console.log(`Instances recieved: ${message.data}`)
            path = '/instances';
            data = <Instances socket={socket} instances={message.data} />
            break;

          case 'all-incidents':
            console.log(`Incidents recieved: ${message.data}`)
            path = '/incidents';
            data = <Incidents socket={socket} incidents={message.data} />
            break;
          
          default: return;
        }
        console.log('Trying to render content')
        window.history.pushState({}, '', path);
        ReactDOM.createRoot(document.getElementById('content') as HTMLElement).render(data);
        
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
      </div>
    </div>
  );
};

export default NavBar;