import React, { useEffect, useState } from 'react';
import './Instances.css';
import ReactDOM from 'react-dom/client';
import Instanceview from './Instanceview.tsx';
import { format } from 'date-fns';

interface InstanceProps {
  socket: WebSocket | null;
  instances: string | null;
}
interface InstanceProps {
  socket: WebSocket | null;
  instances: string | null;
}

const Instances: React.FC<InstanceProps> = ({socket, instances}) => {
  const [bpmnData, setBpmnData] = useState<string | null>(null);
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
  const getComponentForPath = (path: string, id: string) => {
    switch (path) {
      case '/Instanceview':
        fetchInstance(id);
        return instanceData ? <Instanceview instance={instanceData} /> : <div>Loading...</div>;
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

  let formattedDate = '';
  const date = new Date(instanceData[0].value.Timestamp);
  formattedDate = format(date, 'dd-MM-yyyy HH:mm:ss');

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
              <span>{item.bpmnProcessId}</span>
            </div>
        ))}
        
      </div>
      <div className="instance-item">
        <span>State</span>
        {instancesData.map((item, index) => (
            <div className="instance-info" key={index}>
              <span>{item.state}</span>
            </div>
        ))}
      </div>
      <div className="instance-item">
        <span>Time</span>
        {instancesData.map((item, index) => ( 
            <div className="instance-info" key={index}>
              <span>{format(new Date(item.timestamp), 'dd-MM-yyyy HH:mm:ss')}</span>
            </div>
        ))}
      </div>

      <div className="instance-item">
        <span>Process Definition Key</span>
        {instancesData.map((item, index) => (
            <div className="definition-key" key={index}>
              <span>{item.processDefinitionKey}</span>
            </div>
        ))}
    </div>
    </div>
  );
};

export default Instances;