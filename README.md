<br><br>

<h1 align="center">Flog</h1>

<p align="center">
  <a href="/LICENSE"><img src="https://img.shields.io/badge/license-MIT-blue.svg"/></a>
  <a href="https://goreportcard.com/report/github.com/mingrammer/flog"><img src="https://goreportcard.com/badge/github.com/mingrammer/flog"/></a>
  <a href="https://travis-ci.org/mingrammer/flog"><img src="https://travis-ci.org/mingrammer/flog.svg?branch=master"/></a>
  <a href="https://hub.docker.com/r/mingrammer/flog"><img src="https://img.shields.io/docker/pulls/mingrammer/flog.svg"/></a>
</p>

<p align="center">
A fake log generator for common log formats
</p>

<br><br><br>

flog is a fake log generator for common log formats such as apache-common, apache error and RFC3164 syslog.

It is useful for testing some tasks which require log data like amazon kinesis log stream test.

> Thanks to [gofakeit](https://github.com/brianvoe/gofakeit) 😘

## Installation

### Using go get

```bash
go get github.com/mingrammer/flog
```

It is recommended to also run `dep ensure` to make sure that the dependencies are in the correct versions.

### Using [homebrew](https://brew.sh)

```
brew tap mingrammer/flog https://github.com/mingrammer/flog.git
brew install flog
```

### Using .tar.gz archive

Download gzip file from [Github Releases](https://github.com/mingrammer/flog/releases/latest) according to your OS. Then, copy the unzipped executable to under system path.

### Using [docker](https://www.docker.com)

```
docker run -it --rm mingrammer/flog
```

## Usage

There are useful options. (`flog --help`)

```
Options:
  -f, --format string      Choose log format. ("apache_common"|"apache_combined"|"apache_error"|"rfc3164") (default "apache_common")
  -o, --output string      Output filename. Path-like is allowed. (default "generated.log")
  -t, --type string        Log output type. ("stdout"|"log"|"gz") (default "stdout")
  -n, --number integer     Number of lines to generate.
  -b, --bytes integer      Size of logs to generate (in bytes).
                           "bytes" will be ignored when "number" is set.
  -s, --sleep numeric      Sleep interval time between lines (in seconds). It does not actually sleep every log generation.
  -p, --split-by integer   Split the logs by this value in lines or bytes.
                           When "number" is set, it specifies the maximum number of lines for a log file.
                           When "bytes" is set, it specifies the maximum size of a log file.
  -w, --overwrite          [Warning] This will overwrite the existing file with new created logs.
```

```bash
# Generate 1000 lines of logs to stdout
flog

# Generate a single log file with 1000 lines of logs, then overwrite existing log file
flog -t log -w

# Generate a single log gzip file with 3000 lines of logs every 10 seconds
flog -t gz -o log.gz -n 3000 -s 10

# Generate logs up to 10MB and split the log files every 1MB in "web/log/apache.log" path with apache combined format
flog -t log -f apache_combined -o web/log/apache.log -b 10485760 -p 1048576
```

## Features

* [X] Completely random log generator
* [ ] Contextual random log generator
* [ ] Statistical random log generator
* [x] Support `apache common`, `apache combined`, `apache error` and `rfc3164` log format
* [ ] Support some other syslog formats
* [x] Support stdout, file and gzip type of logging

## License

MIT
