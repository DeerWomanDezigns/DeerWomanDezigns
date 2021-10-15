import 'bootstrap/dist/css/bootstrap.min.css';
import logo from './assets/logo.png';
import './index.css';
import Image from 'react-bootstrap/Image';
import React from 'react';
import './index.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import Navigation from './components/Navigation';

function App() {
  return (
    <div>
      <div className="body">
        <Navigation />
        <Image src={logo} className="logoMain"/>
      </div>
    </div>
  );
}

export default App;
