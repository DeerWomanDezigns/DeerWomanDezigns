import React from 'react';
import Spinner from 'react-bootstrap/Spinner';

class EtsyAuth extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            error: null,
            isLoaded: false,
            users: []
        };
    }


    componentDidMount() {
        fetch(`${process.env.REACT_APP_BACKEND_SERVER_URL}/api/v1/etsy/login`, {
            "method": "GET",
            "headers": {
                "Authorization": process.env.REACT_APP_BACKEND_API_KEY
            }
        })
            .then(res => res.json())
            .then(
                (result) => {
                    this.setState({
                        isLoaded: true,
                        results: result
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
        const { error, isLoaded, results } = this.state;
        if (error) {
            console.log("Error: " + error.message)
            return <div><strong>...</strong></div>;
        } else if (!isLoaded) {
            return <div>
                <Spinner animation="grow" />
            </div>;
        } else {
            console.log(results)
            return (
                <div>
                    {results.map(item => (
                        <div>
                            <h1>---</h1>
                        </div>
                    ))}
                </div>
            );
        }
    }
}

export default EtsyAuth;