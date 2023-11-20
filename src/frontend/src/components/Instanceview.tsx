import React, {useRef} from 'react';

const Instanceview: React.FC = () => {
     const containerRef = useRef<HTMLDivElement>(null);
    
    return <div className="bpmn-container" ref={containerRef}>BMPN here</div>;
    }
  
export default Instanceview;