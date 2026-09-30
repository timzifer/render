# Stilus in go-pdfkit/render: end-to-end benchmark

2026-09-30 · Apple M4 (darwin/arm64, 10 logical CPUs) · Go 1.26.4 · 150 dpi.

41 PDFs, 54 pages (up to the first three pages per document), one warm-up and five timed renders per page/backend. Each number below is the median of those five renders. Both backends run the same PDF interpreter, fonts, images, shadings and background initialization. PDF parsing/opening is outside the page timer, while content decoding, geometry construction and drawing are inside. Go GC remains enabled. Engines run sequentially; this is one machine and one corpus, not a confidence interval or a general PDF performance guarantee.

Upstream: `ee4e85419eb648d94a73b389e5bc5cfe5c5af0b5` (render v0.67.0). Integrated source: `1b5ffd9` in the render fork. Sparse core: `11ece9a90ee57ffc9a709e330a22a38b44abda7a`. The baseline is the selectable `BackendGFX` in the fork; it retains the upstream rasterisation/compositing algorithms behind the common path adapter.

Sum of page medians: **43.892 s → 4.420 s (9.93×, 89.9% less time)**. Median of the 54 per-page speedup ratios: **1.57×**. Long hatch pages dominate the summed result, so the categories below are more useful than a single overall factor.

| Workload | Pages | go-gfx sum (ms) | Stilus sum (ms) | Sum speedup | Median page speedup |
|---|---:|---:|---:|---:|---:|
| Synthetic drawings | 8 | 39693.5 | 610.1 | 65.06× | 7.36× |
| arXiv papers | 12 | 436.0 | 232.9 | 1.87× | 2.20× |
| PDF.js feature corpus | 34 | 3762.8 | 3576.8 | 1.05× | 1.16× |

| Synthetic PDF page | go-gfx (ms) | Stilus (ms) | Speedup | Time reduction |
|---|---:|---:|---:|---:|
| contours-3000 | 567.70 | 52.80 | 10.8× | 90.7% |
| glyphs-30000-uncached | 320.26 | 256.26 | 1.2× | 20.0% |
| hatch-2000-0.35mm | 12773.96 | 33.60 | 380.2× | 99.7% |
| hatch-2000-hairline | 12748.72 | 29.38 | 434.0× | 99.8% |
| hatch-2000-maskclip | 12854.73 | 80.41 | 159.9× | 99.4% |
| mixed-drawing | 213.37 | 53.85 | 4.0× | 74.8% |
| short-20000-0.35mm | 111.03 | 53.18 | 2.1× | 52.1% |
| short-20000-hairline | 103.72 | 50.62 | 2.0× | 51.2% |

## Allocations and image quality

- `pdfkit`: median 92,430 allocations/page and 90.04 MB allocated/page; total 59.433 GB allocated across the 54 single measured renders. These are allocation volumes, not peak resident memory.
- `pdfkit-stilus`: median 56,432 allocations/page and 22.38 MB allocated/page; total 4.426 GB allocated across the 54 single measured renders. These are allocation volumes, not peak resident memory.

Against PDFium (WASM), 15/54 go-gfx pages and 16/54 Stilus pages pass the existing target: equal dimensions, mean grayscale difference < 1/255, and < 0.5% of pixels differing by more than 32/255. 0 pages change from passing to failing. The existing PDF pipeline still has font, transparency, shading and page-dimension differences from PDFium; the majority of pages do not meet the strict reference target. This integration changes antialiasing and float32 outline rounding, and does not fix those front-end/compositing limitations.

0 rendering errors across all engines. 5 pages are slower with Stilus in this run; image/shading-dominated pages have little benefit and small differences should be treated as measurement noise.

## Validation

`go test -covermode=set ./...` passes with exact 100% statement coverage, including both backends, transparent backgrounds, alpha coverage recording, nested fractional clip multiplication and geometry budgets. `go test -race ./...` and `go vet ./...` pass. Cross-compilation covers the upstream target matrix (see validation.txt). The timeout test now calibrates its budget because the sparse backend completes its original workload within the old fixed budget.

## Reproduce

Check out Stilus at `11ece9a90ee57ffc9a709e330a22a38b44abda7a`, then apply [harness.patch](harness.patch) with `git -C /path/to/stilus apply /path/to/render/benchmarks/harness.patch`. The patch adds the backend selection and pins the published fork revision. From the resulting Stilus repository:

```sh
cd harness
go run ./cmd/pdfbench fetch -dir cmd/pdfbench/corpus
go run ./cmd/pdfbench scenes -dir cmd/pdfbench/corpus
go run ./cmd/pdfbench run -corpus cmd/pdfbench/corpus \
  -engines pdfkit,pdfkit-stilus,pdfium -ref pdfium \
  -dpi 150 -runs 5 -pages 3 -split=false -out report
```

`results.csv` contains per-page minimum/median times, allocation counts and PDFium comparison metrics. `pdfbench-report.md` contains the full per-page report. `corpus-sha256.txt` identifies every input. To use this fork in another application, retain the upstream import and replace its module with `github.com/timzifer/render` at the pinned integration revision.
