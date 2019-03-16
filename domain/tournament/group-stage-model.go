package tournament

import (
	"sort"

	"github.com/jinzhu/gorm"
)

type GroupStage struct {
	// stage.Stage
	gorm.Model

	Tournament      Tournament `gorm:"foreignkey:TournamentRefer"`
	TournamentRefer uint
	// Groups []*teamgroup.TeamGroup `json:"teamGroups"`
}

type groupStageTable struct {
	GroupName string
	Table     []groupLine
}

type groupLine struct {
	TeamID        uint
	TeamColor     string
	TeamName      string
	TeamTag       string
	TeamRoundWins int
	TeamPoints    int
}

// GetID retorna o id do model
func (gs *GroupStage) GetID() uint {
	return gs.ID
}

// Validate retorna um erro caso o schema
// não for válido
func (gs *GroupStage) Validate() error {

	// if gs.Groups == nil || gs.Matches == nil || gs.Teams == nil {
	// 	return errors.BadRequest
	// }

	return nil
}

func (gs *GroupStage) Start() (err error) {

	// simular as partidas
	matchesOfGs, err := NewMatchService().GetAllByGroupStageID(gs.ID)

	if err != nil {
		return err
	}

	for _, match := range matchesOfGs {
		err = match.Play()

		if err != nil {
			return err
		}
	}

	return err
}

func (gs *GroupStage) GetTable() (resultTable []*groupStageTable, err error) {

	// obter os grupos
	groups, err := NewTeamGroupService().GetTeamGroupsByGroupStageID(gs.GetID())

	if err != nil {
		return nil, err
	}

	// montar a estrutura de retornos
	for _, group := range groups {

		groupTeamList := []groupLine{}

		for _, team := range group.Teams {

			groupTeamList = append(groupTeamList, groupLine{
				TeamID:    team.ID,
				TeamColor: team.Color,
				TeamName:  team.Name,
				TeamTag:   team.Tag,
			})

		}

		resultTable = append(resultTable, &groupStageTable{GroupName: group.Group, Table: groupTeamList})

	}

	// obter as partidas
	matchesOfGs, err := NewMatchService().GetAllByGroupStageID(gs.ID)

	if err != nil {
		return nil, err
	}

	resultList := resultTable

	for _, match := range matchesOfGs {

		winnerTeamID, winnerScore := match.GetWinner()
		loserTeamID, loserScore := match.GetLoser()

		for i, groupTable := range resultTable {

			for j, line := range groupTable.Table {

				if line.TeamID == winnerTeamID {

					resultList[i].Table[j].TeamPoints++
					resultList[i].Table[j].TeamRoundWins += winnerScore
				}

				if line.TeamID == loserTeamID {
					resultList[i].Table[j].TeamRoundWins += loserScore
				}

			}

		}
	}

	for z := 0; z < len(resultList); z++ {

		sort.Slice(resultList[z].Table, func(i int, j int) bool {

			if resultList[z].Table[i].TeamPoints > resultList[z].Table[j].TeamPoints {
				return true
			} else if resultList[z].Table[i].TeamPoints < resultList[z].Table[j].TeamPoints {
				return false
			}

			if resultList[z].Table[i].TeamRoundWins > resultList[z].Table[j].TeamPoints {
				return true
			}

			return false

		})

	}

	return resultList, nil
}
