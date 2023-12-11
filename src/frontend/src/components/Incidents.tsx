import React, { useEffect} from 'react';
import Instanceview from './Instanceview.tsx';
import './component styles/Incidents.css';
import { format } from 'date-fns';

interface IncidentProps {
  socket: WebSocket | null; // current websocket connection with backend
  incidents: string | null; // message from backend  containing all incidents in database
  setContent: React.Dispatch<React.SetStateAction<JSX.Element | null>>; // function to set next render
}

const Incidents: React.FC<IncidentProps> = ({socket, incidents, setContent}) => {
  const incidentdata = incidents ? JSON.parse(incidents) : [];

  //if there is no incidents, print only the headers and text "no incidents"
  if(!incidentdata){
    return (
      <div>
      <div className="incident-container">
        <div className="incident-item">
          <span>BPMN process id</span>
        </div>
        <div className="incident-item">
          <span>Process Definition Key</span>
        </div>
        <div className="incident-item">
          <span>Error code</span>
        </div>
        <div className="incident-item">
          <span>Time</span>
        </div>
        <div className="incident-item">
          <span>Error message</span>
        </div>
      </div>
      <div className="not-found">
          <span>No incidents</span>
      </div>
      </div>
    );
  }

  /**
   * Send a request to backend for a specific instance using process instance key.
   * @param processInstanceKey 
   */
  const fetchInstance = (id: string | undefined) => {
    if (socket && socket.readyState === WebSocket.OPEN) {
      const messageObject = `{ "instance": "${id}" }`;
      socket.send(messageObject);
    }
  };

  // Handle the incoming Websocket messages from backend 
  useEffect(() => {
    if (socket && socket.readyState === WebSocket.OPEN) {
      socket.addEventListener('message', (event) => {
        const message = JSON.parse(event.data);
        const type = message.type;
        let path;
        let data;

        switch(type) {
          case 'instances':
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
    
  //if there is incidents, print the headers and the data
  return(
    <div className="incident-container">
      <div className="incident-item">
      <span>Process Instance Key</span>
        {incidentdata && incidentdata.map((incidentdata, index) => (
          <div className="instance-key" key={index}>
            <span onClick={() => fetchInstance(incidentdata.ProcessInstanceKey)}>{incidentdata.ProcessInstanceKey}</span>
      </div>
      ))}
      </div>
      <div className="incident-item">
            <span>BPMN process id</span>
        {incidentdata && incidentdata.map((incidentdata, index) => (
          <div className="incident-info" key={index}>
            <span>{incidentdata.BpmnProcessId}</span>
          </div>
      ))}
      </div>
      <div className="incident-item">
        <span>Error code</span>
        {incidentdata && incidentdata.map((incidentdata, index) => (
          <div className="incident-info" key={index}>
            <span>{incidentdata.ErrorType}</span>
      </div>
      ))}
      </div>
      <div className="incident-item">
        <span>Time</span>
        {incidentdata && incidentdata.map((incidentdata) => (
          <div className="incident-info">
            <span>{format(new Date(incidentdata.Timestamp), 'dd-MM-yyyy HH:mm:ss')}</span>
      </div>
      ))}
      </div>
      <div className="incident-item">
        <span>Error message</span>
        {incidentdata && incidentdata.map((incidentdata, index) => (
          <div className="incident-info" key={index}>
            <span>{incidentdata.ErrorMessage}</span>
      </div>
      ))}
      </div>
    </div>
  );
  };

export default Incidents;
