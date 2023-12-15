package hascii

import "slices"

type Hashmap [256][]Lens

func (hashmap *Hashmap) Add(lens Lens) {
	hash := Hash(lens.Label)
	box := hashmap[hash]

	for s, slot := range box {
		if slot.Label == lens.Label {
			box[s] = lens
			return
		}
	}
	hashmap[hash] = append(box, lens)
}

func (hashmap *Hashmap) Delete(label string) {
	hash := Hash(label)
	box := hashmap[hash]

	hashmap[hash] = slices.DeleteFunc(box, func(l Lens) bool {
		return l.Label == label
	})
}
