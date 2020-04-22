# Trace Parent

minimal trace parent library to parse and propergate to new span

## Usage

```go
    import "github.com/pallat/traceparent"

    func main() {
        tp := traceparent.Parse("00-0af7651916cd43dd8448eb211c80319c-b7ad6b7169203331-01")

        span := tp.NewSpan()

        fmt.Println(tp, span)
    }
```