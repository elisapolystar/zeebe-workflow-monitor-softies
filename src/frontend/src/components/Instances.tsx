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

const Instances: React.FC<InstanceProps> = ({socket}) => {

  const navigate = (path: string) => {
    window.history.pushState({}, '', path);
    ReactDOM.createRoot(document.getElementById('content') as HTMLElement).render(getComponentForPath(path));
  };

  const getComponentForPath = (path: string) => {
    switch (path) {
      case '/Instanceview':
        return <Instanceview />;
    }
  };

  const [instancedata, setInstanceData] = useState(data);

  useEffect(() => {
    setInstanceData(data);
  }, []);

  let formattedDate = '';
  const date = new Date(instancedata[0].value.Timestamp);
  formattedDate = format(date, 'dd-MM-yyyy HH:mm:ss');

  return (
    <div className="instance-container">
      <div className="instance-item">
        <span>Process Instance Key</span>
        {instancedata &&
          instancedata.map((instancedata, index) => (
            <div className="definition-key" key={index}>
              <span onClick={() => navigate('/Instanceview')}>{instancedata.value.processInstanceKey}</span>
            </div>
        ))}
      </div>

      <div className="instance-item">
        <span>BPMN Process Id</span>
        {instancedata &&
          instancedata.map((instancedata) => (
            <div className = "instance-info">
              <span>{instancedata.value.bpmnProcessId}</span>
            </div>
        ))}
        
      </div>
      <div className="instance-item">
        <span>State</span>
        {instancedata &&
          instancedata.map((instancedata) => (
            <div className="instance-info">
              <span>{instancedata.value.state}</span>
            </div>
        ))}
      </div>
      <div className="instance-item">
        <span>Time</span>
        {instancedata &&
          instancedata.map((instancedata) => ( 
            <div className="instance-info">
              <span>{formattedDate}</span>
            </div>
        ))}
      </div>

      <div className="instance-item">
        <span>Process Definition Key</span>
        {instancedata &&
          instancedata.map((instancedata) => (
            <div className="definition-key">
              <span>{instancedata.value.parentProcessInstanceKey}</span>
            </div>
        ))}
    </div>
    </div>
  );
};

export default Instances;