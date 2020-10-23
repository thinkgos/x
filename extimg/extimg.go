// Package extimg extend image
package extimg

import (
	"encoding/base64"
	"errors"
	"net/http"
	"os"
	"strings"
)

var ext = []string{
	"ase", "art", "bmp", "blp", "cd5", "cit", "cpt", "cr2", "cut", "dds", "dib",
	"djvu", "egt", "exif", "gif", "gpl", "grf", "icns", "ico", "iff", "jng", "jpeg",
	"jpg", "jfif", "jp2", "jps", "lbm", "max", "miff", "mng", "msp", "nitf", "ota",
	"pbm", "pc1", "pc2", "pc3", "pcf", "pcx", "pdn", "pgm", "PI1", "PI2", "PI3",
	"pict", "pct", "pnm", "pns", "ppm", "psb", "psd", "pdd", "psp", "px", "pxm",
	"pxr", "qfx", "raw", "rle", "sct", "sgi", "rgb", "int", "bw", "tga", "tiff",
	"tif", "vtf", "xbm", "xcf", "xpm", "3dv", "amf", "ai", "awg", "cgm", "cdr",
	"cmx", "dxf", "e2d", "egt", "eps", "fs", "gbr", "odg", "svg", "stl", "vrml",
	"x3d", "sxd", "v2d", "vnd", "wmf", "emf", "art", "xar", "png", "webp", "jxr",
	"hdp", "wdp", "cur", "ecw", "iff", "lbm", "liff", "nrrd", "pam", "pcx", "pgf",
	"sgi", "rgb", "rgba", "bw", "int", "inta", "sid", "ras", "sun", "tga",
}

// GetExts get image ext slice
func GetExts() []string {
	return ext
}

// GetType returns the type of image (like image/jpeg)
func GetType(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	buf := make([]byte, 512)
	_, err = file.Read(buf)
	if err != nil {
		return "", err
	}

	filetype := http.DetectContentType(buf)
	for _, ext := range ext {
		if strings.Contains(ext, filetype[6:]) { // like image/jpeg
			return filetype, nil
		}
	}

	return "", errors.New("invalid image type")
}

// EncodeToBase64 image encode to base64,
// format like: data:image/png;base64,xxxxxxxxxxxxxx
// ext: png,jpg... or image/png,image/jpg
// value: image raw value
func EncodeToBase64(ext string, value []byte) string {
	n := 5 + 8 + len(ext) + base64.StdEncoding.EncodedLen(len(value))
	has := strings.HasPrefix(ext, "image/")
	if !has {
		n += 6
	}
	builder := strings.Builder{}
	builder.Grow(n)
	builder.WriteString("data:")
	if !has {
		builder.WriteString("image/")
	}
	builder.WriteString(ext)
	builder.WriteString(";base64,")
	builder.WriteString(base64.StdEncoding.EncodeToString(value))
	return builder.String()
}

// DecodeBase64 decode base64 image which format is like: data:image/png;base64,xxxxxxxxxxxxxx
func DecodeBase64(img string) (string, []byte, error) {
	ss := strings.Split(img, ",")
	if len(ss) != 2 {
		return "", nil, errors.New("invalid base64 image")
	}

	tp := strings.TrimSuffix(strings.TrimPrefix(ss[0], "data:"), ";base64")
	bv, err := base64.StdEncoding.DecodeString(ss[1])
	return tp, bv, err
}
