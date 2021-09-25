package fs

var CurrentDir = &root
var Path []*Directory

var root = newDir("root", &[]Unit{
	newFile("test", "this is the content"),
	newFile("test2", "this is the content"),
	newDir("test_dir", &[]Unit{
		newDir("another_dir", &[]Unit{
			newFile("found_me", "congratulations!"),
		}),
	}),
})

func Find(name string) *Unit {
	var search *Unit
	for _, u := range *CurrentDir.Contents {
		if u.Name() == name {
			search = &u
			break
		}
	}
	return search
}

func ChDir(d *Directory) {
	if d == nil {
		if len(Path) > 0 {
			parent := Path[len(Path) - 1]
			Path = Path[:len(Path) - 1]
			CurrentDir = parent
		} else {
			CurrentDir = &root
		}
	} else {
		if CurrentDir != nil {
			Path = append(Path, CurrentDir)
		}
		CurrentDir = d
	}
}

func MkDir(name string) {
	dir := newDir(name, &[]Unit{})
	*CurrentDir.Contents = append(*CurrentDir.Contents, dir)
}
