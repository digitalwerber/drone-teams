# drone-teams

[![Build Status](http://drone.digitalwerber.com/api/badges/digitalwerber/drone-teams/status.svg)](http://drone.digitalwerber.com/digitalwerber/drone-teams)
[![](https://images.microbadger.com/badges/image/digitalwerber/drone-teams.svg)](https://microbadger.com/images/digitalwerber/drone-teams "Get your own image badge on microbadger.com")

:star: Forked from: `https://github.com/drone-plugins/drone-slack/`

Drone plugin for sending simple Microsoft Teams notifications.

Build the binary with the following commands:

```
export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0
export GO111MODULE=on

go build -v -a -tags netgo -o release/linux/amd64/drone-teams
```

## Docker

Build the Docker image with the following commands:

```
docker build \
  --label org.label-schema.build-date=$(date -u +"%Y-%m-%dT%H:%M:%SZ") \
  --label org.label-schema.vcs-ref=$(git rev-parse --short HEAD) \
  --file docker/Dockerfile.linux.amd64 --tag digitalwerber/drone-teams .
```

## Usage

Execute from the working directory:

```
docker run --rm \
  -e SLACK_WEBHOOK=https://outlook.office.com/webhook/... \
  -e DRONE_REPO_OWNER=octocat \
  -e DRONE_REPO_NAME=hello-world \
  -e DRONE_COMMIT_SHA=7fd1a60b01f91b314f59955a4e4d4e80d8edf11d \
  -e DRONE_COMMIT_BRANCH=master \
  -e DRONE_COMMIT_AUTHOR=octocat \
  -e DRONE_BUILD_NUMBER=1 \
  -e DRONE_BUILD_STATUS=success \
  -e DRONE_BUILD_LINK=http://github.com/octocat/hello-world \
  -e DRONE_TAG=1.0.0 \
  digitalwerber/drone-teams
```
