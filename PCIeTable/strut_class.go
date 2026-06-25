package pcietable

type Class struct {
	ClassID    string
	Name       string
	SubClasses []*Class
}

func (c *Class) AddChild(line string) Child {
	class := lineToClass(line)
	c.SubClasses = append(c.SubClasses, class)
	return class
}
