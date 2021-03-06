package utilities
import (
    "github.com/stretchr/testify/assert"
    "testing"
    "card"
)

func TestGettingCombinationOfCardsByOne(t *testing.T){
    threeCardList := []card.Card{card.Card{"color1",1,0,0}, card.Card{"color1",2,0,0}, card.Card{"color1",3,0,0}}
    actualCombos := CardCombinations (threeCardList, 1)
    expectedCombos := [][]card.Card{}
    expectedCombos = append(expectedCombos,[]card.Card{card.Card{"color1",1,0,0}})
    expectedCombos = append(expectedCombos,[]card.Card{card.Card{"color1",2,0,0}})
    expectedCombos = append(expectedCombos,[]card.Card{card.Card{"color1",3,0,0}})
    assert.Equal(t, expectedCombos, actualCombos)

}

func TestGettingCombinationOfCardsByTwo(t *testing.T){
    threeCardList := []card.Card{card.Card{"color1",1,0,0}, card.Card{"color1",2,0,0}, card.Card{"color1",3,0,0}}
    actualCombos := CardCombinations (threeCardList, 2)
    expectedCombos := [][]card.Card{}
    expectedCombos = append(expectedCombos,[]card.Card{card.Card{"color1",1,0,0},card.Card{"color1",2,0,0}})
    expectedCombos = append(expectedCombos,[]card.Card{card.Card{"color1",1,0,0},card.Card{"color1",3,0,0}})
    expectedCombos = append(expectedCombos,[]card.Card{card.Card{"color1",2,0,0},card.Card{"color1",3,0,0}})
    assert.Equal(t, expectedCombos, actualCombos)

}

func TestGettingCombinationOfCardsByThree(t *testing.T){
    threeCardList := []card.Card{card.Card{"color1",1,0,0}, card.Card{"color1",2,0,0}, card.Card{"color1",3,0,0}}
    actualCombos := CardCombinations (threeCardList, 3)
    expectedCombos := [][]card.Card{}
    expectedCombos = append(expectedCombos,[]card.Card{card.Card{"color1",1,0,0},card.Card{"color1",2,0,0},card.Card{"color1",3,0,0}})
    assert.Equal(t, expectedCombos, actualCombos)
}

func TestCopySliceToSlice(t *testing.T){
    myHand := []card.Card{card.Card{"color1",1,0,0}, card.Card{"color1",2,0,0}, card.Card{"color1",3,0,0}}
    tempHand := []card.Card{}
    tempHand = CopySliceToSlice(tempHand, myHand)
    assert.Equal(t, myHand, tempHand)
}

func TestFindElementInSlice(t *testing.T) {
    myHand := []card.Card{card.Card{"color1",1,0,0}, card.Card{"color1",2,0,0}, card.Card{"color1",3,0,0}}
    index := FindElementInSlice(myHand, card.Card{"color1",2,0,0})
    assert.Equal(t, 1, index)
}