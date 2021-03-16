package process

import (
	"fmt"

	"github.com/lakiluki1/lfmbg/config"

	"gopkg.in/gographics/imagick.v3/imagick"
)

func Process(path string, config *config.Config) error {
	imagick.Initialize()
	defer imagick.Terminate()
	var width uint = config.Width
	var height uint = config.Height
	var coverSmallHeight uint = uint(float32(width) / 5)
	var offsetX int = int(float32(0.0520833) * float32(width))

	cover := imagick.NewMagickWand()
	coverSmall := cover.Clone()

	err := cover.ReadImage(path)
	if err != nil {
		return err
	}

	err = coverSmall.ReadImage(path)
	if err != nil {
		return err
	}

	err = cover.BlurImage(10, 5)
	if err != nil {
		return err
	}

	err = cover.ResizeImage(width, width, 0)
	if err != nil {
		return err
	}

	err = cover.CropImage(width, height, 0, int(width-height)/2)
	if err != nil {
		return err
	}

	err = coverSmall.ResizeImage(coverSmallHeight, coverSmallHeight, imagick.FILTER_CUBIC)
	if err != nil {
		return err
	}

	err = addShadow(&coverSmall, 1.025, 60, 2.75, 10, 10)
	if err != nil {
		return err
	}

	err = dimImage(&cover, 0.3)
	if err != nil {
		return err
	}

	err = cover.CompositeImage(coverSmall, imagick.COMPOSITE_OP_OVER, true, offsetX, int(height/2-(coverSmallHeight/2)))
	if err != nil {
		return err
	}

	err = cover.WriteImage(config.SavePath)

	return err

}

func dimImage(mw **imagick.MagickWand, dimPercent float32) error {
	overlay := imagick.NewMagickWand()
	height := (*mw).GetImageHeight()
	width := (*mw).GetImageWidth()
	pw := imagick.NewPixelWand()
	pw.SetColor(fmt.Sprintf("rgba(0, 0, 0, %f)", dimPercent))
	err := overlay.NewImage(uint(float32(width)), uint(float32(height)), pw)
	if err != nil {
		return err
	}

	err = overlay.SetImageBackgroundColor(pw)
	if err != nil {
		return err
	}

	err = overlay.CompositeImage(*mw, imagick.COMPOSITE_OP_DARKEN, true, 0, 0)
	*mw = overlay
	return err

}

func addShadow(mw **imagick.MagickWand, shadowSize float32, opacity, sigma float64, x, y int) error {
	shadow := imagick.NewMagickWand()
	pw := imagick.NewPixelWand()
	pw.SetColor("black")

	height := (*mw).GetImageHeight()
	width := (*mw).GetImageWidth()
	shadowHeight := uint(float32((*mw).GetImageHeight()) * shadowSize)
	shadowWidth := uint(float32((*mw).GetImageWidth()) * shadowSize)
	err := shadow.NewImage(uint(float32(width)*shadowSize), uint(float32(height)*shadowSize), pw)
	if err != nil {
		return err
	}

	err = shadow.SetImageBackgroundColor(pw)
	if err != nil {
		return err
	}

	err = shadow.ShadowImage(opacity, sigma, x, y)
	if err != nil {
		return err
	}

	err = shadow.CompositeImage(*mw, imagick.COMPOSITE_OP_OVER, true, int(shadowWidth-width), int(shadowHeight-height))
	*mw = shadow
	return err

}
