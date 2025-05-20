# Forecast Service

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

5. Curl Sample Request
    ```bash 
    curl localhost:5000
    ```

    ### Sample Response
    A successful response from the service might look like this:

    ```json
    {
        "A chance of showers and thunderstorms before 4pm, then showers and thunderstorms likely. Mostly sunny. High near 54, with temperatures falling to around 49 in the afternoon. West wind around 14 mph, with gusts as high as 21 mph. Chance of precipitation is 70%."
    }
    ```

## Contributing
See CONTRIBUTING.md

## License
This project is licensed under the MIT License.