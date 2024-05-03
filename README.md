# Certchecker

Simplest possible TLS cert checker. Displays an info about all certificate chain including root CA. Also works for untrusted certificates

## Run locally

```sh
go run main.go -d facebook.com
```

## Run as distributed binary

```sh
go run github.com/hrvadl/certchecker@latest -d google.com
```

## Examples
