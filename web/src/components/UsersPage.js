import Container from 'react-bootstrap/Container';
import Users from './Users';
import React, { Component } from 'react';

export default class UsersPage extends Component {
    render() {
        return (
          <Container>
          <Users />
        </Container> 
        );
    }
}