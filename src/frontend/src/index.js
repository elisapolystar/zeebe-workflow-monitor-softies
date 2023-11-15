import React from 'react';
import { createRoot } from "react-dom/client";
import Header from './components/Header.tsx'
import App from './App.tsx';

const root = createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <div>
      <Header />
      <App />
    </div>
  </React.StrictMode>,
); export {root};