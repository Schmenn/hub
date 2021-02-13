package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	//"time"

	"github.com/Schmenn/hub/modules"
	"github.com/Schmenn/hub/modules/debug"
	"github.com/Schmenn/hub/modules/features"
	//"github.com/Schmenn/hub/structs"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func main() {

	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	debug.Check(err)

	size := strings.Split(string(out[:len(string(out))-1]), " ")
	//fmt.Printf("out: %#v\n", size)
	//fmt.Printf("err: %#v\n", err)
	screenHeight, _ := strconv.Atoi(size[0])
	screenWidth, _ := strconv.Atoi(size[1])

	//fmt.Printf("height: %v", int(screenHeight/2))
	//fmt.Printf("width: %v", int(screenWidth/2))

	conf := modules.GetConfig()
	fmt.Println(conf)

	//time.Sleep(time.Second * 20)

	go features.Play(conf.StartUpSound)

	//os.Exit(0)
	//time.Sleep(time.Second * 10)


	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	p := widgets.NewParagraph()
	p.Title = "hub..."
	p.Text = "Welcome Home " + conf.Name
	p.BorderStyle.Fg = ui.ColorCyan
	p.BorderStyle.Bg = ui.ColorWhite
	p.SetRect(0 , 0, int(screenWidth), int(screenHeight/7))
	ui.Render(p)

	

	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			break
		}
	}
}

