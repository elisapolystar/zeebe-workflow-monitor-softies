import React from 'react';
import ReactDOM from 'react-dom/client';
import './NavBar.css';
import Processes from './Processes.tsx';
import Instances from './Instances.tsx';
import Incidents from './Incidents.tsx';

const NavBar: React.FC = () => {
  const navigate = (path: string) => {
    window.history.pushState({}, '', path);
    ReactDOM.createRoot(document.getElementById('content') as HTMLElement).render(getComponentForPath(path));
  };

  const getComponentForPath = (path: string) => {
    switch (path) {
      case '/processes':
        return <Processes />;
      case '/instances':
        return <Instances />;
      case '/incidents':
        return <Incidents />;
      default:
        return <Processes />;
    }
  };

  return (
    <div>
      <nav>
        <ul>
          <li onClick={() => navigate('/processes')}>Processes</li>
          <li onClick={() => navigate('/instances')}>Instances</li>
          <li onClick={() => navigate('/incidents')}>Incidents</li>
        </ul>
      </nav>
      <div id="content">
        {getComponentForPath(window.location.pathname)}
        </div>
    </div>
  );
};

export default NavBar;
