import React from 'react';
import ReactDOM from 'react-dom';
import './NavBar.css';
import Processes from './Processes';
import Instances from './Instances';
import Incidents from './Incidents';

const NavBar = () => {
  const navigate = (path) => {
    window.history.pushState({}, null, path);
    ReactDOM.render(getComponentForPath(path), document.getElementById('content'));
  };

  const getComponentForPath = (path) => {
    switch (path) {
      case '/processes':
        return <Processes />;
      case '/instances':
        return <Instances />;
      case '/incidents':
        return <Incidents />;
      default:
        return <div>Not Found</div>;
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
