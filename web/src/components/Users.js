import React from 'react';
import Container from 'react-bootstrap/Container';
import Spinner from 'react-bootstrap/Spinner';
import Card from 'react-bootstrap/Card';
import EtsyAuth from './EtsyAuth';

class Users extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      error: null,
      isLoaded: false,
      isRedirecting: false,
      users: []
    };
  }

  componentDidMount() {
    var fetchUrl = `${process.env.REACT_APP_BACKEND_SERVER_URL}/api/v1/etsy/test`
    fetch(fetchUrl, {
      "method": "GET",
      "headers": {
        "Authorization": process.env.REACT_APP_BACKEND_API_KEY
      }
    })
      .then(res => {
        console.log(fetchUrl)
        console.log(res.status)
        console.log(res.body)
        if (res.status === 401) {
          EtsyAuth.InitAuth("shops_r");
          this.setState({
            isRedirecting: true
          });
        } else {
          res.json()
        }
      })
      .then(
        (result) => {
          this.setState({
            isLoaded: true,
            result: result
          });
        })
      .catch(
        (error) => {
          this.setState({
            isLoaded: true,
            error: error
          });
        }
      )
  }

  render() {
    const { error, isLoaded, result, isRedirecting } = this.state;
    if (error) {
      console.log("Error: " + error.message)
      console.log(this.state.error)
      return <div>{error.message}</div>
    } else if (!isLoaded || isRedirecting) {
      return <div>
        <Spinner animation="grow" />
      </div>;
    } else {
      console.log(result)
      return (
        <Container className="orderList">
          <p><strong>Users:</strong></p>
          {result.map(user => (
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
