import 'bootstrap/dist/css/bootstrap.min.css';
import logo from './assets/logo_full.png';
import './index.css';
import Image from 'react-bootstrap/Image';
import React from 'react';
import Navigation from './components/Navigation';
import Backdrop from './assets/Back.jpg';
import UsersPage from './components/UsersPage';

var sectionStyle = {
  backgroundImage: `url(${Backdrop})`
}

function App() {
  return (
    <div style={sectionStyle}>
      <div className="body">
        <Navigation />
        <Image src={logo} className="logoMain" />
        <UsersPage />
      </div>
    </div>
  );
}

export default App;
