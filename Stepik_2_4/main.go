package main

type test_Struct struct {
	On    bool
	Ammo  int
	Power int
}

func (s *test_Struct) Shoot() bool {
	var a bool
	if s.On == false {
		a = false
	} else if s.Ammo > 0 {
		a = true
		s.Ammo--
	}
	return a
}

func (s *test_Struct) RideBike() bool {
	var a bool
	if s.On == false {
		a = false
	} else if s.Power > 0 {
		s.Power--
		a = true
	}
	return a
}
