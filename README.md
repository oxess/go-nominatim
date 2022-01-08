# GoLang Nominatim - simple lib

Simple library for Open Street Maps Nominatim service.

## Install 

```bash
go get -v github.com/oxess/go-nominatim
```

## Usage 

Import library:
```go
import "github.com/oxess/go-nominatim"
```

Create service object (public api or own self-hosted instance):
```go
geocoder := nominatim.NewPublicApi() // Default public 
// or 
geocoder := nominatim.New("https://nominatim.self.hosted") // Default public
```

Search place by address query:
```go
// Limit set to -1 disable limit ion query
geocoder.Search("Warszawa, plac zbawiciela 1", -1)
```

## Tests

Run tests with code coverage:
```bash
go test -v -cover
```