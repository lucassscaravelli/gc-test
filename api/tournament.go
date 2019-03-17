package api

import (
	"encoding/json"
	"gctest/api/helper"
	"gctest/domain/tournament"
	"gctest/errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (s *Server) getAllTournaments(w http.ResponseWriter, r *http.Request) {

	var tournaments []*tournament.Tournament
	var err error

	tournaments, err = tournament.NewTournamentService().GetAllTournaments()

	if err != nil {
		helper.HandleError("/tournaments/ [GET]", err, w)
		return
	}

	json.NewEncoder(w).Encode(tournaments)
}

func (s *Server) createTournament(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var newTournament tournament.Tournament
	decodeError := decoder.Decode(&newTournament)

	if decodeError != nil {
		helper.HandleError("/tournaments/ [POST]", errors.BadRequest, w)
		return
	}

	newTournamentResult, bdError := tournament.NewTournamentService().CreateTournament(&newTournament)

	if bdError != nil {
		helper.HandleError("/tournaments/ [POST]", bdError, w)
		return
	}

	json.NewEncoder(w).Encode(newTournamentResult)
}

func (s *Server) getTournament(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, parseError := strconv.ParseUint(vars["id"], 10, 32)

	if parseError != nil {
		helper.HandleError("/tournaments/<id> [GET]", errors.BadRequest, w)
		return
	}

	tournament, err := tournament.NewTournamentService().GetTournamentByID(uint(id))

	if err != nil {
		helper.HandleError("/tournaments/<id> [GET]", err, w)
		return
	}

	json.NewEncoder(w).Encode(tournament)
}

func (s *Server) startGroupStage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, parseError := strconv.ParseUint(vars["id"], 10, 32)

	if parseError != nil {
		helper.HandleError("/tournaments/<id>/group_stage/start [POST]", errors.BadRequest, w)
		return
	}

	t, errFindTournament := tournament.NewTournamentService().GetTournamentByID(uint(id))

	if errFindTournament != nil {
		helper.HandleError("/tournaments/<id>/group_stage/start [POST]", errFindTournament, w)
		return
	}

	groupStage, errGetGroupStage := tournament.NewGroupStageService().GetGroupStageByTournamentID(t.GetID())

	if errGetGroupStage != nil {
		helper.HandleError("/tournaments/<id>/group_stage/start [POST]", errGetGroupStage, w)
		return
	}

	if err := tournament.NewGroupStageService().RunGroupStage(groupStage); err != nil {
		helper.HandleError("RunGroupStage", err, w)
		return
	}

	json.NewEncoder(w).Encode(groupStage)
}

func (s *Server) getGroupStageTable(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, parseError := strconv.ParseUint(vars["id"], 10, 32)

	if parseError != nil {
		helper.HandleError("/tournaments/<id>/group_stage/table [GET]", errors.BadRequest, w)
		return
	}

	t, errFindTournament := tournament.NewTournamentService().GetTournamentByID(uint(id))

	if errFindTournament != nil {
		helper.HandleError("/tournaments/<id>/group_stage/table [GET]", errFindTournament, w)
		return
	}

	groupStage, errGetGroupStage := tournament.NewGroupStageService().GetGroupStageByTournamentID(t.GetID())

	if errGetGroupStage != nil {
		helper.HandleError("/tournaments/<id>/group_stage/table [GET]", errGetGroupStage, w)
		return
	}

	table, errGetTable := tournament.NewGroupStageService().GetTableResults(groupStage)

	if errGetTable != nil {
		helper.HandleError("/tournaments/<id>/group_stage/table [GET]", errGetTable, w)
		return
	}

	json.NewEncoder(w).Encode(table)
}

func (s *Server) getPlayoffTable(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, parseError := strconv.ParseUint(vars["id"], 10, 32)

	if parseError != nil {
		helper.HandleError("/tournaments/<id>/group_stage/table [GET]", errors.BadRequest, w)
		return
	}

	t, errFindTournament := tournament.NewTournamentService().GetTournamentByID(uint(id))

	if errFindTournament != nil {
		helper.HandleError("/tournaments/<id>/group_stage/table [GET]", errFindTournament, w)
		return
	}

	playoffStage, errGetGroupStage := tournament.NewPlayoffStageService().GetPlayoffStageByTournamentID(t.GetID())

	if errGetGroupStage != nil {
		helper.HandleError("get playoff stage by Tournament ID", errGetGroupStage, w)
		return
	}

	playOffTable, err := tournament.NewPlayoffStageService().GetTable(playoffStage)

	if err != nil {
		helper.HandleError("get playoff table", err, w)
	}

	json.NewEncoder(w).Encode(playOffTable)
}

func (s *Server) runPlayoffNextPhase(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, parseError := strconv.ParseUint(vars["id"], 10, 32)

	if parseError != nil {
		helper.HandleError("/tournaments/<id>/group_stage/table [GET]", errors.BadRequest, w)
		return
	}

	t, errFindTournament := tournament.NewTournamentService().GetTournamentByID(uint(id))

	if errFindTournament != nil {
		helper.HandleError("/tournaments/<id>/group_stage/table [GET]", errFindTournament, w)
		return
	}

	playoffStage, errGetGroupStage := tournament.NewPlayoffStageService().GetPlayoffStageByTournamentID(t.GetID())

	if errGetGroupStage != nil {
		helper.HandleError("get playoff by tournament", errGetGroupStage, w)
		return
	}

	err := tournament.NewPlayoffStageService().RunNextPhase(playoffStage)

	if err != nil {
		helper.HandleError("get playoff table", err, w)
	}

	json.NewEncoder(w).Encode(playoffStage)
}
