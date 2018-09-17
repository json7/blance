package strategy

import (
	"math"
	"github.com/pkg/errors"
)

type blance struct {
	counter int //序列计数器
	curweight int //当前权重
	servers []Servers //节点列表
}

type Servers struct {
	Service string
	Weight int
	Provider string
}

func NewBlance(sers []Servers) (*blance, error) {
	b := blance{counter: -1, curweight: 0, servers: make([]Servers, len(sers))}
	l := copy(b.servers, sers)
	if len(sers)<=0 || l!=len(sers) {
		return nil, errors.New("the numbers of servers is 0")
	}
	return &b, nil
}

//获取权重总和
func (self *blance) getsum(s []Servers) int {
	r := 0
	for _, server := range s {
		r += server.Weight
	}
	return r
}

//获取最大公约数
func (self *blance) gcd(a int, b int) int {
	c := 0
	for b>0 {
		c = b
		b = a%b
		a = c
	}
	return a
}

//获取最大公约数
func (self *blance) getgcd(s []Servers) int {
	res := s[0].Weight
	for i:=1; i<len(s); i++ {
		max := int(math.Max(float64(res), float64(s[i].Weight)))
		min := int(math.Min(float64(res), float64(s[i].Weight)))
		res = self.gcd(max, min)
	}
	return res
}

//获取最大的权重
func (self *blance) getmax(s []Servers) int {
	m := 0
	for _, server := range s {
		if server.Weight > m {
			m = server.Weight
		}
	}
	return m
}

func (self *blance) lb_wrr__getwrr(s []Servers, gcd int, maxweight int, i *int, cw *int) int {
	for {
		*i = (*i+1)%len(s)
		if *i == 0{
			*cw = *cw-gcd
			if *cw <= 0 {
				*cw = maxweight
				if *cw == 0 {
					return -1
				}
			}
		}
		if s[*i].Weight >= *cw {
			return *i
		}
	}
}

func (self *blance) GetServer() Servers {
	gcd := self.getgcd(self.servers)
	max := self.getmax(self.servers)
	self.lb_wrr__getwrr(self.servers, gcd, max, &self.counter, &self.curweight)
	return self.servers[self.counter]
}