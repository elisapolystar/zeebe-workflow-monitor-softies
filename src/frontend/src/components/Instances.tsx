import React, { useEffect, useState } from 'react';
import './Instances.css';
import ReactDOM from 'react-dom/client';
import Instanceview from './Instanceview.tsx';
import { format } from 'date-fns';

interface InstanceProps {
  socket: WebSocket | null;
  instances: string | null;
}

const Instances: React.FC<InstanceProps> = ({socket, instances}) => {
  console.log('Instances prop:', instances);
  const [bpmnData, setBpmnData] = useState<string | null>(null);
  const instancesData = instances ? JSON.parse(instances) : [];

  const navigate = (path: string) => {
    const view = path.split('/');
    console.log(view);
    window.history.pushState({}, '', view[1]);
    ReactDOM.createRoot(document.getElementById('content') as HTMLElement).render(getComponentForPath(`/${view[1]}`, view[2]));
  };

  const fetchBpmn = (id: string | undefined) => {
    if (socket && socket.readyState === WebSocket.OPEN) {
      const messageObject = `{ "process_instance": "${id}" }`;
      socket.send(messageObject);
      console.log(`Process_instance request ${messageObject} sent from frontend`);
    }
  };

  const getComponentForPath = (path: string, id: string) => {
    switch (path) {
      case '/Instanceview':
        fetchBpmn(id);
        return bpmnData ? <Instanceview process_instance={bpmnData} /> : <div>Loading...</div>;

      default:
        return <div>Not Found</div>;
    }
  };

  useEffect(() => {
    if (socket && socket.readyState === WebSocket.OPEN) {
      socket.addEventListener('message', (event) => {
        const message = JSON.parse(event.data);

        if(message.type === 'all-instances' ){
          setBpmnData(message.data);
        }
      });
    }
  }, [socket]);
  return (
    <div className="instance-container">
      <div className="instance-item">
        <span>Process Instance Key</span>
        {instancesData.map((item, index) => (
            <div className="definition-key" key={index}>
              <span onClick={() => navigate('/Instanceview')}>{item.ProcessInstanceKey}</span>
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