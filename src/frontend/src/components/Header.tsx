import React from 'react';
import './Header.css';

const Header: React.FC = () => {
  return (
    <>
      <div className="header-container">
        <div className="logo-item">
          <span>Logo</span>
        </div>
        <div className="header-item">
          <span>Zeebe Workflow Monitor</span>
        </div>
      </div>
    </>
  );
};

export default Header;
