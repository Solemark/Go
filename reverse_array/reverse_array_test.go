package reverse_array

import "testing"

func Test_reverse_array(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	answer := reverse_array(input)

	for i:= 0; i <= len(input)-1; i++ {
		if(answer[i] != expected[i]){
			t.Error("Answer does not equal expected array!", answer, expected)
			return
		}
	} 
}