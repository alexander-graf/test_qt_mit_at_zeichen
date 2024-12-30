package main

import (
	"fmt"
	"time"

	"github.com/therecipe/qt/widgets"
)

func main() {
	app := widgets.NewQApplication(0, nil)
	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("Eingabeformular")
	window.SetMinimumSize2(400, 300)

	// Hauptwidget und Layout
	centralWidget := widgets.NewQWidget(nil, 0)
	layout := widgets.NewQVBoxLayout()
	centralWidget.SetLayout(layout)

	// Eingabefelder erstellen
	inputs := make([]*widgets.QLineEdit, 5)
	for i := 0; i < 5; i++ {
		inputs[i] = widgets.NewQLineEdit(nil)
		layout.AddWidget(inputs[i], 0, 0)
	}

	// Email-Feld vorausfüllen
	inputs[0].SetPlaceholderText("Email eingeben")
	inputs[0].SetText("fantasia@traumwelt.de")

	// Button erstellen
	button := widgets.NewQPushButton2("Absenden", nil)
	layout.AddWidget(button, 0, 0)

	// Statuszeile erstellen
	statusLabel := widgets.NewQLabel(nil, 0)
	statusLabel.SetText("Hallo")
	layout.AddWidget(statusLabel, 0, 0)

	// Button-Klick-Handler
	button.ConnectClicked(func(bool) {
		fmt.Println("Eingegebene Daten:")
		for i, input := range inputs {
			fmt.Printf("Feld %d: %s\n", i+1, input.Text())
		}
	})

	// Animation für Statuszeile in Goroutine
	go func() {
		text := "Hallo"
		direction := 1
		position := 0
		maxPosition := 30

		for {
			spaces := make([]byte, position)
			for i := range spaces {
				spaces[i] = ' '
			}

			displayText := string(spaces) + text

			statusLabel.SetText(displayText)

			if position >= maxPosition {
				direction = -1
			} else if position <= 0 {
				direction = 1
			}

			position += direction
			time.Sleep(100 * time.Millisecond)
		}
	}()

	window.SetCentralWidget(centralWidget)
	window.Show()
	app.Exec()
}
