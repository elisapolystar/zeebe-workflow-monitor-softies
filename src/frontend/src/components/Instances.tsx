import React, { useEffect, useState } from 'react';
import './Instances.css';
import Instanceview from './Instanceview.tsx';
import { format } from 'date-fns';

interface InstanceProps {
  socket: WebSocket | null;
  instances: string | null;
  setContent: React.Dispatch<React.SetStateAction<JSX.Element | null>>;
}

const Instances: React.FC<InstanceProps> = ({socket, instances, setContent}) => {
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

  const fetchBpmn = (id: string | undefined) => {
    if (socket && socket.readyState === WebSocket.OPEN) {
      const messageObject = `{ "process": "${id}" }`;
      socket.send(messageObject);
      console.log(`Process request ${messageObject} sent from frontend`);
    }
  };

  const getComponentForPath = (path: string, id: string) => {
    switch (path) {
      case '/Instanceview':
        fetchInstance(id);
        return;

      case '/BPMNView':
        fetchBpmn(id);
        return;
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
            path = '/Instanceview';
            data = <Instanceview process_instance={instanceData} />;
            break;

          default: return;
        }
        setContent(data);
      });
    }
  }, [socket]);

  return (
    <div className="instance-container">
      <div className="instance-item">
        <span>Process Instance Key</span>
        {instancesData.map((item, index) => (
            <div className="definition-key" key={index}>
              <span onClick={() => navigate(`/Instanceview/${item.ProcessInstanceKey}`)}>{item.ProcessInstanceKey}</span>
            </div>
        ))}
      </div>

      <div className="instance-item">
        <span>BPMN Process Id</span>
        {instancesData.map((item,index) => (
            <div className = "instance-info" key={index}>
              <span>{item.BpmnProcessId}</span>
            </div>
        ))}
        
      </div>
      <div className="instance-item">
        <span>Active</span>
        {instancesData.map((item, index) => (
            <div className="instance-info" key={index}>
              <span>{item.Active.toString()}</span>
            </div>
        ))}
      </div>
      <div className="instance-item">
        <span>Time</span>
        {instancesData.map((item, index) => ( 
            <div className="instance-info" key={index}>
              <span>{format(new Date(item.Timestamp), 'dd-MM-yyyy HH:mm:ss')}</span>
            </div>
        ))}
      </div>

      <div className="instance-item">
        <span>Process Definition Key</span>
        {instancesData.map((item, index) => (
            <div className="definition-key" key={index}>
              <span onClick={() => navigate(`/BPMNView/${item.ProcessDefinitionKey}`)}>{item.ProcessDefinitionKey}</span>
            </div>
        ))}
    </div>
    </div>
  );
};

export default Instances;