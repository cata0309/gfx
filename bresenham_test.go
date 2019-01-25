package gfx

func ExampleDrawBresenhamLine() {
	dst := NewPaletted(10, 5, Palette1Bit, ColorWhite)

	DrawBresenhamLine(dst, V(1, 1), V(8, 3), ColorBlack)

	for y := 0; y < dst.Bounds().Dy(); y++ {
		for x := 0; x < dst.Bounds().Dx(); x++ {
			if dst.Index(x, y) == 0 {
				Printf("▓")
			} else {
				Printf("░")
			}
		}
		Printf("\n")
	}

	// Output:
	//
	//░░░░░░░░░░
	//░▓▓░░░░░░░
	//░░░▓▓▓▓░░░
	//░░░░░░░▓▓░
	//░░░░░░░░░░
	//
}