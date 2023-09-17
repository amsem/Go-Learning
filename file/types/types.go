
package types

import "fmt"

type Mdistance float64 //miles
type Kdistance float64 //KM

type location string

func (origin location) DistanceTo(dest location) Kdistance {
	//TODO calc ......
	fmt.Printf("Origin %v Destination %v", origin, dest)
	return 10
}

func LocationTest() {
	nyc := location("21.321, 312.3412")
	nyc.DistanceTo(location("32.3241, 43.43213"))
}

// golang methods
func (miles Mdistance) ToKm() Kdistance {
	return Kdistance(1.609 * miles)
}


func (km Kdistance) ToMile() Mdistance {
	return Mdistance (km / 1.609)
}

