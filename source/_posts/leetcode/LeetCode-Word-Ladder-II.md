---
title: 'LeetCode: Word Ladder II'
date: 2020-03-15 15:03:56
tags:
---
```
const WildCard = ","
const InfinityDistance = math.MaxInt32

type Ladder string

func (l Ladder) AddWord(w string) Ladder {
	return Ladder(fmt.Sprintf("%s%s%s", l, WildCard, w))
}

func (l Ladder) Duplicated(l1 Ladder) bool {
	return l == l1
}

func (l Ladder) ToResult() []string {
	return strings.Split(string(l), WildCard)
}

type Word struct {
	Content   string
	Nearbys   map[string]*Word
	Ladders   map[Ladder]bool
	Distance  int32
	Estimate  int32
	Available bool
}

func CalculateEstimate(w1, w2 *Word) int32 {
	if w1 == nil || w2 == nil || len(w1.Content) != len(w2.Content) {
		return InfinityDistance
	}

	var result int32

	for i, c := range w1.Content {
		if c != rune(w2.Content[i]) {
			result++
		}
	}
	return result
}

func Link(w1, w2 *Word) {
	if CalculateEstimate(w1, w2) == 1 {
		w1.Nearbys[w2.Content] = w2
		w2.Nearbys[w1.Content] = w1
	}
}

type Words []*Word

func (ws *Words) Size() int {
	if ws == nil {
		return 0
	}
	return len(*ws)
}

func (ws Words) Len() int {
	return len(ws)
}

func (ws Words) Less(i, j int) bool {
	return ws[i].Estimate+ws[i].Distance < ws[j].Estimate+ws[j].Distance
}

func (ws Words) Swap(i, j int) {
	ws[i], ws[j] = ws[j], ws[i]
}

func (ws *Words) Push(w interface{}) {
	if ws == nil {
		return
	}
	//TODO 探索一下
	*ws = append(*ws, w.(*Word))
}

func (ws *Words) Pop() interface{} {
	old := *ws
	n := len(old)
	item := old[n-1]
	*ws = old[0 : n-1]
	return item
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	MinDistance := int32(math.MaxInt32)
	wordMap := make(map[string]*Word)
	origin := &Word{
		Content: beginWord,
		Nearbys: make(map[string]*Word),
		Ladders: map[Ladder]bool{
			Ladder(beginWord): true,
		},
		Distance:  0,
		Estimate:  InfinityDistance,
		Available: false,
	}

	destination := &Word{
		Content:   endWord,
		Nearbys:   make(map[string]*Word),
		Ladders:   make(map[Ladder]bool),
		Distance:  InfinityDistance,
		Estimate:  0,
		Available: true,
	}
	origin.Estimate = CalculateEstimate(origin, destination)
	wordMap[beginWord] = origin

	for _, w := range wordList {
		if wordMap[w] != nil {
			continue
		}
		word := &Word{
			Content:   w,
			Nearbys:   make(map[string]*Word),
			Ladders:   make(map[Ladder]bool),
			Available: true,
			Distance:  InfinityDistance,
		}
		word.Estimate = CalculateEstimate(word, destination)
		wordMap[w] = word

	}

	if wordMap[endWord] == nil {
		return nil
	}
	wordMap[endWord] = destination

	for _, w1 := range wordMap {
		for _, w2 := range wordMap {
			Link(w1, w2)
		}
	}

	wordHeap := &Words{}

	heap.Push(wordHeap, origin)

	for {
		if wordHeap.Size() == 0 {
			break
		}
		word := heap.Pop(wordHeap).(*Word)
		for _, nearby := range word.Nearbys {
			if nearby.Distance >= word.Distance+1 {
				nearby.Available = true
			}
			if !nearby.Available {
				continue
			}
			if nearby.Distance > word.Distance+1 {
				nearby.Distance = word.Distance + 1
				nearby.Ladders = make(map[Ladder]bool)
				for ladder, _ := range word.Ladders {
					nearby.Ladders[ladder.AddWord(nearby.Content)] = true
				}
			} else if nearby.Distance == word.Distance+1 {
				for ladder, _ := range word.Ladders {
					nearby.Ladders[ladder.AddWord(nearby.Content)] = true
				}
			}
			if nearby.Content == destination.Content {
				if nearby.Distance < MinDistance {
					MinDistance = nearby.Distance
				}
			}
			if nearby.Distance+nearby.Estimate <= MinDistance {
				nearby.Available = false
				heap.Push(wordHeap, nearby)
			} else {
				nearby.Available = false
			}
		}
	}

	var result [][]string
	for ladder, _ := range destination.Ladders {
		result = append(result, ladder.ToResult())
	}
	return result

}
```
