package convert_distance

import "testing"

func Test_convert_distance_10km_to_miles(t *testing.T){
	km := 10.0
	ktm := 0.6213712
	answer := convert_distance(km, ktm)
	
	if answer != 6.213712 {
		t.Errorf("convert_distance(%g, %g) = 6.213712, recieved: %g", km, ktm, answer)
	}

}