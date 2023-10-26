import React from 'react';
import './Processes.css'; 
/*   
https://blog.logrocket.com/creating-react-sortable-table/     
https://mui.com/material-ui/react-accordion/

*/
const Processes = () => {
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