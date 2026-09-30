# pdfbench report

2026-09-30 08:20 · 150 dpi · 5 timed runs/page after warm-up · max 3 pages/document · darwin/arm64, 10 CPUs · reference: pdfium

## Speed and allocations

| engine | pages | unsupported | errors | median ms/page | p90 ms/page | total ms | median allocs/page | median MB alloc/page |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| pdfkit | 54 | 0 | 0 | 32.2 | 318.2 | 43745 | 94477 | 88.6 |
| pdfkit-stilus | 54 | 0 | 0 | 14.9 | 79.9 | 4329 | 58627 | 22.0 |
| pdfium | 54 | 0 | 0 | 8.8 | 164.5 | 4565 | n/a (wasm) | n/a |

## Accuracy against pdfium

Target: same size, mean < 1/255 and < 0.5 % of pixels deviating by more than 32/255.

| engine | category | pages | pass | size mismatch | mean dev. (median) | >32/255 (median) | SSIM (median) |
|---|---|---:|---:|---:|---:|---:|---:|
| pdfkit | (all) | 54 | 15 | 1 | 2.64 | 1.52 % | 0.9799 |
| pdfkit | drawing | 8 | 5 | 0 | 0.69 | 0.04 % | 0.9991 |
| pdfkit | image | 4 | 2 | 0 | 6.91 | 3.47 % | 0.9948 |
| pdfkit | paper | 12 | 0 | 0 | 4.20 | 5.16 % | 0.9704 |
| pdfkit | scan | 3 | 1 | 0 | 0.59 | 0.58 % | 0.9908 |
| pdfkit | shading | 9 | 2 | 0 | 1.58 | 0.61 % | 0.9871 |
| pdfkit | text | 9 | 1 | 1 | 3.09 | 2.48 % | 0.9628 |
| pdfkit | transparency | 7 | 3 | 0 | 2.44 | 0.51 % | 0.9846 |
| pdfkit | vector | 2 | 1 | 0 | 1.10 | 0.86 % | 0.9997 |
| pdfkit-stilus | (all) | 54 | 16 | 1 | 2.49 | 1.52 % | 0.9803 |
| pdfkit-stilus | drawing | 8 | 6 | 0 | 0.63 | 0.01 % | 0.9993 |
| pdfkit-stilus | image | 4 | 2 | 0 | 6.80 | 3.47 % | 0.9948 |
| pdfkit-stilus | paper | 12 | 0 | 0 | 4.16 | 5.11 % | 0.9712 |
| pdfkit-stilus | scan | 3 | 1 | 0 | 0.59 | 0.58 % | 0.9908 |
| pdfkit-stilus | shading | 9 | 2 | 0 | 1.50 | 0.61 % | 0.9875 |
| pdfkit-stilus | text | 9 | 1 | 1 | 3.05 | 2.50 % | 0.9637 |
| pdfkit-stilus | transparency | 7 | 3 | 0 | 2.49 | 0.51 % | 0.9850 |
| pdfkit-stilus | vector | 2 | 1 | 0 | 1.10 | 0.86 % | 0.9998 |

## Pages

| file | page | engine | size | ms (min) | allocs | mean dev. | >32/255 | SSIM | note |
|---|---:|---|---|---:|---:|---:|---:|---:|---|
| arxiv/1406.2661v1.pdf | 1 | pdfium | 1275×1651 | 7.4 |  |  |  |  |  |
| arxiv/1406.2661v1.pdf | 1 | pdfkit | 1275×1651 | 25.1 | 90384 | 3.47 | 4.19 % | 0.9742 |  |
| arxiv/1406.2661v1.pdf | 1 | pdfkit-stilus | 1275×1651 | 11.6 | 47108 | 3.41 | 4.15 % | 0.9748 |  |
| arxiv/1406.2661v1.pdf | 2 | pdfium | 1275×1651 | 8.9 |  |  |  |  |  |
| arxiv/1406.2661v1.pdf | 2 | pdfkit | 1275×1651 | 33.7 | 121105 | 4.55 | 5.57 % | 0.9665 |  |
| arxiv/1406.2661v1.pdf | 2 | pdfkit-stilus | 1275×1651 | 14.9 | 61201 | 4.50 | 5.52 % | 0.9678 |  |
| arxiv/1406.2661v1.pdf | 3 | pdfium | 1275×1651 | 18.8 |  |  |  |  |  |
| arxiv/1406.2661v1.pdf | 3 | pdfkit | 1275×1651 | 42.2 | 257897 | 3.50 | 4.28 % | 0.9742 |  |
| arxiv/1406.2661v1.pdf | 3 | pdfkit-stilus | 1275×1651 | 22.5 | 193288 | 3.44 | 4.16 % | 0.9751 |  |
| arxiv/1512.03385v1.pdf | 1 | pdfium | 1275×1651 | 9.7 |  |  |  |  |  |
| arxiv/1512.03385v1.pdf | 1 | pdfkit | 1275×1651 | 36.4 | 121312 | 4.25 | 5.17 % | 0.9640 |  |
| arxiv/1512.03385v1.pdf | 1 | pdfkit-stilus | 1275×1651 | 16.1 | 65589 | 4.22 | 5.14 % | 0.9647 |  |
| arxiv/1512.03385v1.pdf | 2 | pdfium | 1275×1651 | 10.6 |  |  |  |  |  |
| arxiv/1512.03385v1.pdf | 2 | pdfkit | 1275×1651 | 44.1 | 147897 | 5.52 | 6.74 % | 0.9572 |  |
| arxiv/1512.03385v1.pdf | 2 | pdfkit-stilus | 1275×1651 | 20.8 | 77271 | 5.46 | 6.71 % | 0.9586 |  |
| arxiv/1512.03385v1.pdf | 3 | pdfium | 1275×1651 | 10.4 |  |  |  |  |  |
| arxiv/1512.03385v1.pdf | 3 | pdfkit | 1275×1651 | 38.1 | 141977 | 5.58 | 6.85 % | 0.9548 |  |
| arxiv/1512.03385v1.pdf | 3 | pdfkit-stilus | 1275×1651 | 16.4 | 73595 | 5.52 | 6.81 % | 0.9563 |  |
| arxiv/1706.03762v7.pdf | 1 | pdfium | 1275×1651 | 6.9 |  |  |  |  |  |
| arxiv/1706.03762v7.pdf | 1 | pdfkit | 1275×1651 | 23.2 | 83312 | 3.13 | 3.75 % | 0.9716 |  |
| arxiv/1706.03762v7.pdf | 1 | pdfkit-stilus | 1275×1651 | 10.6 | 45412 | 3.10 | 3.71 % | 0.9722 |  |
| arxiv/1706.03762v7.pdf | 2 | pdfium | 1275×1651 | 8.8 |  |  |  |  |  |
| arxiv/1706.03762v7.pdf | 2 | pdfkit | 1275×1651 | 31.1 | 117861 | 4.20 | 5.16 % | 0.9704 |  |
| arxiv/1706.03762v7.pdf | 2 | pdfkit-stilus | 1275×1651 | 14.1 | 62676 | 4.16 | 5.11 % | 0.9712 |  |
| arxiv/1706.03762v7.pdf | 3 | pdfium | 1275×1651 | 26.3 |  |  |  |  |  |
| arxiv/1706.03762v7.pdf | 3 | pdfkit | 1275×1651 | 60.4 | 52103 | 2.64 | 3.10 % | 0.9773 |  |
| arxiv/1706.03762v7.pdf | 3 | pdfkit-stilus | 1275×1651 | 52.5 | 29019 | 2.62 | 3.09 % | 0.9780 |  |
| arxiv/2010.11929v2.pdf | 1 | pdfium | 1275×1651 | 7.8 |  |  |  |  |  |
| arxiv/2010.11929v2.pdf | 1 | pdfkit | 1275×1651 | 26.1 | 94477 | 3.76 | 4.60 % | 0.9693 |  |
| arxiv/2010.11929v2.pdf | 1 | pdfkit-stilus | 1275×1651 | 11.8 | 49029 | 3.70 | 4.54 % | 0.9703 |  |
| arxiv/2010.11929v2.pdf | 2 | pdfium | 1275×1651 | 9.2 |  |  |  |  |  |
| arxiv/2010.11929v2.pdf | 2 | pdfkit | 1275×1651 | 32.2 | 117811 | 4.73 | 5.79 % | 0.9625 |  |
| arxiv/2010.11929v2.pdf | 2 | pdfkit-stilus | 1275×1651 | 13.6 | 58627 | 4.66 | 5.75 % | 0.9640 |  |
| arxiv/2010.11929v2.pdf | 3 | pdfium | 1275×1651 | 17.3 |  |  |  |  |  |
| arxiv/2010.11929v2.pdf | 3 | pdfkit | 1275×1651 | 38.2 | 192582 | 2.78 | 3.23 % | 0.9796 |  |
| arxiv/2010.11929v2.pdf | 3 | pdfkit-stilus | 1275×1651 | 24.7 | 155050 | 2.69 | 3.21 % | 0.9803 |  |
| pdfjs/TrueType_without_cmap.pdf | 1 | pdfium | 417×105 | 0.1 |  |  |  |  |  |
| pdfjs/TrueType_without_cmap.pdf | 1 | pdfkit | 417×105 | 0.1 | 190 | 1.83 | 1.10 % | 0.9592 |  |
| pdfjs/TrueType_without_cmap.pdf | 1 | pdfkit-stilus | 417×105 | 0.1 | 191 | 1.83 | 1.10 % | 0.9592 |  |
| pdfjs/annotation-caret-ink.pdf | 1 | pdfium | 1241×1754 | 3.7 |  |  |  |  |  |
| pdfjs/annotation-caret-ink.pdf | 1 | pdfkit | 1241×1754 | 4.4 | 29690 | 0.03 | 0.04 % | 0.9997 |  |
| pdfjs/annotation-caret-ink.pdf | 1 | pdfkit-stilus | 1241×1754 | 4.2 | 29323 | 0.03 | 0.04 % | 0.9998 |  |
| pdfjs/bitmap-composite-and-xnor-text.pdf | 1 | pdfium | 832×834 | 4.4 |  |  |  |  |  |
| pdfjs/bitmap-composite-and-xnor-text.pdf | 1 | pdfkit | 832×834 | 5.1 | 108 | 0.59 | 0.58 % | 0.9908 |  |
| pdfjs/bitmap-composite-and-xnor-text.pdf | 1 | pdfkit-stilus | 832×834 | 5.2 | 109 | 0.59 | 0.58 % | 0.9908 |  |
| pdfjs/blendmode.pdf | 1 | pdfium | 1241×1754 | 38.7 |  |  |  |  |  |
| pdfjs/blendmode.pdf | 1 | pdfkit | 1241×1754 | 19.4 | 8555 | 11.98 | 12.41 % | 0.8670 |  |
| pdfjs/blendmode.pdf | 1 | pdfkit-stilus | 1241×1754 | 18.7 | 6797 | 11.98 | 12.41 % | 0.8670 |  |
| pdfjs/ccitt_EndOfBlock_false.pdf | 1 | pdfium | 1241×1754 | 5.7 |  |  |  |  |  |
| pdfjs/ccitt_EndOfBlock_false.pdf | 1 | pdfkit | 1241×1754 | 16.9 | 223 | 2.05 | 0.95 % | 0.9599 |  |
| pdfjs/ccitt_EndOfBlock_false.pdf | 1 | pdfkit-stilus | 1241×1754 | 6.4 | 219 | 2.06 | 0.95 % | 0.9597 |  |
| pdfjs/cid_cff.pdf | 1 | pdfium | 1575×1170 | 2.9 |  |  |  |  |  |
| pdfjs/cid_cff.pdf | 1 | pdfkit | 1575×1171 | 1.7 | 1624 | 0.01 | 0.00 % | 1.0000 | size (1575,1171) vs ref (1575,1170) |
| pdfjs/cid_cff.pdf | 1 | pdfkit-stilus | 1575×1171 | 1.5 | 1212 | 0.00 | 0.00 % | 1.0000 | size (1575,1171) vs ref (1575,1170) |
| pdfjs/clippath.pdf | 1 | pdfium | 417×209 | 0.1 |  |  |  |  |  |
| pdfjs/clippath.pdf | 1 | pdfkit | 417×209 | 0.7 | 70 | 1.10 | 0.86 % | 0.9905 |  |
| pdfjs/clippath.pdf | 1 | pdfkit-stilus | 417×209 | 0.3 | 58 | 1.10 | 0.86 % | 0.9909 |  |
| pdfjs/cmykjpeg.pdf | 1 | pdfium | 1275×1651 | 3.8 |  |  |  |  |  |
| pdfjs/cmykjpeg.pdf | 1 | pdfkit | 1275×1651 | 5.2 | 89 | 0.37 | 0.05 % | 0.9948 |  |
| pdfjs/cmykjpeg.pdf | 1 | pdfkit-stilus | 1275×1651 | 5.2 | 90 | 0.37 | 0.05 % | 0.9948 |  |
| pdfjs/devicen.pdf | 1 | pdfium | 1275×1651 | 3.2 |  |  |  |  |  |
| pdfjs/devicen.pdf | 1 | pdfkit | 1275×1651 | 2.5 | 406 | 0.01 | 0.00 % | 1.0000 |  |
| pdfjs/devicen.pdf | 1 | pdfkit-stilus | 1275×1651 | 1.7 | 391 | 0.00 | 0.00 % | 1.0000 |  |
| pdfjs/function_based_shading.pdf | 1 | pdfium | 1275×1651 | 110.9 |  |  |  |  |  |
| pdfjs/function_based_shading.pdf | 1 | pdfkit | 1275×1651 | 157.8 | 4520022 | 0.79 | 0.34 % | 0.9891 |  |
| pdfjs/function_based_shading.pdf | 1 | pdfkit-stilus | 1275×1651 | 156.8 | 4520022 | 0.79 | 0.34 % | 0.9891 |  |
| pdfjs/function_based_shading_cmyk.pdf | 1 | pdfium | 605×605 | 40.4 |  |  |  |  |  |
| pdfjs/function_based_shading_cmyk.pdf | 1 | pdfkit | 605×605 | 40.0 | 398821 | 9.44 | 11.21 % | 0.9683 |  |
| pdfjs/function_based_shading_cmyk.pdf | 1 | pdfkit-stilus | 605×605 | 39.1 | 398822 | 9.44 | 11.21 % | 0.9683 |  |
| pdfjs/function_based_shading_cmyk.pdf | 2 | pdfium | 3917×2625 | 2335.4 |  |  |  |  |  |
| pdfjs/function_based_shading_cmyk.pdf | 2 | pdfkit | 3917×2625 | 2601.9 | 18750220 | 31.28 | 41.08 % | 0.9437 |  |
| pdfjs/function_based_shading_cmyk.pdf | 2 | pdfkit-stilus | 3917×2625 | 2602.6 | 18750220 | 31.28 | 41.08 % | 0.9437 |  |
| pdfjs/issue7847_radial.pdf | 1 | pdfium | 1001×501 | 3.0 |  |  |  |  |  |
| pdfjs/issue7847_radial.pdf | 1 | pdfkit | 1001×501 | 17.4 | 1151257 | 1.15 | 0.00 % | 0.9963 |  |
| pdfjs/issue7847_radial.pdf | 1 | pdfkit-stilus | 1001×501 | 17.1 | 1151254 | 1.15 | 0.00 % | 0.9963 |  |
| pdfjs/jbig2_symbol_offset.pdf | 1 | pdfium | 1241×1754 | 3.1 |  |  |  |  |  |
| pdfjs/jbig2_symbol_offset.pdf | 1 | pdfkit | 1241×1754 | 2.2 | 104 | 0.35 | 0.14 % | 0.9943 |  |
| pdfjs/jbig2_symbol_offset.pdf | 1 | pdfkit-stilus | 1241×1754 | 2.4 | 105 | 0.35 | 0.14 % | 0.9943 |  |
| pdfjs/jp2k-resetprob.pdf | 1 | pdfium | 63×43 | 0.3 |  |  |  |  |  |
| pdfjs/jp2k-resetprob.pdf | 1 | pdfkit | 63×43 | 0.6 | 2913 | 10.08 | 3.47 % | 0.9283 |  |
| pdfjs/jp2k-resetprob.pdf | 1 | pdfkit-stilus | 63×43 | 0.6 | 2914 | 10.08 | 3.47 % | 0.9283 |  |
| pdfjs/jpx_smaskindata.pdf | 1 | pdfium | 292×84 | 0.2 |  |  |  |  |  |
| pdfjs/jpx_smaskindata.pdf | 1 | pdfkit | 292×84 | 0.4 | 5133 | 6.91 | 9.49 % | 0.8608 |  |
| pdfjs/jpx_smaskindata.pdf | 1 | pdfkit-stilus | 292×84 | 0.3 | 5129 | 6.80 | 9.49 % | 0.8615 |  |
| pdfjs/knockout_blend_multiply.pdf | 1 | pdfium | 417×334 | 1.0 |  |  |  |  |  |
| pdfjs/knockout_blend_multiply.pdf | 1 | pdfkit | 417×334 | 1.5 | 98 | 11.08 | 0.51 % | 0.9846 |  |
| pdfjs/knockout_blend_multiply.pdf | 1 | pdfkit-stilus | 417×334 | 0.6 | 89 | 11.06 | 0.51 % | 0.9850 |  |
| pdfjs/knockout_nested.pdf | 1 | pdfium | 417×334 | 0.4 |  |  |  |  |  |
| pdfjs/knockout_nested.pdf | 1 | pdfkit | 417×334 | 1.2 | 120 | 0.35 | 0.42 % | 0.9903 |  |
| pdfjs/knockout_nested.pdf | 1 | pdfkit-stilus | 417×334 | 0.9 | 96 | 0.29 | 0.12 % | 0.9924 |  |
| pdfjs/mmtype1.pdf | 1 | pdfium | 417×105 | 0.1 |  |  |  |  |  |
| pdfjs/mmtype1.pdf | 1 | pdfkit | 417×105 | 0.2 | 1168 | 2.08 | 2.23 % | 0.9851 |  |
| pdfjs/mmtype1.pdf | 1 | pdfkit-stilus | 417×105 | 0.2 | 1030 | 2.09 | 2.22 % | 0.9851 |  |
| pdfjs/nonisolated_blend_smask.pdf | 1 | pdfium | 667×501 | 2.2 |  |  |  |  |  |
| pdfjs/nonisolated_blend_smask.pdf | 1 | pdfkit | 667×501 | 4.9 | 2937 | 2.44 | 2.49 % | 0.9662 |  |
| pdfjs/nonisolated_blend_smask.pdf | 1 | pdfkit-stilus | 667×501 | 3.3 | 1736 | 2.49 | 2.47 % | 0.9627 |  |
| pdfjs/pattern_text_embedded_font.pdf | 1 | pdfium | 1241×1754 | 5.5 |  |  |  |  |  |
| pdfjs/pattern_text_embedded_font.pdf | 1 | pdfkit | 1241×1754 | 12.1 | 242918 | 3.09 | 2.48 % | 0.9583 |  |
| pdfjs/pattern_text_embedded_font.pdf | 1 | pdfkit-stilus | 1241×1754 | 7.7 | 235392 | 3.05 | 2.50 % | 0.9583 |  |
| pdfjs/radial_gradients.pdf | 1 | pdfium | 1240×1755 | 13.0 |  |  |  |  |  |
| pdfjs/radial_gradients.pdf | 1 | pdfkit | 1240×1755 | 70.1 | 2327065 | 1.35 | 0.54 % | 0.9911 |  |
| pdfjs/radial_gradients.pdf | 1 | pdfkit-stilus | 1240×1755 | 63.2 | 2324940 | 1.27 | 0.54 % | 0.9914 |  |
| pdfjs/radial_gradients.pdf | 2 | pdfium | 1240×1755 | 13.4 |  |  |  |  |  |
| pdfjs/radial_gradients.pdf | 2 | pdfkit | 1240×1755 | 86.8 | 3691951 | 3.46 | 1.52 % | 0.9799 |  |
| pdfjs/radial_gradients.pdf | 2 | pdfkit-stilus | 1240×1755 | 77.6 | 3690078 | 3.38 | 1.52 % | 0.9802 |  |
| pdfjs/radial_gradients.pdf | 3 | pdfium | 1240×1755 | 13.8 |  |  |  |  |  |
| pdfjs/radial_gradients.pdf | 3 | pdfkit | 1240×1755 | 78.2 | 2683469 | 1.58 | 0.61 % | 0.9871 |  |
| pdfjs/radial_gradients.pdf | 3 | pdfkit-stilus | 1240×1755 | 66.4 | 2681185 | 1.50 | 0.61 % | 0.9875 |  |
| pdfjs/simpletype3font.pdf | 1 | pdfium | 1275×1651 | 2.7 |  |  |  |  |  |
| pdfjs/simpletype3font.pdf | 1 | pdfkit | 1275×1651 | 1.5 | 261 | 0.00 | 0.00 % | 1.0000 |  |
| pdfjs/simpletype3font.pdf | 1 | pdfkit-stilus | 1275×1651 | 1.5 | 222 | 0.00 | 0.00 % | 1.0000 |  |
| pdfjs/smask_alpha_bc.pdf | 1 | pdfium | 459×334 | 0.6 |  |  |  |  |  |
| pdfjs/smask_alpha_bc.pdf | 1 | pdfkit | 459×334 | 1.7 | 175 | 0.10 | 0.00 % | 0.9999 |  |
| pdfjs/smask_alpha_bc.pdf | 1 | pdfkit-stilus | 459×334 | 0.8 | 171 | 0.05 | 0.00 % | 1.0000 |  |
| pdfjs/smaskdim.pdf | 1 | pdfium | 1275×1651 | 2.9 |  |  |  |  |  |
| pdfjs/smaskdim.pdf | 1 | pdfkit | 1275×1651 | 1.9 | 2906 | 0.04 | 0.04 % | 0.9996 |  |
| pdfjs/smaskdim.pdf | 1 | pdfkit-stilus | 1275×1651 | 1.9 | 2671 | 0.04 | 0.04 % | 0.9996 |  |
| pdfjs/text_clip_cff_cid.pdf | 1 | pdfium | 1209×417 | 1.4 |  |  |  |  |  |
| pdfjs/text_clip_cff_cid.pdf | 1 | pdfkit | 1209×417 | 1.5 | 544 | 58.20 | 26.23 % | 0.7217 |  |
| pdfjs/text_clip_cff_cid.pdf | 1 | pdfkit-stilus | 1209×417 | 0.5 | 577 | 58.20 | 26.23 % | 0.7216 |  |
| pdfjs/tiling-pattern-box.pdf | 1 | pdfium | 1241×1754 | 4.9 |  |  |  |  |  |
| pdfjs/tiling-pattern-box.pdf | 1 | pdfkit | 1241×1754 | 314.3 | 121624 | 0.03 | 0.00 % | 0.9996 |  |
| pdfjs/tiling-pattern-box.pdf | 1 | pdfkit-stilus | 1241×1754 | 209.6 | 78583 | 0.01 | 0.00 % | 0.9999 |  |
| pdfjs/tiling_patterns_variations.pdf | 1 | pdfium | 1250×1667 | 5.4 |  |  |  |  |  |
| pdfjs/tiling_patterns_variations.pdf | 1 | pdfkit | 1250×1667 | 26.2 | 11710 | 3.52 | 2.40 % | 0.9580 |  |
| pdfjs/tiling_patterns_variations.pdf | 1 | pdfkit-stilus | 1250×1667 | 5.7 | 9516 | 3.45 | 2.42 % | 0.9584 |  |
| pdfjs/tracemonkey.pdf | 1 | pdfium | 1275×1651 | 9.1 |  |  |  |  |  |
| pdfjs/tracemonkey.pdf | 1 | pdfkit | 1275×1651 | 36.9 | 136449 | 4.64 | 5.73 % | 0.9628 |  |
| pdfjs/tracemonkey.pdf | 1 | pdfkit-stilus | 1275×1651 | 16.2 | 70306 | 4.59 | 5.65 % | 0.9637 |  |
| pdfjs/tracemonkey.pdf | 2 | pdfium | 1275×1651 | 14.1 |  |  |  |  |  |
| pdfjs/tracemonkey.pdf | 2 | pdfkit | 1275×1651 | 47.8 | 170536 | 5.35 | 6.54 % | 0.9609 |  |
| pdfjs/tracemonkey.pdf | 2 | pdfkit-stilus | 1275×1651 | 22.7 | 88524 | 5.27 | 6.44 % | 0.9617 |  |
| pdfjs/tracemonkey.pdf | 3 | pdfium | 1275×1651 | 8.3 |  |  |  |  |  |
| pdfjs/tracemonkey.pdf | 3 | pdfkit | 1275×1651 | 32.3 | 119708 | 4.21 | 5.16 % | 0.9682 |  |
| pdfjs/tracemonkey.pdf | 3 | pdfkit-stilus | 1275×1651 | 14.3 | 61563 | 4.13 | 5.08 % | 0.9692 |  |
| pdfjs/transparency_group.pdf | 1 | pdfium | 1667×1250 | 14.8 |  |  |  |  |  |
| pdfjs/transparency_group.pdf | 1 | pdfkit | 1667×1250 | 142.9 | 1643012 | 4.93 | 5.01 % | 0.9684 |  |
| pdfjs/transparency_group.pdf | 1 | pdfkit-stilus | 1667×1250 | 139.7 | 1643791 | 4.92 | 5.01 % | 0.9684 |  |
| synthetic/contours-3000.pdf | 1 | pdfium | 2481×1754 | 164.5 |  |  |  |  |  |
| synthetic/contours-3000.pdf | 1 | pdfkit | 2481×1754 | 563.6 | 292896 | 1.46 | 0.40 % | 0.9984 |  |
| synthetic/contours-3000.pdf | 1 | pdfkit-stilus | 2481×1754 | 51.7 | 272658 | 0.92 | 0.04 % | 0.9992 |  |
| synthetic/glyphs-30000-uncached.pdf | 1 | pdfium | 2481×1754 | 322.4 |  |  |  |  |  |
| synthetic/glyphs-30000-uncached.pdf | 1 | pdfkit | 2481×1754 | 318.2 | 6944324 | 2.72 | 0.39 % | 0.9965 |  |
| synthetic/glyphs-30000-uncached.pdf | 1 | pdfkit-stilus | 2481×1754 | 255.8 | 6414482 | 1.47 | 0.01 % | 0.9988 |  |
| synthetic/hatch-2000-0.35mm.pdf | 1 | pdfium | 2481×1754 | 327.5 |  |  |  |  |  |
| synthetic/hatch-2000-0.35mm.pdf | 1 | pdfkit | 2481×1754 | 12750.6 | 67785 | 0.41 | 0.04 % | 0.9873 |  |
| synthetic/hatch-2000-0.35mm.pdf | 1 | pdfkit-stilus | 2481×1754 | 33.1 | 60129 | 0.49 | 0.04 % | 0.9837 |  |
| synthetic/hatch-2000-hairline.pdf | 1 | pdfium | 2481×1754 | 307.1 |  |  |  |  |  |
| synthetic/hatch-2000-hairline.pdf | 1 | pdfkit | 2481×1754 | 12729.9 | 61779 | 1.23 | 0.04 % | 0.9991 |  |
| synthetic/hatch-2000-hairline.pdf | 1 | pdfkit-stilus | 2481×1754 | 28.4 | 54124 | 1.20 | 0.04 % | 0.9993 |  |
| synthetic/hatch-2000-maskclip.pdf | 1 | pdfium | 2481×1754 | 307.6 |  |  |  |  |  |
| synthetic/hatch-2000-maskclip.pdf | 1 | pdfkit | 2481×1754 | 12789.0 | 61162 | 0.67 | 0.00 % | 0.9997 |  |
| synthetic/hatch-2000-maskclip.pdf | 1 | pdfkit-stilus | 2481×1754 | 79.9 | 54237 | 0.63 | 0.00 % | 0.9999 |  |
| synthetic/mixed-drawing.pdf | 1 | pdfium | 2481×1754 | 81.2 |  |  |  |  |  |
| synthetic/mixed-drawing.pdf | 1 | pdfkit | 2481×1754 | 211.7 | 101276 | 0.69 | 0.08 % | 0.9977 |  |
| synthetic/mixed-drawing.pdf | 1 | pdfkit-stilus | 2481×1754 | 53.3 | 87676 | 0.50 | 0.01 % | 0.9987 |  |
| synthetic/short-20000-0.35mm.pdf | 1 | pdfium | 2481×1754 | 125.9 |  |  |  |  |  |
| synthetic/short-20000-0.35mm.pdf | 1 | pdfkit | 2481×1754 | 110.2 | 720235 | 0.55 | 0.01 % | 0.9994 |  |
| synthetic/short-20000-0.35mm.pdf | 1 | pdfkit-stilus | 2481×1754 | 52.8 | 660245 | 0.28 | 0.01 % | 0.9997 |  |
| synthetic/short-20000-hairline.pdf | 1 | pdfium | 2481×1754 | 117.8 |  |  |  |  |  |
| synthetic/short-20000-hairline.pdf | 1 | pdfkit | 2481×1754 | 102.3 | 660245 | 0.59 | 0.01 % | 0.9992 |  |
| synthetic/short-20000-hairline.pdf | 1 | pdfkit-stilus | 2481×1754 | 49.5 | 600255 | 0.28 | 0.01 % | 0.9997 |  |
