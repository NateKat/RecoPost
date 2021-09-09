package parcel

type Parcel struct {
	uid    string
	weight int
}

func New(uid string, weight int) Parcel {
	p := Parcel{uid, weight}
	return p
}
