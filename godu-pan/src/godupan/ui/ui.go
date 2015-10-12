package ui

import (
	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/samples/flags"
    "github.com/google/gxui/math"
    "log"
    "io/ioutil"
)

func Start() {
	gl.StartDriver(appMain)
}

// Create a PanelHolder with a 3 panels
func panelHolder(name string, theme gxui.Theme) gxui.PanelHolder {
	label := func(text string) gxui.Label {
		label := theme.CreateLabel()
		label.SetText(text)
		return label
	}

	holder := theme.CreatePanelHolder()
	holder.AddPanel(label(name+" 0 content"), name+" 0 panel")
	holder.AddPanel(label(name+" 1 content"), name+" 1 panel")
	holder.AddPanel(label(name+" 2 content"), name+" 2 panel")
	return holder
}

func appMain(driver gxui.Driver) {
	theme := flags.CreateTheme(driver)

	// ┌───────┐║┌───────┐
	// │       │║│       │
	// │   A   │║│   B   │
	// │       │║│       │
	// └───────┘║└───────┘
	// ═══════════════════
	// ┌───────┐║┌───────┐
	// │       │║│       │
	// │   C   │║│   D   │
	// │       │║│       │
	// └───────┘║└───────┘
	
	ftData, err := ioutil.ReadFile("static/font/simhei.ttf")
	if err != nil {
		log.Println(err)
	}
	ft, err := driver.CreateFont(ftData, 20)
	if err != nil {
		log.Println(err)
	}
	menu := theme.CreateLinearLayout()
	menu.SetDirection(gxui.LeftToRight)
	logo_label:=theme.CreateLabel()
	logo_label.SetColor(gxui.White)
	logo_label.SetFont(ft)	
	logo_label.SetText("百度盘")
	logo_label.SetSize(math.Size{300,200})
	menu.AddChild(logo_label)
	
	
	
	splitterCD := theme.CreateSplitterLayout()
	splitterCD.SetOrientation(gxui.Horizontal)
	splitterCD.AddChild(panelHolder("C", theme))
	splitterCD.AddChild(panelHolder("D", theme))

	vSplitter := theme.CreateSplitterLayout()
	vSplitter.SetOrientation(gxui.Vertical)
	vSplitter.AddChild(menu)
	vSplitter.AddChild(splitterCD)

	window := theme.CreateWindow(800, 600, "百度盘")
	window.SetScale(flags.DefaultScaleFactor)
	window.AddChild(vSplitter)
	window.OnClose(driver.Terminate)
}