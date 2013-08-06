# Package `color`

Colorize your terminal strings.

```go
yel := color.Yellow()
fmt.Printf("This is %s\n", yel.Get("yellow"))

redBg := yel.WithBackground(color.RedPaint)
fmt.Printf("This is %s\n", redBg.Get("yellow with red background"))
```

That's it for now.
