package shup

type FileRegistry struct {
	files []string
}

func (fr *FileRegistry) Register(filepath string) (bool, string, error) {
	abspath, err := AbsPath(filepath)
	if err != nil {
		return false, "", err
	}

	if fr.Has(abspath) {
		return false, "", nil
	}

	fr.insert(abspath)
	return true, abspath, nil
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