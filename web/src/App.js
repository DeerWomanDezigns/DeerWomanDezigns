import 'bootstrap/dist/css/bootstrap.min.css';
import logo from './assets/logo_full.png';
import './index.css';
import Image from 'react-bootstrap/Image';
import React from 'react';
import Navigation from './components/Navigation';
import UsersPage from './components/UsersPage';
import Home from './components/Home';
import Backdrop from './assets/Back.jpg';
import { BrowserRouter as Router, Route, Switch} from 'react-router-dom';

var sectionStyle = {
  backgroundImage: `url(${Backdrop})`
}

function App() {
  return (
    <div style={sectionStyle}>
      <div className="body">
        <Router>
          <Navigation />
          <Image src={logo} className="logoMain" />
          <Switch>
            <Route exact path='/' component={Home} />
            <Route exact path='/Home' component={Home} />
            <Route path='/UsersPage' component={UsersPage} />
          </Switch>
        </Router>
      </div>
    </div>
  );
}

export default App;
