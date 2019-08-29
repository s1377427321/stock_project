package main

import (
	"fmt"
	mediainfo "github.com/jie123108/go_mediainfo"
	"os"
)

func main() {
	filename := os.Args[1]
	info, err := mediainfo.GetMediaInfo(filename)
	if err != nil {
		fmt.Printf("open failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%v\n", info)

	if len(os.Args) >= 3 && os.Args[2] == "full" {
		mi := mediainfo.NewMediaInfo()
		err = mi.OpenFile(filename)
		if err != nil {
			fmt.Printf("open failed: %v\n", err)
			os.Exit(1)
		}
		defer mi.Close()
		fullinfo := mi.Inform()
		fmt.Println(fullinfo)
	}
}
