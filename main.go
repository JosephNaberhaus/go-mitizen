package main

import (
	"github.com/JosephNaberhaus/go-mitizen/prompt"
	"github.com/eiannone/keyboard"
	"log"
	"os"
)

func main()  {
	f, err := os.OpenFile("testlogfile", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)

	test := prompt.YesNo{Description: "testing"}

	test.Show()

	keyboard.Open()
	defer keyboard.Close()

	for true {
		rune, key, _ := keyboard.GetKey()

		if key == keyboard.KeyCtrlC {
			keyboard.Close()
			test.Finish()
			os.Exit(1)
		}

		test.HandleInput(prompt.ToKey(rune, key))
	}
}