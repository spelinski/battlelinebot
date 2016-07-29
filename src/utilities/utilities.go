package utilities

import (
    "card"
)

func CardCombinations(cardList []card.Card, numberOfCards int) ([][]card.Card) {
    combinationsOfCards := [][]card.Card{}
    if numberOfCards > 1 {
        combinationsOfCards = combinations(cardList, numberOfCards)
    } else {
        for index:=0; index < len(cardList); index++ {
            combinationsOfCards = append(combinationsOfCards,[]card.Card{cardList[index]})
        }
    }

    return combinationsOfCards
}

//Pulled from https://play.golang.org/p/JEgfXR2zSH
func combinations(iterable []card.Card, r int) ([][]card.Card) {

    outerCombination := [][]card.Card{}
    pool := iterable
    n := len(pool)

    if r > n {
        return outerCombination
    }

    indices := make([]int, r)
    for i := range indices {
        indices[i] = i
    }

    result := make([]card.Card, r)
    for i, el := range indices {
        result[i] = pool[el]
    }
    outerCombination = copySliceToEndOfOtherSliceOfSlices(outerCombination, result)
    for {
        i := r - 1
        for ; i >= 0 && indices[i] == i+n-r; i -= 1 {
        }

        if i < 0 {
            return outerCombination
        }

        indices[i] += 1
        for j := i + 1; j < r; j += 1 {
            indices[j] = indices[j-1] + 1
        }

        for ; i < len(indices); i += 1 {
            result[i] = pool[indices[i]]
        }
        outerCombination = copySliceToEndOfOtherSliceOfSlices(outerCombination, result)
    }
}

func copySliceToEndOfOtherSliceOfSlices( destination [][]card.Card, stuffToCopy []card.Card) ([][]card.Card) {
    tempSlice := make([]card.Card, len(stuffToCopy))
    for index := range stuffToCopy {
        a := stuffToCopy[index]
        tempSlice[index] = a
    }
    destination = append(destination, tempSlice)

    return destination
}