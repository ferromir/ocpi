# OCPI
Go types for the [OCPI protocol](https://evroaming.org/ocpi/).

This library provides types for the following OCPI versions:
* 2.1.1
* 2.2.1

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
)

var token1 ocpi211.Token
var token2 ocpi221.Token
// same for Location, Tariff, CDR, Session, etc...
```