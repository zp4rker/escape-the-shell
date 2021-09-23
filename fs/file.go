package fs

type File struct {
	name string
	contents string
}

func (f File) Name() string {
	return f.name
}

func (f File) Contents() string {
	return f.contents
}

func newFile(name string, contents string) File {
	return File{name, contents}
}
