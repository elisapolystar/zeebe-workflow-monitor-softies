import React, { useEffect, useState } from 'react';
import './Instances.css';
import ReactDOM from 'react-dom/client';
import Instanceview from './Instanceview.tsx';
import data from "./instance.json";
import { format } from 'date-fns';

interface InstanceProps {
  socket: WebSocket | null;
  instances: string | null;
}

const Instances: React.FC<InstanceProps> = ({socket, instances}) => {
  const [instanceData, setInstanceData] = useState<string | null>(null);
  const instancesData = instances ? JSON.parse(instances) : [];

  const navigate = (path: string) => {
    const view = path.split('/');
    console.log(view);
    getComponentForPath(`/${view[1]}`, view[2])
  };

  const fetchInstance = (id: string | undefined) => {
    if (socket && socket.readyState === WebSocket.OPEN) {
      const messageObject = `{ "instance": "${id}" }`;
      socket.send(messageObject);
      console.log(`Instance request for process ${messageObject} sent from frontend`);
    }
  };

  const getComponentForPath = (path: string, id: string) => {
    switch (path) {
      case '/Instanceview':
        fetchInstance(id);
        return instanceData ? <Instanceview process_instance={instanceData} /> : <div>Loading...</div>;
    }
  };

  useEffect(() => {
    if (socket && socket.readyState === WebSocket.OPEN) {
      socket.addEventListener('message', (event) => {
        const message = JSON.parse(event.data);
        const type = message.type;
        let path;
        let data;

        switch(type) {
          case 'instances':
            console.log(`Data for an instance recieved: ${message.data}`)
            setInstanceData(message.data);
            path = '/instances';
            data = <Instances socket={socket} instances={instanceData} />
            break;

          default: return;
        }
        window.history.pushState({}, '', path);
        ReactDOM.createRoot(document.getElementById('content') as HTMLElement).render(data);
      });
    }
  }, [socket]);

  return (
    <div className="instance-container">
      <div className="instance-item">
        <span>Process Instance Key</span>
        {instancesData.map((instancedata, index) => (
            <div className="definition-key" key={index}>
              <span onClick={() => navigate(`/Instanceview/${instancedata.ProcessInstanceKey}`)}>{instancedata.ProcessInstanceKey}</span>
            </div>
        ))}
      </div>

      <div className="instance-item">
        <span>BPMN Process Id</span>
        {instancesData.map((instancedata) => (
            <div className = "instance-info">
              <span>{instancedata.BpmnProcessId}</span>
            </div>
        ))}
        
      </div>
      <div className="instance-item">
        <span>Active</span>
        {instancesData.map((instancedata) => (
            <div className="instance-info">
              <span>{instancedata.Active.toString()}</span>
            </div>
        ))}
      </div>
      <div className="instance-item">
        <span>Time</span>
        {instancesData.map((instancedata) => ( 
            <div className="instance-info">
              <span>{format(new Date(instancedata.Timestamp), 'dd-MM-yyyy HH:mm:ss')}</span>
            </div>
        ))}
      </div>

      <div className="instance-item">
        <span>Process Definition Key</span>
        {instancesData.map((instancedata) => (
            <div className="definition-key">
              <span>{instancedata.ProcessDefinitionKey}</span>
            </div>
        ))}
    </div>
    </div>
  );
};

export default Instances;