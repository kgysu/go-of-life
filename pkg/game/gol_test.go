package game

import (
	"reflect"
	"testing"
)

func TestPlayOneRound_Rules(t *testing.T) {
	input := [][]int{
		{0, 0, 0, 0, 0},
		{0, 1, 1, 1, 0},
		{0, 0, 1, 1, 0},
		{0, 1, 0, 1, 0},
		{0, 0, 0, 0, 0},
	}
	want := [][]int{
		{0, 0, 1, 0, 0},
		{0, 1, 0, 1, 0},
		{0, 0, 0, 0, 1},
		{0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0},
	}
	got := PlayRound(input)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("PlayOneRound failed, want=%v, got=%v", want, got)
	}
}

func TestGetNeighborCount_0(t *testing.T) {
	state := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	got := getNeighborCount(state, 1, 1)
	want := 0

	if got != want {
		t.Errorf("GetNeighborCount failed, want=%d, got=%d", want, got)
	}
}

func TestGetNeighborCount_1(t *testing.T) {
	state := [][]int{
		{0, 0, 0},
		{1, 1, 0},
		{0, 0, 0},
	}
	got := getNeighborCount(state, 1, 1)
	want := 1

	if got != want {
		t.Errorf("GetNeighborCount failed, want=%d, got=%d", want, got)
	}
}

func TestGetNeighborCount_8(t *testing.T) {
	state := [][]int{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	}
	got := getNeighborCount(state, 1, 1)
	want := 8

	if got != want {
		t.Errorf("GetNeighborCount failed, want=%d, got=%d", want, got)
	}
}
