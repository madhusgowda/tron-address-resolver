# Tron Address Resolver

A Golang package that resolves Tron addresses by converting them from base58 format to the standard hexadecimal (Ethereum-like) format. This package is useful when you need to work with Tron addresses in their hexadecimal form.

## Installation

To install the package, you can simply run:

```bash
go get github.com/madhusgowda/tron-address-resolver
```

## Usage
Once installed, you can import and use the package to resolve Tron addresses

## Functions
ResolveAddress(address string) string

**Parameters:**

address: The Tron address to be resolved (in base58 format).

Returns:

A hexadecimal (Ethereum-like) address if the input starts with "T", otherwise it returns the original address.

## Dependencies
This package relies on the following Go module:
```
github.com/btcsuite/btcutil/base58: For base58 decoding.
```

## Contributing
Feel free to open issues or submit pull requests if you find bugs or have suggestions for improvement!
