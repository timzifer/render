package render

import (
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/go-pdfkit/reader"
)

// slowPage builds a page that asks for a great deal of drawing: many thousands
// of filled rectangles, which is what a plotting tool writes and what takes a
// renderer a long time.
func slowPage(t *testing.T, rects int) *reader.Document {
	t.Helper()
	var content strings.Builder
	for i := 0; i < rects; i++ {
		fmt.Fprintf(&content, "%d %d 40 40 re f\n", i%60, (i*7)%60)
	}
	return shadedPage(t, content.String(), func(w *reader.Writer) reader.Dict {
		return reader.Dict{}
	})
}

func TestAPageMayBeGivenOnlySoLong(t *testing.T) {
	// Some pages take a very long time, and a caller drawing somebody else's
	// file cannot afford to wait for the worst of them. What comes back is
	// how far it got, and an error saying so.
	const rects = 40000

	// Calibrate the budget against this backend on this machine. Sparse fills
	// finish this page in less than the old fixed fifty millisecond budget.
	// Decode time is paid before the drawing deadline starts.
	d := slowPage(t, rects)
	whole := time.Now()
	if _, err := Page(d, 1, Options{Scale: 1}); err != nil {
		t.Fatal(err)
	}
	unbounded := time.Since(whole)

	budget := unbounded / 20
	start := time.Now()
	img, err := Page(d, 1, Options{Scale: 1, MaxDuration: budget})
	took := time.Since(start)
	if !errors.Is(err, ErrTimedOut) {
		t.Fatalf("a page given %s came back with %v after %s", budget, err, took)
	}
	if img == nil {
		t.Fatal("nothing came back at all; half a page is worth more than none")
	}
	if img.W == 0 || img.H == 0 {
		t.Fatalf("what came back is %dx%d", img.W, img.H)
	}
	// It has to stop near when it was told to, not merely eventually.
	// Decoding happens before the drawing deadline starts and dominates the
	// sparse backend on this tiny page. Include that fixed cost in the bound.
	if took > unbounded+budget {
		t.Errorf("it took %s to give up on %s, against %s to draw the same page whole",
			took, budget, unbounded)
	}
}

func TestAPageGivenNoLimitIsDrawnWhole(t *testing.T) {
	// Zero means as long as it takes, which is what this did before there was
	// anywhere to say otherwise.
	d := slowPage(t, 200)
	img, err := Page(d, 1, Options{Scale: 1})
	if err != nil {
		t.Fatal(err)
	}
	if isWhite(img, 20, 30) {
		t.Error("the page came back blank")
	}
}

func TestAPageThatFinishesInTimeSaysNothingAboutIt(t *testing.T) {
	d := slowPage(t, 50)
	img, err := Page(d, 1, Options{Scale: 1, MaxDuration: time.Minute})
	if err != nil {
		t.Fatalf("a page with a minute to draw in came back with %v", err)
	}
	if isWhite(img, 20, 30) {
		t.Error("the page came back blank")
	}
}

func TestLookingAtTheClock(t *testing.T) {
	// The clock is looked at once every so many operations, because asking the
	// machine the time is dear beside drawing a line. Once the time has gone,
	// it is gone: the answer does not depend on being asked again.
	r := &renderer{}
	if r.overrun() {
		t.Error("a page with no deadline said it had run out of time")
	}
	r.deadline = time.Now().Add(-time.Second)
	for i := 0; i < timeCheckEvery-1; i++ {
		if r.overrun() {
			t.Fatalf("it looked at the clock after %d operations, not %d", i+1, timeCheckEvery)
		}
	}
	if !r.overrun() {
		t.Fatal("it never looked at the clock at all")
	}
	if !r.overrun() {
		t.Error("having run out of time, it changed its mind")
	}

	fresh := &renderer{deadline: time.Now().Add(time.Hour)}
	for i := 0; i < timeCheckEvery+2; i++ {
		if fresh.overrun() {
			t.Fatal("a page with an hour to draw in said its time had gone")
		}
	}
}
