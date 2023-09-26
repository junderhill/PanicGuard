# PanicGuard
Go Package to safely call a function which may panic. The returned error can then be handled gracefully.

## Usage
```go
str, err := preventpanic.RunWithReturn[string](func() string {
    xraysegment.AddAnnotation("token", token)
    return "Hello World"
})

err := preventpanic.Run(func() {
    xraysegment.AddAnnotation("token", token)
})
```