package main

import (
	"strconv"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type Booking struct{
	booking_id int
	booking_date string
	num_weeks int
	property_owner_name string
	contact_number string
	address string
	rooms int
	rooms_cost int
	garden_area int
	garden_area_cost int
	security_alarm_check bool
	pool_maintenance bool
}

func get_luxury_cost(luxury_cost int, security_alarm_check bool, pool_maintenance bool) int{
	if(security_alarm_check){
		luxury_cost += 50
	}
	if(pool_maintenance){
		luxury_cost += 50
	}
	return luxury_cost
}

func save(bookings[]Booking, booking_date string, num_weeks int, property_owner_name string, contact_number string, address string, rooms int, garden_area int, security_alarm_check bool, pool_maintenance bool) []Booking{
	booking := Booking{(len(bookings) + 1), booking_date, num_weeks, property_owner_name, contact_number, address, rooms, (rooms * 5), garden_area, (garden_area * 2), security_alarm_check, pool_maintenance}
	bookings = append(bookings, booking)
	return bookings
}

func booking_gui(){
	app := app.New()
	window := app.NewWindow("Booking GUI")

	var bookings = []Booking{}
	s_booking_date := widget.NewEntry()
	s_booking_date.SetPlaceHolder("Booking Date")
	s_num_weeks := widget.NewEntry()
	s_num_weeks.SetPlaceHolder("Num Weeks")
	s_property_owner_name := widget.NewEntry()
	s_property_owner_name.SetPlaceHolder("Property Owner Name")
	s_contact_number := widget.NewEntry()
	s_contact_number.SetPlaceHolder("Contact Number")
	s_address := widget.NewEntry()
	s_address.SetPlaceHolder("Address")
	s_rooms := widget.NewEntry()
	s_rooms.SetPlaceHolder("Rooms")
	s_garden_area := widget.NewEntry()
	s_garden_area.SetPlaceHolder("Garden Area")
	security_alarm_check := false
	s_security_alarm_check := widget.NewCheck("Security Alarm Check", func(b_security_alarm_check bool){security_alarm_check = b_security_alarm_check})
	pool_maintenance := false
	s_pool_maintenance := widget.NewCheck("Pool Maintenance", func(b_pool_maintenance bool){pool_maintenance = b_pool_maintenance})
	output := binding.NewString()
	s_output := ""
	output.Set(s_output)

	save_button := widget.NewButton("Save", func(){
		num_weeks, _ := strconv.Atoi(s_num_weeks.Text)
		rooms, _ := strconv.Atoi(s_rooms.Text)
		garden_area, _ := strconv.Atoi(s_garden_area.Text)
		bookings = save(bookings, s_booking_date.Text, num_weeks, s_property_owner_name.Text, s_contact_number.Text, s_address.Text, rooms, garden_area, security_alarm_check, pool_maintenance)
		s_booking_date.SetText("")
		s_num_weeks.SetText("")
		s_property_owner_name.SetText("")
		s_contact_number.SetText("")
		s_address.SetText("")
		s_rooms.SetText("")
		s_garden_area.SetText("")
		s_security_alarm_check.SetChecked(false)
		s_pool_maintenance.SetChecked(false)
	})
	display_button := widget.NewButton("Display", func(){
		s_output = ""
		output.Set(s_output)
		for i := 0; i < len(bookings); i++ {
			s_output += "ID: " + strconv.Itoa(bookings[i].booking_id) + " Booking Date: " + bookings[i].booking_date + " Num. Weeks: " + strconv.Itoa(bookings[i].num_weeks) + " Property Owner: " + bookings[i].property_owner_name + " Contact Number: " + bookings[i].contact_number + " Address: " + bookings[i].address + " Rooms: " + strconv.Itoa(bookings[i].rooms) + " Garden Area: " + strconv.Itoa(bookings[i].garden_area) + " Total Cost: " + strconv.Itoa(get_luxury_cost((bookings[i].rooms_cost + bookings[i].garden_area_cost), bookings[i].security_alarm_check, bookings[i].pool_maintenance)) + "\n"
		}
		output.Set(s_output)
	})

	input_container := container.NewVBox(s_booking_date, s_num_weeks, s_property_owner_name, s_contact_number, s_address, s_rooms, s_garden_area, s_security_alarm_check, s_pool_maintenance)
	button_container := container.NewHBox(save_button, display_button)
	container := container.NewVBox(input_container, button_container, widget.NewLabelWithData(output))

	window.SetContent(container)
	window.ShowAndRun()
}

func main(){
	booking_gui()
}
