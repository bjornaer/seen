# Google Pub/Sub Client in Golang
[![Go Report Card](https://goreportcard.com/badge/github.com/bjornaer/seen)](https://goreportcard.com/report/github.com/bjornaer/seen)
[![](https://tokei.rs/b1/github/bjornaer/seen?category=code)](https://github.com/bjornaer/seen)
![Test](https://github.com/bjornaer/seen/actions/workflows/test_and_lint.yml/badge.svg)



## Introduction

Welcome to this rather remarkable journey through asynchronous messaging using Google Cloud Pub/Sub, elegantly scribed in the Go language. This repository harbours a lightweight client designed to interact with Google Pub/Sub services to pull messages, both for mainline execution and testing.

In the fashion of a tale's opening, let us step through the key elements that constitute this repository.

## Table of Contents

- [Features](#features)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Environment Variables](#environment-variables)
- [Architecture](#architecture)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

## Features

- Efficient message pulling from Google Pub/Sub.
- Support for Google Cloud Pub/Sub authentication.
- Unit tests for core functionalities.
- Exemplary use of Docker for containerisation.
- Streamlined CI/CD pipeline with GitHub Actions.

## Getting Started

### Prerequisites

- Go 1.19 or higher
- Docker (for containerisation)
- Google Cloud SDK (for interacting with Google Pub/Sub)

### Installation

Clone the repository into your desired location:

```sh
git clone https://github.com/your_project.git
```

Navigate to the directory and run:

```sh
go build
```

### Environment Variables

Certain environment variables must be set for proper execution. Define these either in your shell or a `.env` file.

- `PROJECT_ID`: Your Google Cloud Project ID.
- `SUB_ID`: Your Google Cloud Pub/Sub Subscription ID.

## Architecture

The code is separated into different directories, each serving a specific purpose:

- `cmd/app/`: Entry point for the application.
- `internal/pubsubclient/`: Core logic for interacting with Google Pub/Sub.

## Testing

To run the unit tests, simply execute:

```sh
go test ./...
```

## Contributing

We welcome contributions from all who find this project of interest. Kindly consult the [CONTRIBUTING.md](CONTRIBUTING.md) file for guidelines on how to contribute.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

A tale is only as good as its telling, and this README aims to narrate the essence of this code repository with the gravity it deserves. Should you have questions or suggestions, feel free to raise an issue or submit a pull request.

Happy Coding! 📚👨‍💻👩‍💻