package tree

type child struct {
	Index *ParentChildTree
	child *child
	data  string
}

type ParentChildTree struct {
	parentIndex *ParentChildTree
	firstChild  *child
	data        string
}

//树表头
type ParentChildTreeBox []*ParentChildTree
