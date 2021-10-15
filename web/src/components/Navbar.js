import Container from 'react-bootstrap/Container';
import Navbar from 'react-bootstrap/Navbar';
import Nav from 'react-bootstrap/Nav';
import Form from 'react-bootstrap/Form';
import FormControl from 'react-bootstrap/FormControl';
import Modal from 'react-bootstrap/Modal';
import Button from 'react-bootstrap/Button';
import React, { useState } from "react";

function Navigation() {

  const [show, setShow] = useState(false);
  const handleClose = () => setShow(false);
  const handleShow = () => setShow(true);

    return (
      <>
            <div>
        <Navbar collapseOnSelect bg="dark" expand="md" className="mb-3" variant="dark">
        <Container>
        <Nav className="me-auto">
        <Navbar.Toggle />
        <Navbar.Collapse className="justify-content-end">
          <Nav>
            <Nav.Link>Home</Nav.Link>
            <Nav.Link>Orders</Nav.Link>
            <Nav.Link onClick={handleShow}>Sign In</Nav.Link>
            <Nav.Link>Users</Nav.Link>
          </Nav>
        </Navbar.Collapse>
        </Nav>
        <Form inline>
           <FormControl type="text" placeholder="Search" className="mr-sm-2" />
          </Form>
        </Container>
        </Navbar>
        </div>  
      </>
    );
  }
  

export default Navigation;
