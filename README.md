**Clog is a simple leveled golang logger, which adds context info into log.**

Clog log api receives a `context.Context` variable as the first parameter:

```
clog.Debug(ctx, "hello word")
clog.Info(ctx, "this is %s", "clog")
```

## Context Fommater
Context info can be request id, message id, current user id, and so on.
You must custom your context formatter through function `clog.SetContextFommatter`, eg:

```golang
clog.SetContextFommatter(func(ctx context.Context) string {
	if rid, ok := ctx.Value("request_id").(string); ok {
		return s
	}
	return "-"
})
```

## Log Format

Clog format is fixed currently as follow:

```
yyyyMMdd HH:mm:ss.SSSSS <LEVEL> [<CONTEXT_INFO>] <LOG_MESSAGE> (<FILE_NAME>:<FILE_LINE>)
```

- `yyyyMMdd HH:mm:ss.SSSSS`: log time
- `<LEVEL>`: log level
- `<CONTEXT_INFO>`: context formmater output , eg, request id
- `<LOG_MESSAGE>`: log message
- `<FILE_NAME>`: the file to call Clog
- `<FILE_LINE>`: the line of file to call Clog

Log output sample:
```
20181018 17:12:40.97451   DEBUG [4a44f3d3-d1ab-47a8-92af-ca88102705fe] hello world (server.go:128)
20181018 17:12:40.97458   INFO  [4a44f3d3-d1ab-47a8-92af-ca88102705fe] this is clog (server.go:138)
```

## Output Writer
Clog receive a `io.Writer` to write output, default write to stdout, you can write log to a rolling file:

```golang
import	"gopkg.in/natefinch/lumberjack.v2"

clog.SetOutput(&lumberjack.Logger{
	Filename:   file.Name(),
	MaxSize:    10,       // megabytes
	MaxBackups: 3,        // backup number
	MaxAge:     30,       // days
	Compress:   true,     // disabled by default
})
```

## Log Level

Change global log level:

```golang
clog.SetLevelByString("info") // debug/info/warn/error/fatal
```

And `clog.GlobalLevel()` return the current global log level.



