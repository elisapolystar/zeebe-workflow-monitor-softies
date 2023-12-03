import React, { useEffect, useState } from 'react';
import ReactDOM from 'react-dom/client';
import './Processes.css'; 
import BPMNView from './BPMNView.tsx';
import {format} from 'date-fns'

interface ProcessProps {
  socket: WebSocket | null;
  processes: string | null;
}

const Processes: React.FC<ProcessProps> = ({socket, processes}) => {
  const [bpmnData, setBpmnData] = useState<string | null>(null);
  const processesData = processes ? JSON.parse(processes) : [];

  const navigate = (path: string) => {
    const view = path.split('/');
    console.log(view);
    window.history.pushState({}, '', view[1]);
    ReactDOM.createRoot(document.getElementById('content') as HTMLElement).render(getComponentForPath(`/${view[1]}`, view[2]));
  };

  const fetchBpmn = (id: string | undefined) => {
    if (socket && socket.readyState === WebSocket.OPEN) {
      const messageObject = `{ "process": "${id}" }`;
      socket.send(messageObject);
      console.log(`Process request ${messageObject} sent from frontend`);
    }
  };

  const getComponentForPath = (path: string, id: string) => {
    switch (path) {
      case '/BPMNView':
        fetchBpmn(id);
        return bpmnData ? <BPMNView process={bpmnData} /> : <div>Loading...</div>;

      default:
        return <div>Not Found</div>;
    }
  };

  useEffect(() => {
    if (socket && socket.readyState === WebSocket.OPEN) {
      socket.addEventListener('message', (event) => {
        const message = JSON.parse(event.data);

        if(message.type === 'all-processes' ){
          setBpmnData(message.data);
        }
      });
    }
  }, [socket]);

  return (
    <div className="process-container">
      <div className="process-item">
        <span>Process Definition Key</span>
        {processesData.map((item, index) => (
          <div className="process-key" key={index}>
            <span onClick={() => navigate(`/BPMNView/${item.key}`)}>{item.key}</span>
          </div>
        ))}
      </div>
      <div className="process-item">
        <span>BPMN process id</span>
        {processesData.map((item, index) => (
          <div className="process-info" key={index}>
            <span>{item.bpmnProcessId}</span>
          </div>
        ))}
      </div>
      <div className="process-item">
        <span>Instances</span>
        {processesData.map((item, index) => (
          <div className="process-info" key={index}>
            <span>{item.instances}</span>
          </div>
        ))}
      </div>
      <div className="process-item">
        <span>Version</span>
        {processesData.map((item, index) => (
          <div className="process-info" key={index}>
            <span>{item.version}</span>
          </div>
        ))}
      </div>
      <div className="process-item">
        <span>Time</span>
        {processesData.map((item, index) => (
          <div className="process-info" key={index}>
            <span>{format(new Date(item.timestamp), 'dd-MM-yyyy HH:mm:ss')}</span>
          </div>
        ))}
      </div>

    </div>
    
  );
};

export default Processes;