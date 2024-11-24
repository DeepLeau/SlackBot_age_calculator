# Slack Bot with Slacker

A Slack bot built using the Slacker library in Go. This bot calculates a user's age based on their year of birth.

## Features

- Responds to the command: `my year of birth is <year>`.
- Calculates the user's age and replies with it.
- Logs command events for analytics.

## Prerequisites

- Go version 1.16 or higher installed.
- A Slack workspace and Slack App configured with appropriate tokens.
- `.env` file containing your Slack tokens.

## Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/username/repository-name.git
   ```

2. Navigate to the project directory:
   ```bash
   cd repository-name
   ```

3. Install dependencies:
   ```bash
   go get -u github.com/shomali11/slacker
   go get -u github.com/joho/godotenv
   ```

4. Create a `.env` file in the project directory with the following content:
   ```env
   SLACK_BOT_TOKEN=your-slack-bot-token
   SLACK_APP_TOKEN=your-slack-app-token
   ```

   Replace `your-slack-bot-token` and `your-slack-app-token` with your actual Slack tokens.

## Usage

1. Run the bot:
   ```bash
   go run main.go
   ```

2. In your Slack workspace, send the command:
   ```
   my year of birth is <year>
   ```

   Example:
   ```
   my year of birth is 1990
   ```

   The bot will respond with your calculated age.

## Command Events

The bot logs command events for analytics, including the timestamp, command, parameters, and the event details.

## Example Output

### Slack Command
Input:
```
my year of birth is 1990
```

Response:
```
Your age is 34
```

### Console Log
```text
Commande Events
2024-11-24T12:00:00Z
my year of birth is <year>
map[year:1990]
<Slack Event Details>
```

## Contributing

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-name`).
3. Make your changes and commit them (`git commit -m "Add feature"`).
4. Push to the branch (`git push origin feature-name`).
5. Open a pull request.

## License

This project is licensed under the MIT License.

## Acknowledgements

- [Slacker](https://github.com/shomali11/slacker) for simplifying Slack bot development.
- [godotenv](https://github.com/joho/godotenv) for environment variable management in Go.
