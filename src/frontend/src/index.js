import React from 'react';
import { createRoot } from "react-dom/client";
import MessageListener from './messageListener.tsx';
import NavBar from './components/NavBar.tsx';
import Header from './components/Header.tsx'

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