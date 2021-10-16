import React from 'react';
import Container from 'react-bootstrap/Container';
import Spinner from 'react-bootstrap/Spinner';
import Card from 'react-bootstrap/Card';
import configData from '../config.json'

class Users extends React.Component {
    constructor(props) {
      super(props);
      this.state = {
        error: null,
        isLoaded: false,
        users: []
      };
    }


    componentDidMount() {
        fetch(`${configData.SERVER_URL}/api/v1/users`, {
            "method": "GET",
            "headers": {
                "Authorization": configData.API_KEY
            }
        })
        .then(res => res.json())
        .then(
          (result) => {
            this.setState({
              isLoaded: true,
              users: result
            });
          },
          (error) => {
            this.setState({
              isLoaded: true,
              error
            });
          }
        )
    }
    
    render() {
      const { error, isLoaded, users} = this.state;
      if (error) {
        console.log("Error: " + error.message)
        return <div><strong>...</strong></div>;
      } else if (!isLoaded) {
        return <div>
        <Spinner animation="grow" />    
      </div>;
      } else {
        console.log(users)
        return (
            <Container className="orderList">
                <p><strong>Users:</strong></p>
                {users.map(user => (
                  <Card>
                  <Card.Body>
                      <Card.Title>
                        {user.firstName} {user.lastName}
                      </Card.Title>
                      <Card.Text>
                        <strong>ID: </strong> {user.id}
                        <br />
                        <strong>Address: </strong> {user.address}
                      </Card.Text>
                  </Card.Body>
              </Card>
                ))}
            </Container>
        );
      }
    }
  }

export default Users;