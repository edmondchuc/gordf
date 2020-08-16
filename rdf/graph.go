package rdf

import (
	"github.com/edmondchuc/gordf/set"
)

// A Graph struct with three indexes of the different possible triple permutations for fast query access.
type Graph struct {
	spo map[Node]map[Node]set.Set
	pos map[Node]map[Node]set.Set
	osp map[Node]map[Node]set.Set
}

func (g *Graph) Add(s, p, o Node) {
	addToIndex(&g.spo, s, p, o)
	addToIndex(&g.pos, p, o, s)
	addToIndex(&g.osp, o, s, p)
}

func addToIndex(index *map[Node]map[Node]set.Set, a, b, c Node) {
	// Initialise the current index if it is nil.
	if *index == nil {
		*index = make(map[Node]map[Node]set.Set)
	}

	// If a not in index
	if _, ok := (*index)[a]; !ok {
		(*index)[a] = make(map[Node]set.Set)
		s := set.Set{}
		s.Add(c)
		(*index)[a][b] = s
	} else {
		// if b not in index
		if _, ok := (*index)[a][b]; !ok {
			s := set.Set{}
			s.Add(c)
			(*index)[a][b] = s
		} else {
			s := (*index)[a][b]
			s.Add(c)
		}
	}
}

func (g *Graph) Triples(s, p, o Node) chan Triple {
	triples := make(chan Triple)

	go func() {
		// if s is not None
		if _, ok := s.(None); !ok {
			// if P is not None
			if _, ok := p.(None); !ok {
				// if O is not None
				// sub pred obj
				if _, ok := o.(None); !ok {
					// if O in spo index
					if (*g).spo[s][p].Exists(o) {
						triples <- Triple{S: s, P: p, O: o}
					}
				} else {
					// sub pred None
					retObjects := (*g).spo[s][p]
					for retObj := range retObjects.Range() {
						retObjNode, _ := retObj.(Node)
						triples <- Triple{S: s, P: p, O: retObjNode}
					}
				}
			} else {
				// sub None obj
				if _, ok := o.(None); !ok {
					retPredicates := (*g).osp[o][s]
					for retPred := range retPredicates.Range() {
						retPredNode, _ := retPred.(Node)
						triples <- Triple{S: s, P: retPredNode, O: o}
					}
				} else {
					// sub None None
					retPredicates := (*g).spo[s]
					for retPred := range retPredicates {
						retObjects := (*g).spo[s][retPred]
						for retObj := range retObjects.Range() {
							retObjNode, _ := retObj.(Node)
							triples <- Triple{S: s, P: retPred, O: retObjNode}
						}
					}
				}
			}
		} else {
			if _, ok := p.(None); !ok {
				// None pred obj
				if _, ok := o.(None); !ok {
					retSubjects := (*g).pos[p][o]
					for retSub := range retSubjects.Range() {
						retSubNode, _ := retSub.(Node)
						triples <- Triple{S: retSubNode, P: p, O: o}
					}
				} else {
					// None pred None
					retObjects := (*g).pos[p]
					for retObj := range retObjects {
						retSubjects := (*g).pos[p][retObj]
						for retSub := range retSubjects.Range() {
							retSubNode, _ := retSub.(Node)
							triples <- Triple{S: retSubNode, P: p, O: retObj}
						}
					}
				}
			} else {
				// None None obj
				if _, ok := o.(None); !ok {
					retSubjects := (*g).osp[o]
					for retSub := range retSubjects {
						retPredicates := (*g).osp[o][retSub]
						for retPred := range retPredicates.Range() {
							retPredNode, _ := retPred.(Node)
							triples <- Triple{S: retSub, P: retPredNode, O: o}
						}
					}
				} else {
					// None None None
					retSubjects := (*g).spo
					for retSub := range retSubjects {
						retPredicates := (*g).spo[retSub]
						for retPred := range retPredicates {
							retObjects := (*g).spo[retSub][retPred]
							for retObj := range retObjects.Range() {
								retObjNode, _ := retObj.(Node)
								triples <- Triple{S: retSub, P: retPred, O: retObjNode}
							}
						}
					}
				}
			}
		}

		close(triples)
	}()

	return triples
}