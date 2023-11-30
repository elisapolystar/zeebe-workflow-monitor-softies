import React from 'react';
import { createRoot } from "react-dom/client";
<<<<<<< HEAD
import Header from './components/Header.tsx'
import App from './App.tsx';
=======
import MessageListener from './messageListener';
import NavBar from './components/NavBar';
import Header from './components/Header'
>>>>>>> main

const root = createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <div>
      <Header />
      <App />
    </div>
  </React.StrictMode>,
<<<<<<< HEAD
); export {root};
=======
);
>>>>>>> main
