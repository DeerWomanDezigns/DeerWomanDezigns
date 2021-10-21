import React, { useState } from 'react';
import { CognitoUser, AuthenticationDetails } from 'amazon-cognito-identity-js';
import Form from 'react-bootstrap/Form';
import Button from 'react-bootstrap/Button';
import UserPool from '../UserPool';

const Login = () => {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");

    const onSubmit = (event) => {
        event.preventDefault();

        const user = new CognitoUser({
            Username: email,
            Pool: UserPool
        });
        const authDetails = new AuthenticationDetails({
            Username: email,
            Password: password
        });
        
        user.authenticateUser(authDetails, {

            onSuccess: (data) => {
                console.log("onSuccess: ", data);
            },
            onFailure: (err) => {
                console.error("onFailure: ", err);
            },
            newPasswordRequired: (data) => {
                console.log("newPasswordRequired: ", data);
            },
        });
    };
    return (
        <div>
            <Form className="col-lg-6 offset-lg-3 ">
                <div className="row justify-content-center">
                <Form.Group className="mb-3">
                    <Form.Label>Email Address</Form.Label>
                    <Form.Control type="email" placeholder="Enter email" value={email} onChange={(event) => setEmail(event.target.value)}/>
                </Form.Group>

                <Form.Group className="mb-3">
                    <Form.Label>Password</Form.Label>
                    <Form.Control type="password" placeholder="Password" value={password} onChange={(event) => setPassword(event.target.value)}/>
                    <Form.Text className="text-muted">
                        Contact site admins for support.
                    </Form.Text>
                </Form.Group>
                <Form.Group className="mb-3">
                    <Form.Check type="checkbox" label="Remember Me" />
                </Form.Group>
                <Button variant="primary" type="submit" onSubmit={onSubmit}>
                    Login
                </Button>
                </div>
            </Form>
        </div>
    )
}

export default Login;