package pcietable

type Parent interface {
	AddChild(line string) Child
}
type Child interface {
	Parent
}
