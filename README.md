<img src="web/src/assets/logo.png" alt="logo" width="200"/>
<br /><br />

# DeerWomanDezigns
This repository hosts the source code for the administration website for the [DeerWomanDezigns Etsy shop](https://www.etsy.com/shop/DeerwomanDezigns).

## Languages
The frontend is a react app hosted on an nginx server while the backend is written in golang.

## Infrastructure
All data is stored in AWS DynamoDB and the servers are hosted in AWS Lightsail.

## Running the App

### Using Docker
1. Clone the repository
2. Install Docker
3. Ensure the Docker agent is running locally
4. Navigate to the root directory of the repository
5. Run `docker-compose up`
6. Navigate to http://localhost and ensure the site loads
