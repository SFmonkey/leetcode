package LinkedList

import "fmt"

type BrowserHistory struct {
	homePage		*node			// 根页面
	currPage		*node			// 当前所在页面
}

type node struct {
	url				string
	prevNode		*node
	nextNode		*node
}

// 创建root对象
func Constructor(homepage string) BrowserHistory {
	root := new(BrowserHistory)
	root.homePage = &node{
		url: homepage,
	}
	root.currPage = root.homePage
	return *root
}

// todo 这里的this和调用它的obj其实是两个变量，同 l1:=l2
func (this *BrowserHistory) Visit(url string)  {
	// 删除前进记录
	this.currPage.nextNode = nil
	// 尾部添加新页面
	this.currPage.nextNode = &node{
		url:      url,
		prevNode: this.currPage,
		nextNode: nil,
	}
	this.currPage = this.currPage.nextNode
}

// 回退n个节点
func (this *BrowserHistory) Back(steps int) string {
	for i:=0; i<steps; i++ {
		if this.currPage.prevNode == nil {
			return this.currPage.url
		}
		this.currPage = this.currPage.prevNode
	}
	return this.currPage.url
}

// 前进n个节点
func (this *BrowserHistory) Forward(steps int) string {
	for i:=0; i<steps; i++ {
		if this.currPage.nextNode == nil {
			return this.currPage.url
		}
		this.currPage = this.currPage.nextNode
	}
	return this.currPage.url
}

func printHistory(b *BrowserHistory)  {
	head := b.homePage
	for head != nil {
		fmt.Print(fmt.Sprintf("%s->", head.url))
		head = head.nextNode
	}
	fmt.Println("nil")
}