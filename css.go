package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func Tea() {
	var title = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		Blink(true).
		Align(lipgloss.Center).
		//

		PaddingTop(0).
		PaddingBottom(0).
		PaddingRight(1).
		PaddingLeft(1).
		//
		MarginTop(2).
		MarginRight(2).
		MarginLeft(2).
		MarginBottom(1).
		Width(100)

	fmt.Println(title.Render("WazaMan"))

	var icon = lipgloss.NewStyle().SetString(`
 
  \   /     
   .-.      
― (   ) ―   
   ` + "`-’`" + `     
  /   \     

`).
		Bold(true).
		Foreground(lipgloss.Color("#FFFF00")).
		//Background(lipgloss.Color("201")).
		PaddingTop(0).
		PaddingBottom(0).
		PaddingLeft(0).
		Width(0)

	// fmt.Println(icon.Render(`

	//   \   /
	//    .-.
	// ― (   ) ―
	//    ` + "`-’`" + `
	//   /   \

	// `))

	var info = lipgloss.NewStyle().SetString("35C").
		Bold(false).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		Blink(true).
		Align(lipgloss.Left).
		//

		PaddingTop(0).
		PaddingBottom(0).
		PaddingRight(0).
		PaddingLeft(0).
		//
		MarginTop(0).
		MarginRight(0).
		MarginLeft(0).
		MarginBottom(0).
		Width(10)

	// fmt.Println(info.Render("35C"))

	fmt.Println(lipgloss.JoinHorizontal(lipgloss.Center, icon.Render(), info.Render()))
}
