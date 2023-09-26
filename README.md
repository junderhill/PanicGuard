# PanicGuard
Go Package to safely call a function which may panic. The returned error can then be handled gracefully.

## Usage
```go
preventpanic.Run[any](func() any {
    xraysegment.AddAnnotation("token", token)
    return nil
})
```