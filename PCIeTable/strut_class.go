package pcietable

type Class struct {
	ClassID    string
	Name       string
	SubClasses map[string]*Class
}

func (c *Class) AddChild(line string) Child {
	if c.SubClasses == nil {
		c.SubClasses = make(map[string]*Class)
	}
	class := lineToClass(line)
	c.SubClasses[class.ClassID] = class
	return class
}
