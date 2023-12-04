import React, {useEffect} from 'react';
import Instanceview from './Instanceview.tsx';
import './Incidents.css';
import { format } from 'date-fns';

interface IncidentProps {
  socket: WebSocket | null;
  incidents: string | null;
  setContent: React.Dispatch<React.SetStateAction<JSX.Element | null>>;
}

const Incidents: React.FC<IncidentProps> = ({socket, incidents, setContent}) => {
  //const [incidentdata, setincidentData] = useState(info.data.incidents);
  const incidentdata = incidents ? JSON.parse(incidents) : [];


  useEffect(() => {
    //setincidentData(info.data.incidents);
  }, []);

  const fetchInstance = (id: string | undefined) => {
    if (socket && socket.readyState === WebSocket.OPEN) {
      const messageObject = `{ "instance": "${id}" }`;
      socket.send(messageObject);
      console.log(`Instance request for process ${messageObject} sent from frontend`);
    }
  };

  const navigate = (navData: string) => {
    const view = navData.split('/');
    const path = `/${view[1]}`;
    const id = view[2];
    getComponentForPath(path, id)
  };

  const getComponentForPath = (path: string, id: string) => {
    switch (path) {
      case '/Instanceview':
        fetchInstance(id)
        return 
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
            path = '/Instanceview';
            data = <Instanceview process_instance={message.data} />;
            break;

          default: return;
        }
        setContent(data);
      });
    }
  }, [socket]);
  
  if(incidentdata.length > 0){

    let formattedDate = '';
    const date = new Date(incidentdata[0].Timestamp);
    formattedDate = format(date, 'dd-MM-yyyy HH:mm:ss');
    
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
              <span>{formattedDate}</span>
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
