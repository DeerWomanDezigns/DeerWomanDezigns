import Users from './Users';
import React, { Component } from 'react';
import '../index.css';

export default class UsersPage extends Component {
    render() {
        return (
          <div className="UsersPage">
            <h1>Users</h1>
          <Users />
        </div> 
        );
    }
}