import React from 'react';
import ReactDOM from 'react-dom/client';
import './Processes.css'; 
import data from "./test.json";
import BPMNView from './BPMNView.tsx';

/*   
https://blog.logrocket.com/creating-react-sortable-table/     
https://mui.com/material-ui/react-accordion/*/


const Processes: React.FC = () => {
  const navigate = (path: string) => {
    window.history.pushState({}, '', path);
    ReactDOM.createRoot(document.getElementById('content') as HTMLElement).render(getComponentForPath(path));
  };

  const getComponentForPath = (path: string) => {
    switch (path) {
      case '/BPMNView':
        return <BPMNView />;

      default:
        return <div>Not Found</div>;
    }
  };

  return (
    <div className="process-container">
      <div className="process-item">
        <span>Process Definition Key</span>
        {data.map((item, index) => (
          <div className="process-key" key={index}>
            <span onClick={() => navigate('/BPMNView')}>{item.processDefinitionKey}</span>
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
            <span>{item.instances}</span>
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

