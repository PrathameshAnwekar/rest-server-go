package api

import (
	"bytes"
	"crypto/rand"
	"image"
	"image/color"
	"image/png"
	"io"
	"math/big"
	"net/http"
	"os"
	"time"

	cs "github.com/PrathameshAnwekar/rest-server-go/constants"
	"github.com/gin-gonic/gin"
	"gocv.io/x/gocv"
)

type MediaHandler struct{}

func (h *MediaHandler) RegisterRoutes(server *gin.Engine) {
	userGroup := server.Group("/media")

	userGroup.GET("/image", h.GetImage)
	userGroup.GET("/video-stream", h.GetVideoStream)
	userGroup.GET("/camera-stream", h.GetCameraStream)
}

func (h *MediaHandler) GetImage(c *gin.Context) {
	img := image.NewRGBA(image.Rect(0, 0, getRandomInt(), getRandomInt()))
	for x := 0; x < img.Bounds().Max.X; x++ {
		for y := 0; y < 1000; y++ {
			img.Set(x, y, color.RGBA{uint8(getRandomInt()), uint8(getRandomInt()), uint8(getRandomInt()), uint8(getRandomInt())})
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

func (h *MediaHandler) GetVideoStream(c *gin.Context) {
	filePath := "/Users/anwprath/go/pkg/mod/github.com/gabriel-vasile/mimetype@v1.4.2/testdata/mp4.mp4"
	file, err := os.Open(filePath)
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	defer file.Close()

	c.Header("Content-Type", "video/mp4")
	const bufferLimit int = 1024
	buffer := make([]byte, bufferLimit)

	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		} else if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}

		c.Data(http.StatusOK, "video/mp4", buffer[:n])

		c.Writer.Flush()
	}
}

// GetCameraStream uses the opencv library, however simple libraries like
// `github.com/blackjack/webcam` should work too.
func (h *MediaHandler) GetCameraStream(c *gin.Context) {
	c.Header("Content-Type", "multipart/x-mixed-replace; boundary=frame")

	camera, err := gocv.OpenVideoCapture(0)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	defer camera.Close()

	frame := gocv.NewMat()
	defer frame.Close()

	for {
		if ok := camera.Read(&frame); !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err)
			return
		}

		gocv.Rectangle(&frame,
			image.Rect(cs.RectBegin, cs.RectBegin, cs.RectEnd, cs.RectEnd),
			color.RGBA{255, 0, 0, 0}, cs.FrameThickness)

		imgBytes, err := gocv.IMEncode(".jpg", frame)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err)
			return
		}
		data := "--frame\r\n  Content-Type: image/jpeg\r\n\r\n" + string(imgBytes.GetBytes()) + "\r\n\r\n"
		c.Data(http.StatusOK, "image/jpeg", []byte(data))

		c.Writer.Flush()

		time.Sleep(time.Millisecond)

		// stops the stream if client disconnects
		select {
		case <-c.Request.Context().Done():
			return
		default:
			continue
		}
	}
}

func getRandomInt() int {
	const maxH int64 = int64(256)
	n, err := rand.Int(rand.Reader, big.NewInt(maxH))
	if err != nil {
		return 0
	}
	return int(n.Int64())
}
