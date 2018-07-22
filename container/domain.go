package container

//Domain
type Domain []Interval

//NewDomain
func NewDomain() *Domain {
	return new(Domain)
}

//Append
func (d *Domain) Append(i Interval) {
	*d = append(*d, i)
}

//First
func (d Domain) First() Interval {
	return d[0]
}

//Last
func (d Domain) Last() {

}

//Len
func (d Domain) Len() int {
	return len(d)
}
