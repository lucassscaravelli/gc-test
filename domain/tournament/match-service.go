package tournament

import "math/rand"

// MatchService representa o service de um torneio
type MatchService struct {
	rep *MatchRepository
}

// NewService cria um novo serviço de torneios
func NewMatchService() *MatchService {
	service := MatchService{NewMatchRepository()}

	return &service
}

func (s *MatchService) CreateMatch(newMatch *Match) (m *Match, err error) {
	m, err = s.rep.Insert(newMatch)
	return
}

func (s *MatchService) GetAllByBracketID(bracketID uint) (m []*Match, err error) {
	m, err = s.rep.FindAll("team_bracket_refer = ?", bracketID)

	// obter as partidas
	for _, match := range m {

		err = s.rep.IRepository.Preload(match, "HostTeam")

		if err != nil {
			return nil, err
		}

		err = s.rep.IRepository.Preload(match, "VisitorTeam")

		if err != nil {
			return nil, err
		}

	}

	return
}

func (s *MatchService) GetAllByGroupStageID(ID uint) (m []*Match, err error) {
	m, err = s.rep.FindAll("group_stage_refer = ?", ID)

	if err != nil {
		return nil, err
	}

	for _, match := range m {

		err = s.rep.IRepository.Preload(match, "HostTeam")

		if err != nil {
			return nil, err
		}

		err = s.rep.IRepository.Preload(match, "VisitorTeam")

		if err != nil {
			return nil, err
		}

	}

	return
}

func (s *MatchService) UpdateMatch(match *Match) (m *Match, err error) {
	return s.rep.Update(match)
}

func (s *MatchService) RunMatch(match *Match) (m *Match, err error) {
	hostPoints := 0
	visitorPoints := 0

	for i := 0; i < 30 && hostPoints < 16 && visitorPoints < 16; i++ {

		if rand.Intn(2) == 0 {
			hostPoints++
		} else {
			visitorPoints++
		}

	}

	// no caso de 15 a 15 para nao impatar :P
	if hostPoints == visitorPoints {
		hostPoints = 14
		visitorPoints = 16
	}

	match.HostScore = hostPoints
	match.VisitorScore = visitorPoints

	m, err = s.UpdateMatch(match)

	return
}

func (s *MatchService) GetMatchByBracketID(ID uint) (m *Match, err error) {
	m, err = s.rep.FindFirst("team_bracket_refer = ?", ID)

	err = s.rep.IRepository.Preload(m, "HostTeam")

	if err != nil {
		return nil, err
	}

	err = s.rep.IRepository.Preload(m, "VisitorTeam")

	if err != nil {
		return nil, err
	}

	return
}
