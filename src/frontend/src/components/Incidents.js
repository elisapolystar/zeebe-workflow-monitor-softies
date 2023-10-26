import React from 'react';
import './Incidents.css'; 

const Incidents = () => {
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
    </div>
  );
};

export default Incidents;