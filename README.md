[![Build Status](https://travis-ci.org/hanks/log_emitter.svg?branch=master)](https://travis-ci.org/hanks/log_emitter)

# Log Emitter

A tiny tool to auto-generate logs to help to test the whole pipeline of ELK stack

The core feature is trivial, to generate the new message with the latest timestamp repeatedly based on
the source log message

# Implementation

You can define source log message as a list in `conf.yml`, and the related timestamp regular expression and format used 
in generating the latest log message for you repeatedly

```bash
- log:
      - '2017/04/26 19:57:00 [error] 22422#22422: *18419270 upstream prematurely closed connection while reading response header from upstream, client: 172.22.2.48, server: adhoc.com, request: "GET /adhoc HTTP/1.1", upstream: "uwsgi://172.22.4.200", host: "adhoc.com", referrer: "https://adhoc.com"'
    dest: /var/log/nginx/error.log
    category: nginx-error
    ts_re: "[0-9]{4}/[0-9]{2}/[0-9]{2} [0-9]{2}:[0-9]{2}:[0-9]{2}"
    ts_format: "2006/01/02 15:04:05"
    interval: 4
    enable: true
```

Also use `goroutine`, `channel` and `Ticker` to implement repeated log generation.

# Compile

```bash
# Build
> make

# Test
> make test

# Benchmark Test
> make benchmark
```

# Usage

```bash
cd bin/
./log_emitter.linux64 --help
./log_emitter.linux64 --config ../conf.yml
```
