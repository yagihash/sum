# SUM: Site Update Monitor
[![CircleCI](https://circleci.com/gh/yagihashoo/sum.svg?style=svg&circle-token=a4bde268d2cc1780b8ea03035b3200d6fc5da1de)](https://circleci.com/gh/yagihashoo/sum) [![codecov](https://codecov.io/gh/yagihashoo/sum/branch/master/graph/badge.svg?token=qGtqMkqT2E)](https://codecov.io/gh/yagihashoo/sum)

World Cheapest Site Monitor. It may be useful for...

- Detecting illegal tamper
- Keeping up to date for info about concert of your favarite rock band

## Deployment

### Example service file

```
[Unit]
Description=Site Update Monitor

[Service]
EnvironmentFile=/etc/sum.config
ExecStart=/opt/sum
ExecStop=/bin/kill ${MAINPID}
KillMode=process
User=yagihash
Group=yagihash

[Install]
WantedBy=multi-user.target
```

### Example config file

```
SLACK_TOKEN="xoxp-XXXXXXXXXXX"
SLACK_CHANNEL="CHANNEL_ID"
SLACK_NAME="sum"
SLACK_EMOJI=":robot_face:"
URL="https://sqli.moe"
```

### Once you register service

```bash
sudo systemctl start sum
sudo systemctl enable sum
```
