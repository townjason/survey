package photo

import (
	"api/util/log"
	"github.com/disintegration/imaging"
	"github.com/nfnt/resize"
	"image/jpeg"
	"io"
	"os"
)

const (
	imageEdgeUpperBound = 1280
	imageCompressEdgeFix = 960
)

// 輸入圖片路徑，取得圖片的寬高資訊
func getJpegWidthHeight(path string) (width int, height int, err error) {
	reader, err := os.Open(path)
	defer log.Error(reader.Close())
	if err != nil {
		log.Error(err)
		return 0, 0, err
	}

	cfg, err := jpeg.DecodeConfig(reader)
	if err != nil {
		log.Error(err)
		return 0, 0, err
	}

	width = cfg.Width
	height = cfg.Height
	return width, height, nil
}

// 取得要壓縮的寬高
func getCompressWidthHeight(srcWidth int, srcHeight int) (compressWidth int, compressHeight int) {
	if srcWidth >= srcHeight{
		compressWidth = imageCompressEdgeFix
		compressHeight = srcHeight*imageCompressEdgeFix/srcWidth
	}else{
		compressHeight = imageCompressEdgeFix
		compressWidth = srcWidth*imageCompressEdgeFix/srcHeight
	}
	return compressWidth, compressHeight
}

// 取得要進行影像寬高縮圖的分母
func getPhotoDivider(width int, height int)(divider int){
	var maxEdge int
	if width >= height{
		maxEdge = width
	}else{
		maxEdge = height
	}

	divider = 1
	for{
		if maxEdge <= imageEdgeUpperBound{
			return divider
		}
		maxEdge = maxEdge/2
		divider = divider*2
	}
}

func SaveUploadFileCompress(srcPath string, destPath string) error {

	// 取得圖檔寬高
	width, height, err := getJpegWidthHeight(srcPath)
	if err != nil{
		return err
	}

	jpgFile, err := os.Open(srcPath)
	defer jpgFile.Close()
	if err != nil{
		return err
	}

	out, err := os.Create(destPath)
	if err != nil {
		log.Error(err)
	}
	defer out.Close()

	// 取得要進行影像寬高resize的分母
	divider := getPhotoDivider(width, height)

	// 進行圖片resize
	if divider > 1{
		// 讀取圖片
		imageJpg, err := jpeg.Decode(jpgFile)
		if err != nil{
			return err
		}
		// 計算resize的寬高
		resizeWidth := uint(width/divider)
		resizeHeight := uint(height/divider)
		// resize image
		resizeImg := resize.Resize(resizeWidth, resizeHeight, imageJpg, resize.Lanczos3)

		// save image
		err = jpeg.Encode(out, resizeImg, &jpeg.Options{100})
		if err != nil{
			return err
		}
	}else{
		// 圖片不用進行resize，直接複製到目的地
		io.Copy(out, jpgFile)
	}

	return nil
}

func SaveUploadFileCompress2(srcPath string, destPath string) error {
	srcImage, err := imaging.Open(srcPath, imaging.AutoOrientation(true))
	if err != nil{
		return err
	}
	r := srcImage.Bounds()
	width := r.Max.X
	height := r.Max.Y

	if width > imageCompressEdgeFix || height > imageCompressEdgeFix{
		// 取得等比壓縮要的寬高
		resizeWidth, resizeHeight := getCompressWidthHeight(width, height)
		// resize image
		detImage := imaging.Resize(srcImage, resizeWidth, resizeHeight, imaging.Lanczos)
		// save image
		if err = imaging.Save(detImage, destPath, imaging.JPEGQuality(100));err != nil{
			return err
		}
	}else{
		// 圖片不用進行resize，直接複製到目的地
		// save image
		if err = imaging.Save(srcImage, destPath, imaging.JPEGQuality(100));err != nil{
			return err
		}
	}

	return nil
}
