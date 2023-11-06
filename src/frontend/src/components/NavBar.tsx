import React, { useState } from 'react';
import ReactDOM from 'react-dom/client';
import './NavBar.css';
import Processes from './Processes.tsx';
import Instances from './Instances.tsx';
import Incidents from './Incidents.tsx';


/*

{
process
[

{
Process-definition-key: 1111111111
BPMN-process-id: order1
instances: 1
Version: 1/1
Time: 11/11/1111
}

{
Process-definition-key: 222222222222
BPMN-process-id: order2
instances: 2
Version: 2/2
Time: 22/22/2222
}
"timestamp":1698688538626,
]
}
*/
interface NavBarProps {
  socket: WebSocket | null;
}


const NavBar: React.FC<NavBarProps> = ({socket}) => {
  const [processes, setProcesses] = useState([]); // State for processes

  const navigate = (path: string) => {
    window.history.pushState({}, '', path);
    ReactDOM.createRoot(document.getElementById('content') as HTMLElement).render(getComponentForPath(path));
    return getComponentForPath(path) ;
  };

  const getComponentForPath = (path: string) => {
    switch (path) {
      case '/processes':
        fetchProcesses(socket, undefined);
        return <Processes /*processes={processes}*/ />;
      case '/instances':
        return <Instances />;
      case '/incidents':
        return <Incidents />;
      default:
        return <div>Not Found</div>;
    }
  };

  const fetchProcesses = (socket, id) => {
    console.log("fetchProocesses");
    if (socket) {
      let messageObject;
      
      if(id === undefined){
        messageObject = `{
          process: "", 
        }`;
      }else{
        messageObject = `{
          process: "${id}", 
        }`;
      }

      // Send a WebSocket message to request processes from the backend
      socket.send(messageObject);
      console.log("Process request sent from frontend");

      // Handle the response from the backend
      socket.addEventListener('message', (event) => {
        const data = JSON.parse(event.data);
        setProcesses(data);
      });
    }
  };

  return (
    <div>
      <nav>
        <ul>
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
