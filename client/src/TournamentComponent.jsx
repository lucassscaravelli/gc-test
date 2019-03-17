import React from "react";
import axios from "axios";

import { Collapse, Spin, Card, Button, Popover, notification } from "antd";
import Matches from "./Matches";

const Panel = Collapse.Panel;

export default class Tournaments extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      selectedID: -1,
      isLoading: false,
      groupStage: null,
      playoffStage: null
    };

    this.onChange = this.onChange.bind(this);
    this.runGroupStage = this.runGroupStage.bind(this);
    this.runPlayoff = this.runPlayoff.bind(this);
  }

  onChange(selectedID) {
    if (!selectedID) {
      return;
    }

    const id = selectedID[0];
    this.getTournamentDetails(id);
  }

  getTournamentDetails(id) {
    this.setState({
      selectedID: id,
      isLoading: true
    });

    axios.get(`/tournaments/${id}/group_stage/table`).then(({ data }) => {
      const groupStage = data;

      this.getPlayoffStage(groupStage);
    });
  }

  getPlayoffStage(groupStage) {
    axios
      .get(`/tournaments/${this.state.selectedID}/playoff_stage/table`)
      .then(({ data }) => {
        this.setState({
          isLoading: false,
          groupStage,
          playoffStage: data
        });
      });
  }

  renderPlayoffs() {
    const { selectedID, playoffStage } = this.state;

    console.log("playoffStage", playoffStage);
    return (
      <div>
        <h4>Playoffs</h4>

        <div className="actions">
          <Button onClick={() => this.runPlayoff(selectedID)}>
            Avançar playoff
          </Button>
        </div>

        {/* <p>{JSON.stringify(playoffStage)}</p> */}

        <div className="playoff-table">
          <div className="bracket">
            <h4>Primeira Fase</h4>
            {playoffStage.FirstPhase && (
              <Matches matches={playoffStage.FirstPhase} />
            )}
          </div>

          <div className="bracket">
            <h4>Oitavas</h4>
            {playoffStage.Octaves && <Matches matches={playoffStage.Octaves} />}
          </div>

          <div className="bracket">
            <h4>Quartas</h4>
            {playoffStage.Quarter && <Matches matches={playoffStage.Quarter} />}
          </div>

          <div className="bracket">
            <h4>Semi final</h4>
            {playoffStage.Semi && <Matches matches={playoffStage.Semi} />}
          </div>

          <div className="bracket">
            <h4>Final</h4>
            {playoffStage.Final && <Matches matches={playoffStage.Final} />}
          </div>
        </div>
      </div>
    );
  }

  runGroupStage(id) {
    this.setState({
      selectedID: id,
      isLoading: true
    });

    axios
      .post(`/tournaments/${id}/group_stage/start`)
      .then(() => {
        this.getTournamentDetails(id);
      })
      .catch(err => {
        this.setState({ isLoading: false });

        notification.open({
          message: "Ocorreu um erro",
          description: err.response.data.msg,
          type: "error"
        });
      });
  }

  runPlayoff(id) {
    this.setState({
      selectedID: id,
      isLoading: true
    });

    axios
      .post(`/tournaments/${id}/playoff_stage/run_next_phase`)
      .then(() => {
        this.getTournamentDetails(id);
      })
      .catch(err => {
        this.setState({ isLoading: false });

        notification.open({
          message: "Ocorreu um erro",
          description: err.response.data.msg,
          type: "error"
        });
      });
  }

  renderGroups() {
    const { selectedID, groupStage } = this.state;

    return (
      <div>
        <h4>Fase de grupos</h4>

        <div className="actions">
          <Button onClick={() => this.runGroupStage(selectedID)}>
            Simular fase de grupos
          </Button>
        </div>

        <div className="group-list">
          {groupStage.map(group => {
            return (
              <div className="group-container" key={group.GroupName}>
                <Card title={group.GroupName}>
                  <table style={{ width: "100%" }}>
                    <tr>
                      <th>COR</th>
                      <th style={{ width: "60%" }}>Nome</th>
                      <th>TAG</th>
                      <th>SR</th>
                      <th>P</th>
                    </tr>
                    {group.Table.map(teamLine => {
                      return (
                        <tr key={teamLine.TeamID}>
                          <td
                            style={{
                              background: teamLine.TeamColor
                            }}
                          />
                          <td>{teamLine.TeamName}</td>
                          <td>{teamLine.TeamTag}</td>

                          <td>{teamLine.TeamRoundBalance}</td>
                          <td>{teamLine.TeamPoints}</td>
                        </tr>
                      );
                    })}
                  </table>

                  {group.Matches && (
                    <Popover
                      trigger="click"
                      content={
                        <table>
                          {group.Matches.map(
                            ({
                              HostTag,
                              HostScore,
                              HostColor,
                              VisitorTag,
                              VisitorScore,
                              VisitorColor
                            }) => {
                              return (
                                <tr key={HostTag + VisitorTag}>
                                  <td
                                    style={{
                                      background: HostColor,
                                      width: "5%"
                                    }}
                                  />
                                  <td>{HostTag}</td>
                                  <td>{HostScore}</td>
                                  <td>X</td>
                                  <td>{VisitorScore}</td>
                                  <td>{VisitorTag}</td>
                                  <td
                                    style={{
                                      background: VisitorColor,
                                      width: "5%"
                                    }}
                                  />
                                </tr>
                              );
                            }
                          )}
                        </table>
                      }
                      title="Partidas"
                    >
                      <Button>Exibir partidas</Button>
                    </Popover>
                  )}
                </Card>
              </div>
            );
          })}
        </div>
      </div>
    );
  }

  renderPanelBody(tournament) {
    const { groupStage, playoffStage, isLoading } = this.state;

    if (isLoading) {
      return <Spin />;
    }

    return (
      <div className="tournament-detail">
        {groupStage ? this.renderGroups() : null}
        {playoffStage ? this.renderPlayoffs() : null}
      </div>
    );
  }

  render() {
    const { selectedID } = this.state;
    const { tournaments } = this.props;

    return (
      <Collapse accordion onChange={this.onChange}>
        {tournaments.map(tournament => {
          const header = (
            <div className="tournament-item">
              <h2>{tournament.name}</h2>
              <h3>{tournament.description}</h3>
            </div>
          );

          return (
            <Panel key={tournament.ID} header={header}>
              {selectedID == tournament.ID
                ? this.renderPanelBody(tournament)
                : null}
            </Panel>
          );
        })}
      </Collapse>
    );
  }
}
