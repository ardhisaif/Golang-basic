package src

type Hitung interface {
	Hitung2d
	Hitung3d
}

type Hitung2d interface {
	Luas() float64
	Keliling() float64
}

type Hitung3d interface {
	Volume() float64
}

type Data struct {
	Sisi *float64
}

func (d *Data) Luas() float64 {
	return *d.Sisi * *d.Sisi
}

func (d *Data) Keliling() float64 {
	return 4 * *d.Sisi
}

func (d *Data) Volume() float64 {
	return *d.Sisi * *d.Sisi * *d.Sisi
}
