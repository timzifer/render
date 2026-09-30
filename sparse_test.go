package render

import (
	"errors"
	"image/color"
	"testing"

	"github.com/go-gfx/gfx/raster"
	"github.com/go-gfx/gfx/vector"
	"github.com/timzifer/stilus"
)

func TestBackendsAndTransparentPaper(t *testing.T) {
	d := onePage(t, [4]float64{0, 0, 20, 20}, "1 0 0 rg 2 2 8 8 re f 0 0 1 RG 2 w 2 15 m 5 12 10 18 15 15 c S", nil)
	for _, backend := range []Backend{BackendStilus, BackendGFX} {
		t.Run(map[Backend]string{BackendStilus: "stilus", BackendGFX: "gfx"}[backend], func(t *testing.T) {
			img, err := Page(d, 1, Options{Backend: backend, Background: &color.RGBA{}})
			if err != nil {
				t.Fatal(err)
			}
			if c := img.At(5, 15); c.R != 255 || c.A != 255 {
				t.Fatalf("filled interior: %v", c)
			}
			if c := img.At(19, 19); c.A != 0 {
				t.Fatalf("paper gained alpha: %v", c)
			}
			r := &renderer{img: raster.New(20, 20)}
			if backend == BackendStilus {
				r.sparse = &sparseRaster{}
			}
			p := r.newPath()
			p.MoveTo(2, 2)
			p.QuadTo(8, 1, 12, 8)
			p.LineTo(2, 8)
			p.Close()
			g := gstate{clip: &clip{ox: 0, oy: 0, w: 20, h: 20, rect: &vector.Rect{X0: 0, Y0: 0, X1: 20, Y1: 20}}, softMask: make([]uint8, 400)}
			for i := range g.softMask {
				g.softMask[i] = 128
			}
			r.paint(&g, nil, 0, 0, 0, 0, color.RGBA{A: 255}, 1)
			r.painted = make([]float32, 400)
			for i := 0; i < 2; i++ {
				r.drawPath(&g, p, vector.NonZero, nil, nil, color.RGBA{R: 255, A: 255}, 0.5, nil)
			}
			if c := r.img.At(4, 5); c.R != 255 || c.A < 100 || c.A > 115 {
				t.Fatalf("masked alpha: %v", c)
			}
			if r.painted[5*20+4] < 0.4 {
				t.Fatal("alpha mask did not record painted coverage")
			}
		})
	}
}

func TestSparseGeometryBudgetsReturnPartialImage(t *testing.T) {
	r := &renderer{img: raster.New(20, 20), sparse: &sparseRaster{}}
	r.sparse.rz.MaxEdges = 1
	p := r.newPath()
	p.MoveTo(2, 2)
	p.LineTo(18, 2)
	p.LineTo(18, 18)
	p.LineTo(2, 18)
	p.Close()
	r.drawPath(&gstate{}, p, vector.NonZero, nil, nil, color.RGBA{A: 255}, 1, nil)
	if !errors.Is(r.rasterErr, ErrRasterBudget) {
		t.Fatalf("edge budget: %v", r.rasterErr)
	}
	r = &renderer{img: raster.New(20, 20), sparse: &sparseRaster{opaque: true}}
	r.sparse.canvas.MaxClipDepth = 1
	p = r.newPath()
	p.MoveTo(2, 2)
	p.LineTo(18, 2)
	p.LineTo(18, 18)
	p.Close()
	r.drawPath(&gstate{clip: &clip{ox: 0, oy: 0, w: 20, h: 20, rect: &vector.Rect{X0: 0, Y0: 0, X1: 20, Y1: 20}}}, p, vector.NonZero, nil, nil, color.RGBA{A: 255}, 1, nil)
	if !errors.Is(r.rasterErr, stilus.ErrClipDepth) {
		t.Fatalf("clip budget: %v", r.rasterErr)
	}
}

func TestSparseNestedFractionalClipsMultiply(t *testing.T) {
	d := onePage(t, [4]float64{0, 0, 20, 20}, "2.5 2 10 10 re W n 2.5 2 10 10 re W n 0 0 20 20 re f", nil)
	sparse, err := Page(d, 1, Options{})
	if err != nil {
		t.Fatal(err)
	}
	gfx, err := Page(d, 1, Options{Backend: BackendGFX})
	if err != nil {
		t.Fatal(err)
	}
	a, b := sparse.At(2, 12), gfx.At(2, 12)
	if abs(int(a.R)-int(b.R)) > 1 {
		t.Fatalf("repeated fractional clip: stilus %v, gfx %v", a, b)
	}
}
