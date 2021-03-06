package main

import (
	"github.com/giorgisio/goav/avcodec"
	"github.com/giorgisio/goav/avdevice"
	"github.com/giorgisio/goav/avfilter"
	"github.com/giorgisio/goav/avformat"
	"github.com/giorgisio/goav/avutil"
	"github.com/giorgisio/goav/swresample"
	"github.com/giorgisio/goav/swscale"
	"log"
)

func main() {

	// Register all formats and codecs
	avformat.Av_register_all()

	log.Printf("AVFilter Version:\t%v", avfilter.Avfilter_version())
	log.Printf("AVDevice Version:\t%v", avdevice.Avdevice_version())
	log.Printf("SWScale Version:\t%v", swscale.Swscale_version())
	log.Printf("AVUtil Version:\t%v", avutil.AvutilVersion())
	log.Printf("AVCodec Version:\t%v", avcodec.Avcodec_version())
	log.Printf("Resample Version:\t%v", swresample.Swresample_license())

}
