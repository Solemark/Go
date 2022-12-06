package main

import (
	"strconv"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type Student struct{
	name string
	mark int
}

func get_grade(mark int) string{
	switch {
		case mark < 50:
			return "F"
		case mark < 65:
			return "P"
		case mark < 75:
			return "C"
		case mark < 85:
			return "D"
		default:
			return "HD"
	}
}

func save(students[]Student, name string, s_mark string) []Student{
	mark, _ := strconv.Atoi(s_mark)
	student := Student{name, mark}
	students = append(students, student)
	
	return students
}

func main() {
	app := app.New()
	window := app.NewWindow("Marks GUI")

	var students = []Student{}
	name_input := widget.NewEntry()
	name_input.SetPlaceHolder("Name")
	mark_input := widget.NewEntry()
	mark_input.SetPlaceHolder("Mark")
	output := binding.NewString()

	s_output := ""
	output.Set("")
	
	save_button := widget.NewButton("Save", func(){
		students = save(students, name_input.Text, mark_input.Text)
		name_input.SetText("")
		mark_input.SetText("")
	})
	display_button := widget.NewButton("Display", func(){
		s_output = ""
		output.Set(s_output)
		for i := 0; i < len(students); i++ {
			s_output += students[i].name + " " + get_grade(students[i].mark) + "\n"
		}
		output.Set(s_output)
	})
	input_container := container.New(layout.NewGridLayout(2), name_input, mark_input)
	button_container := container.NewHBox(save_button, display_button)
	container := container.NewVBox(input_container, button_container, widget.NewLabelWithData(output))
	window.SetContent(container)
	window.ShowAndRun()
}
