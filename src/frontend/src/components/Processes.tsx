import React from 'react';
import './Processes.css';

const Processes: React.FC = () => {
  return (
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
    </div>
  );
};

export default Processes;
