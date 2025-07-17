package collectcrds

type Options struct {
	Directory string
}

func NewOptions() *Options {
	return &Options{}
}

func NewOptionsWithDirectory(directory string) *Options {
	return &Options{
		Directory: directory,
	}
}
