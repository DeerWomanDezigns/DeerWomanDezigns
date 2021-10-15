<img src="web/src/assets/logo.png" alt="logo" width="200"/>
<br /><br />

[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=DeerWomanDezigns_DeerWomanDezigns&metric=bugs)](https://sonarcloud.io/dashboard?id=DeerWomanDezigns_DeerWomanDezigns)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=DeerWomanDezigns_DeerWomanDezigns&metric=vulnerabilities)](https://sonarcloud.io/dashboard?id=DeerWomanDezigns_DeerWomanDezigns)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=DeerWomanDezigns_DeerWomanDezigns&metric=code_smells)](https://sonarcloud.io/dashboard?id=DeerWomanDezigns_DeerWomanDezigns)
[![CodeQL](https://github.com/DeerWomanDezigns/DeerWomanDezigns/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/DeerWomanDezigns/DeerWomanDezigns/actions/workflows/codeql-analysis.yml)

[![tf-apply](https://github.com/DeerWomanDezigns/DeerWomanDezigns/actions/workflows/tf-apply.yml/badge.svg)](https://github.com/DeerWomanDezigns/DeerWomanDezigns/actions/workflows/tf-apply.yml)
[![build-backend](https://github.com/DeerWomanDezigns/DeerWomanDezigns/actions/workflows/build-backend.yml/badge.svg)](https://github.com/DeerWomanDezigns/DeerWomanDezigns/actions/workflows/build-backend.yml)
[![build-frontend](https://github.com/DeerWomanDezigns/DeerWomanDezigns/actions/workflows/build-frontend.yml/badge.svg)](https://github.com/DeerWomanDezigns/DeerWomanDezigns/actions/workflows/build-frontend.yml)
[![deploy-backend](https://github.com/DeerWomanDezigns/DeerWomanDezigns/actions/workflows/deploy-backend.yml/badge.svg)](https://github.com/DeerWomanDezigns/DeerWomanDezigns/actions/workflows/deploy-backend.yml)
[![deploy-frontend](https://github.com/DeerWomanDezigns/DeerWomanDezigns/actions/workflows/deploy-frontend.yml/badge.svg)](https://github.com/DeerWomanDezigns/DeerWomanDezigns/actions/workflows/deploy-frontend.yml)

# DeerWomanDezigns
This repository hosts the source code for the administration website for the [DeerWomanDezigns Etsy shop](https://www.etsy.com/shop/DeerwomanDezigns).

## Languages
The frontend is a react app hosted on an nginx server while the backend is written in golang.

## Infrastructure
All data is stored in AWS DynamoDB and the servers are hosted in AWS Lightsail.

## CI/CD
All integration and deployments are handled through GitHub Actions.

## Running the App

### Using Docker
1. Clone the repository
2. Install Docker
3. Ensure the Docker agent is running locally
4. Navigate to the root directory of the repository
5. Run `docker-compose up`
6. Navigate to http://localhost and ensure the site loads
