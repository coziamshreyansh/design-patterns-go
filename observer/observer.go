package observer

import "fmt"

type BaseObserver interface {
	Update(news string)
}

type Subject interface {
	AddObserver(o BaseObserver)
	RemoveObserver(o BaseObserver)
	NotifyObservers()
}

type NewsAgency struct {
	O    []BaseObserver
	news string
}

func (n *NewsAgency) SetNews(news string) {
	n.news = news
	n.NotifyObservers()
}

func (n *NewsAgency) AddObserver(o BaseObserver) {
	n.O = append(n.O, o)
}

func (n *NewsAgency) RemoveObserver(o BaseObserver) {
	for i, observer := range n.O {
		if o == observer {
			n.O = append(n.O[:i], n.O[i+1:]...)
			break
		}
	}
}
func (n *NewsAgency) NotifyObservers() {
	for _, o := range n.O {
		o.Update(n.news)
	}
}

type NewsChannel struct {
	Name string
}

func NewNewsChannel(name string) BaseObserver {
	return &NewsChannel{Name: name}
}

func (n *NewsChannel) Update(news string) {
	fmt.Printf("We are watching news channel: %s, with news: %s", n.Name, news)
}
