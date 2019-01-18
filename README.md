# gfx

[![Build Status](https://travis-ci.org/peterhellberg/gfx.svg?branch=master)](https://travis-ci.org/peterhellberg/gfx)
[![Go Report Card](https://goreportcard.com/badge/github.com/peterhellberg/gfx?style=flat)](https://goreportcard.com/report/github.com/peterhellberg/gfx)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/peterhellberg/gfx)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](https://github.com/peterhellberg/gfx#license-mit)

#### :warning: NO STABILITY GUARANTEES AS OF YET

Convenience package for dealing with graphics in my pixel drawing experiments.

## Geometry

The geometry types is based on those found in <https://github.com/faiface/pixel> (but indended for use without Pixel)

### `Vec` a 2D vector type

### `Rect` is a 2D rectangle

## Matrix

Matrix is a 2x3 affine matrix that can be used for all kinds of spatial transforms, such as movement, scaling and rotations.

## Line algorithms

### Bresenham's line algorithm

<http://en.wikipedia.org/wiki/Bresenham's_line_algorithm>

## Polygon

A Polygon is represented by a list of vectors.

### Polyline

A Polyline is a slice of polygons forming a line.

## `Turtle` drawing :turtle:

A small Turtle inspired drawing type. (`Resize`, `Turn`, `Move`, `Forward`, `Draw`)

<https://www.cse.wustl.edu/~taoju/research/TurtlesforCADRevised.pdf>

## Animation

There is rudimentary support for making animations using this package, the animations can then be encoded into GIF.

## Colors

There are a few default colors in this package, convenient when you just want to experiment,
for more ambitious projects I suggest using the Palette support (or even one of the included palettes).

### Default colors


| Variable               | Color
|------------------------|---------------------------------------------------------
| `gfx.ColorBlack`       | ![gfx.ColorBlack](https://fakeimg.pl/128x32/000000?text=+)
| `gfx.ColorWhite`       | ![gfx.ColorWhite](https://fakeimg.pl/128x32/FFFFFF?text=+)
| `gfx.ColorTransparent` | ![gfx.ColorTransparent](https://fakeimg.pl/128x32/000000,0/?text=+)
| `gfx.ColorOpaque`      | ![gfx.ColorOpaque](https://fakeimg.pl/128x32/FFFFFF?text=+)
| `gfx.ColorRed`         | ![gfx.ColorRed](https://fakeimg.pl/128x32/FF0000?text=+)
| `gfx.ColorGreen`       | ![gfx.ColorGreen](https://fakeimg.pl/128x32/00FF00?text=+)
| `gfx.ColorBlue`        | ![gfx.ColorBlue](https://fakeimg.pl/128x32/0000FF?text=+)
| `gfx.ColorCyan`        | ![gfx.ColorCyan](https://fakeimg.pl/128x32/00FFFF?text=+)
| `gfx.ColorMagenta`     | ![gfx.ColorMagenta](https://fakeimg.pl/128x32/FF00FF?text=+)
| `gfx.ColorYellow`      | ![gfx.ColorYellow](https://fakeimg.pl/128x32/FFFF00?text=+)

### Palettes

There are a number of palettes in the `gfx` package,
most of them are found in the [Lospec Palette List](https://lospec.com/palette-list/).

- `gfx.Palette1Bit`
- `gfx.Palette2BitGrayScale`
- `gfx.Palette3Bit`
- `gfx.PaletteCGA`
- `gfx.Palette15PDX`
- `gfx.Palette20PDX`
- `gfx.PaletteAAP16`
- `gfx.PaletteAAP64`
- `gfx.PaletteSplendor128`
- `gfx.PaletteArne16`
- `gfx.PaletteFamicube`
- `gfx.PaletteEDG16`
- `gfx.PaletteEDG32`
- `gfx.PaletteEDG36`
- `gfx.PaletteEDG64`
- `gfx.PaletteEDG8`
- `gfx.PaletteEN4`
- `gfx.PaletteARQ4`
- `gfx.PaletteInk`
- `gfx.PaletteAmmo8`
- `gfx.PaletteNYX8`
- `gfx.PaletteNight16`
- `gfx.PalettePICO8`

### Errors

There is a `gfx.Error` type.

> If you are using [Ebiten](https://github.com/hajimehoshi/ebiten) then you can return the provided `gfx.ErrDone` error to exit its run loop.

### HTTP

You can use `gfx.GetPNG` to download and decode a PNG given an URL.

### Log

I find that it is fairly common for me to do some logging driven development
when experimenting with graphical effects, so I've included `gfx.Log`,
`gfx.Dump`, `gfx.Printf` and `gfx.Sprintf` in this package.

### Resizing images

You can use `gfx.ResizeImage` to resize an image. (nearest neighbor, mainly useful for pixelated graphics)

## License (MIT)

Copyright (c) 2019 [Peter Hellberg](https://c7.se)

> Permission is hereby granted, free of charge, to any person obtaining
> a copy of this software and associated documentation files (the
> "Software"), to deal in the Software without restriction, including
> without limitation the rights to use, copy, modify, merge, publish,
> distribute, sublicense, and/or sell copies of the Software, and to
> permit persons to whom the Software is furnished to do so, subject to
> the following conditions:
>
> The above copyright notice and this permission notice shall be
> included in all copies or substantial portions of the Software.

<img src="https://data.gopher.se/gopher/viking-gopher.svg" align="right" width="230" height="230">

> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
> EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
> MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
> NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
> LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
> OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
> WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
