package types

import (
	"fmt"

	"github.com/alice/checkers/x/checkers/rules"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (storedGame StoredGame) GetBlackAddress() (black sdk.AccAddress, err error) {
	black, errBlack := sdk.AccAddressFromBech32(storedGame.Black)
	return black, errorsmod.Wrap(errBlack, ErrInvalidBlack.Error())
}

func (storedGame StoredGame) GetRedAddress() (red sdk.AccAddress, err error) {
	red, errRed := sdk.AccAddressFromBech32(storedGame.Red)
	return red, errorsmod.Wrap(errRed, ErrInvalidRed.Error())
}

func (storedGame StoredGame) ParseGame() (game *rules.Game, err error) {
	board, errBoard := rules.Parse(storedGame.Board)
	if errBoard != nil {
		return nil, errorsmod.Wrap(errBoard, ErrGameNotParseable.Error())
	}
	board.Turn = rules.StringPieces[storedGame.Turn].Player
	if board.Turn.Color == "" {
		return nil, errorsmod.Wrap(fmt.Errorf("turn: %s", storedGame.Turn), ErrGameNotParseable.Error())
	}
	return board, nil
}

func (storedGame StoredGame) Validate() (err error) {
	_, err = storedGame.GetBlackAddress()
	if err != nil {
		return err
	}
	_, err = storedGame.GetRedAddress()
	if err != nil {
		return err
	}
	_, err = storedGame.ParseGame()
	return err
}
