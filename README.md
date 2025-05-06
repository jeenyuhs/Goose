# Goose 📞

A Discord bot reference implementation that creates anonymous "telephone" connections between users across different servers.

![billede](https://github.com/user-attachments/assets/660ec8d8-c661-41d1-b1f8-864dcc95ccd8)


## Overview

Goose is a fun and simple Discord bot framework that simulates phone calls between users in a server. This code is designed as a reference implementation for developers to integrate into their own Discord bots, not as a hosted service. Therefore, there is no invite link.

## Features

- **Anonymous Phone Calls**: Connect with other users without knowing who they are
- **Simple Interface**: Start calls with `!call` and end them with `!end`
- **Thread-Based**: Uses Discord threads for clean organization
- **Real-Time Communication**: Messages in one thread are forwarded to the connected thread
- **Auto-Archiving**: Calls are automatically archived when ended

## How It Works

1. A user types `!call` in any channel
2. The bot creates a thread for that user's "phone call"
3. When another user also initiates a call, the bot connects the two threads
4. Messages sent in one thread appear in the connected thread
5. Either user can end the call with `!end`

## Commands

- `!call` - Start a new phone call
- `!end` - End the current call (only works in call threads)

## Installation

### Prerequisites

- Go 1.24.2 or higher
- Discord Bot Token with required permissions

### Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/jeenyuhs/Goose.git
   cd Goose
   ```

2. Set your Discord bot token as an environment variable:
   ```bash
   export BOT_TOKEN=your_bot_token_here
   ```

3. Build and run the bot:
   ```bash
   go build -o goose ./cmd/main.go
   ./goose
   ```

### Required Bot Permissions

The bot requires the following permissions:
- Read Messages/View Channels
- Send Messages
- Create Public Threads
- Send Messages in Threads
- Manage Threads

## Project Structure

The project is organized as follows:

```
Goose/
├── cmd/
│   └── main.go               # Entry point for the bot
├── internal/
│   ├── handlers/
│   │   └── messages.go       # Message handling logic
│   ├── models/
│   │   └── threads.go        # Thread representation and status tracking
│   └── repository/
│       └── repository.go     # Manages call state and connections
├── go.mod                    
├── LICENSE                   
└── README.md                 # Project documentation
```

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## License

This project is licensed under the [MIT License](LICENSE).
