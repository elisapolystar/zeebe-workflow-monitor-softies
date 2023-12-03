import { BpmnVisualization } from 'bpmn-visualization';
import React, { useEffect, useState } from 'react';
import info from './testinstance.json';
import './instanceview.css';

interface InstanceViewProps {
  instance: string | null;
}

//change base64 coded resource to xml
const encodedBpmn = info.data.process.Resource;
const xml = atob(encodedBpmn);
console.log(xml);

const Instanceview: React.FC<InstanceViewProps> = () => {
  const [diagramData, setDiagramData] = useState<string | null>(null);
  const [instanceData, setInstanceData] = useState(info.data.variables);

  useEffect(() => {
    setInstanceData(info.data.variables);
  }, []);

  const bpmnContainerElt = window.document.getElementById('bpmn-container');
  useEffect(() => {
      async function fetchData() {
          try {
              const response = await fetch(encodedBpmn);
              const data = await response.text();
              setDiagramData(data);
          } catch (error) {
              console.error('Error fetching diagram:', error);
          }
      }
      fetchData();
  }, []);
  useEffect(() => {
    const bpmnContainerElt = window.document.getElementById('bpmn-container');
    if (diagramData  && bpmnContainerElt) {
        const bpmnVisualization = new BpmnVisualization({ container: bpmnContainerElt as HTMLElement, navigation: { enabled: false } });
        bpmnVisualization.load(xml);

        //Change the color to green for completed elements
        info.data.elements.forEach((element) => {
            if (element.Intent === "ELEMENT_COMPLETED" || element.Intent === "SEQUENCE_FLOW_TAKEN") {
              bpmnVisualization.bpmnElementsRegistry.updateStyle(element.ElementId, 
                { stroke: {
                color: 'green', opacity: 80
              }});
            }
          });
      }
        
  }, [bpmnContainerElt, diagramData]);
  return (
    <div className='instanceview'>
      <h2>{info.data.process.BpmnProcessId}</h2>
      <br/>

      <div id="bpmn-container"></div>
      <br/>
      
      <div className='variables'>
        <b>All variables</b>
      </div>

      <div className='variable-container'>    
        <div className='variable-item'>
          <span>Name</span>
          {instanceData &&
          instanceData.map((instanceData, index) => (
          <div className="variable-info" key={index}>
            <span>{instanceData.Name}</span>
          </div>
          ))}
        </div>
        

        <div className='variable-item'>
          <span>Value</span>
          {instanceData && instanceData.map((instanceData, index) => (
          <div className="variable-info" key={index}>
            <span>{instanceData.Value}</span>
          </div>
        ))}
      </div>
    </div>
</div>  
  );  
}
export default Instanceview;