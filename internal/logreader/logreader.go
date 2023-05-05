package logreader

var r *reader

func Init() error {
	r = new(reader)
	return r.run()
}

func GetLine() string {
	return r.read()
}
