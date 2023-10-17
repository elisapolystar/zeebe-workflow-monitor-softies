import React from 'react';
import ReactDOM from 'react-dom/client';
import MessageListener from './messageListener';

const root = ReactDOM.createRoot(document.getElementById("root"));
root.render(
  <React.StrictMode>
    <MessageListener />
  </React.StrictMode>
);