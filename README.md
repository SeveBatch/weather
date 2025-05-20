# Go Service Starter Guide

## Introduction
This is a simple readme to start the weather service and make generic requests

## Getting Started

1. **Clone the Repository**
    ```bash
    git clone git@github.com:SeveBatch/weather.git
    cd poc
    ```

2. **Install Dependencies**
    ```bash
    go mod tidy
    ```

3. **Run the Service**
    ```bash
    source deploy.sh
    ```

4. Port forward in a seperate tab
    ```bash 
    kubectl port-forward service/weather 5000:5000
    ```

## Contributing
See CONTRIBUTING.md

## License
This project is licensed under the MIT License.