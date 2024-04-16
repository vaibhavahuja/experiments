package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

//func main() {
//	//test one -> successful
//	//downloadTestAkamaiMp4()
//
//	//trying to download dash now - will heavily modify the other files ehh
//	//test two -> successful
//	downloadDashFiles()
//
//	//able to play the downloaded files on my local dash player
//}

func downloadDashFiles() {
	//initUrl := "https://ak-mediavod.jiocinema.com/bpkvod/jcvod/default/653ac68da6d5148cc0dda999_v4/653ac68da6d5148cc0dda999_v4/653ac68da6d5148cc0dda999_v4-video=200000.dash?hdntl=exp=1713127903~acl=%2fbpkvod%2fjcvod%2fdefault%2f653ac68da6d5148cc0dda999_v4%2f653ac68da6d5148cc0dda999_v4%2f*~id=0352f4ee36d84d0bbd5866754be8ff1d~data=hdntl~hmac=c6616ea9c9459638e8515a96bc83c2907b6ead7cb2af2b6fc295aa50d86eeb02"
	url := "https://ak-mediavod.jiocinema.com/bpkvod/jcvod/default/653ac68da6d5148cc0dda999_v4/653ac68da6d5148cc0dda999_v4/653ac68da6d5148cc0dda999_v4-video=200000-$Time$.dash?hdntl=exp=1713127903~acl=%2fbpkvod%2fjcvod%2fdefault%2f653ac68da6d5148cc0dda999_v4%2f653ac68da6d5148cc0dda999_v4%2f*~id=0352f4ee36d84d0bbd5866754be8ff1d~data=hdntl~hmac=c6616ea9c9459638e8515a96bc83c2907b6ead7cb2af2b6fc295aa50d86eeb02"

	//download init url separately
	wg := &sync.WaitGroup{}
	//wg.Add(1)
	//download(initUrl, 0, nil)
	for i := 0; i < 1200*50; i += 1200 {
		wg.Add(1)
		urlFinal := strings.Replace(url, "$Time$", strconv.Itoa(i), 1)
		go download(urlFinal, i, wg)
	}
	wg.Wait()
}

func downloadTestAkamaiMp4() {
	url := "https://dash.akamaized.net/dash264/TestCasesUHD/2b/11/video_8000k_{number}.mp4"
	//initStreamUrl := "https://dash.akamaized.net/dash264/TestCasesUHD/2b/11/video_8000k_init.mp4"
	wg := &sync.WaitGroup{}
	//wg.Add(1)
	//go download(initStreamUrl, 0, wg)
	//wg.Wait()
	for i := 400; i <= 420; i++ {
		//urlComplete := fmt.Sprintf(url, i)
		urlComplete := strings.Replace(url, "{number}", strconv.Itoa(i), 1)
		wg.Add(1)
		go download(urlComplete, i, wg)
	}
	wg.Wait()
}

func download(url string, i int, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	// Create a GET request to fetch the video
	log.Printf("downloading %s", url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching video:", err)
		return
	}
	defer resp.Body.Close()

	//file name should be same as the file you download
	//parsedUrl, _ := url2.Parse(url)
	//parsedUrl.RawQuery = ""
	//fileName := parsedUrl.Path
	//fileName := fmt.Sprintf("653ac68da6d5148cc0dda999_v4-video=200000-%d.dash", i)

	//extract fileName need to write code for the same
	var fileName string
	if i == -1 {
		fileName = "testing_init.dash"
	} else {
		fileName = fmt.Sprintf("testing_%d.dash", i)
	}

	out, err := os.Create(fmt.Sprintf("test-server/test-dash-akamai-stream/%s", fileName))
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer out.Close()

	// Write the response body to the output file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println("Error writing to output file:", err)
		return
	}

	fmt.Println("Video downloaded successfully! ", i)
}
