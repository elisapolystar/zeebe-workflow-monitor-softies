import React, {useState, useEffect} from 'react';
import ReactDOM from 'react-dom/client';
import Instanceview from './Instanceview.tsx';
import './Incidents.css';
import info from './testinstance.json';
import { format } from 'date-fns';

interface IncidentProps {
  socket: WebSocket | null;
  incidents: string | null;
}

const Incidents: React.FC<IncidentProps> = ({socket}) => {

  const [incidentdata, setincidentData] = useState(info.data.incidents);
  const [bpmnData, setBpmnData] = useState<string | null>(null);

  useEffect(() => {
    setincidentData(info.data.incidents);
  }, []);

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
  
  if(incidentdata.length > 0){
    
    return(
      <div className="incident-container">
        <div className="incident-item">
        <span>Process Instance Key</span>
          {incidentdata && incidentdata.map((incidentdata, index) => (
            <div className="instance-key" key={index}>
              <span onClick={() => navigate(`/Instanceview/${incidentdata.ProcessInstanceKey}`)}>{incidentdata.ProcessInstanceKey}</span>
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
          {incidentdata && incidentdata.map((incidentdata, index) => (
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
  }

  else{
    return (
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
    );
  }
};

export default Incidents;
