package mediainfo_test

import (
	"fmt"
	"github.com/jie123108/go_mediainfo"
	"io/ioutil"
	"os"
	"testing"
)

const (
	ogg       = "testdata/test.ogg"
	mp3       = "testdata/test.mp3"
	nonExists = "testdata/non_exists.ogg"
)

func TestOpenWithOgg(t *testing.T) {
	mi := mediainfo.NewMediaInfo()
	error := mi.OpenFile(ogg)
	if error != nil {
		t.Fail()
	}
}

func TestOpenWithMp3(t *testing.T) {
	mi := mediainfo.NewMediaInfo()
	error := mi.OpenFile(mp3)
	if error != nil {
		t.Fail()
	}
}

func TestOpenWithUnexistsFile(t *testing.T) {
	mi := mediainfo.NewMediaInfo()
	error := mi.OpenFile(nonExists)
	if error == nil {
		t.Fail()
	}
}

func TestOpenMemoryWithOgg(t *testing.T) {
	mi := mediainfo.NewMediaInfo()
	f, _ := os.Open(ogg)
	bytes, _ := ioutil.ReadAll(f)

	error := mi.OpenMemory(bytes)
	if error != nil {
		t.Fail()
	}
}

func TestOpenMemoryWithMp3(t *testing.T) {
	mi := mediainfo.NewMediaInfo()
	f, _ := os.Open(mp3)
	bytes, _ := ioutil.ReadAll(f)

	error := mi.OpenMemory(bytes)
	if error != nil {
		t.Fail()
	}
}

func TestOpenMemoryWithEmptyArray(t *testing.T) {
	mi := mediainfo.NewMediaInfo()
	error := mi.OpenMemory([]byte{})
	if error == nil {
		t.Fail()
	}
}

func TestInformWithOgg(t *testing.T) {
	mi := mediainfo.NewMediaInfo()
	mi.OpenFile(ogg)

	if len(mi.Inform()) < 10 {
		t.Fail()
	}
}

func TestInformWithMp3(t *testing.T) {
	mi := mediainfo.NewMediaInfo()
	mi.OpenFile(mp3)

	if len(mi.Inform()) < 10 {
		t.Fail()
	}
}

func TestAvailableParametersWithOgg(t *testing.T) {
	mi := mediainfo.NewMediaInfo()
	mi.OpenFile(ogg)

	if len(mi.AvailableParameters()) < 10 {
		t.Fail()
	}
}

func TestAvailableParametersWithMp3(t *testing.T) {
	mi := mediainfo.NewMediaInfo()
	mi.OpenFile(mp3)

	if len(mi.AvailableParameters()) < 10 {
		t.Fail()
	}
}

func TestDurationWithOgg(t *testing.T) {
	info, _ := mediainfo.GetMediaInfo(ogg)
	if info.General.Duration != 3494 {
		t.Fail()
	}
}

func TestDurationWithMp3(t *testing.T) {
	info, _ := mediainfo.GetMediaInfo(mp3)
	if info.General.Duration != 87771 {
		t.Fail()
	}
}

// func TestCodecWithOgg(t *testing.T) {
// 	info, _ := mediainfo.GetMediaInfo(ogg)
// 	if info.General.Codec != "OGG" {
// 		fmt.Printf("info.General.Codec: %s", info.General.Codec)
// 		t.Fail()
// 	}
// }

// func TestCodecWithMp3(t *testing.T) {
// 	info, _ := mediainfo.GetMediaInfo(mp3)
// 	if info.General.Codec != "MPEG Audio" {
// 		t.Fail()
// 	}
// }

// func TestFormatWithOgg(t *testing.T) {
// 	info, _ := mediainfo.GetMediaInfo(ogg)
// 	if info.General.Format != "OGG" {
// 		t.Fail()
// 	}
// }

// func TestFormatWithMp3(t *testing.T) {
// 	info, _ := mediainfo.GetMediaInfo(mp3)
// 	if info.General.Format != "MPEG Audio" {
// 		t.Fail()
// 	}
// }

//----------------------------------------------------------------------------------------------------------------------
func BenchmarkOpenAndDurationWithOgg(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mi := mediainfo.NewMediaInfo()
		mi.OpenFile(ogg)

		mi.GetInt(mediainfo.MediaInfo_Stream_General, "Duration")
	}
}

func BenchmarkOpenAndDurationWithMp3(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mi := mediainfo.NewMediaInfo()
		mi.OpenFile(mp3)

		mi.GetInt(mediainfo.MediaInfo_Stream_General, "Duration")
	}
}

func BenchmarkOpenMemoryAndDurationWithOgg(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mi := mediainfo.NewMediaInfo()
		f, _ := os.Open(ogg)
		bytes, _ := ioutil.ReadAll(f)

		mi.OpenMemory(bytes)
		mi.GetInt(mediainfo.MediaInfo_Stream_General, "Duration")
	}
}

func BenchmarkOpenMemoryAndDurationWithMp3(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mi := mediainfo.NewMediaInfo()
		f, _ := os.Open(mp3)
		bytes, _ := ioutil.ReadAll(f)

		mi.OpenMemory(bytes)
		mi.GetInt(mediainfo.MediaInfo_Stream_General, "Duration")
	}
}

//----------------------------------------------------------------------------------------------------------------------

func ExampleUsage() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	mi := mediainfo.NewMediaInfo()
	err = mi.OpenMemory(bytes)
	if err != nil {
		panic(err)
	}
	fmt.Println(mi.AvailableParameters())                                 // Print all supported params for Get
	fmt.Println(mi.GetInt(mediainfo.MediaInfo_Stream_General, "BitRate")) // Print bitrate
}
