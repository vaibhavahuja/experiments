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
			if *adaptSet.ContentType == "video" {
				downloadAllVideo(completeUrl, adaptSet)
			}
			if *adaptSet.ContentType == "audio" {
				//downloadAllAudio(completeUrl, adaptSet)
			}
		}
	}

}

func downloadAllAudio(url string, adaptSet *mpd.AdaptationSet) {
	segmentTemplate := adaptSet.SegmentTemplate

	for _, representation := range adaptSet.Representations {
		id := representation.ID
		if id != nil {
			fmt.Println(*id)
		}
		if *id != "audio_hin=45377" {
			continue
		}

		//only downloading init right now
		var initialisationSegment string
		if segmentTemplate != nil && segmentTemplate.Initialization != nil {
			initialisationSegment = *segmentTemplate.Initialization
		}
		modifiedString := strings.Replace(initialisationSegment, "$RepresentationID$", *id, 1)
		download(url+modifiedString, "audio", -1, nil)

		//downloading all segments now
		mediaString := *segmentTemplate.Media
		//modifiedMediaString will further modify this by replacing the other placeholder
		modifiedMediaString := strings.Replace(mediaString, "$RepresentationID$", *id, 1)
		fmt.Println(modifiedMediaString)
		var listOfSegments []string
		var lastVal uint64
		for _, timeline := range segmentTemplate.SegmentTimeline.S {
			var repeat int64
			if timeline.R == nil {
				repeat = 0
			} else {
				repeat = *timeline.R
			}

			//var start uint64
			//if timeline.T == nil {
			//	start = lastVal
			//} else {
			//	start = *timeline.T
			//}
			if timeline.T != nil {
				lastVal = *timeline.T
			}
			//lastVal = *timeline.T

			//0, 96256, 192512, 288768, 384000, 480256, 576512, 672768, 768000, 864256, 960512, 1056768
			//r+1 -> times repeat
			//lastVal = start
			for val2 := 0; val2 <= int(repeat); val2++ {
				modifiedMediaString2 := strings.Replace(modifiedMediaString, "$Time$", strconv.Itoa(int(lastVal)), 1)
				lastVal += timeline.D
				listOfSegments = append(listOfSegments, modifiedMediaString2)
			}

		}
		//my prepared list of segments is ready
		fmt.Println(len(listOfSegments))

		//I will download only first 50?
		wg := &sync.WaitGroup{}
		for i := 0; i < len(listOfSegments); i++ {
			wg.Add(1)
			go download(url+listOfSegments[i], "audio", i, wg)
		}
		wg.Wait()
	}
}

func downloadAllVideo(completeUrl string, adaptSet *mpd.AdaptationSet) {
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
		download(completeUrl+modifiedString, "video", -1, nil)

		//downloading all segments now
		mediaString := *segmentTemplate.Media
		//modifiedMediaString will further modify this by replacing the other placeholder
		modifiedMediaString := strings.Replace(mediaString, "$RepresentationID$", *id, 1)
		var listOfSegments []string
		var lastVal uint64
		for _, timeline := range segmentTemplate.SegmentTimeline.S {
			var repeat int64
			if timeline.R == nil {
				repeat = 0
			} else {
				repeat = *timeline.R
			}

			if timeline.T != nil {
				lastVal = *timeline.T
			}

			for val2 := 0; val2 <= int(repeat); val2++ {
				modifiedMediaString2 := strings.Replace(modifiedMediaString, "$Time$", strconv.Itoa(int(lastVal)), 1)
				lastVal += timeline.D
				listOfSegments = append(listOfSegments, modifiedMediaString2)
			}

		}

		//for _, timeline := range segmentTemplate.SegmentTimeline.S {
		//	//append R times to listOfSegments
		//	//p := 1
		//	for val := *timeline.T; val <= uint64(*timeline.R)*timeline.D; val += timeline.D {
		//		//append all segments in listOfSegments
		//		modifiedMediaString2 := strings.Replace(modifiedMediaString, "$Time$", strconv.Itoa(int(val)), 1)
		//		//p += 1
		//		listOfSegments = append(listOfSegments, modifiedMediaString2)
		//	}
		//}
		//my prepared list of segments is ready
		fmt.Println(len(listOfSegments))

		//I will download only first 50?
		wg := &sync.WaitGroup{}
		for i := 0; i < len(listOfSegments); i++ {
			wg.Add(1)
			go download(completeUrl+listOfSegments[i], "video", i, wg)
		}
		wg.Wait()
	}
}
