# ClickUp CLI

This command-line interface (CLI) allows you to interact with ClickUp, a project management tool, using the Go programming language. With this CLI, you can perform various operations on ClickUp tickets.

## Prerequisites

Before using this CLI, make sure you have the following installed:

- Go programming language (version 1.13 or later)
- ClickUp API token

## Installation

To install the ClickUp CLI, follow these steps:

1. Clone the repository:

   ```shell
   git clone https://github.com/gabrieldebem/clickup-cli.git
   ```

2. Change to the project directory:

   ```shell
   cd clickup
   ```

3. Install the dependencies:

   ```shell
   go mod download
   ```

4. Build the CLI:

   ```shell
   go build -o clickup-cli
   ```

## Usage

The CLI supports the following commands:

### List Tickets

To list all the tickets from the folder provided, use the `list` command:

```shell
./clickup-cli list
```

You can add the `--only-mine` flag to list only your assigned tickets:

```shell
./clickup-cli list --only-mine
```

### Show Ticket

To view the details of a specific ticket, use the `show` command followed by the ticket ID:

```shell
./clickup-cli show <ticket-id>
```

Replace `<ticket-id>` with the actual ID of the ticket you want to show.

### Update Ticket

To update the status of a ticket, use the `update` command followed by the ticket ID and the new status:

```shell
./clickup-cli update <ticket-id> <new-status>
```

Replace `<ticket-id>` with the actual ID of the ticket you want to update, and `<new-status>` with the desired status.

## Configuration

Before using the CLI, make sure to set your ClickUp API token. Create a `/.clickup/.env` file on your `$HOME` and add the following line:

```
CLICKUP_BASE_URL="https://api.clickup.com/api"
CLICKUP_API_TOKEN=<your-api-token>
CLICKUP_SPACE_ID=<your-space-id>
CLICKUP_TEAM_ID=<your-team-id>
CLICKUP_FOLDER_ID<your-folder-id>=
CLICKUP_LIST_ID=<your-list-id>
CLICKUP_USER_ID=<your-user-id>
```

Replace all values with your actual ClickUp API credentials.

## Contributions

Contributions to this CLI are welcome. Feel free to submit bug reports, feature requests, or pull requests on the [GitHub repository](https://github.com/gabrieldebem/clickup-cli).

## License

This project is licensed under the [MIT License](LICENSE).
