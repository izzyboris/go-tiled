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
	"testing"

	"github.com/lafriks/go-tiled"
	"github.com/stretchr/testify/assert"
)

func TestIsometricImageSize(t *testing.T) {
	type test struct {
		name string
		tmap tiled.Map
		size image.Rectangle
	}

	tests := []test{
		{
			name: "Isometric size 0,0",
			tmap: tiled.Map{
				TileHeight:  64,
				TileWidth:   128,
				Width:       0,
				Height:      0,
				Orientation: "isometric",
				RenderOrder: "right-down",
			},
			size: image.Rect(0, 0, 0, 0),
		},
		{
			name: "Isometric size 1,1",
			tmap: tiled.Map{
				TileHeight:  64,
				TileWidth:   128,
				Width:       1,
				Height:      1,
				Orientation: "isometric",
				RenderOrder: "right-down",
			},
			size: image.Rect(0, 0, 128, 64),
		},
		{
			name: "Isometric size 4,8",
			tmap: tiled.Map{
				TileHeight:  64,
				TileWidth:   128,
				Width:       4,
				Height:      8,
				Orientation: "isometric",
				RenderOrder: "right-down",
			},
			size: image.Rect(0, 0, 768, 384),
		},
		{
			name: "Isometric size 8,6",
			tmap: tiled.Map{
				TileHeight:  64,
				TileWidth:   128,
				Width:       8,
				Height:      6,
				Orientation: "isometric",
				RenderOrder: "right-down",
			},
			size: image.Rect(0, 0, 896, 448),
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			r, err := NewRenderer(&test.tmap)
			if err != nil {
				t.Error(err)
			}
			rect := r.engine.GetFinalImageSize()
			assert.Equal(t, test.size, rect)
		})
	}
}

func TestIsometricTilePos(t *testing.T) {
	type test struct {
		name string
		tmap tiled.Map
		x    int
		y    int
		pos  image.Rectangle
	}

	tile := &tiled.LayerTile{
		Tileset: &tiled.Tileset{
			TileWidth:  128,
			TileHeight: 64,
		},
	}

	tests := []test{
		{
			name: "Isometric pos top of 1,1",
			tmap: tiled.Map{
				TileHeight:  64,
				TileWidth:   128,
				Width:       1,
				Height:      1,
				Orientation: "isometric",
				RenderOrder: "right-down",
			},
			x:   0,
			y:   0,
			pos: image.Rect(0, 0, 128, 64),
		},
		{
			name: "Isometric pos top of 2,2",
			tmap: tiled.Map{
				TileHeight:  64,
				TileWidth:   128,
				Width:       2,
				Height:      2,
				Orientation: "isometric",
				RenderOrder: "right-down",
			},
			x:   0,
			y:   0,
			pos: image.Rect(64, 0, 192, 64),
		},
		{
			name: "Isometric pos bottom of 2,2",
			tmap: tiled.Map{
				TileHeight:  64,
				TileWidth:   128,
				Width:       2,
				Height:      2,
				Orientation: "isometric",
				RenderOrder: "right-down",
			},
			x:   1,
			y:   1,
			pos: image.Rect(64, 64, 192, 128),
		},
		{
			name: "Isometric pos left of 2,2",
			tmap: tiled.Map{
				TileHeight:  64,
				TileWidth:   128,
				Width:       2,
				Height:      2,
				Orientation: "isometric",
				RenderOrder: "right-down",
			},
			x:   0,
			y:   1,
			pos: image.Rect(0, 32, 128, 96),
		},
		{
			name: "Isometric pos right of 2,2",
			tmap: tiled.Map{
				TileHeight:  64,
				TileWidth:   128,
				Width:       2,
				Height:      2,
				Orientation: "isometric",
				RenderOrder: "right-down",
			},
			x:   1,
			y:   0,
			pos: image.Rect(128, 32, 256, 96),
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			r, err := NewRenderer(&test.tmap)
			if err != nil {
				t.Error(err)
			}
			rect := r.engine.GetTilePosition(test.x, test.y, tile)
			assert.Equal(t, test.pos, rect)
		})
	}
}
