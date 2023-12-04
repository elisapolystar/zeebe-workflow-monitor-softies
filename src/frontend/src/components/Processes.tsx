import React, { useEffect, useState } from 'react';
import ReactDOM from 'react-dom/client';
import './Processes.css'; 
import BPMNView from './BPMNView.tsx';
import {format} from 'date-fns'import Instances from './Instances.tsx';


interface ProcessProps {
  socket: WebSocket | null;
  processes: string | null;
}

const Processes: React.FC<ProcessProps> = ({socket, processes}) => {
  const [bpmnData, setBpmnData] = useState<string | null>(null);
  const processesData = processes ? JSON.parse(processes) : [];

  const navigate = (navData: string) => {
    const view = navData.split('/');
    const path = `/${view[1]}`;
    const id = view[2];
    getComponentForPath(path, id)
  };

  const fetchBpmn = (id: string | undefined) => {
    if (socket && socket.readyState === WebSocket.OPEN) {
      const messageObject = `{ "process": "${id}" }`;
      socket.send(messageObject);
      console.log(`Process request ${messageObject} sent from frontend`);
    }
  };

  const fetchInstancesForProcess = (id: string | undefined) => {
    if (socket && socket.readyState === WebSocket.OPEN) {
      const messageObject = `{ "instances-for-process": "${id}" }`;
      socket.send(messageObject);
      console.log(`Instance request for process ${messageObject} sent from frontend`);
    }
  };

  const getComponentForPath = (path: string, id: string) => {
    switch (path) {

      case '/BPMNView':
        fetchBpmn(id);
        return;

      case '/instances':
        fetchInstancesForProcess(id);
        return;
    }
  };

  useEffect(() => {
    if (socket && socket.readyState === WebSocket.OPEN) {
      socket.addEventListener('message', (event) => {
        const message = JSON.parse(event.data);
        const type = message.type;
        let data;
        let path;

        switch(type) {
          case 'process':
            console.log(`Process recieved: ${message.data}`)
            setBpmnData(message.data);
            data = <BPMNView process={bpmnData}/>;
            path = '/BPMNView'
            break;
          
          case 'instances-for-process':
            console.log(`Instances for a process recieved: ${message.data}`)
            setInstances(message.data);
            data = <Instances socket={socket} instances={instancesData} />;
            path = '/instances';
            break;
          
          default: return;
        }
        window.history.pushState({}, '', path);
        const root = ReactDOM.createRoot(document.getElementById('content') as HTMLElement);
        root.render(data);
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
            <span onClick={() => navigate(`/Instances/${item.processDefinitionKey}`)}>{item.instances}</span>
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