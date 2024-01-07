# goalarm-cli

CLI with simple alarm functionality built in Go language.

## Installation

### Prerequisites

Make sure you have Go installed on your system. If not, follow the official Go installation guide: [Go Installation Guide](https://go.dev/doc/install).

### Installing goalarm-cli

```shell
$ go install github.com/tomo-kn/goalarm-cli/cmd/goalarm@latest
```

## Usage

1. To set a timer, enter a command like `goalarm set [TIME]`. For example:

   ```
   goalarm set 15:30
   ```

   This will set the alarm for 15:30.

2. When the alarm triggers at the set time, press `Enter` to stop it.

If you need to cancel the alarm before it goes off, simply press `Ctrl + C`.

## Notice

- Please be aware that the alarm emits sound; adjust your PC's volume settings accordingly.
- For safety and to avoid accidental prolonged use, `goalarm-cli` is designed to automatically turn off 15 minutes after the alarm has been activated.

## Supported Environment

`goalarm-cli` is tested and guaranteed to work on M1 Mac.

## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.
