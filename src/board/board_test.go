package board
import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestHandleFlagClaimCommand(t *testing.T) {
  testBoard := Board{}
  commandArray := []string{"unclaimed","north","south","unclaimed","north","south","unclaimed","north","south"}
	testBoard.HandleFlagClaimCommand(commandArray)
	assert.Equal(t, testBoard.Flags[0].Claimer, "unclaimed")
	assert.Equal(t, testBoard.Flags[1].Claimer, "north")
	assert.Equal(t, testBoard.Flags[2].Claimer, "south")
	assert.Equal(t, testBoard.Flags[3].Claimer, "unclaimed")
	assert.Equal(t, testBoard.Flags[4].Claimer, "north")
	assert.Equal(t, testBoard.Flags[5].Claimer, "south")
	assert.Equal(t, testBoard.Flags[6].Claimer, "unclaimed")
	assert.Equal(t, testBoard.Flags[7].Claimer, "north")
	assert.Equal(t, testBoard.Flags[8].Claimer, "south")
}
