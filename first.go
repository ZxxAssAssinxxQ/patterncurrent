package main

import (
	"os"

	"github.com/therecipe/qt/gui"

	"github.com/therecipe/qt/widgets"

        "patcurr/tssMath"
)

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)
	debug := false // check compiled debug mode or not
	// setup window
	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowIcon(gui.NewQIcon5(":qml/tss_logo.png"))
	// window.SetFixedSize2(290, 265) // (290, 265)
	window.SetWindowTitle("Patturn current calculator")
	// setup master layout for sub layouts
	masterVertWidget := widgets.NewQWidget(nil, 0)
	masterVertWidget.SetLayout(widgets.NewQVBoxLayout())
	window.SetCentralWidget(masterVertWidget)
	// setup sub layouts
	inputHorWidget := widgets.NewQWidget(nil, 0)
	ptHorWidget := widgets.NewQWidget(nil, 0)
	cHorWidget := widgets.NewQWidget(nil, 0)
	idepHorWidget := widgets.NewQWidget(nil, 0)
	wHorWidget := widgets.NewQWidget(nil, 0)

	inputHorWidget.SetLayout(widgets.NewQHBoxLayout())
	ptHorWidget.SetLayout(widgets.NewQHBoxLayout())
	cHorWidget.SetLayout(widgets.NewQHBoxLayout())
	idepHorWidget.SetLayout(widgets.NewQHBoxLayout())
	wHorWidget.SetLayout(widgets.NewQHBoxLayout())
	// define interactive elements
	input := widgets.NewQSpinBox(nil)
	input.SetMaximum(10000)
	inputLabel := widgets.NewQLabel2("Pattern area in square microns", nil, 0)
	// initialize output display elements
	ptLabel := widgets.NewQLabel2("Platinum", nil, 0)
	ptOutputBox := widgets.NewQLCDNumber(nil)
	ptOutputBox.SetFixedHeight(40)
	ptOutputBox.SetDigitCount(6)
	ptOutputBox.Display2Default(0)

	cLabel := widgets.NewQLabel2("Carbon", nil, 0)
	cOutputBox := widgets.NewQLCDNumber(nil)
	cOutputBox.SetFixedHeight(40)
	cOutputBox.SetDigitCount(6)
	cOutputBox.Display2Default(0)

	idepLabel := widgets.NewQLabel2("IDEP(2)", nil, 0)
	idepOutputBox := widgets.NewQLCDNumber(nil)
	idepOutputBox.SetFixedHeight(40)
	idepOutputBox.SetDigitCount(6)
	idepOutputBox.Display2Default(0)

	wLabel := widgets.NewQLabel2("Tungsten", nil, 0)
	wOutputBox := widgets.NewQLCDNumber(nil)
	wOutputBox.SetFixedHeight(40)
	wOutputBox.SetDigitCount(6)
	wOutputBox.Display2Default(0)

	// layout layouts on the page
	masterVertWidget.Layout().AddWidget(inputHorWidget)
	inputHorWidget.Layout().AddWidget(inputLabel)
	inputHorWidget.Layout().AddWidget(input)

	masterVertWidget.Layout().AddWidget(ptHorWidget)
	ptHorWidget.Layout().AddWidget(ptLabel)
	ptHorWidget.Layout().AddWidget(ptOutputBox)

	masterVertWidget.Layout().AddWidget(cHorWidget)
	cHorWidget.Layout().AddWidget(cLabel)
	cHorWidget.Layout().AddWidget(cOutputBox)

	masterVertWidget.Layout().AddWidget(idepHorWidget)
	idepHorWidget.Layout().AddWidget(idepLabel)
	idepHorWidget.Layout().AddWidget(idepOutputBox)

	masterVertWidget.Layout().AddWidget(wHorWidget)
	wHorWidget.Layout().AddWidget(wLabel)
	wHorWidget.Layout().AddWidget(wOutputBox)

	// update outputs on input change
	input.ConnectValueChanged(func(int) {
		ptOutputBox.Display2(tssMath.PtCalculate(input.Value()))
		cOutputBox.Display2(tssMath.CCalculate(input.Value()))
		idepOutputBox.Display2(tssMath.IdepCalculate(input.Value()))
		wOutputBox.Display2(tssMath.WCalculate(input.Value()))
	})

	window.Show()
	if debug == true { // debug window with current debug toolkit
		debugwindow := widgets.NewQMainWindow(nil, 0)
		debugwindow.SetWindowTitle("Debug")
		debugwindow.Show()
		mainWinHLabel := widgets.NewQLabel2("Horizontal Pixels in main window", nil, 0)
		mainWinVLabel := widgets.NewQLabel2("Vertical Pixels in main window", nil, 0)
		mainWinHPix := widgets.NewQLCDNumber(nil)
		mainWinVPix := widgets.NewQLCDNumber(nil)
		mainWinHPix.Display2Default(window.Width())
		mainWinVPix.Display2Default(window.Height())
		debugVirtWidget := widgets.NewQWidget(nil, 0)
		debugVirtWidget.SetLayout(widgets.NewQVBoxLayout())
		debugVirtWidget.Layout().AddWidget(mainWinHLabel)
		debugVirtWidget.Layout().AddWidget(mainWinHPix)
		debugVirtWidget.Layout().AddWidget(mainWinVLabel)
		debugVirtWidget.Layout().AddWidget(mainWinVPix)
		debugwindow.SetCentralWidget(debugVirtWidget)
		window.ConnectResizeEvent(func(*gui.QResizeEvent) {
			mainWinHPix.Display2(window.Width())
			mainWinVPix.Display2(window.Height())
		})
	}
	app.Exec()
}
