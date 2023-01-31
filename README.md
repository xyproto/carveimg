# Carve

Two image viewing utilities for the teriminal. Both of them uses 16 colors.

* `carve` - uses content-aware image resizing before displaying the image
* `img` - uses regular image resizing before displaying the image

## Screenshots

| Original PNG image                    | In a VT100 compatible terminal emulator, using 16 colors |
|---------------------------------------|----------------------------------------------------------|
| <img src=img/grumpycat.png width=512> |            <img src=img/grumpycat16colors.png width=512> |

## Installation

The `preview` utility can be installed with Go 1.17 or later:

    go install github.com/xyproto/preview/cmd/preview@latest

## The `carve` utility

* The image resizing is done with [`github.com/esimov/caire`](https://github.com/esimov/caire).
* The palette reduction is done with [`github.com/xyproto/palgen`](https://github.com/xyproto/palgen).
* The image reszing may be very slow for larger images.

## The `img` utilitiy

* The image resizing is done with [`golang.org/x/image/draw`](https://golang.org/x/image/draw).
* The palette reduction is done with [`github.com/xyproto/palgen`](https://github.com/xyproto/palgen).

## General info

* Version: 1.2.0
* License: BSD-3
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
