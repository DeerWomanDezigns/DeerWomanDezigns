import Orders from './Orders';
import React from "react";
import '../index.css';
import Form from 'react-bootstrap/Form';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import RangeSlider from 'react-bootstrap-range-slider';


class OrdersPage extends React.Component {
    constructor(props) {
        super(props);
        this.state = { value: '' };
        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }

    handleChange(event) {    this.setState({value: event.target.value});  }
    handleSubmit(event) {
      console.log(this.state.value);
      event.preventDefault();
    }

render() {
    return (
        <div>
            <h1 className="OrdersPage">Orders</h1>
            <br />
            <div>
                <Col className="container">
                    <Row>
                        <Form>
                            <Form.Check
                                type="switch"
                                id="paid"
                                label="Paid in Full"
                            />
                            <Form.Check
                                disabled
                                type="switch"
                                label="Fullfilled"
                                id="fullfilled"
                            />
                        </Form>
                    </Row>
                    <Row>
                        <p>&thinsp;</p>
                        <p>&thinsp;</p>
                        <p>&thinsp;</p>
                    </Row>
                    <Row className="slider">
                        Amount:
                        <RangeSlider
                            value={this.state.value}
                            onChange={this.handleChange}
                        />
                        <Form.Control placeholder="0" value={this.state.value} onChange={this.handleChange} />
                    </Row>
                </Col>
            </div>
            <Orders />
        </div>
    );
}
}

export default OrdersPage;