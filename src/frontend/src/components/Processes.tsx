import React, { useEffect } from 'react';
import './component styles/Processes.css'; 
import BPMNView from './BPMNView.tsx';
import { format } from 'date-fns';
import Instances from './Instances.tsx';
import ErrorDisplay from './ErrorDisplay.tsx'


interface ProcessProps {
  socket: WebSocket | null; // current websocket connection with backend
  processes: string | null; // message from backend  containing all processes in database
  setContent: React.Dispatch<React.SetStateAction<JSX.Element | null>>; // function to set next render
}

const Processes: React.FC<ProcessProps> = ({ socket, processes, setContent }) => {
    const processesData = processes ? JSON.parse(processes) : [];

    if(!processesData) {
      return (
        <div>
        <div className="process-container">
          <div className="process-item">
            <span>Process Definition Key</span>
          </div>
          <div className="process-item">
            <span>BPMN process id</span>
          </div>
          <div className="process-item">
            <span>Instances</span>
          </div>
          <div className="process-item">
            <span>Version</span>
          </div>
          <div className="process-item">
            <span>Time</span>
          </div>
        </div>
        <div className="not-found">
              <span>No processes found</span>
        </div>
        </div>
      );
    }

  /**
   * Send data request to backend for a specific process. 
   * @param processDefinitionKey of the process to request.
   */
  const fetchBpmn = (processDefinitionKey: string | undefined) => {
    if (socket && socket.readyState === WebSocket.OPEN) {
      const messageObject = `{ "process": "${processDefinitionKey}" }`;
      socket.send(messageObject);
    }
  };

  /**
   * Send data request to backend for all instances of specific process. 
   * @param processDefinitionKey of the process to request all instances of.
   */
  const fetchInstancesForProcess = (processDefinitionKey: string | undefined) => {
    if (socket && socket.readyState === WebSocket.OPEN) {
      const messageObject = `{ "process-instances": "${processDefinitionKey}" }`;
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

          case 'process':
            d = JSON.parse(message.data);
            // If/else to check for error messages and get the render data accordingly
            if(d.Error){ data = <ErrorDisplay errorMessage={message.data}/>}
            else data = <BPMNView process={message.data}/>;
            path = '/BPMNView'
            break;
     
          case 'instances-for-process':
            data = <Instances socket={socket} instances={message.data} setContent={setContent} />;
            path = '/instances';
            break;
          
          default: return;
        }
        window.history.pushState({}, '', path);
        setContent(data);
      });
    }
  }, [socket]);

  return (
    <div className="process-container">
      <div className="process-item">
        <span>Process Definition Key</span>
        {processesData.map((item, index) => (
          <div className="process-key" key={index}>
            <span onClick={() => fetchBpmn(item.key)}>{item.key}</span>
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
          <div className="process-key" key={index}>
            <span onClick={() => fetchInstancesForProcess(item.key)}>{item.instances}</span>
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