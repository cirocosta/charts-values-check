package pkg

type Finder interface {
	Find([]byte) (results []string, err error)
}
