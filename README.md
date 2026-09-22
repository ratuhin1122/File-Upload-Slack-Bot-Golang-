# Slack File Upload Bot (Golang)

A lightweight and efficient Slack bot built in Go that automates uploading files and images directly to specified Slack channels using the official [`slack-go/slack`](https://github.com/slack-go/slack) SDK.

---

## Features

- **Automated File Uploading**: Upload multiple documents, images, or PDFs to any Slack channel.
- **Slack V2 Upload API**: Compatible with Slack's modern external upload flow (`files.getUploadURLExternal`).
- **Secure Configuration**: Reads sensitive tokens and channel IDs from environment variables to prevent accidental credential leaks.
- **Error Handling**: Graceful error handling for file system operations and Slack API responses.

---

## Prerequisites

Before running this bot, make sure you have:

1. **Go installed**: Version 1.21 or higher ([Download Go](https://golang.org/dl/)).
2. **A Slack Workspace**: With administrative or app installation permissions.
3. **A Slack App**: Created at [Slack API Portal](https://api.slack.com/apps).

---

## Slack App Setup & Scopes

1. Go to [api.slack.com/apps](https://api.slack.com/apps) and select **Create New App** > **From scratch**.
2. Navigate to **OAuth & Permissions** in the sidebar.
3. Under **Bot Token Scopes**, add the following permission:
   - `files:write` (Allows the bot to upload files)
4. Scroll to the top of the page and click **Install to Workspace**.
5. Copy the **Bot User OAuth Token** (starts with `xoxb-...`).
6. In your Slack workspace, go to the target channel and invite your bot:
   ```text
   /invite @YourBotName
   ```
7. Copy the **Channel ID** (Right-click channel > *View channel details* > scroll to bottom to find Channel ID, e.g., `C0123456789`).

---

## Installation & Setup

1. **Clone the repository:**
   ```bash
   git clone https://github.com/ratuhin1122/File-Upload-Slack-Bot-Golang-.git
   cd File-Upload-Slack-Bot-Golang-
   ```

2. **Download dependencies:**
   ```bash
   go mod download
   ```

---

## Configuration & Usage

### 1. Set Environment Variables

#### On Windows (PowerShell):
```powershell
$env:SLACK_BOT_TOKEN="xoxb-your-slack-bot-token"
$env:CHANNEL_ID="your-channel-id"
```

#### On Linux / macOS (Bash):
```bash
export SLACK_BOT_TOKEN="xoxb-your-slack-bot-token"
export CHANNEL_ID="your-channel-id"
```

#### On Windows (Command Prompt):
```cmd
set SLACK_BOT_TOKEN=xoxb-your-slack-bot-token
set CHANNEL_ID=your-channel-id
```

### 2. Configure Files to Upload

In `main.go`, specify the path to your file(s) in `fileArr`:

```go
fileArr := []string{"git-cheat-sheet.pdf"}
```

### 3. Run the Bot

```bash
go run main.go
```

Upon successful upload, you will see output similar to:
```text
ID : F0C4FH8EP32, Title : git-cheat-sheet.pdf
```

---

## Project Structure

```text
├── .gitignore          # Git ignore rules for binaries and sensitive files
├── go.mod              # Go module definition and dependencies
├── go.sum              # Checksums for module dependencies
├── main.go             # Main application logic
└── README.md           # Project documentation
```

---

## Troubleshooting

- **`file.upload.v2: filename cannot be empty` / `file size cannot be 0`**:
  Make sure the file specified in `fileArr` exists locally and is not empty.
- **`CompleteUploadExternal: not_in_channel`**:
  The bot is not invited to the channel. Run `/invite @BotName` inside the channel.
- **`invalid_auth`**:
  Check that `SLACK_BOT_TOKEN` is correctly set and starts with `xoxb-`.

---

## License

This project is open source and available under the [MIT License](LICENSE).
