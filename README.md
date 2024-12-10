# Flag API

This API provides a simple service to retrieve a mapping of country flags to their respective languages. The main functionality of the API is to return a JSON object that maps country flag emojis to their respective languages.

## Features

- Retrieves a list of country flags with their respective languages
- Supports multiple languages for some countries (e.g., Canada)

## Getting Started

### Prerequisites

- Docker installed on your system
- Go installed (if you want to run locally without Docker)

### Running the API with Docker

1. Build the Docker image:

    ```sh
    docker build -t go-translation-api .
    ```

2. Run the Docker container:

    ```sh
    docker run -p 65000:65000 go-translation-api
    ```

The API will be available at `http://localhost:65000/flags`.

### Running the API Locally

1. Clone the repository:

    ```sh
    git clone https://github.com/wadedesign/flagapi.git
    cd flagapi
    ```

2. Install dependencies:

    ```sh
    go mod tidy
    ```

3. Run the application:

    ```sh
    go run main.go
    ```

The API will be available at `http://localhost:65000/flags`.

## API Endpoints

### GET /flags

Retrieves a list of country flags with their respective languages.

- **URL:** `/flags`
- **Method:** `GET`
- **Success Response:**
  - **Code:** 200
  - **Content:** 

    ```json
    {
        "🇺🇸": "English",
        "🇫🇷": "French",
        "🇪🇸": "Spanish",
        "🇩🇪": "German",
        "🇮🇹": "Italian",
        "🇨🇳": "Chinese",
        "🇯🇵": "Japanese",
        "🇰🇷": "Korean",
        "🇷🇺": "Russian",
        "🇨🇦": ["English", "French"],
        "🇦🇺": "Australian",
        "🇳🇿": "New Zealand",
        "🇧🇷": "Brazilian"
    }
    ```

## Contribution Guidelines

### Adding New Flags

1. Open the `data/flags.go` file.
2. Add your new flag and language(s) to the `FlagToLanguage` map.

Example:

```go
var FlagToLanguage = map[string]interface{}{
    "🇺🇸": "English",
    "🇫🇷": "French",
    "🇪🇸": "Spanish",
    "🇩🇪": "German",
    "🇮🇹": "Italian",
    "🇨🇳": "Chinese",
    "🇯🇵": "Japanese",
    "🇰🇷": "Korean",
    "🇷🇺": "Russian",
    "🇨🇦": []string{"English", "French"},
    "🇦🇺": "Australian",
    "🇳🇿": "New Zealand",
    "🇧🇷": "Brazilian",
    "🇲🇽": "Spanish" // Example of adding a new flag
}
```