// main is a package for a command-line tool to demonstrate "random" selection
// algorithms. There are many ways to randomly select values in a set, but the
// original playground involves a solution to a "weighted" random selection.
//
// Problem statement:
//
// Given an input of a list of pairs, where the first element of each pair is
// a "label" and the second element an associated "weight", write a function
// which chooses and returns a single label in a correctly random fashion where
// the probability of returning any of the particular labels is distributed
// according to their weights. For example, given input:
//
// [("a", 1), ("b", 2), ("c", 3)]
//
// A single invocation of the function would return a label a, b or c with the
// following probabilities:
//
// a: 1/3 (16.6%)
// b: 2/3 (33.3%)
// c: 3/3 (50.0%)
//
package main

import (
	"fmt"
	"math/rand"
)

type WeightedSelection struct {
	Label  string
	Weight int
}

type Selection struct {
	Label string
}

type PartialSum struct {
	PrevTotal int
	Total     int
	Label     string
}

const (
	WeightedNaive    = "weighted-naive"
	WeightedLinear   = "weighted-linear"
	WeightedConstant = "weighted-constant"
)

var weightedInput = []*WeightedSelection{
	{
		Label:  "a",
		Weight: 1,
	},
	{
		Label:  "b",
		Weight: 2,
	},
	{
		Label:  "c",
		Weight: 3,
	},
}

// this should be built once for any given input
var naiveWeightedIdx []string

func main() {
	// naive solution to weighted random selection
	fmt.Printf("Naive weighted selection: %s\n", doSelection(WeightedNaive))
	fmt.Printf("Linear (time/space) weighted selection: %s\n", doSelection(WeightedLinear))
	fmt.Printf("Constant (space), linear (time) weighted selection: %s\n", doSelection(WeightedConstant))
}

func doSelection(kind string) {
	switch kind {
	case WeightedNaive:
		return selectWeightedNaive(weightedInput)
	case WeightedLinear:
		return selectWeightedLinear(weightedInput)
	case WeightedConstant:
		return selectWeightedConstant(weightedInput)
	default:
		return "Unknown kind"
	}
}

// O(n+m) time and space where n is len(input) and m is sum of all
// WeightedInput.Weight values.
func selectWeightedNaive(input []*WeightedSelection) string {
	// build the index once, small optimization that will yield O(m) for
	// subsequent calls
	if naiveWeightedIdx == nil {
		naiveWeightedIdx := []string{}
		sum := 0
		for _, s := range input {
			for i := 0; i < input.Weight; i++ {
				naiveWeightedIdx = append(naiveWeightedIdx, input.Label)
			}
		}
	}
	fmt.Printf(naiveWeightedIdx)
	return naiveWeightedIdx[rand.Intn(len(naiveWeightedIdx))]
}

// O(n) time and space where n is len(input).
func selectWeightedLinear(input []*WeightedSelection) string {
	total := 0
	sums := make([]PartialSum, len(input))

	for idx, s := range input {
		newTotal := total + input.Weight
		sums[idx] = PartialSum{
			PrevTotal: total,
			Total:     newTotal,
			Label:     input.Label,
		}
		total = newTotal
	}

	rnd = rand.Intn(total)
	for _, s := range sums {
		if rnd >= s.before && rnd < s.total {
			return s.label
		}
	}
	return "??"
}

// O(1) space, O(n) time where n is len(input).
func selectWeightedConstant(input []*WeightedSelection) string {
	// TODO
	return "not implemented"
}
