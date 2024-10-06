# Blu Installment App

Blu Installment is a project owned by [Wigata Intech](https://wigataintech.com). This repository provides the code for a Go-based HTTP server, using Docker for deployment. It includes a multi-stage Docker build for optimized image size and a `Makefile` to simplify building, running, and cleaning up the project.

## Requirements

- [Go 1.23.1](https://golang.org/dl/)
- [Docker](https://www.docker.com/get-started)
- [Make](https://www.gnu.org/software/make/)
- [Goose](https://github.com/pressly/goose)

## Project Structure

- `Dockerfile`: Defines a multi-stage build for compiling the Go binary and running it in a minimal Docker image.
- `docker-compose.yml`: Manages multi-container Docker applications, including the Taurus app and MySQL.
- `Makefile`: Provides commands for building, running, and cleaning the project.
- `migration/`: Contains SQL migration files for managing the database schema using goose.
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

### 3. Build and Run the Application Using Docker Compose

#### Build the Docker Images
To build the Docker images for the application and MySQL using Docker Compose, run:

``` make build-compose ```

#### Start the Application and Database
To start the Go application and MySQL database in containers, use:

``` make up-compose ```

This will start both services in the background.

##### Stop the Application and Database
To stop the running containers:

``` make down-compose ```

### 4. Manage Database Migrations with goose
We use goose to manage database schema migrations. All migration files should be placed in the migrations/ directory.

#### Install goose command-line tool
To install goose, run:

``` go install github.com/pressly/goose/v3/cmd/goose@latest ```

#### Create a New Migration
To create a new migration, run:

``` make create-migration name=<migration_name> ```

This will generate two files in the migrations/ folder (one for applying the migration and one for rolling it back).

#### Apply Migrations
To apply all pending migrations to the MySQL database, use:

``` make migrate ```

This will connect to the MySQL database running inside the Docker container and apply the migration files.

#### Rollback Last Migration
To rollback the last migration that was applied:

``` make rollback ```

#### Reset the Database
To reset the database (rollback all migrations):

``` make reset ```

### 5. Local Development

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

### 6. Running Unit Tests and Generating Test Coverage Report
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

### 7. Clean Up

#### Clean Docker Resources
To remove all Docker containers, images, and volumes, run:

''' make clean '''

#### Clean Local Build
If you built the Go binary locally and want to clean up:

''' make clean-local '''

## Docker Workflow Summary

- **Build** the Docker image: `make build`
- **Run** the application in a Docker container: `make run`
- **Run with arguments**: `make run-args`
- **Clean** Docker resources: `make clean`
- **Rebuild** the image: `make rebuild`

## Summary of Makefile Commands
- Build the Docker images: `make build`
- Start the application and database: `make up`
- Stop the application and database: `make down`
- Create a new migration: `make create-migration name=<migration_name>`
- Run database migrations: `make migrate`
- Rollback last migration: `make rollback`
- Reset the database: `make reset`
- Run tests and generate a coverage report: `make test`
- Clean up Docker resources: `make clean`
- Clean up local build: `make clean-local`

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Author

This project is maintained by **Dhira Wigata**. For more information, visit [Wigata Intech](https://wigataintech.com).
