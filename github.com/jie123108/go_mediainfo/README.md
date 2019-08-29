# mediainfo
Golang binding for [libmediainfo](https://mediaarea.net/en/MediaInfo)

Duration, Bitrate, Codec, Streams and a lot of other meta-information about media files can be extracted through it.

For now supports only media files with one stream. Bindings for MediaInfoList is not provided. It can be easy fixed if anybody ask me.

Works through MediaInfoDLL/MediaInfoDLL.h(dynamic load and so on), so your mediainfo installation should has it.

Supports direct reading files by name and reading data from []byte buffers(without copying it for C calls).

Documentation for libmediainfo is poor and ascetic, can be found [here](https://mediaarea.net/en/MediaInfo/Support/SDK).

Your advices and suggestions are welcome!

## Simple Example
```go
package main

import (
    "fmt"
    mediainfo "github.com/jie123108/go_mediainfo"
    "os"
)

func main() {
    info, err := mediainfo.GetMediaInfo(os.Args[1])
    if err != nil {
        fmt.Printf("open failed: %v\n", err)
        os.Exit(1)
    }
    fmt.Printf("%v\n", info)
}
```

## Example 
```go
package main

import (
    "fmt"
    mediainfo "github.com/jie123108/go_mediainfo"
    "os"
)

func main() {
    mi := mediainfo.NewMediaInfo()
    err := mi.OpenFile(os.Args[1])
    if err != nil {
        fmt.Printf("open failed: %v\n", err)
        os.Exit(1)
    }
    defer mi.Close()
    video := &mediainfo.SimpleMediaInfo{}
    g := &video.General
    g.DurationStr = mi.Get(mediainfo.MediaInfo_Stream_General, "Duration/String3")
    g.Duration = mi.GetInt(mediainfo.MediaInfo_Stream_General, "Duration")
    g.Start = mi.GetInt(mediainfo.MediaInfo_Stream_General, "Start")
    g.BitRate = mi.GetInt(mediainfo.MediaInfo_Stream_General, "OverallBitRate")
    g.FrameRate = mi.GetInt(mediainfo.MediaInfo_Stream_General, "FrameRate")
    g.FileSize = mi.GetInt(mediainfo.MediaInfo_Stream_General, "FileSize")

    g.GeneralCount = int(mi.GetInt(mediainfo.MediaInfo_Stream_General, "GeneralCount"))
    g.VideoCount = int(mi.GetInt(mediainfo.MediaInfo_Stream_General, "VideoCount"))
    g.AudioCount = int(mi.GetInt(mediainfo.MediaInfo_Stream_General, "AudioCount"))
    g.TextCount = int(mi.GetInt(mediainfo.MediaInfo_Stream_General, "TextCount"))
    g.OtherCount = int(mi.GetInt(mediainfo.MediaInfo_Stream_General, "OtherCount"))
    g.ImageCount = int(mi.GetInt(mediainfo.MediaInfo_Stream_General, "ImageCount"))
    g.MenuCount = int(mi.GetInt(mediainfo.MediaInfo_Stream_General, "MenuCount"))

    v := &video.Video
    v.CodecID = mi.Get(mediainfo.MediaInfo_Stream_Video, "CodecID")
    v.BitRate = mi.GetInt(mediainfo.MediaInfo_Stream_Video, "BitRate")
    v.Width = mi.GetInt(mediainfo.MediaInfo_Stream_Video, "Width")
    v.Height = mi.GetInt(mediainfo.MediaInfo_Stream_Video, "Height")
    v.Resolution = fmt.Sprintf("%dx%d", v.Width, v.Height)
    v.DAR = mi.Get(mediainfo.MediaInfo_Stream_Video, "DisplayAspectRatio/String")

    a := &video.Audio
    a.CodecID = mi.Get(mediainfo.MediaInfo_Stream_Audio, "CodecID")
    a.SamplingRate = mi.GetInt(mediainfo.MediaInfo_Stream_Audio, "SamplingRate")
    a.BitRate = mi.GetInt(mediainfo.MediaInfo_Stream_Audio, "BitRate")

    video.SubtitlesCnt = int(mi.GetInt(mediainfo.MediaInfo_Stream_Text, "StreamCount"))
    for i := 0; i < video.SubtitlesCnt; i++ {
        Format := mi.GetIdx(mediainfo.MediaInfo_Stream_Text, i, "Format")
        CodecID := mi.GetIdx(mediainfo.MediaInfo_Stream_Text, i, "CodecID")
        Title := mi.GetIdx(mediainfo.MediaInfo_Stream_Text, i, "Title")
        subtitles := mediainfo.SubtitlesInfo{Format: Format, CodecID: CodecID, Title: Title}
        video.Subtitles = append(video.Subtitles, subtitles)
    }
    fmt.Printf("%v\n", video)
}

```

Read the [documentation](https://godoc.org/github.com/zhulik/go_mediainfo) for other functions

## Contributing
You know=)