# OpenAI CLI

OpenAI CLI is a command-line interface (CLI) written in Golang that allows users to interact with the OpenAI API using the chat option. Users can pass prompts using the `--prompt` flag or through STDIN, allowing for a flexible and convenient way to interact with the OpenAI API.

## Installation

To use the OpenAI CLI, you don't need to have Golang installed. You can download the pre-built binaries for your operating system from the [Releases](https://github.com/renanmac/openai_cli/releases) page.

1. Go to the [Releases](https://github.com/renanmac/openai_cli/releases) page of the GitHub repository.
2. Download the appropriate binary for your operating system (e.g., `openai_cli-windows-amd64.exe` for Windows, `openai_cli-darwin-amd64` for macOS, or `openai_cli-linux-amd64` for Linux).
3. Place the downloaded binary in a directory that's included in your system's `PATH` environment variable.

## Getting an API Key

To use the OpenAI CLI, you'll need an API key from OpenAI. If you don't have one yet, you can sign up for an account and obtain an API key from the [OpenAI website](https://openai.com/).

## Usage

### Using the `chat` command and `--prompt` flag

You can use the `chat` command along with the `--prompt` flag to interact with the OpenAI API. Here's an example of how to use it:

```bash
$ OPENAI_API_KEY=your_token_here openai_cli chat --prompt='Something interesting'
```

Replace your_token_here with your actual OpenAI API key.

### Using STDIN
Alternatively, you can pass the prompt through STDIN. Create a file (e.g., prompt_file.txt) containing your prompt text, and then use the following command:

```bash
$ cat prompt_file.txt | openai_cli
```

This will read the contents of prompt_file.txt and use it as the prompt for interacting with the OpenAI API.

### Contributing
Contributions to the OpenAI CLI are welcome! If you'd like to contribute, please follow these steps:

1. Fork the repository and clone it to your local machine.
2. Create a new branch for your feature or bug fix.
3. Make your changes and test thoroughly.
4. Commit your changes with clear commit messages.
5. Push your branch to your forked repository.
6. Create a pull request to the main repository.
7. Before submitting a pull request, please ensure that your code adheres to the project's coding standards and practices.

### License
This project is licensed under the MIT License.

❗️This README was generated with the assistance of an AI language model 😄