import React, { useEffect, useState } from 'react';
import ReactDOM from 'react-dom/client';
import './Processes.css'; 
import data from "./test.json";
import BPMNView from './BPMNView.tsx';
import Instances from './Instances.tsx';


interface ProcessProps {
  socket: WebSocket | null;
  processes: string | null;
}

const Processes: React.FC<ProcessProps> = ({socket}) => {
  const [bpmnData, setBpmnData] = useState<string | null>(null);
  const [instancesData, setInstances] = useState<string | null>(null);

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
        return bpmnData ? <BPMNView process={bpmnData} /> : <div>Loading...</div>;

      case '/instances':
        fetchInstancesForProcess(id);
        return instancesData ? <Instances socket={socket} instances={instancesData} /> : <div>Loading...</div>;

      default:
        return <div>Not Found</div>;
    }
  };

  useEffect(() => {
    if (socket && socket.readyState === WebSocket.OPEN) {
      socket.addEventListener('message', (event) => {
        const message = JSON.parse(event.data);
        const type = message.type;

        switch(type) {
          case 'process':
            console.log(`Process recieved: ${message.data}`)
            setBpmnData(message.data);
            return;
          
          case 'instances-for-process':
            console.log(`Instances for a process recieved: ${message.data}`)
            setInstances(message.data);
            return;
        }
      });
    }
  }, [socket]);

  return (
    <div className="process-container">
      <div className="process-item">
        <span>Process Definition Key</span>
        {data.map((item, index) => (
          <div className="process-key" key={index}>
            <span onClick={() => navigate(`/BPMNView/${item.processDefinitionKey}`)}>{item.processDefinitionKey}</span>
          </div>
        ))}
      </div>
      <div className="process-item">
        <span>BPMN process id</span>
        {data.map((item, index) => (
          <div className="process-info" key={index}>
            <span>{item.bpmnProcessId}</span>
          </div>
        ))}
      </div>
      <div className="process-item">
        <span>Instances</span>
        {data.map((item, index) => (
          <div className="process-info" key={index}>
            <span onClick={() => navigate(`/Instances/${item.processDefinitionKey}`)}>{item.instances}</span>
          </div>
        ))}
      </div>
      <div className="process-item">
        <span>Version</span>
        {data.map((item, index) => (
          <div className="process-info" key={index}>
            <span>{item.version}</span>
          </div>
        ))}
      </div>
      <div className="process-item">
        <span>Time</span>
        {data.map((item, index) => (
          <div className="process-info" key={index}>
            <span>{item.time}</span>
          </div>
        ))}
      </div>

    </div>
    
  );
};

export default Processes;

