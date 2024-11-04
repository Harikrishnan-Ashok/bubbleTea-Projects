package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {}

func (m model) Init() tea.Cmd{
	return nil
}

func (m model) Update(msg tea.Msg)(tea.Model,tea.Cmd){
	switch msg:= msg.(type){
		case tea.KeyMsg:
			switch msg.String(){
				case "q","ctrl+q":
					return m,tea.Quit
			}
	}
	return m,nil
}

func (m model) View() string {
	s:="hello world"
	s += fmt.Sprintf("\npress q or ctrl+c to quit\n ")
	return s
}

func main(){
	p:=tea.NewProgram(model{},tea.WithAltScreen())
	p.Run()
}
