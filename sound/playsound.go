package sound

import (
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)

func PlaySound() {
	f, err := os.Open("./sound/bell.wav")
	// f, err := os.Open("YOUR NEW MORNING ALARM.mp3")

	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := wav.Decode(f)
	//streamer, format, err := mp3.Decode(f)

	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))
	<-done
}
