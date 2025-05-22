package shup

type FileRegistry struct {
	files []string
}

func (fr *FileRegistry) Register(filepath string) (bool, string, error) {
	// resolve absolute path
	// if file already registered, return `false, "", nil`
	// else register, and return `true, abspath, nil`
	// on error return `false, "", error{}`
}

func (fr *FileRegistry) insert(filepath string) {
	// Isolated. Optimization likely unnecessary, but it's here.
	fr.files = append(fr.files, filepath)
}

func (fr FileRegistry) Has(abspath string) bool {
	// Update in conjunction with `insert()` function
	for _, path := range(fr.files) {
		if path == abspath {
			return true
		}
	}
	return false
}