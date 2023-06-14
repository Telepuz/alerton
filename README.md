# AlertOn

A simple alerting daemon

# Features

Its main features are:

- Easily extensible using scripts
- Single-host system
- Send messages to telegram

# Installation

For DEB-based distr use packages from releases

Building with Go 1.20+

```sh
make install
```

# Configuration

Configuration example

```yaml
---
# How offen run scripts
check_interval: 15m
# Cooldown alert
cooldown_duration: 60m
# Script's timeout
alert_timeout: 60s
# Credentials for telegram
telegram_token: XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
telegram_chatid: XXXXXXXXXXXXXXXXXXXX
alerts:
- name: Ping google.com
  command: pingHost.sh
  params:
    - google.com
- name: Resolve DNS-name google.com
  command: dnsCheck.sh
  params:
    - google.com
- name: Check free disk space
  command: diskFreeSpace.sh
```
