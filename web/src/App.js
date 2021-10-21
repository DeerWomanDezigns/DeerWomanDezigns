import 'bootstrap/dist/css/bootstrap.min.css';
import logo from './assets/logo_full.png';
import './index.css';
import Image from 'react-bootstrap/Image';
import UsersPage from './components/UsersPage';
import Login from './components/Login';
import Home from './components/Home';
import Backdrop from './assets/Back.jpg';
import { BrowserRouter as Router, Route, Switch, Link} from 'react-router-dom';
import Nav from 'react-bootstrap/Nav';
import React, { useState } from "react";
import { RiCloseFill, RiMore2Fill } from 'react-icons/ri';
import { IconContext } from 'react-icons';
import { SidebarData } from './components/sidebarData';
import './components/navbar.css'


var sectionStyle = {
  backgroundImage: `url(${Backdrop})`
}



function App() {
  const [sidebar, setSidebar] = useState(false);
  const showSidebar = () => setSidebar(!sidebar)

  return (
    <div style={sectionStyle} onClick={showSidebar}>
      <div className="body">
        <Router>
          <IconContext.Provider value={{ color: "white" }}>
          <div>
            <div className="navbar">
              <Link to='#' className='menu-bars'>
                <RiMore2Fill onClick={showSidebar} />
              </Link>
            </div>
            <Nav className={sidebar ? 'nav-menu active' : 'nav-menu'}>
              <ul className='nav-menu-items'>
                <li className='navbar-toggle'>
                  <Link to="/" className='menu-bars'>
                    <RiCloseFill onClick={showSidebar} />
                  </Link>
                </li>
                {SidebarData.map((item, index) => {
                  return (
                    <li key={index} className={item.cName}>
                      <Link to={item.path}>
                        {item.icon}
                        <span>{item.title}</span>
                      </Link>
                    </li>
                  );
                })}
              </ul>
            </Nav>
          </div>
        </IconContext.Provider>
          <Image src={logo} className="logoMain" alt="Logo"/>
          <Switch>
            <Route exact path='/' component={Home} />
            <Route exact path='/Home' component={Home} />
            <Route path='/UsersPage' component={UsersPage} />
          </Switch>
        </Router>
        
        <Login />
      </div>
    </div>
  );
}

export default App;
