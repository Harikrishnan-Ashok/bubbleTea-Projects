package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choices []string
	cursor int
	selected map[int]struct{}
}

func initalModel() model {
	return model{
		choices: []string{"test1","test2","test3"},
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd{
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd){
	switch msg := msg.(type){
		case tea.KeyMsg:
			switch msg.String(){
				case "q","ctrl+c":
					return m,tea.Quit
				case "up":
					if m.cursor>0{
						m.cursor--
					}
				case "down":
					if m.cursor<len(m.choices)-1{
						m.cursor++
					}
				case "enter":
					_,ok:=m.selected[m.cursor]
					if ok{
						delete(m.selected,m.cursor)
					}else{
						m.selected[m.cursor]= struct{}{}
					}
			}
	}
	return m,nil
}

func (m model) View() string {
	s:="pick a choice nigga\n\n"

	for i,choice := range m.choices{
		
		cursor:=" "
		if m.cursor==i{
			cursor=">"
		}
		
		checked:=" "
		_,ok:=m.selected[i] 
		if ok{
			checked="x"
		}

		s+=fmt.Sprintf("%s [%s] %s \n",cursor,checked,choice)
	}
	s+="\npress q or ctrl+c to quit , enter to check , up and down to navigate \n"
	return s
}

func main(){
	p:= tea.NewProgram(initalModel())
	p.Start()
}

