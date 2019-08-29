package mediainfo

// #cgo CFLAGS: -DUNICODE
// #cgo windows LDFLAGS:
// #cgo linux LDFLAGS: -lz -lpthread -lstdc++ -ldl
// #include "go_mediainfo.h"
import "C"

import (
	"fmt"
	"runtime"
	"strconv"
	"unsafe"
	"strings"
	"github.com/jie123108/glog"
)

// MediaInfo - represents MediaInfo class, all interaction with libmediainfo through it
type MediaInfo struct {
	handle unsafe.Pointer
}

func init() {
	C.setlocale(C.LC_CTYPE, C.CString(""))
	C.MediaInfoDLL_Load()

	if C.MediaInfoDLL_IsLoaded() == 0 {
		fmt.Println("Cannot load mediainfo")
	}
}

// NewMediaInfo - constructs new MediaInfo
func NewMediaInfo() *MediaInfo {
	result := &MediaInfo{handle: C.GoMediaInfo_New()}
	runtime.SetFinalizer(result, func(f *MediaInfo) {
		f.Close()
		C.GoMediaInfo_Delete(f.handle)
	})
	return result
}

// OpenFile - opens file
func (mi *MediaInfo) OpenFile(path string) error {
	p := C.CString(path)
	s := C.GoMediaInfo_OpenFile(mi.handle, p)
	if s == 0 {
		return fmt.Errorf("MediaInfo can't open file: %s", path)
	}
	C.free(unsafe.Pointer(p))
	return nil
}

// OpenMemory - opens memory buffer
func (mi *MediaInfo) OpenMemory(bytes []byte) error {
	if len(bytes) == 0 {
		return fmt.Errorf("Buffer is empty")
	}
	s := C.GoMediaInfo_OpenMemory(mi.handle, (*C.char)(unsafe.Pointer(&bytes[0])), C.size_t(len(bytes)))
	if s == 0 {
		return fmt.Errorf("MediaInfo can't open memory buffer")
	}
	return nil
}

// Close - closes file
func (mi *MediaInfo) Close() {
	C.GoMediaInfo_Close(mi.handle)
}

const (
	MediaInfo_Stream_General = iota
	MediaInfo_Stream_Video
	MediaInfo_Stream_Audio
	MediaInfo_Stream_Text
	MediaInfo_Stream_Other
	MediaInfo_Stream_Image
	MediaInfo_Stream_Menu
)

// GetIdx - allow to read info from file
func (mi *MediaInfo) GetIdx(stream_type int, stream_idx int, param string) (result string) {
	p := C.CString(param)
	st := C.int(stream_type)
	si := C.int(stream_idx)
	r := C.GoMediaInfoGetIdx(mi.handle, si, st, p)
	result = C.GoString(r)
	C.free(unsafe.Pointer(p))
	C.free(unsafe.Pointer(r))
	return
}

// Get - allow to read info from file
func (mi *MediaInfo) Get(stream_type int, param string) (result string) {
	result = mi.GetIdx(stream_type, 0, param)
	return
}

// GetIntIdx - allow to read info as int from file
func (mi *MediaInfo) GetIntIdx(stream_type int, stream_idx int, param string) (ivalue int64) {
  result := mi.GetIdx(stream_type, stream_idx, param)
  if result != "" {
    var err error
    tmp := strings.Split(result, ".")
    ivalue, err = strconv.ParseInt(tmp[0], 10, 64)
    if err != nil {
      glog.Errorf("GetIntIdx ParseInt error:%s", err.Error())
    }
  }
  return
}

func (mi *MediaInfo) GetFloatIdx(stream_type int, stream_idx int, param string) (ivalue float64) {
	result := mi.GetIdx(stream_type, stream_idx, param)
	if result != "" {
		ivalue, _ = strconv.ParseFloat(result, 64)
	}
	return
}

// GetInt - allow to read info as int from file
func (mi *MediaInfo) GetInt(stream_type int, param string) (ivalue int64) {
	ivalue = mi.GetIntIdx(stream_type, 0, param)
	return
}

// Inform returns string with summary file information, like mediainfo util
func (mi *MediaInfo) Inform() (result string) {
	r := C.GoMediaInfoInform(mi.handle)
	result = C.GoString(r)
	C.free(unsafe.Pointer(r))
	return
}

// Option configure or get information about MediaInfoLib
func (mi *MediaInfo) Option(option string, value string) (result string) {
	o := C.CString(option)
	v := C.CString(value)
	r := C.GoMediaInfoOption(mi.handle, o, v)
	C.free(unsafe.Pointer(o))
	C.free(unsafe.Pointer(v))
	result = C.GoString(r)
	C.free(unsafe.Pointer(r))
	return
}

// AvailableParameters returns string with all available Get params and it's descriptions
func (mi *MediaInfo) AvailableParameters() string {
	return mi.Option("Info_Parameters", "")
}

type GeneralInfo struct {
	Codec       string //Codec
	Format      string //Format
	DurationStr string //Duration/String3 : Play time in format : HH:MM:SS.MMM
	Duration    int64  //Duration Play time of the stream in ms
	Start       int64  //Start
	BitRate     int64  //OverallBitRate  Bit rate of all streams in bps
	FrameRate   int64  //FrameRate Frames per second
	FileSize    int64  //File size in bytes

	GeneralCount int //Number of general streams
	VideoCount   int //Number of video streams
	AudioCount   int //Number of audio streams
	TextCount    int //Number of text streams
	OtherCount   int //Number of other streams
	ImageCount   int //Number of image streams
	MenuCount    int //Number of menu streams
}

type VideoInfo struct {
	StreamID 	  int64//stream id of this file
	CodecID       string // CodecID
	BitRate       int64  // Bit rate in bps
	FrameRate	  float64//帧率
	Resolution    string //WxH
	Width, Height int64  //Width,Height
	DAR           string //DisplayAspectRatio/String
	Default		  string   //默认视频，只有一个
}

type AudioInfo struct {
	StreamID 				int64//stream id of this file
	CodecID      			string //CodecID
	channel, SamplingRate 	int64  //channel SamplingRate
	BitRate      			int64  //BitRate: Bit rate in bps
	Title, Language			string //标题，语言
	Default					string   //默认音频，只有一个
}

type SubtitlesInfo struct {
	StreamID 	int64//stream id of this file
	Format  	string //Format UTF-8
	CodecID 	string //CodecID S_TEXT/UTF8
	Title   	string //字幕标题，如：英文字幕
	Language	string //语言
	Default		string   //默认字幕，只有一个
}

type SimpleMediaInfo struct {
	General      GeneralInfo
	//VideoCount	 int
	Video        []VideoInfo
	//AudioCount	 int
	Audio        []AudioInfo
	SubtitlesCnt int
	Subtitles    []SubtitlesInfo
}

// GetMediaInfo returns SimpleMediaInfo of given file
func GetMediaInfo(filename string) (info *SimpleMediaInfo, err error) {
	mi := NewMediaInfo()
	err = mi.OpenFile(filename)
	if err != nil {
		return
	}
	defer mi.Close()
	info = &SimpleMediaInfo{}
	g := &info.General
	g.Codec = mi.Get(MediaInfo_Stream_General, "CodecID")
	g.Format = mi.Get(MediaInfo_Stream_General, "Format")
	g.DurationStr = mi.Get(MediaInfo_Stream_General, "Duration/String3")
	g.Duration = mi.GetInt(MediaInfo_Stream_General, "Duration")
	g.Start = mi.GetInt(MediaInfo_Stream_General, "Start")
	g.BitRate = mi.GetInt(MediaInfo_Stream_General, "OverallBitRate")
	g.FrameRate = mi.GetInt(MediaInfo_Stream_General, "FrameRate")
	g.FileSize = mi.GetInt(MediaInfo_Stream_General, "FileSize")

	g.GeneralCount = int(mi.GetInt(MediaInfo_Stream_General, "GeneralCount"))
	g.VideoCount = int(mi.GetInt(MediaInfo_Stream_General, "VideoCount"))
	g.AudioCount = int(mi.GetInt(MediaInfo_Stream_General, "AudioCount"))
	g.TextCount = int(mi.GetInt(MediaInfo_Stream_General, "TextCount"))
	g.OtherCount = int(mi.GetInt(MediaInfo_Stream_General, "OtherCount"))
	g.ImageCount = int(mi.GetInt(MediaInfo_Stream_General, "ImageCount"))
	g.MenuCount = int(mi.GetInt(MediaInfo_Stream_General, "MenuCount"))

	var StreamIDOffSet int64
	if "RealMedia" == g.Format {
		StreamIDOffSet = 0
	} else {
		StreamIDOffSet = 1
	}
	glog.Infof("GetMediaInfo g.VideoCount:%d, g.AudioCount:%d, g.TextCount, StreamIDOffSet:%d\n", g.VideoCount, g.AudioCount, g.TextCount, StreamIDOffSet)
	//v := &info.Video
	for i := 0; i < g.VideoCount; i++ {
	vinfo := VideoInfo{}
	vinfo.StreamID  = mi.GetIntIdx(MediaInfo_Stream_Video, i, "ID") - StreamIDOffSet
	vinfo.CodecID 	= mi.GetIdx(MediaInfo_Stream_Video, i, "CodecID")
	vinfo.BitRate 	= mi.GetIntIdx(MediaInfo_Stream_Video, i, "BitRate")
	vinfo.FrameRate = mi.GetFloatIdx(MediaInfo_Stream_Video, i, "FrameRate")
	vinfo.Width 	= mi.GetIntIdx(MediaInfo_Stream_Video, i, "Width")
	vinfo.Height 	= mi.GetIntIdx(MediaInfo_Stream_Video, i, "Height")
	vinfo.Resolution = fmt.Sprintf("%dx%d", vinfo.Width, vinfo.Height)
	vinfo.Default 	= mi.GetIdx(MediaInfo_Stream_Video, i, "Default")
	vinfo.DAR 		= mi.GetIdx(MediaInfo_Stream_Video, i, "DisplayAspectRatio/String")
	glog.Infof("GetMediaInfo vinfo.StreamID:%d vinfo.CodecID:%s, BitRate:%d, FrameRate:%f, Width:%d, Height:%d, Resolution:%s, Default:%s\n", vinfo.StreamID, vinfo.CodecID, vinfo.BitRate, vinfo.FrameRate, vinfo.Width, vinfo.Height, vinfo.Resolution, vinfo.Default)
	info.Video 		= append(info.Video, vinfo)
	}

	//a := &info.Audio
	for i := 0; i < g.AudioCount; i++ {
	ainfo := AudioInfo{}
	ainfo.StreamID  = mi.GetIntIdx(MediaInfo_Stream_Audio, i, "ID") - StreamIDOffSet
	ainfo.CodecID = mi.GetIdx(MediaInfo_Stream_Audio, i, "CodecID")
	ainfo.channel = mi.GetIntIdx(MediaInfo_Stream_Audio, i, "Channel")
	ainfo.SamplingRate = mi.GetIntIdx(MediaInfo_Stream_Audio, i, "SamplingRate")
	ainfo.BitRate = mi.GetIntIdx(MediaInfo_Stream_Audio, i, "BitRate")
	ainfo.Title = mi.GetIdx(MediaInfo_Stream_Audio, i, "Title")
	ainfo.Language = mi.GetIdx(MediaInfo_Stream_Audio, i, "Language")
	ainfo.Default = mi.GetIdx(MediaInfo_Stream_Audio, i, "Default")
	info.Audio = append(info.Audio, ainfo)

	glog.Infof("GetMediaInfo ainfo.StreamID:%d, ainfo.CodecID:%s, ainfo.channel:%d, ainfo.SamplingRate:%d, ainfo.BitRate:%d, ainfo.Title:%s, ainfo.Language:%s, ainfo.Default:%s\n", ainfo.StreamID, ainfo.CodecID, ainfo.channel, ainfo.SamplingRate, ainfo.BitRate, ainfo.Title, ainfo.Language, ainfo.Default)
	}

	info.SubtitlesCnt = int(mi.GetInt(MediaInfo_Stream_Text, "StreamCount"))
	glog.Infof("GetMediaInfo info.SubtitlesCnt:%d, g.TextCount:%d\n", info.SubtitlesCnt, g.TextCount)
	for i := 0; i < g.TextCount; i++ {
		subtitles := SubtitlesInfo{}
		subtitles.StreamID  = mi.GetIntIdx(MediaInfo_Stream_Text, i, "ID") - StreamIDOffSet
		subtitles.Format = mi.GetIdx(MediaInfo_Stream_Text, i, "Format")
		subtitles.CodecID = mi.GetIdx(MediaInfo_Stream_Text, i, "CodecID")
		subtitles.Title = mi.GetIdx(MediaInfo_Stream_Text, i, "Title")
		subtitles.Language = mi.GetIdx(MediaInfo_Stream_Text, i, "Language")
		subtitles.Default = mi.GetIdx(MediaInfo_Stream_Text, i, "Default")
		//subtitles := SubtitlesInfo{Format: Format, CodecID: CodecID, Title: Title}
		glog.Infof("GetMediaInfo subtitles.StreamID:%d, subtitles.Format:%s, subtitles.CodecID:%s, subtitles.Title:%s, subtitles.Language:%s, subtitles.Default:%s\n", subtitles.StreamID, subtitles.Format, subtitles.CodecID, subtitles.Title, subtitles.Language, subtitles.Default)
		info.Subtitles = append(info.Subtitles, subtitles)
		//subtitles := SubtitlesInfo{Format: Format, CodecID: CodecID, Title: Title}
		//info.Subtitles = append(info.Subtitles, subtitles)
	}

	return
}
