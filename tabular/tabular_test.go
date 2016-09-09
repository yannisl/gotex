package tabular

//import "gotex/tabular"

import "testing"

var z = new(Tabular)

// 
func TestTabular(t *testing.T) {
  //z := new(Tabular)
  expected:= false
  actual := z.HasHead
  if actual != expected {
    t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
  }
}

// checkSpecification list is set to true as a default
func TestCheckSpecificationList(t *testing.T) {
	expected:=true
	actual := z.checkSpecificationList()
	if actual != expected {
    	t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
  	}
  }

//Check that we are inserting in Data
func TestInsert(t *testing.T) {
	z.Insert("Qty", "Apr", "May", "Jun", "Jul")
	expected := 1
	actual := len(z.Data)
	if actual != expected {
    	t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
  	}
}

  
