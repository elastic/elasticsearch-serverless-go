# Elasticsearch Serverless Client

[![main](https://github.com/elastic/elasticsearch-serverless-go/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/elastic/elasticsearch-serverless-go/actions/workflows/tests.yml)

This is the official Elastic client for the **Elasticsearch Serverless** service. If you're looking to develop your Go application with the Elasticsearch Stack, you should look at the [Elasticsearch Client](https://github.com/elastic/elasticsearch-go) instead.


## Installation

You can install the Elasticsearch Serverless Go client with the following
commands:

```bash
go get -u github.com/elastic/elasticsearch-serverless-go@latest
```

### Instantiate a Client

You can instantiate a client by running the following command:

```go
client, err := elasticsearch.NewClient(elasticsearch.Config{
	APIKey: "you_api_key",
	Address: "https://my-project-url",
})
```

### Usage

You can read to the [official documentation]()
page for a getting started guide.

## Development

See [CONTRIBUTING](./CONTRIBUTING.md).

### Docs

TBD

## License 📗

[Apache 2.0](LICENSE) © [Elastic](https://www.elastic.co/)