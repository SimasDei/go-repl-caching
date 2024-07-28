
# Go REPL Caching

This project is a Go implementation of a REPL (Read-Eval-Print Loop) that includes caching functionality. It is designed to interact with the PokeAPI and cache responses for more efficient and faster access.

## Features

- **REPL Interface**: Interactive command-line interface for user inputs.
- **Caching**: Caches responses from the PokeAPI to improve performance.
- **Commands**:
  - `explore`: Explore different endpoints of the PokeAPI, including Pokémon details, locations, and other resources.
  - `map`: Show cached data.
  - `help`: Display available commands.
  - `exit`: Exit the REPL.

## Getting Started

### Prerequisites

- Go 1.16 or later

### Installation

Clone the repository:

```sh
git clone https://github.com/SimasDei/go-repl-caching.git
cd go-repl-caching
```

Build the project:

```sh
go build
```

### Usage

Run the REPL:

```sh
./go-repl-caching
```

#### Commands

- **explore**: This command allows users to interactively explore different endpoints of the PokeAPI. You can retrieve details about various Pokémon, their abilities, types, and other characteristics. Additionally, you can explore locations within the Pokémon world, retrieving information about different areas, including regions, towns, and specific locations.

  Example:
  ```sh
  > explore pokemon/pikachu
  > explore location/1
  ```

- **map**: Displays the currently cached data. This includes any previously retrieved Pokémon information or locations, providing quick access without needing to hit the API again.

  Example:
  ```sh
  > map
  ```

- **help**: Lists all available commands and their descriptions to assist users in navigating the REPL.

  Example:
  ```sh
  > help
  ```

- **exit**: Exits the REPL session.

  Example:
  ```sh
  > exit
  ```

## Project Structure

- `main.go`: Entry point of the application.
- `repl.go`: Contains the REPL logic.
- `command_*.go`: Command implementations.
- `internal/pokeapi`: Handles interactions with the PokeAPI.
- `pokecache`: Manages cached data.

