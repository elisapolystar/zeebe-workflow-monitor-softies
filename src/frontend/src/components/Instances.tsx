import React, { useEffect } from 'react';
import './component styles/Instances.css';
import Instanceview from './Instanceview.tsx';
import ErrorDisplay from './ErrorDisplay.tsx';
import { format } from 'date-fns';

interface InstanceProps {
  socket: WebSocket | null; // current websocket connection with backend
  instances: string | null; // message from backend  containing data of instances 
  setContent: React.Dispatch<React.SetStateAction<JSX.Element | null>>; // function to set next render
}

const Instances: React.FC<InstanceProps> = ({socket, instances, setContent}) => {
  const instancesData = instances ? JSON.parse(instances) : {};

  //If there is no instances, print only the headers and text "no instances found"
  if(!instancesData) {
    return (
      <div>
      <div className="instance-container">
        <div className="process-item">
          <span>Process Definition Key</span>
        </div>
        <div className="instance-item">
          <span>BPMN process id</span>
        </div>
        <div className="instance-item">
          <span>Active</span>
        </div>
        <div className="instance-item">
          <span>Time</span>
        </div>
        <div className="instance-item">
          <span>Process Definition Key</span>
        </div>
      </div>
      <div className="not-found">
            <span>No instances found</span>
      </div>
      </div>
      
    );
    }

  /**
   * Send a request to backend for a specific instance using process instance key.
   * @param processInstanceKey 
   */
  const fetchInstance = (processInstanceKey: string | undefined) => {
    if (socket && socket.readyState === WebSocket.OPEN) {
      const messageObject = `{ "instance": "${processInstanceKey}" }`;
      socket.send(messageObject);
    }
  };

  /**
   * Send a request to backend for a specific process using process definition key.
   * @param processDefinitionKey 
   */
  const fetchBpmn = (processDefinitionKey: string | undefined) => {
    if (socket && socket.readyState === WebSocket.OPEN) {
      const messageObject = `{ "process": "${processDefinitionKey}" }`;
      socket.send(messageObject);
    }
  };

  // Handle the incoming Websocket messages from backend 
  useEffect(() => {
    if (socket && socket.readyState === WebSocket.OPEN) {
      socket.addEventListener('message', (event) => {
        const message = JSON.parse(event.data);
        const type = message.type;
        let data;
        let path;
        let d;

        switch(type) {
          case 'instance':
            d = JSON.parse(message.data);
            if(d.Error){ data = <ErrorDisplay errorMessage={message.data}/>}
            path = '/Instanceview';
            data = <Instanceview instance={message.data} />;
            break;

          default: return;
        }
        window.history.pushState({}, '', path);
        setContent(data);
      });
    }
  }, [socket]);

  //If there is instances, print the headers and instance data
  return (
    <div className="instance-container">
      <div className="instance-item">
        <span>Process Instance Key</span>
        {instancesData.map((item, index) => (
            <div className="definition-key" key={index}>
              <span onClick={() => fetchInstance(item.ProcessInstanceKey)}>{item.ProcessInstanceKey}</span>
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
              <span onClick={() => fetchBpmn(item.ProcessDefinitionKey)}>{item.ProcessDefinitionKey}</span>
            </div>
        ))}
    </div>
    </div>
  );
};

export default Instances;