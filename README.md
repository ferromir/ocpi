# OCPI
Go types for the [OCPI protocol](https://evroaming.org/ocpi/).

This library provides types for the following OCPI versions:
* 2.1.1
* 2.2.1
* 2.3

## Install
```
go get github.com/ferromir/ocpi
```

## Usage
```Go
package main

import (
	"github.com/ferromir/ocpi/ocpi211"
	"github.com/ferromir/ocpi/ocpi221"
	"github.com/ferromir/ocpi/ocpi230"
)

var token1 ocpi211.Token
var token2 ocpi221.Token
var token3 ocpi230.Token
// same for Location, Tariff, CDR, Session, etc...
```

## Notes
* Case insensitive and sensitive strings are typed the same (as string).
* Validation of structs is out of the scope of this library.
* An **optional** "metadata" field has been added to the entities despite this field not being specified in the protocol. The intention is to allow parties to exchange data beyond the protocol.
* The type described as "number" in the protocol is represented by a float32.