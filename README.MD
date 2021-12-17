# Micro Logger

## package for easy logging to logstash http input from microservices

First call InitLogger

```go
micrologger.InitLogger()
```

Example:

```go
micrologger.InitLogger(
    viper.GetString("app.name"),
    viper.GetString("app.log.logstash.url"),
    true,
)
```

Then you can log errors:

```go
microlgger.LogError(reqID, action, file, data string, err error)
microlgger.LogInfo(reqID, action, file, data string)
microlgger.LogDebug(reqID, action, file, data string)
microlgger.LogWarning(reqID, action, file, data string)
microlgger.LogErrorStrict(err error)
```

###### io.Writer recently use for log panics in recovery middlewares like gin-framefork
