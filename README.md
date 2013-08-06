# Package `color`

Colorize your terminal strings.

![A coloured terminal](https://s3-us-west-2.amazonaws.com/aybabtme/color_demo.png "A fine terminal")

```go
yel := color.Yellow()
fmt.Printf("This is %s\n", yel.Get("yellow"))

redBg := yel.WithBackground(color.RedPaint)
fmt.Printf("This is %s\n", redBg.Get("yellow with red background"))

// yel is not changed! Styles are immutable!

yel.Print("HELLO STDOUT!!! CAN YOU SEE ME???")
```

That's it!

# It's spelled "colour"

NO!

# You're canadian, spell it "colour"

NO!
