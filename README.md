# Blu Installment App

Blu Installment is a project owned by [Wigata Intech](https://wigataintech.com). This repository provides the code for a Go-based HTTP server, using Docker for deployment. It includes a multi-stage Docker build for optimized image size and a `Makefile` to simplify building, running, and cleaning up the project.

## Requirements

- [Go 1.23.1](https://golang.org/dl/)
- [Docker](https://www.docker.com/get-started)
- [Make](https://www.gnu.org/software/make/)

## Project Structure

- `Dockerfile`: Defines a multi-stage build for compiling the Go binary and running it in a minimal Docker image.
- `Makefile`: Provides commands for building, running, and cleaning the project.
- `main.go`: Contains the main application logic.

## Getting Started

### 1. Clone the Repository

```
git clone https://blu-installment.git
cd blu-installment
```

### 2. Build and Run Using Docker

#### Build the Docker Image

To build the Docker image using the provided `Makefile`, run the following command:

```
make build
```

#### Run the Docker Container

To run the application inside the Docker container:

```
make run
```

#### Run with Custom Arguments

You can pass custom configuration and secret files using the following command:

```
make run-args
```

For example:

```
make run-args CONFIG=/path/to/config.yaml SECRET=/path/to/secret.yaml
```

#### Clean Docker Resources

To clean up any unused Docker resources (images, containers):

```
make clean
```

#### Rebuild the Docker Image

To force a rebuild of the Docker image:

```
make rebuild
```

### 3. Local Development

You can also build the Go binary locally (without Docker) for local development or testing.

#### Build Locally

To build the Go binary locally:

```
make build-local
```

#### Run the Binary Locally

Once built, you can run the binary directly:

```
./taurus
```

#### Clean Up Local Build

To clean up the locally built binary:

```
make clean-local
```

### 4. Running Unit Tests and Generating Test Coverage Report
You can generate a unit test coverage report in HTML format using the Makefile.

#### Run Unit Tests and Generate Coverage Report
To run the tests and generate a coverage report:

``` make test ```

This will:

- Run the unit tests.
- Generate a coverage report (coverage.out).
- Convert the coverage report into an HTML file (coverage.html), which you can open in a browser to view.

#### View the Coverage Report
After running make test, you can view the report by opening the coverage.html file in your browser:

- On macOS: ``` open coverage.html ```

- On Linux: ``` xdg-open coverage.html ```

- On Windows: ``` start coverage.html ```

#### Clean Test Coverage Reports
To clean up the generated test coverage files:

``` make clean-test ```

## Docker Workflow Summary

- **Build** the Docker image: `make build`
- **Run** the application in a Docker container: `make run`
- **Run with arguments**: `make run-args`
- **Clean** Docker resources: `make clean`
- **Rebuild** the image: `make rebuild`

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Author

This project is maintained by **Dhira Wigata**. For more information, visit [Wigata Intech](https://wigataintech.com).
