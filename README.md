# GAIT
Gait - an ✨AI✨ powered CLI tool for generating commit messages based on staged files.

## Building
To build **gait** from source:

1) If not installed, install **go 1.22 or newer** on your machine. You can get **go** from [the official website](https://go.dev/doc/install).
2) If not installed, install **Taskfile**. You can get **Taskfile** from [the official website](https://taskfile.dev/installation/).
3) Clone this repo: `git clone https://github.com/gfxv/gait.git`
4) In the `gait` directory run `task` or `task build`, which will result in a binary file named `gait` in `bin` directory.

## Usage
> For now, only Google's Gemini is supported. Other models will be available in the near future.

### Configuration
To use **gait**, you need to prorivde an API key using the `GAIT_API_KEY` environment variable.

1. Obtain your API key by following these steps:
- For Gemini, please visit [this link](https://aistudio.google.com/app/apikey).

2. Set the Environment Variable

Write you API key to env file like so:
```sh
echo "GAIT_API_KEY=<YOUR_API_KEY>" > <PATH_TO_FILE>/local.env
```

Add the following line to your shell configuration file (`~/.bashrc`, `~/.zshrc`, etc.) to load env variable:
```sh
# example for zshrc
if [ -f <PATH_TO_FILE>/local.env ]; then
  export $(cat <PATH_TO_FILE>/local.env | xargs)
fi
```

3. Add **gait** to your PATH

Add the following line to your shell configuration file, replacing `<path-to-binary>` with the actual path to compiled binary:
```sh
export PATH="<path-to-binary>:$PATH"
```

4. Reload shell configuration or restard the terminal.

Here is example of how you can reload shell configuration for `~/.zshrc`:
```sh
source ~/.zshrc
```

### Available commands

| Command | Description |
| ------- | ----------- |
| `gait commit` (or `gait c`) | Generates a commit message based on the changes in staged files. If the generated message is acceptable, it creates a commit and opens it in edit mode for any final adjustments before confirming the commit. |

#### Planned Features
- Add a command to generate multiple commit message suggestions, allowing the user to choose the most suitable one.
- Add commands that simplify common git tasks, such as unstaging files, reverting commits, renaming commits, etc.

