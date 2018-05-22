package main

import (
	"fmt"
	"os"
	"time"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func main() {
	if (len(os.Args)) == 1 {
		//Print help & exit
		fmt.Println("Use: ")
		fmt.Println(os.Args[0], "MP3FILE")
		return
	}


	//Open file
	file, err := os.Open(os.Args[1])
	if err != nil { fmt.Println(err); return }
	//Close file
	defer func() {
		err = file.Close()
		if err != nil {	fmt.Println(err) }
	}()

	//Decode
	streamer, format, err := mp3.Decode(file)
	if err != nil { fmt.Println(err); return }


	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	done := make(chan struct{})
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {close(done)})))

	fmt.Println(os.Args)

	<-done
}
