# Package `color`

Colorize your terminal strings.

# Usage

```go
// Default Brush are available for your convenience.  You can invoke
// them directly
yel := color.Yellow()
fmt.Printf("This is %s\n", yel("yellow"))

// or you can create new ones!
weird := color.NewBrush(color.PurplePaint, color.CyanPaint)
fmt.Printf("This color is %s\n", weird("weird"))

// Create a Style, which has convenience methods
redBg := color.NewStyle(color.RedPaint, color.YellowPaint)

// Style.WithForeground or WithBackground returns a new Style, with the applied
// Paint.  Styles are immutable so the original one is left unchanged
greenFg := redBg.WithForeground(color.GreenPaint)

// Style.Brush gives you a Brush that you can invoke directly to colorize strings.
green := greenFg.Brush()
fmt.Printf("This is %s\n", green("green"))

```

That's it!

# Demo

![A coloured terminal](https://s3-us-west-2.amazonaws.com/aybabtme/color_demo.png "A fine terminal")

# Docs

[GoDoc!](http://godoc.org/github.com/aybabtme/color) (↫ this is a link)

# FAQ

> It's spelled "colour"

NO!

> You're canadian, spell it "colour"

NO!
