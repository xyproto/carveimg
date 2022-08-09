# Preview

Approximate image viewer for the terminal that uses content-aware image resizing.

## Go packages in use

* The image resizing is done by the [`caire`](https://github.com/esimov/caire) package.
* The palette reduction is done by the [`palgen`](https://github.com/xyproto/palgen) package.

## Example preview

| Original PNG image                    | In a VT100 compatible terminal emulator, using 16 colors |
|---------------------------------------|----------------------------------------------------------|
| <img src=img/grumpycat.png width=512> |            <img src=img/grumpycat16colors.png width=512> |

## Installation

The `preview` utility can be installed with Go 1.17 or later:

    go install github.com/xyproto/preview/cmd/preview@latest

## Performance

* Note that the image reszing is very slow for larger images (wallpaper size and up).

## General info

* Version: 1.1.0
* License: BSD-3
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
