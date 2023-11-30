package services

import (
    "bytes"
    "encoding/base64"
    "image"
    "image/color"
    "image/draw"
    "image/gif"
    "image/jpeg"
    "image/png"
    "io/ioutil"
    "log"
    "net/http"

    "golang.org/x/image/bmp"
    "rsc.io/rsc/qr/web/resize"

    "github.com/haierspi/gateway_full/utils/rpc"
)

var bgImage = getNetImage("https://img6.lovelywholesale.com/images/duopingtai/202004/202004_D_1587967670_65923.jpg")

// Image Image
func (Examples) Image(param rpc.BodyArgs, reply *rpc.BodyReply) error {
    data, err := base64.StdEncoding.DecodeString(param.Body)
    if err != nil {
        log.Println(err)
        return nil
    }

    imgBounds := bgImage.Bounds()
    img := image.NewRGBA(imgBounds)
    draw.Draw(img, imgBounds, bgImage, image.Pt(0, 0), draw.Src)

    headImage := getNetImage(string(data))
    headImage = resize.Resample(headImage, headImage.Bounds(), 128, 128)
    sp := image.Pt(88, 235)
    rect := headImage.Bounds().Add(sp)
    draw.DrawMask(img, rect, headImage, image.ZP, &circle{image.Pt(64, 64), 64}, image.ZP, draw.Over)

    b := bytes.NewBuffer(nil)
    jpeg.Encode(b, img, &jpeg.Options{Quality: 90})

    reply.ContentType = "image/jpeg"
    reply.Body = b.Bytes()
    return nil
}

func getNetImage(url string) image.Image {
    resp, err := http.Get(url)
    if err != nil {
        log.Println(err)
        return image.NewNRGBA(image.Rect(0, 0, 600, 400))
    }
    defer resp.Body.Close()
    data, _ := ioutil.ReadAll(resp.Body)
    contentType := http.DetectContentType(data)
    var imgBG image.Image
    if contentType == "image/png" {
        imgBG, err = png.Decode(bytes.NewReader(data))
    } else if contentType == "image/jpeg" {
        imgBG, err = jpeg.Decode(bytes.NewReader(data))
    } else if contentType == "image/gif" {
        imgBG, err = gif.Decode(bytes.NewReader(data))
    } else if contentType == "image/bmp" {
        imgBG, err = bmp.Decode(bytes.NewReader(data))
    }
    if err != nil {
        log.Println(err)
        return image.NewNRGBA(image.Rect(0, 0, 600, 400))
    }
    return imgBG
}

type circle struct {
    p image.Point
    r int
}

func (c *circle) ColorModel() color.Model {
    return color.AlphaModel
}

func (c *circle) Bounds() image.Rectangle {
    return image.Rect(c.p.X-c.r, c.p.Y-c.r, c.p.X+c.r, c.p.Y+c.r)
}

func (c *circle) At(x, y int) color.Color {
    xx, yy, rr := float64(x-c.p.X)+0.5, float64(y-c.p.Y)+0.5, float64(c.r)
    if xx*xx+yy*yy < rr*rr {
        return color.Alpha{255}
    }
    return color.Alpha{0}
}
