# LogSpy


LogSpy is a command-line tool written in Go that connects to a remote server and tails log files. It provides functionality to monitor and filter log entries based on specified criteria. LogSpy supports connecting using either a private key or a Kerberos ticket.

## Installation

To install LogSpy, follow these steps:

1. Make sure you have Go installed on your system. If not, you can download it from the official Go website: [https://golang.org/](https://golang.org/)

2. Clone the LogSpy repository to your local machine:

```bash
git clone https://github.com/a-dakani/logSpy.git
```

3. Change into the LogSpy directory:

```bash
cd logSpy
```

4. Build the LogSpy binary using the Go compiler:

```bash
go build
```

5. Optionally, you can install LogSpy globally on your system by moving the binary to a directory in your system's PATH:

```bash
sudo mv LogSpy /usr/local/bin
```

## Usage

LogSpy can be used to monitor log files on a remote server. It provides various command-line options to configure the connection and filtering options. Here are the available command-line flags:

| Flag     | Description                                                              |
| -------- | ------------------------------------------------------------------------ |
| `-srv`   | Specify a predefined service name from the `config.services.yaml` file.  |
| `-fs`    | Provide a list of log file paths to tail, separated by commas.            |
| `-h`     | Set the host to connect to.                                              |
| `-u`     | Set the user to connect to the host.                                      |
| `-p`     | Set the port to connect to the host. Defaults to `22`.                    |
| `-pk`    | Set the path to the private key for authentication.                       |
| `-krb5`  | Set the location of the `krb5.conf` file for Kerberos authentication.     |
| `-f` (WIP) | Set filter words for log file entries. Multiple words should be separated by commas. |

### Examples

Here are a few examples to demonstrate how LogSpy can be used:

1. Tailing log files using predefined service:

```bash
logSpy -srv myService
```

2. Tailing log files using command-line arguments:

```bash
logSpy -fs=/var/log/app.log,/var/log/error.log -h=192.168.1.1 -u=admin -p=22 -pk=/path/to/private/key -f=ERROR,WARN
```

## Configuration

LogSpy relies on configuration files to define services and their respective settings. The configuration files should be placed in the same directory as the LogSpy binary and have the following formats:

- `config.yaml`: Contains general configurations for LogSpy.
- `config.services.yaml`: Defines predefined services and their connection details.

Please make sure to configure these files correctly before running LogSpy.

## Known Issues:
- Error logging for Kerberos tickets is currently buggy and may not provide accurate error messages or handle certain scenarios correctly. This can affect the authentication process when connecting to a remote server using a Kerberos ticket.
