package tag

type Tag string

func (t *Tag) New(s string) {
	*t = Tag(s)
}
