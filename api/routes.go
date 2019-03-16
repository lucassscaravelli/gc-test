package api

func (s *Server) routes() {
	s.mux.HandleFunc("/tournaments", s.getAllTournaments).Methods("GET")
	s.mux.HandleFunc("/tournaments", s.createTournament).Methods("POST")

	s.mux.HandleFunc("/tournaments/{id}", s.getTournament).Methods("GET")
	s.mux.HandleFunc("/tournaments/{id}/group_stage/start", s.startGroupStage).Methods("POST")
	s.mux.HandleFunc("/tournaments/{id}/group_stage/table", s.getGroupStageTable).Methods("GET")
}
