package main

type NestedInteger struct {
}

func (n *NestedInteger) isInteger() bool {
	return true
}

func (n *NestedInteger) getInteger() int {
	return 0
}

func (n *NestedInteger) getList() []NestedInteger {
	return nil
}

type Algorithm22 struct {
	ResultList []int
}

func (this *Algorithm22) flatten(nestList []NestedInteger) ([]int) {
	this.doFlatten(nestList)
	return this.ResultList
}

func (this *Algorithm22) doFlatten(nestList []NestedInteger) {
	if nestList != nil {
		for i := 0; i < len(nestList); i++ {
			child := nestList[i]
			if child.isInteger() {
				this.ResultList = append(this.ResultList,child.getInteger())
			} else {
				t:=child.getList()
				this.doFlatten(t)
			}
		}
	}
}


