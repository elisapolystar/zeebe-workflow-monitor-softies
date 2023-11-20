import React, { useEffect, useState } from 'react';
import './Instances.css';
import ReactDOM from 'react-dom/client';
import Instanceview from './Instanceview.tsx';
import data from "./instance.json";


const Instances: React.FC = () => {

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
  /*const [selectedkey, setSelectedkey] = useState<number | null>(null);*/

  useEffect(() => {
    setInstanceData(data);
  }, []);

  /*const handleInstanceClick = (processDefinitionKey: number) => {
    // Update the selected BPMN when an instance definition key is clicked
    setSelectedkey(processDefinitionKey);
  };*/

  return (
    <div className="instance-container">
      <div className="instance-item">
        <span>Instance Definition Key</span>
        {instancedata &&
          instancedata.map((instancedata, index) => (
            <div className="definition-key" key={index} 
            onClick={() => navigate('/Instanceview')}>{instancedata.value.processDefinitionKey}
            </div>
        ))}

      </div>
      <div className="instance-item">
        <span>BPMN Process Id</span>
        {instancedata &&
          instancedata.map((instancedata, index) => (
            <div className = "instance-info">
              <span>{instancedata.value.bpmnProcessId}</span>
            </div>
        ))}
        
      </div>
      <div className="instance-item">
        <span>State</span>
        {instancedata &&
          instancedata.map((instancedata, index) => (
            <div className="instance-info">
              <span>{instancedata.value.state}</span>
            </div>
        ))}
      </div>
      <div className="instance-item">
        <span>Time</span>
        {instancedata &&
          instancedata.map((instancedata, index) => ( 
            <div className="instance-info">
              <span>{instancedata.value.time}</span>
            </div>
        ))}
      </div>
    </div>
  );
};

export default Instances;