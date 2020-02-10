### echo-logrus

Middleware echo-logrus is a [logrus](https://github.com/sirupsen/logrus) logger support for [echo](https://github.com/labstack/echo).

`v3.0` tag supports v3.
`v4.0` tag supports v4.

#### Install

```sh
go get -u github.com/plutov/echo-logrus
```

#### Usage

import package

```go
echologrus "github.com/plutov/echo-logrus"
```

define new logrus

```go
echologrus.Logger = logrus.New()
e.Logger = echologrus.GetEchoLogger()
e.Use(echologrus.Hook())
```
