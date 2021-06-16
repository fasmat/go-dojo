# ZCamp 2021 Coding Dojo - Concurrency in Go

[![GitHub Actions CI](https://github.com/fasmat/go-dojo/actions/workflows/main.yml/badge.svg)](https://github.com/fasmat/go-dojo/actions/workflows/main.yml)

## Introduction

![go Logo](https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/1920px-Go_Logo_Blue.svg.png)

This coding dojo was created to have some fun during the #ZCamp2021 and learn Go together.

## How to participate

1. Download and install [VSCode](https://code.visualstudio.com/)
2. Download and install [Docker](https://www.docker.com/)
3. Install the extension [Remote Containers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)

When you clone the repository and open the directory inside VSCode it should prompt you to re-open it in a docker container. If it doesn't
click the green Icon in the bottom left of the VSCode window and select "Reopen in Container". The container that is started by VSCode already is configured
for the challenge. You can test if it works by:

```bash
go build -v ./...
```

Which should pass without any errors.

The first start of the Container takes ~1 minute with fast internet connection, after that you are ready to go!

## Challenges

### [Challenge 01: WebCrawler](challenge01)

**Note:** This example was taken and extended from <https://tour.golang.org/concurrency/10>

The ``Crawl`` function uses a ``Fetcher`` to fetch a website from a given URL and recursively continue to crawl URLs until a certain depth.

The goal of this challenge is to extend this program to a more full-featured tool that works more efficiently can deal with uncooperative servers. Only modify
the contents of [challenge01.go](challenge01/challenge01.go)! You have completed a task when the corresponding test passes. To start the tests use

```bash
go test -v -race ./challenge01
```

This ensures that the race detector is active during the tests. Take a look at [challenge01_test.go](challenge01/challenge01_test.go) to start working on the
tasks.

|#|Summary|
|-|:--------|
|1|Fetch each URL only once to save bandwidth|
|2|Fetch URLs concurrently to speed up the crawler|
|3|Add the option to limit the rate at which the crawler fetches the server|

## Contributing

Feel free to submit a PR. I'm happy to extend and improve these challenges.
