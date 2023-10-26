import React from 'react';
import { createRoot } from "react-dom/client";
import MessageListener from './messageListener';
import NavBar from './components/NavBar';
import Header from './components/Header'

const root = createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <div>
      <Header />
      <NavBar />
      <MessageListener />
    </div>
  </React.StrictMode>,
);