# Logger

Fast, simple, structured, leveled logger.

## Installation

```shell
go get -u github.com/outsidedigital/logger
```

## Quick Start

The following example demonstrates how to use the logger:

```golang
log := logger.NewLogger().With().String("url", url).Span().Logger()
res, err := client.Do(req)
if err != nil {
  log.Error(err).Message("failed to fetch url")
}
```

## Development

The project contains the [Development Container](.devcontainer) configuration
with well-defined tools and its prerequisites. Using this configuration along
with [Visual Studio Code](https://code.visualstudio.com) and its
[Remote Containers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)
extension is a preferred way to set up a local development environment.

## Contributing

We will do our best to keep [main branch](../../tree/master) in good shape,
with tests passing at all times.

If you intend to make breaking changes, we recommend [filling an issue](../../issues).
If you're only fixing a bug, it's fine to submit a merge request right away but
we still recommend to fill an issue detailing what you're fixing. This is helpful
in case we don't accept that specific fix but want to keep track of the issue.
