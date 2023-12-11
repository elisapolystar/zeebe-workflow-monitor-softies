import { BpmnVisualization } from 'bpmn-visualization';
import React, { useEffect, useState } from 'react';
import './component styles/instanceview.css';

interface InstanceProps {
  instance: string | null;
}

const Instanceview: React.FC<InstanceProps> = ({instance}) => {
  const [diagramData, setDiagramData] = useState<string | null>(null);
  const instancesData = instance ? JSON.parse(instance) : {};
  const bpmnContainerElt = window.document.getElementById('bpmn-container');

  //change base64 encoded resource to xml
  const encodedBpmn = instancesData.process.resource;
  const xml = atob(encodedBpmn);

  //fetch the diagram data
  useEffect(() => {
    async function fetchData() {
        try {
            const response = await fetch(encodedBpmn);
            if (!response.ok) {
              throw new Error('Failed to fetch diagram');
            }
            const data = await response.text();
            setDiagramData(data);
        } catch (error) {
            console.error('Error fetching diagram:', error);
        }
    }
    fetchData();
}, []);

  /**
   * Render the BPMN diagram
   * Change the color to green for completed elements so the execution path can be seen
   */
  useEffect(() => {
    if (diagramData  && bpmnContainerElt) {
        const bpmnVisualization = new BpmnVisualization({ container: bpmnContainerElt as HTMLElement, navigation: { enabled: false } });
        bpmnVisualization.load(xml);

        if(instancesData && instancesData.Elements){
          instancesData.Elements.forEach((Elements) => {
            if (Elements.Intent === "ELEMENT_COMPLETED" || Elements.Intent === "SEQUENCE_FLOW_TAKEN") {
              bpmnVisualization.bpmnElementsRegistry.updateStyle(Elements.ElementId, 
                { stroke: {
                color: 'green', opacity: 80
              }});
            }
          });
      }
    }
  }, [bpmnContainerElt, diagramData]);

  //Show all the variable names and values.
  return (
  <div className='instanceview'>
    <h2>{instancesData.process.bpmnProcessId}</h2>
    <div id="bpmn-container"></div>
    <div className='variables'>
      <b>All variables</b>
    </div>

    <div className='variable-container'>    
      <div className='variable-item'>
        <span>Name</span>
        {instancesData.Variables &&
          instancesData.Variables.map((instanceData, index) => (
          <div className="variable-info" key={index}>
            <span>{instanceData.Name}</span>
          </div>
        ))}   
      </div>

      <div className='variable-item'>
        <span>Value</span>
        {instancesData.Variables &&
          instancesData.Variables.map((instanceData, index) => (
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