package types

import (
	"fmt"
	"sort"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (leaderboard Leaderboard) Validate() error {
	// Check for duplicated player address in winners
	winnerInfoIndexMap := make(map[string]struct{})

	for _, elem := range leaderboard.Winners {
		index := string(PlayerInfoKey(elem.PlayerAddress))
		if _, ok := winnerInfoIndexMap[index]; ok {
			return fmt.Errorf("duplicated playerAddress for winner")
		}
		winnerInfoIndexMap[index] = struct{}{}
	}
	return nil
}

type WinningPlayerParsed struct {
	PlayerAddress string
	WonCount      uint64
	DateAdded     time.Time
}

func ParseDateAddedAsTime(dateAdded string) (dateAddedParsed time.Time, err error) {
	dateAddedParsed, errDateAdded := time.Parse(DateAddedLayout, dateAdded)
	return dateAddedParsed, sdkerrors.Wrapf(errDateAdded, ErrInvalidDateAdded.Error(), dateAdded)
}

func (winningPlayer WinningPlayer) GetDateAddedAsTime() (dateAdded time.Time, err error) {
	return ParseDateAddedAsTime(winningPlayer.DateAdded)
}

func GetDateAdded(ctx sdk.Context) time.Time {
	return ctx.BlockTime()
}

func FormatDateAdded(dateAdded time.Time) string {
	return dateAdded.UTC().Format(DateAddedLayout)
}

func (winningPlayer WinningPlayer) Parse() (parsed WinningPlayerParsed, err error) {
	dateAdded, err := winningPlayer.GetDateAddedAsTime()
	if err != nil {
		return WinningPlayerParsed{}, err
	}
	return WinningPlayerParsed{
		PlayerAddress: winningPlayer.PlayerAddress,
		WonCount:      winningPlayer.WonCount,
		DateAdded:     dateAdded,
	}, nil
}

func (parsed WinningPlayerParsed) Stringify() WinningPlayer {
	return WinningPlayer{
		PlayerAddress: parsed.PlayerAddress,
		WonCount:      parsed.WonCount,
		DateAdded:     FormatDateAdded(parsed.DateAdded),
	}
}

func ParseWinners(winners []WinningPlayer) (parsedWinners []WinningPlayerParsed, err error) {
	parsedWinners = make([]WinningPlayerParsed, len(winners))
	var parsed WinningPlayerParsed
	for index, winningPlayer := range winners {
		parsed, err = winningPlayer.Parse()
		if err != nil {
			return nil, err
		}
		parsedWinners[index] = parsed
	}
	return parsedWinners, nil
}

func (leaderboard Leaderboard) ParseWinners() (winners []WinningPlayerParsed, err error) {
	return ParseWinners(leaderboard.Winners)
}

func StringifyWinners(winners []WinningPlayerParsed) []WinningPlayer {
	stringified := make([]WinningPlayer, len(winners))
	for index, winner := range winners {
		stringified[index] = winner.Stringify()
	}
	return stringified
}

func CreateLeaderboardFromParsedWinners(winners []WinningPlayerParsed) Leaderboard {
	return Leaderboard{
		Winners: StringifyWinners(winners),
	}
}

func SortWinners(winners []WinningPlayerParsed) {
	sort.SliceStable(winners[:], func(i, j int) bool {
		if winners[i].WonCount > winners[j].WonCount {
			return true
		}
		if winners[i].WonCount < winners[j].WonCount {
			return false
		}
		return winners[i].DateAdded.After(winners[j].DateAdded)
	})
}

func UpdatePlayerInfoAtNow(winners []WinningPlayerParsed, now time.Time, candidate PlayerInfo) (updated []WinningPlayerParsed) {
	if candidate.WonCount < 1 {
		return winners
	}
	found := false
	for index, winner := range winners {
		if winner.PlayerAddress == candidate.Index {
			if winner.WonCount < candidate.WonCount {
				winners[index] = WinningPlayerParsed{
					PlayerAddress: candidate.Index,
					WonCount:      candidate.WonCount,
					DateAdded:     now,
				}
			}
			found = true
			break
		}
	}
	if !found {
		updated = append(winners, WinningPlayerParsed{
			PlayerAddress: candidate.Index,
			WonCount:      candidate.WonCount,
			DateAdded:     now,
		})
	} else {
		updated = winners
	}
	SortWinners(updated)
	if LeaderboardWinnerLength < uint64(len(updated)) {
		updated = updated[:LeaderboardWinnerLength]
	}
	return updated
}

func (leaderboard *Leaderboard) UpdatePlayerInfoAtNow(now time.Time, candidate PlayerInfo) error {
	winners, err := leaderboard.ParseWinners()
	if err != nil {
		return err
	}
	updated := UpdatePlayerInfoAtNow(winners, now, candidate)
	leaderboard.Winners = StringifyWinners(updated)
	candidate.Index = "fake"
	return nil
}
