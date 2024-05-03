# Certchecker

Simplest possible TLS cert checker. Displays info about all certificate chains including root CA. Also works for untrusted certificates

## Run locally

```sh
go run main.go -d facebook.com
```

## Run as distributed binary

```sh
go run github.com/hrvadl/certchecker@latest -d google.com
```

## Examples
<img width="1551" alt="image" src="https://github.com/hrvadl/certchecker/assets/93580374/90a69488-1bb4-4202-9c5a-1481347cf862">
