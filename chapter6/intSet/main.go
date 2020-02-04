package main

import "strconv"

import "fmt"

type IntSet struct {
	words []uint64
}

func (this *IntSet) Insert(val int) {
	word := val / 64
	bias := val % 64

	for word >= len(this.words) {
		this.words = append(this.words, 0)
	}
	this.words[word] = this.words[word] | (1 << bias)
}

func (this *IntSet) Delete(val int) bool {
	word := val / 64
	bias := val % 64
	if word >= len(this.words) {
		return false
	}
	if this.words[word]&(1<<bias) == 0 {
		return false
	}
	this.words[word] &= (1 << bias)
	return true
}

func (this *IntSet) Has(val int) bool {
	word := val / 64
	bias := val % 64
	if word >= len(this.words) {
		return false
	}
	if this.words[word]&(1<<bias) == 0 {
		return false
	}
	return true
}

func (this *IntSet) Union(t *IntSet) {
	for i := 0; i < len(t.words)*64; i++ {
		if t.Has(i) {
			this.Insert(i)
		}
	}
}

func (this *IntSet) String() string {
	res := "{"
	for i := 0; i < len(this.words); i++ {
		for j := 0; j < 64; j++ {
			if this.words[i]&(1<<j) != 0 {
				res += strconv.FormatInt(int64(i*64+j), 10)
				res += ", "
			}
		}
	}
	res += " }"
	return res
}

func main() {
	s := IntSet{}
	t := IntSet{}
	for i := 0; i < 100; i += 2 {
		s.Insert(i)
		t.Insert(i + 1)
	}
	fmt.Println(s.String())
	fmt.Println(t.String())
	s.Union(&t)
	fmt.Println(s.String())

}
