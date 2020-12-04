package module

import "github.com/fogleman/gg"

type Drawable interface {
	Draw(dc *gg.Context) (err error)
}
