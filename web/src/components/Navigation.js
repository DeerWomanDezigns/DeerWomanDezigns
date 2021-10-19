import Nav from 'react-bootstrap/Nav';
import { Link } from "react-router-dom";
import React, { useState } from "react";
import { RiCloseFill, RiMore2Fill } from 'react-icons/ri';
import { IconContext } from 'react-icons';
import { SidebarData } from './sidebarData';
import './navbar.css'

function Navigation() {

  const [sidebar, setSidebar] = useState(false);
  const showSidebar = () => setSidebar(!sidebar)

  return (
    <>
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
    </>
  );
}


export default Navigation;
