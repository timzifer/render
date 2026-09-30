package render

import (
	"errors"
	"image"
	"image/color"
	"math"

	"github.com/go-gfx/gfx/vector"
	"github.com/go-pdfkit/reader"
	"github.com/timzifer/stilus"
)

// Backend selects the rasteriser used for paths, strokes, glyphs and clips.
type Backend uint8

const (
	BackendStilus Backend = iota // sparse exact-area rasterisation (default)
	BackendGFX                   // upstream go-gfx coverage rasterisation
)

// ErrRasterBudget reports truncated geometry; Page returns the partial image.
var ErrRasterBudget = errors.New("render: rasterisation geometry budget exceeded")

// renderPath records only the selected backend's geometry. Coordinates are
// already in device space; stroke widths retain upstream's scaleOf behaviour.
type renderPath struct {
	gfx      vector.Path
	isSparse bool
	sparse   stilus.Path
}

func (r *renderer) newPath() *renderPath {
	return &renderPath{isSparse: r.sparse != nil}
}
func (p *renderPath) MoveTo(x, y float64) {
	if !p.isSparse {
		p.gfx.MoveTo(x, y)
	} else {
		p.sparse.MoveTo(float32(x), float32(y))
	}
}
func (p *renderPath) LineTo(x, y float64) {
	if !p.isSparse {
		p.gfx.LineTo(x, y)
	} else {
		p.sparse.LineTo(float32(x), float32(y))
	}
}
func (p *renderPath) QuadTo(cx, cy, x, y float64) {
	if !p.isSparse {
		p.gfx.QuadTo(cx, cy, x, y)
	} else {
		p.sparse.QuadTo(float32(cx), float32(cy), float32(x), float32(y))
	}
}
func (p *renderPath) CubicTo(ax, ay, bx, by, x, y float64) {
	if !p.isSparse {
		p.gfx.CubicTo(ax, ay, bx, by, x, y)
	} else {
		p.sparse.CubicTo(float32(ax), float32(ay), float32(bx), float32(by), float32(x), float32(y))
	}
}
func (p *renderPath) Close() {
	if !p.isSparse {
		p.gfx.Close()
	} else {
		p.sparse.Close()
	}
}
func ptrStyle(s vector.StrokeStyle) *vector.StrokeStyle { return &s }
func sparseStyle(s *vector.StrokeStyle) stilus.StrokeStyle {
	return stilus.StrokeStyle{Width: s.Width, Cap: stilus.Cap(s.Cap), Join: stilus.Join(s.Join), MiterLimit: s.MiterLimit, Dash: s.Dash, DashPhase: s.DashPhase}
}

type sparseRaster struct {
	rz      stilus.Rasterizer
	stroker stilus.Stroker
	canvas  stilus.Canvas
	dst     image.RGBA
	opaque  bool
	cov     []float64
	dense   coverageSpans
	paint   paintSpans
}

func (r *renderer) sparseError() {
	if r.sparse.rz.Truncated() || r.sparse.stroker.Truncated() {
		r.rasterErr = ErrRasterBudget
	}
}
func (r *renderer) rasterSparse(p *renderPath, rule vector.FillRule, st *vector.StrokeStyle, box image.Rectangle, b stilus.Blitter) {
	s := r.sparse
	if s.rz.Clip() != box {
		s.rz.SetClip(box)
	} else {
		s.rz.Reset()
	}
	if st == nil {
		s.rz.Fill(&p.sparse, stilus.Identity, stilus.FillRule(rule), b)
	} else {
		style := sparseStyle(st)
		s.stroker.Stroke(&s.rz, &p.sparse, stilus.Identity, &style)
		s.rz.Rasterize(stilus.NonZero, b)
	}
	r.sparseError()
}

// pathCoverage is needed only when the consumer requires a grid (clip or
// pattern). Solid marks stream directly to the destination in drawPath.
func (r *renderer) pathCoverage(p *renderPath, rule vector.FillRule, st *vector.StrokeStyle) (cov []float64, ox, oy, w, h int, ok bool) {
	if r.sparse == nil {
		if st == nil {
			return r.rz.Fill(&p.gfx, rule, r.img.W, r.img.H)
		}
		return r.rz.StrokeWith(&p.gfx, *st, r.img.W, r.img.H)
	}
	bb := p.sparse.Bounds()
	pad := 1.0
	if st != nil {
		pad += max(st.Width, 1) * max(st.MiterLimit, 2) / 2
	}
	// Clamp as floats before converting to int, including hostile coordinates.
	x0 := max(0, math.Floor(bb.X0-pad))
	y0 := max(0, math.Floor(bb.Y0-pad))
	x1 := min(float64(r.img.W), math.Ceil(bb.X1+pad))
	y1 := min(float64(r.img.H), math.Ceil(bb.Y1+pad))
	if !(x0 < x1 && y0 < y1) {
		return
	}
	box := image.Rect(int(x0), int(y0), int(x1), int(y1))
	ox, oy, w, h = box.Min.X, box.Min.Y, box.Dx(), box.Dy()
	s := r.sparse
	if cap(s.cov) < w*h {
		s.cov = make([]float64, w*h)
	} else {
		s.cov = s.cov[:w*h]
		clear(s.cov)
	}
	s.dense = coverageSpans{cov: s.cov, box: box}
	r.rasterSparse(p, rule, st, box, &s.dense)
	return s.cov, ox, oy, w, h, true
}

type coverageSpans struct {
	cov []float64
	box image.Rectangle
}

func (b *coverageSpans) BlitRun(y, x0, x1 int, a uint8) {
	start := (y-b.box.Min.Y)*b.box.Dx() + x0 - b.box.Min.X
	for i := 0; i < x1-x0; i++ {
		b.cov[start+i] = float64(a) / 255
	}
}
func (b *coverageSpans) BlitCoverage(y, x int, cov []uint8) {
	start := (y-b.box.Min.Y)*b.box.Dx() + x - b.box.Min.X
	for i, a := range cov {
		b.cov[start+i] = float64(a) / 255
	}
}

func (r *renderer) drawPath(g *gstate, p *renderPath, rule vector.FillRule, st *vector.StrokeStyle, pat *pattern, c color.RGBA, alpha float64, resources reader.Dict) {
	if r.sparse == nil || pat != nil {
		if cov, ox, oy, w, h, ok := r.pathCoverage(p, rule, st); ok {
			r.paintCoverage(g, pat, cov, ox, oy, w, h, c, alpha, resources)
		}
		return
	}
	if alpha <= 0 || c.A == 0 {
		return
	}
	s := r.sparse
	box := image.Rect(0, 0, r.img.W, r.img.H)
	if g.clip != nil {
		box = box.Intersect(image.Rect(g.clip.ox, g.clip.oy, g.clip.ox+g.clip.w, g.clip.oy+g.clip.h))
	}
	if box.Empty() {
		return
	}
	// On opaque paper, straight and premultiplied storage agree. Canvas enables
	// Stilus' analytic stroke and solid-run kernels without changing pixel format.
	if s.opaque && alpha == 1 && c.A == 255 && g.softMask == nil && r.painted == nil && rectangularClip(g.clip) {
		s.dst = image.RGBA{Pix: r.img.Pix, Stride: r.img.W * 4, Rect: image.Rect(0, 0, r.img.W, r.img.H)}
		s.canvas.Reset(&s.dst, box)
		for clip := g.clip; clip != nil; clip = clip.under {
			v := clip.rect
			s.canvas.ClipRect(stilus.Rect{X0: v.X0, Y0: v.Y0, X1: v.X1, Y1: v.Y1}, stilus.Identity)
		}
		paint := stilus.Paint{Color: c}
		if st == nil {
			s.canvas.Fill(&p.sparse, stilus.Identity, stilus.FillRule(rule), &paint)
		} else {
			style := sparseStyle(st)
			s.canvas.Stroke(&p.sparse, stilus.Identity, &style, &paint)
		}
		if err := s.canvas.Err(); err != nil {
			r.rasterErr = err
		}
		return
	}
	s.paint = paintSpans{r: r, g: g, c: c, alpha: alpha}
	r.rasterSparse(p, rule, st, box, &s.paint)
}
func rectangularClip(c *clip) bool {
	fractional := false
	for ; c != nil; c = c.under {
		if c.rect == nil {
			return false
		}
		v := c.rect
		if v.X0 != math.Trunc(v.X0) || v.Y0 != math.Trunc(v.Y0) || v.X1 != math.Trunc(v.X1) || v.Y1 != math.Trunc(v.Y1) {
			// Canvas intersects rectangles geometrically. PDF coverage must
			// multiply if more than one clip has a fractional border.
			if fractional {
				return false
			}
			fractional = true
		}
	}
	return true
}

// paintSpans preserves go-gfx's straight-alpha output, PDF clip multiplication
// and soft-mask coverage recording without visiting uncovered pixels.
type paintSpans struct {
	r     *renderer
	g     *gstate
	c     color.RGBA
	alpha float64
}

func (b *paintSpans) BlitRun(y, x0, x1 int, a uint8) {
	for x := x0; x < x1; x++ {
		b.pixel(x, y, a)
	}
}
func (b *paintSpans) BlitCoverage(y, x int, cov []uint8) {
	for i, a := range cov {
		b.pixel(x+i, y, a)
	}
}
func (b *paintSpans) pixel(x, y int, a uint8) {
	cov := float64(a) / 255 * b.alpha
	if b.g.clip != nil {
		cov *= b.g.clip.at(x, y)
	}
	if b.g.softMask != nil {
		cov *= maskLevel(b.g.softMask[y*b.r.img.W+x])
	}
	if cov <= 0 {
		return
	}
	pix := b.r.img.Pix[(y*b.r.img.W+x)*4:][:4]
	sa := float64(b.c.A) / 255 * cov
	da := float64(pix[3]) / 255
	out := sa + da*(1-sa)
	inv := da * (1 - sa)
	pix[0] = byteOf((float64(b.c.R)*sa + float64(pix[0])*inv) / out / 255)
	pix[1] = byteOf((float64(b.c.G)*sa + float64(pix[1])*inv) / out / 255)
	pix[2] = byteOf((float64(b.c.B)*sa + float64(pix[2])*inv) / out / 255)
	pix[3] = byteOf(out)
	b.r.markPixel(x, y, cov)
}
