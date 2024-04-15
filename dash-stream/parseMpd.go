package main

import (
	"fmt"
	"github.com/unki2aut/go-mpd"
)

// will do quick POC here
func main() {
	//can use this to parse your file, but will have to write code to convert
	// segmentTemplate to segments on own
	mpd := new(mpd.MPD)
	//mpd.Decode()
	for _, period := range mpd.Period {
		//period.BaseURL
		for _, adaptSet := range period.AdaptationSets {
			//fmt.Println("printing for ", *adaptSet.Codecs)
			//adaptSet.SegmentTemplate.SegmentTimeline
			for _, representation := range adaptSet.Representations {
				id := representation.ID
				if id != nil {
					fmt.Println(*id)
				}

			}
			fmt.Println()
		}
	}

}
