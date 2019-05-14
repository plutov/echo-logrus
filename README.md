### echo-logrus

Middleware echo-logrus is a [logrus](https://github.com/sirupsen/logrus) logger support for [echo](https://github.com/labstack/echo).

#### Usage

import package

```go
echologrus "github.com/cemkiy/echo-logrus"
```

define new logrus

```go
echologrus.Logger = logrus.New()
e.Logger = echologrus.GetEchoLogger()
e.Use(echologrus.Hook())
```
