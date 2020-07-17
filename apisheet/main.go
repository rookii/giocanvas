package main

import (
	"flag"
	"image/color"
	"math"
	"os"

	"gioui.org/app"
	"gioui.org/io/key"
	"gioui.org/io/system"
	"gioui.org/unit"
	"github.com/ajstarks/giocanvas"
)

func ref(title string, width, height float32) {
	defer os.Exit(0)
	win := app.NewWindow(app.Title(title), app.Size(unit.Px(width), unit.Px(height)))
	var col1, col2, col3 float32 = 15, 50, 85
	var top, subtop float32 = 92, 82
	var titlesize, headsize, apisize, dotsize float32 = 4, 3, 0.9, 0.3
	titlecolor := giocanvas.ColorLookup("black")
	subcolor := giocanvas.ColorLookup("gray")
	dotcolor := giocanvas.ColorLookup("rgb(100,100,100,180)")
	bgcolor := giocanvas.ColorLookup("rgb(250,250,250)")
	apicolor := giocanvas.ColorLookup("rgb(75,75,75)")
	shapecolor := color.RGBA{70, 130, 180, 150}
	tcolor := titlecolor
	quote := "If there is no struggle, there is no progress. Those who profess to favor freedom, and yet depreciate agitation, are men who want crops without plowing up the ground. They want rain without thunder and lightning. They want the ocean without the awful roar of its many waters."
	for e := range win.Events() {
		switch e := e.(type) {
		case system.FrameEvent:
			canvas := giocanvas.NewCanvas(width, height, e)
			canvas.Background(bgcolor)
			canvas.CText(50, top, titlesize, "Giocanvas API Reference", titlecolor)
			canvas.CText(col1, subtop, headsize, "Text", subcolor)
			canvas.CText(col2, subtop, headsize, "Graphics", subcolor)
			canvas.CText(col3, subtop, headsize, "Transforms", subcolor)
			canvas.CText(col3, subtop-60, headsize, "Image", subcolor)

			// Text
			y := subtop - 10
			canvas.Text(col1, y, headsize, "hello", tcolor)
			canvas.Circle(col1, y, dotsize, dotcolor)
			canvas.CText(col1, y-5, apisize, "Text(x, y, size float32, s string, c color.RGBA)", apicolor)

			y -= 15
			canvas.EText(col1, y, headsize, "hello", tcolor)
			canvas.Circle(col1, y, dotsize, dotcolor)
			canvas.CText(col1, y-5, apisize, "EText(x, y, size float32, s string, c color.RGBA)", apicolor)

			y -= 15
			canvas.CText(col1, y, headsize, "hello", tcolor)
			canvas.Circle(col1, y, dotsize, dotcolor)
			canvas.CText(col1, y-5, apisize, "CText(x, y, size float32, s string, c color.RGBA)", apicolor)

			y -= 15
			canvas.TextWrap(col1-10, y, headsize*0.4, 18, quote, tcolor)
			canvas.Circle(col1-10, y, dotsize, dotcolor)
			canvas.CText(col1, y-20, apisize, "TextWrap(x, y, size, width float32, s string, c color.RGBA)", apicolor)

			// graphics
			x1 := col2 - 10
			x2 := col2 + 10

			y = subtop - 10
			canvas.Line(x1, y, x2, y, 0.2, shapecolor)
			canvas.Circle(x1, y, dotsize, dotcolor)
			canvas.Circle(x2, y, dotsize, dotcolor)
			canvas.CText(col2, y-5, apisize, "Line(x1, y1, x2, y2, width float32, c color.RGBA)", apicolor)

			y -= 15
			canvas.Circle(x1, y, 2.5, shapecolor)
			canvas.Ellipse(x2, y, 5, 2.5, shapecolor)
			canvas.CText(x1, y-5, apisize, "Circle(x, y, size float32, c color.RGBA)", apicolor)
			canvas.CText(x2, y-5, apisize, "Ellipse(x, y, w, h float32, c color.RGBA)", apicolor)
			canvas.Circle(x1, y, dotsize, dotcolor)
			canvas.Circle(x2, y, dotsize, dotcolor)

			y -= 15
			canvas.Square(x1, y, 5, shapecolor)
			canvas.Rect(x2, y, 10, 5, shapecolor)
			canvas.CText(x1, y-5, apisize, "Square(x, y, size float32, c color.RGBA)", apicolor)
			canvas.CText(x2, y-5, apisize, "Rect(x, y, w, h float32, c color.RGBA)", apicolor)
			canvas.Circle(x1, y, dotsize, dotcolor)
			canvas.Circle(x2, y, dotsize, dotcolor)

			y -= 15
			canvas.Curve(x1-5, y, x1-5, y+7, col2-5, y, shapecolor)
			canvas.CText(x1, y-5, apisize, "QuadCurve(x, y, cx, cy, ex, ey float32, c color.RGBA)", apicolor)
			canvas.Circle(x1-5, y, dotsize, dotcolor)
			canvas.Circle(x1-5, y+7, dotsize, dotcolor)
			canvas.Circle(col2-5, y, dotsize, dotcolor)

			canvas.CubeCurve(col2+5, y, x2, y+5, col2+15, y+7, x2+5, y, shapecolor)
			canvas.CText(x2+5, y-5, apisize, "CubeCurve(x, y, x1, y1, x2, y2, ex, ey float32, c color.RGBA)", apicolor)
			canvas.Circle(col2+5, y, dotsize, dotcolor)
			canvas.Circle(x2, y+5, dotsize, dotcolor)
			canvas.Circle(col2+15, y+7, dotsize, dotcolor)
			canvas.Circle(x2+5, y, dotsize, dotcolor)

			y -= 15
			px := []float32{x1 - 5, x1, x1 + 5}
			py := []float32{y, y + 5, y}
			canvas.Polygon(px, py, shapecolor)
			canvas.CText(x1, y-5, apisize, "Polygon(x, y []float32, c color.RGBA)", apicolor)
			canvas.Circle(x1-5, y, dotsize, dotcolor)
			canvas.Circle(x1, y+5, dotsize, dotcolor)
			canvas.Circle(x1+5, y, dotsize, dotcolor)

			canvas.Arc(x2, y+5, 5, 0, math.Pi/2, shapecolor)
			canvas.Circle(x2, y+5, dotsize, dotcolor)
			canvas.CText(x2, y-5, apisize, "Arc(x, y, radius, a1, a2, c color.RGBA)", apicolor)

			// Transforms
			var rectw, recth, ts, ts2 float32
			var midx float32 = col3
			rectw = 10
			recth = rectw / 4
			ts = 1.5
			ts2 = ts / 3

			y = subtop - 10
			stack := canvas.Scale(midx, y, 1.5)
			canvas.CenterRect(midx, y, rectw, recth, shapecolor)
			canvas.TextMid(midx, y-ts2, ts, "scale", tcolor)
			giocanvas.EndTransform(stack)
			canvas.CText(col3, y-5, apisize, "Scale(x, y, factor float32) op.StackOp", apicolor)

			y -= 15
			stack = canvas.Shear(midx, y, math.Pi/4, 0)
			canvas.CenterRect(midx, y, rectw, recth, shapecolor)
			canvas.TextMid(midx, y-ts2, ts, "shear", tcolor)
			giocanvas.EndTransform(stack)
			canvas.CText(col3, y-5, apisize, "Shear(x, y, ax, ay float32) op.StackOp", apicolor)

			y -= 15
			stack = canvas.Rotate(midx, y, math.Pi/6)
			canvas.CenterRect(midx, y, rectw, recth, shapecolor)
			canvas.TextMid(midx, y-ts2, ts, "rotate", tcolor)
			giocanvas.EndTransform(stack)
			canvas.CText(col3, y-5, apisize, "Rotate(x, y, angle float32) op.StackOp", apicolor)

			y -= 30
			canvas.Image("earth.jpg", midx, y+2, 1000, 1000, 15)
			canvas.CText(midx, y-5, apisize, "Image(file string, x, y float32, w, h int, scale float32)", apicolor)
			canvas.CText(midx, y-7, apisize, "Img(img image.Image, x, y float32, w, h int, scale float32)", apicolor)

			e.Frame(canvas.Context.Ops)
		case key.Event:
			switch e.Name {
			case "Q", key.NameEscape:
				os.Exit(0)
			}
		}
	}
}

func main() {
	var w, h int
	flag.IntVar(&w, "width", 2400, "canvas width")
	flag.IntVar(&h, "height", 1800, "canvas height")
	flag.Parse()
	go ref("API Reference", float32(w), float32(h))
	app.Main()
}