package server

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"strconv"
	"strings"

	"github.com/josephburnett/dsq-golang/pkg/engine"
	"github.com/josephburnett/dsq-golang/pkg/types"
)

func parseMove(move string) ([2]types.Point, error) {
	coord := strings.Split(move, "")
	if len(coord) != 4 {
		return [2]types.Point{}, fmt.Errorf("Invalid move: %v", move)
	}
	fromX, err := strconv.Atoi(coord[0])
	if err != nil {
		return [2]types.Point{}, fmt.Errorf("Invalid move: %v", move)
	}
	fromY, err := strconv.Atoi(coord[1])
	if err != nil {
		return [2]types.Point{}, fmt.Errorf("Invalid move: %v", move)
	}
	toX, err := strconv.Atoi(coord[2])
	if err != nil {
		return [2]types.Point{}, fmt.Errorf("Invalid move: %v", move)
	}
	toY, err := strconv.Atoi(coord[3])
	if err != nil {
		return [2]types.Point{}, fmt.Errorf("Invalid move: %v", move)
	}
	return [2]types.Point{
		types.Point{fromX, fromY},
		types.Point{toX, toY},
	}, nil
}

type searchRequest struct {
	Board *types.Board
	Side  types.Side
	Depth int
}

func (sr *searchRequest) marshal() []byte {
	var blob bytes.Buffer
	enc := gob.NewEncoder(&blob)
	_ = enc.Encode(sr)
	return blob.Bytes()
}

func unmarshalRequest(blob []byte) (*searchRequest, error) {
	sr := &searchRequest{}
	dec := gob.NewDecoder(bytes.NewBuffer(blob))
	err := dec.Decode(sr)
	if err != nil {
		return nil, err
	}
	if sr.Board == nil {
		return nil, fmt.Errorf("Nil board.")
	}
	if sr.Depth < 2 {
		return nil, fmt.Errorf("Invalid depth: %v.", sr.Depth)
	}
	if sr.Side != types.A && sr.Side != types.B {
		return nil, fmt.Errorf("Invalid side: %v.", sr.Side)
	}
	return sr, nil
}

type searchResponse struct {
	BestMove [2]types.Point
	Stat     *engine.Stat
	Ok       bool
	Err      error
}

func (sr *searchResponse) marshal() []byte {
	var blob bytes.Buffer
	enc := gob.NewEncoder(&blob)
	_ = enc.Encode(sr)
	return blob.Bytes()
}

func unmarshalResponse(blob []byte) (*searchResponse, error) {
	sr := &searchResponse{}
	dec := gob.NewDecoder(bytes.NewBuffer(blob))
	err := dec.Decode(sr)
	if err != nil {
		return nil, err
	}
	return sr, nil
}
