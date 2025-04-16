package basics

import "fmt"

/*type Gender string

const (
	Male   Gender = "Male"
	Female Gender = "Female"
)

func ConstAndEnumMethod(gender Gender) {}

func (g *Gender) String() string {
	switch *g {
	case Male:
		return "Male"
	case Female:
		return "Female"
	default:
		return "Unknown"
	}
}

func (g *Gender) isMale() bool {
	return *g == Male
}*/

type Gender int

const (
	Male Gender = iota
	Female
)

func (g *Gender) isMale() bool {
	return *g == Male
}

func (g *Gender) isFemale() bool {
	return *g == Female
}

func mainConstAndEnum() {
	var gender = Male
	fmt.Println(gender)
	fmt.Println(gender.isMale())
	fmt.Println(gender.isFemale())
	//flag := gender.String()
	//fmt.Println(flag)
}
