import React, { Component } from 'react';
import '../index.css';
import Table from 'react-bootstrap/Table';
import Container from 'react-bootstrap/esm/Container';

export default class Orders extends Component {
    render() {
        return (
            <Container>
                <div className="orders-table">
                    <Table striped bordered hover responsive="sm" variant="dark">
                        <thead>
                            <tr>
                                <th>Order ID</th>
                                <th>Product Description</th>
                                <th>Product ID</th>
                                <th>Transaction Date</th>
                                <th>Subtotal Amount</th>
                                <th>Total Amount</th>
                                <th>Seller Name</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr>
                                <td>1</td>
                                <td>Product A</td>
                                <td>12A</td>
                                <td>2021-10-14 06:14</td>
                                <td>$12.99</td>
                                <td>$15.99</td>
                                <td>Jane Doe</td>
                            </tr>
                        </tbody>
                    </Table>
                </div>
            </Container>
        );
    }
}