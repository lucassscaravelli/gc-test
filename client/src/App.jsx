import React from "react";
import ReactDOM from "react-dom";
import axios from "axios";

import { Layout, Spin } from "antd";

import Tournaments from "./TournamentComponent";

import "./index.less";
import CreateTournament from "./CreateTournament";

const { Header, Content, Footer } = Layout;

class App extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      errorMsg: "",
      isLoading: true,
      tournaments: []
    };

    this.loadTournaments = this.loadTournaments.bind(this);
  }

  componentDidMount() {
    this.loadTournaments();
  }

  loadTournaments() {
    this.setState({ isLoading: true });
    axios
      .get("/tournaments")
      .then(({ data }) => {
        this.setState({ isLoading: false, tournaments: data });
      })
      .catch(e => console.log("e", e));
  }

  render() {
    const { isLoading, tournaments } = this.state;

    return (
      <Layout>
        <Header>
          <h1>Teste - Gamers Club</h1>
        </Header>
        <Content>
          <div className="container form">
            <CreateTournament
              isLodaing={isLoading}
              callback={this.loadTournaments}
            />
          </div>
          <div className="container">
            <h2>Torneios</h2>
            {isLoading ? <Spin /> : <Tournaments tournaments={tournaments} />}
          </div>
        </Content>
        <Footer>Footer</Footer>
      </Layout>
    );
  }
}

ReactDOM.render(<App />, document.getElementById("app"));
