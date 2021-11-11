package types

import (
	"time"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type winningPlayerParsed struct {
	PlayerAddress string
	WonCount      uint64
	DateAdded     time.Time
}

func (winningPlayer *WinningPlayer) GetDateAddedAsTime() (dateAdded time.Time, err error) {
	dateAdded, errDateAdded := time.Parse(DateAddedLayout, winningPlayer.DateAdded)
	return dateAdded, sdkerrors.Wrapf(errDateAdded, ErrInvalidDateAdded.Error(), winningPlayer.DateAdded)
}

func GetDateAdded(ctx sdk.Context) time.Time {
	return ctx.BlockTime()
}

func FormatDateAdded(dateAdded time.Time) string {
	return dateAdded.UTC().Format(DateAddedLayout)
}

func (winningPlayer *WinningPlayer) parse() (parsed *winningPlayerParsed, err error) {
	dateAdded, err := winningPlayer.GetDateAddedAsTime()
	if err != nil {
		return nil, err
	}
	return &winningPlayerParsed{
		PlayerAddress: winningPlayer.PlayerAddress,
		WonCount:      winningPlayer.WonCount,
		DateAdded:     dateAdded,
	}, nil
}

func (parsed *winningPlayerParsed) stringify() (stringified *WinningPlayer) {
	return &WinningPlayer{
		PlayerAddress: parsed.PlayerAddress,
		WonCount:      parsed.WonCount,
		DateAdded:     FormatDateAdded(parsed.DateAdded),
	}
}

func (leaderboard *Leaderboard) parseWinners() (winners []*winningPlayerParsed, err error) {
	winners = make([]*winningPlayerParsed, len(leaderboard.Winners))
	var parsed *winningPlayerParsed
	for index, winningPlayer := range leaderboard.Winners {
		parsed, err = winningPlayer.parse()
		if err != nil {
			return nil, err
		}
		winners[index] = parsed
	}
	return winners, nil
}

func stringifyWinners(winners []*winningPlayerParsed) (stringified []*WinningPlayer) {
	stringified = make([]*WinningPlayer, len(winners))
	for index, winner := range winners {
		stringified[index] = winner.stringify()
	}
	return stringified
}
