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

### [Challenge 02: The context package](challenge02)

**Disclaimer:** The idea for this challenge came from [Just for Func](https://www.youtube.com/channel/UC_BzFbxG2za3bp5NRRRXJSw) and more specifically
this [video](https://www.youtube.com/watch?v=8M90t0KvEDY).

The goal of this challenge is to re-implement the official [context](https://golang.org/pkg/context/) package. It is widely used in applications to carry
deadlines, cancellation signals, and other request-scoped values across API boundaries and between processes. For further information see the official
documentation.

In this challenge the foundation for the package is already layed. There is a single test already available, but it is unfortunately failing. Fix this test
and then proceed to implement the whole package in a TDD fashion test after test.

The whole context package without tests is < 350 lines of code and can be used as a reference if you get stuck. Remember to run all tests with the race
detector enabled to ensure the context package is concurrency safe.

```bash
go test -v -race ./challenge02
```

## Contributing

Feel free to submit a PR. I'm happy to extend and improve these challenges.
