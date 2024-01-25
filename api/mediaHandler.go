package api

import (
	"bytes"
	"crypto/rand"
	"image"
	"image/color"
	"image/png"
	"math/big"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MediaHandler struct{}

func (h *MediaHandler) RegisterRoutes(server *gin.Engine) {
	userGroup := server.Group("/media")

	userGroup.GET("/image", h.GetImage)
}

func (h *MediaHandler) GetImage(c *gin.Context) {
	img := image.NewRGBA(image.Rect(0, 0, random(), random()))
	for x := 0; x < img.Bounds().Max.X; x++ {
		for y := 0; y < 1000; y++ {
			img.Set(x, y, color.RGBA{uint8(random()), uint8(random()), uint8(random()), uint8(random())})
		}
	}

	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, img); err != nil {
		c.String(http.StatusInternalServerError, "Error encoding image")
		return
	}

	c.Header("Content-Type", "image/png")

	c.Data(http.StatusOK, "image/png", buffer.Bytes())
}

func random() int {
	const maxH int64 = int64(256)
	n, err := rand.Int(rand.Reader, big.NewInt(maxH))
	if err != nil {
		return 0
	}
	return int(n.Int64())
}
