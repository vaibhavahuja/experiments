package main

import (
	"fmt"
	"github.com/unki2aut/go-mpd"
	"log"
	"os"
)

// will do quick POC here
func main() {
	//can use this to parse your file, but will have to write code to convert
	// segmentTemplate to segments on own
	mpd := new(mpd.MPD)
	myFile, err := os.ReadFile("test.xml")
	if err != nil {
		log.Fatalf("error while opening file %s", err.Error())
	}

	err = mpd.Decode(myFile)
	if err != nil {
		log.Fatalf("error while decoding file %s", err.Error())
	}
	for _, period := range mpd.Period {
		baseUrl := period.BaseURL[0].Value
		fmt.Println(baseUrl)
		//period.BaseURL
		for _, adaptSet := range period.AdaptationSets {
			//adaptSet.SegmentTemplate.Initialization
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
