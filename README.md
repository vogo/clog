Clog is a simple leveled golang logger, which adds context info into log. 

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

Clog writes log to stdout default, you can write log to a rolling file:

```golang
import	"gopkg.in/natefinch/lumberjack.v2"

SetOutput(&lumberjack.Logger{
	Filename:   file.Name(),
	MaxSize:    10, // megabytes
	MaxBackups: 3,
	MaxAge:     30,    //days
	Compress:   true, // disabled by default
})
```
