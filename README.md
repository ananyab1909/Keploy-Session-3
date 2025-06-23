# Keploy-Session-3

## Description

It contains a comprehensive test suite for a custom RESTful API server built with Go and PostgreSQL. The tests verify the correctness, reliability, and behavior of the API's CRUD operations on user data.

## Project Overview

It features :
  - Mocked and non-mocked unit tests covering all four endpoints (CreateUser, GetUsers, UpdateUser, DeleteUser)
  - Validates real interactions between the API server and a test PostgreSQL database
  - End-to-end tests ensuring the API processes HTTP requests correctly and returns expected responses
  - Uses Go’s built-in coverage tools to measure and ensure high test coverage.

## Test Types

  - Unit Tests (mocked & non-mocked) - Validate handler logic in isolation, ensuring edge cases and failures are covered.
  - Integration Tests - Check the interaction between the API and a real database to ensure CRUD operations function as expected.
  - API Tests - Simulate real HTTP requests to verify the full request-response cycle works correctly.
  - Test Coverage - Uses go test -coverprofile to track and improve code coverage for all major components.

## Prerequisites

  - Go 1.20+ installed
  - PostgreSQL (local setup or managed service like Neon DB)

## Implementation

1. Clone the repository
     ```
     https://github.com/ananyab1909/Keploy-Session-3.git
     ```

2. Enter into the directory
     ```
     cd Keploy-Session-3
     ```

3. Run All Tests with Coverage
    ```
    go test ./... -coverprofile=coverage -coverpkg=custom-api-server/handlers
    ```

4. View Coverage Report - Generate and open an HTML report to inspect code coverage visually.
    ```
    go tool cover -html=coverage
    ```

5. Access the application locally
    ```
    Backend API: http://localhost:8080
    Frontend UI: http://localhost:3000
    ```


## User Routes

  The API Documentation is provided in [API.md](https://github.com/ananyab1909/Keploy-Session-2/blob/main/API.md)

## About Me

Hello, my name is Ananya Biswas. I am an Engineering Student at [Kalinga Institute of Industrial Technology](https://kiit.ac.in/). I enjoy making projects and now that my this project is over, I am open-sourcing the project. Hope you like it! Lastly, I would like to put it out there that I have worked on other projects that you may like. You can check them out at my [Github](https://github.com/ananyab1909/). Give it a whirl and let me know your thoughts.

## Socials
  - Portfolio : https://dub.sh/ananyabiswas
  - LinkedIn : https://linkedin.com/in/ananya-biswas-kiit/
  - Mastodon : https://mastodon.social/@captain_obvious/
  - Twitter : https://x.com/not_average_x/
  - Github : https://github.com/ananyab1909/
