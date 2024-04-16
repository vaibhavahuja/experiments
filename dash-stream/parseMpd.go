package main

import (
	"fmt"
	"github.com/unki2aut/go-mpd"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
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

	//trying to download all segments - just a POC
	host := "https://ak-mediavod.jiocinema.com/bpkvod/jcvod/default/64a437de48d27972cc952d8e_v10/64a437de48d27972cc952d8e_v10/"

	for _, period := range mpd.Period {
		baseUrl := period.BaseURL[0].Value
		completeUrl := host + baseUrl
		for _, adaptSet := range period.AdaptationSets {
			if *adaptSet.ContentType != "video" {
				continue
			}
			segmentTemplate := adaptSet.SegmentTemplate
			for _, representation := range adaptSet.Representations {
				id := representation.ID
				if id != nil {
					fmt.Println(*id)
				}
				if *id != "video=200000" {
					continue
				}
				//only downloading init right now
				var initialisationSegment string
				if segmentTemplate != nil && segmentTemplate.Initialization != nil {
					initialisationSegment = *segmentTemplate.Initialization
				}
				modifiedString := strings.Replace(initialisationSegment, "$RepresentationID$", *id, 1)
				download(completeUrl+modifiedString, -1, nil)

				//downloading all segments now
				mediaString := *segmentTemplate.Media
				//modifiedMediaString will further modify this by replacing the other placeholder
				modifiedMediaString := strings.Replace(mediaString, "$RepresentationID$", *id, 1)
				var listOfSegments []string
				for _, timeline := range segmentTemplate.SegmentTimeline.S {
					//append R times to listOfSegments
					//p := 1
					for val := *timeline.T; val <= uint64(*timeline.R)*timeline.D; val += timeline.D {
						//append all segments in listOfSegments
						modifiedMediaString2 := strings.Replace(modifiedMediaString, "$Time$", strconv.Itoa(int(val)), 1)
						//p += 1
						listOfSegments = append(listOfSegments, modifiedMediaString2)
					}
				}
				//my prepared list of segments is ready
				fmt.Println(len(listOfSegments))

				//I will download only first 50?
				wg := &sync.WaitGroup{}
				for i := 0; i < len(listOfSegments); i++ {
					wg.Add(1)
					go download(completeUrl+listOfSegments[i], i, wg)
				}
				wg.Wait()
			}
			fmt.Println()
		}
	}

}
