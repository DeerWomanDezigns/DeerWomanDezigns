import Container from 'react-bootstrap/Container';
import Navbar from 'react-bootstrap/Navbar';
import Nav from 'react-bootstrap/Nav';
import Image from 'react-bootstrap/Image';
import logo from '../assets/logo.png';
import './navbar.css';

function Navigation() {
    return (
      <>
            <div>
        <Navbar bg="white" variant="light">
        <Container>
            <Image src={logo} className="logo"/>
        <Nav className="me-auto">
        <Nav.Link>Sign-In</Nav.Link>
        <Nav.Link>Popular Items</Nav.Link>
        </Nav>
        </Container>
        </Navbar>
        </div>  
      </>
    );
  }
  

export default Navigation;
