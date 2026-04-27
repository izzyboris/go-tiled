/*
Copyright (c) 2017 Lauris Bukšis-Haberkorns <lauris@nix.lv>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package render

import (
	"image"
	"log"

	"github.com/disintegration/imaging"
	tiled "github.com/lafriks/go-tiled"
)

// IsometricRendererEngine represents isometric rendering engine.
type IsometricRendererEngine struct {
	m *tiled.Map
}

// Init initializes rendering engine with provided map options.
func (e *IsometricRendererEngine) Init(m *tiled.Map) {
	e.m = m
}

// GetFinalImageSize returns final image size based on map data.
func (e *IsometricRendererEngine) GetFinalImageSize() image.Rectangle {
	mh := float64(e.m.Height)
	mw := float64(e.m.Width)
	th2 := float64(e.m.TileHeight) / 2
	tw2 := float64(e.m.TileWidth) / 2
	// max cartesian image size from isometric dimensions
	h := int(th2 * (mh + mw))
	w := int(tw2 * (mh + mw))
	log.Println("Getting final image size for", e.m)
	return image.Rect(0, 0, w, h)
}

// RotateTileImage rotates provided tile layer.
func (e *IsometricRendererEngine) RotateTileImage(tile *tiled.LayerTile, img image.Image) image.Image {
	timg := img
	if tile.DiagonalFlip {
		timg = imaging.FlipH(imaging.Rotate270(timg))
	}
	if tile.HorizontalFlip {
		timg = imaging.FlipH(timg)
	}
	if tile.VerticalFlip {
		timg = imaging.FlipV(timg)
	}

	return timg
}

// GetTilePosition returns tile position in image.
func (e *IsometricRendererEngine) GetTilePosition(x, y int, tile *tiled.LayerTile) image.Rectangle {
	fx := float64(x)
	fy := float64(y)
	th2 := float64(e.m.TileHeight) / 2
	tw2 := float64(e.m.TileWidth) / 2
	ix := int(tw2*(fx-fy)) + int(tw2)*(e.m.Height-1)
	iy := int(th2 * (fx + fy))
	return image.Rect(ix, iy,
		ix+tile.Tileset.TileWidth,
		iy+tile.Tileset.TileHeight)
}
