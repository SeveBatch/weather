"""
This module provides functionality for interacting with the weather service API. 
Currently it includes a main enry point to retrieving a random location and it's current weather data, forecasts, and other 
related information.

Contributing:
If you'd like to contribute to the weather service, please ensure your code 
follows the project's style guidelines and includes appropriate tests and coverage. Submit 
pull requests with clear descriptions of your changes and their purpose.
"""

## Auto deploy process
1. **Run deploy.sh**
    ```bash
    source deploy.sh
    ```

## Manual Deployment process
1. **Build Docker Image**
    ```bash
    docker build -t weather:latest .
    ```
2. **Tag and Push GH images**
    ```bash
    docker tag github.com/sevebatch/weather:latest ghcr.io/sevebatch/weather:v1.0.0
    docker push ghcr.io/sevebatch/weather:v1.0.0
    ```

    If using a local image with same name you can 
    ```bash 
    kubectl rollout restart deployment/weather
    ```

3. **Apply Manifest**
    ``` bash
     kubectl apply -f deployment.yml
     ```
