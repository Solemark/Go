package calculator

import "testing"
 
func Test_addition(t *testing.T) {
   want := 12
 
   ans := addition(5, 7)
 
   if ans != want {
       t.Fatalf("got %d, wanted %d", ans, want)
   }
}
 
func Test_subtraction(t *testing.T) {
   want := 0
 
   ans := subtraction(5, 5)
 
   if ans != want {
       t.Fatalf("got %d, wanted %d", ans, want)
   }
}
 
func Test_multiplication(t *testing.T) {
   want := 25
 
   ans := multiplication(5, 5)
 
   if ans != want {
       t.Fatalf("got %d, wanted %d", ans, want)
   }
}
 
func Test_division(t *testing.T) {
   want := 1.0
 
   ans := division(5, 5)
 
   if ans != want {
       t.Fatalf("got %f, wanted %f", ans, want)
   }
}