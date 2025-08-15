# MrAndreID / Go Package

[![Go Reference](https://pkg.go.dev/badge/github.com/MrAndreID/gopackage.svg)](https://pkg.go.dev/github.com/MrAndreID/gopackage) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

The `MrAndreID/GoPackage` is a collection of functions for package in Go Language.

## Table of Contents

* [Installation](#installation)
* [Usage](#usage)
* [Versioning](#versioning)
* [Authors](#authors)
* [Contributing](#contributing)
* [Official Documentation for Go Language](#official-documentation-for-go-language)
* [License](#license)
* [More](#more)

## Installation

To use The `MrAndreID/GoPackage`, you must follow the steps below:
- Get Dependancies
```go
go get -u github.com/MrAndreID/gopackage
```

## Usage

To use The `MrAndreID/GoPackage`, you must ensure that you meet the following requirements:

### Custom Validator

```go
import (
    "github.com/MrAndreID/gopackage"
	"github.com/labstack/echo/v4"
)

e := echo.New()
e.Validator = gopackage.CustomValidator()
```

### Custom HTTP Error for Echo Framework

```go
import (
    "github.com/MrAndreID/gopackage"
	"github.com/labstack/echo/v4"
)

e := echo.New()
e.HTTPErrorHandler = gopackage.EchoCustomHTTPErrorHandler
```

### Custom JSON

```go
import (
    "github.com/MrAndreID/gopackage"
	"github.com/labstack/echo/v4"
)

e := echo.New()
e.JSONSerializer = gopackage.CustomJSON()
```

### GoRM Data Table

```go
import (
    "github.com/MrAndreID/gopackage"
)

gopackage.DataTable(
    ctx,
    queryBuilder,
    []string{"name", "level"},
    "name",
    "asc",
    "id",
    "desc",
    page,
    &limit,
    request.Search,
    false,
)
```

### Binding Request for Echo Framework

```go
import (
    "github.com/MrAndreID/gopackage"
)

if err := gopackage.EchoBindRequest(c, &req); err != nil {
    return err
}
```

### SeaweedFS Client

```go
import (
    "github.com/MrAndreID/gopackage"
)

seaweedFSData, err := gopackage.NewSeaweedFS("127.0.0.1", "9333", false)

if err != nil {
    return err
}

publicURL, err := seaweedFSData.Upload(base64File)

if err != nil {
    return err
}

base64File, err := seaweedFSData.Download(publicURL)

if err != nil {
    return err
}

err := seaweedFSData.Delete(publicURL)

if err != nil {
    return err
}
```

## Versioning

I use [Semantic Versioning](https://semver.org/). For the versions available, see the tags on this repository. 

## Authors

**Andrea Adam** - [MrAndreID](https://github.com/MrAndreID/)

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
Please make sure to update tests as appropriate.

## Official Documentation for Go Language

Documentation for Go Language can be found on the [Go Package website](https://pkg.go.dev/).

## License

The `MrAndreID/GoPackage` is released under the [MIT License](https://opensource.org/licenses/MIT). See the `LICENSE` file for more information.

## More

Documentation can be found [on https://go.dev/](https://pkg.go.dev/github.com/MrAndreID/gopackage).
