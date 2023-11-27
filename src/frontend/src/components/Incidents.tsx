import React, {useState, useEffect} from 'react';
import ReactDOM from 'react-dom/client';
import Instanceview from './Instanceview.tsx';
import './Incidents.css';
import info from './testinstance.json';
import { format } from 'date-fns';

const Incidents: React.FC = () => {

  const [incidentdata, setincidentData] = useState(info.data.incidents);
  useEffect(() => {
    setincidentData(info.data.incidents);
  }, []);

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
              <span onClick={() => navigate('/Instanceview')}>{incidentdata.ProcessInstanceKey}</span>
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
