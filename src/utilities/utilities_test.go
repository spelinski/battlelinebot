package utilities
import (
    "github.com/stretchr/testify/assert"
    "testing"
    "card"
)

func TestGettingCombinationOfCardsByOne(t *testing.T){
    threeCardList := []card.Card{card.Card{"color1",1}, card.Card{"color1",2}, card.Card{"color1",3}}
    actualCombos := CardCombinations (threeCardList, 1)
    expectedCombos := [][]card.Card{}
    expectedCombos = append(expectedCombos,[]card.Card{card.Card{"color1",1}})
    expectedCombos = append(expectedCombos,[]card.Card{card.Card{"color1",2}})
    expectedCombos = append(expectedCombos,[]card.Card{card.Card{"color1",3}})
    assert.Equal(t, expectedCombos, actualCombos)

}

func TestGettingCombinationOfCardsByTwo(t *testing.T){
    threeCardList := []card.Card{card.Card{"color1",1}, card.Card{"color1",2}, card.Card{"color1",3}}
    actualCombos := CardCombinations (threeCardList, 2)
    expectedCombos := [][]card.Card{}
    expectedCombos = append(expectedCombos,[]card.Card{card.Card{"color1",1},card.Card{"color1",2}})
    expectedCombos = append(expectedCombos,[]card.Card{card.Card{"color1",1},card.Card{"color1",3}})
    expectedCombos = append(expectedCombos,[]card.Card{card.Card{"color1",2},card.Card{"color1",3}})
    assert.Equal(t, expectedCombos, actualCombos)

}

func TestGettingCombinationOfCardsByThree(t *testing.T){
    threeCardList := []card.Card{card.Card{"color1",1}, card.Card{"color1",2}, card.Card{"color1",3}}
    actualCombos := CardCombinations (threeCardList, 3)
    expectedCombos := [][]card.Card{}
    expectedCombos = append(expectedCombos,[]card.Card{card.Card{"color1",1},card.Card{"color1",2},card.Card{"color1",3}})
    assert.Equal(t, expectedCombos, actualCombos)

}