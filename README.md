# go-cli-template

This is a `go-cli` template written in Go, providing basic authentication commands for AWS, Cloudflare, GitHub, and GitLab, as well as file reading capabilities for JSON, YAML, and ENV files using Viper.

Check out the release at [GitHub Releases](https://github.com/opslabhqx/go-cli-template/releases).

Edit the `go-cli` name at `/app/internal/config/base`. The version defaults to `dev` and parses the git tag release at the GitHub workflow.

## Getting Started

1. Clone the repository:

```sh
git clone https://github.com/yourusername/go-cli-template.git
cd go-cli-template
```

2. Build the `go-cli`:

```sh
make build
```

3. Run the `go-cli`:

```sh
./go-cli
```

### Usage

#### AWS Commands

Authenticate with AWS and display account details.

```sh
./go-cli aws auth
```

#### Cloudflare Commands

Authenticate with Cloudflare (Placeholder).

```sh
./go-cli cf auth
```

#### GitHub Commands

Authenticate with GitHub (Placeholder).

```sh
./go-cli gh auth
```

#### GitLab Commands

Authenticate with GitLab (Placeholder).

```sh
./go-cli gl auth
```

### Output Formats

The `go-cli` supports multiple output formats for displaying data. The output formats can be specified using the `-o` or `--output` flag.

#### Table Output

The default output format is a table. This format presents data in a well-structured table using the `tablewriter` package.

Example:

```sh
./go-cli aws auth -o table
```

Table Output

```bash
% go-cli gh auth
|-------|---------------------|
|  KEY  |        VALUE        |
|-------|---------------------|
| Email | example@gmail.com   |
|-------|---------------------|
```

#### Raw Output

The raw output format displays data as plain text, with each row of data separated by commas.

Example:

```sh
./go-cli aws auth -o raw
```

`raw` Output

```bash
% go-cli gh auth -o raw
Email, example@gmail.com
```

#### JSON Output

The JSON output format presents data as a JSON object, making it easy to parse programmatically.

Example:

```sh
./go-cli aws auth -o json
```

`json` Output

```bash
% go-cli gh auth -o json
{
"Email": "example@gmail.com"
}
```

## Development

### Adding New Commands

1. Create a new folder under `app/cmd` for your command.
2. Implement your command in a new `.go` file.
3. Register your command in `app/cmd/root.go`.

### Project Dependencies

Dependencies are managed using Go modules. To add a new dependency:

```sh
go get <dependency>
go mod tidy
```

### Building Docker Image

Build a Docker image for the `go-cli`.

```sh
docker build -t go-cli-template .
```

### Docker Compose

A `docker-compose.yml` and `docker-compose.json` are provided for running the application with Docker Compose.

```sh
docker-compose up
```

# License

This project is licensed under the [MIT License](/LICENSE). See the LICENSE file for details.
