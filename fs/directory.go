package fs

type Directory struct {
	name string
	Contents *[]Unit
}

func (d Directory) Name() string {
	return d.name
}

func newDir(name string, contents *[]Unit) Directory {
	return Directory{name, contents}
}