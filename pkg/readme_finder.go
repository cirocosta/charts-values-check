package pkg

import (
	"github.com/pkg/errors"
	bf "gopkg.in/russross/blackfriday.v2"
)

var (
	ErrTableNotFound             = errors.Errorf("couldn't find table")
	ErrUnexpectedTableFormatting = errors.Errorf("unexpected table formatting")
)

type ReadmeFinder struct{}

func (f *ReadmeFinder) Find(input []byte) (parameters []string, err error) {
	markdown := bf.New(bf.WithExtensions(bf.Tables))

	n := markdown.Parse(input)

	table := findTable(n)
	if table == nil {
		err = ErrTableNotFound
		return
	}

	for row := table.FirstChild; row != nil; row = row.Next {
		if row.FirstChild == nil ||
			row.FirstChild.FirstChild == nil ||
			row.FirstChild.FirstChild.Next == nil {
			err = ErrUnexpectedTableFormatting
			return
		}

		parameters = append(parameters,
			string(row.FirstChild.FirstChild.Next.Literal))
	}

	return
}

func findTable(n *bf.Node) (res *bf.Node) {
	n.Walk(func(node *bf.Node, entering bool) bf.WalkStatus {
		if node.Type != bf.TableBody {
			return bf.GoToNext
		}

		res = node
		return bf.Terminate
	})

	return
}
