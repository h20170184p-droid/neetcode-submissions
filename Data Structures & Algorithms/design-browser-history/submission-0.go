type tabNode struct {
	tabName string
	prevtabAddrr *tabNode
	nexttabAddrr *tabNode 
}

type BrowserHistory struct {
    head *tabNode
	tail *tabNode
	currentAddrr *tabNode
}

func Constructor(homepage string) BrowserHistory {
	dummyhead := &tabNode{tabName: ""}
	dummytail := &tabNode{tabName: ""}

	homepageNode := &tabNode{
		tabName : homepage,
		prevtabAddrr : dummyhead,
		nexttabAddrr : dummytail,
	}

	dummyhead.nexttabAddrr = homepageNode
	dummytail.prevtabAddrr = homepageNode
	return BrowserHistory {
		head : dummyhead,
		tail : dummytail,
		currentAddrr : homepageNode,
	}
}


func (this *BrowserHistory) Visit(url string)  {
	// When new visit happens, first need to know where I am so far
	// That is more or less same as Get applied on the index
    presentpageAddrr := this.currentAddrr
	newPageAddrr := &tabNode{
		tabName : url,
		prevtabAddrr : presentpageAddrr,
		nexttabAddrr : this.tail,
	}
	presentpageAddrr.nexttabAddrr = newPageAddrr
	this.tail.prevtabAddrr = newPageAddrr
	this.currentAddrr = newPageAddrr
}


func (this *BrowserHistory) Back(steps int) string {
    for i := 0; i < steps; i++ {
		if this.currentAddrr.prevtabAddrr == this.head {
			break
		}
		this.currentAddrr = this.currentAddrr.prevtabAddrr
	}
	return this.currentAddrr.tabName
}


func (this *BrowserHistory) Forward(steps int) string {
    for i := 0; i < steps; i++ {
		if this.currentAddrr.nexttabAddrr == this.tail {
			break
		}
		this.currentAddrr = this.currentAddrr.nexttabAddrr
	}
	return this.currentAddrr.tabName
}


/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */