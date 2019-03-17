import React, { Component } from "react";
import axios from "axios";
import { notification } from "antd";

export default class CreateTournament extends Component {
  constructor(props) {
    super(props);

    this.state = {
      name: "",
      description: ""
    };

    this.onChange = this.onChange.bind(this);
    this.onSubmit = this.onSubmit.bind(this);
  }

  onChange(e) {
    e.preventDefault();

    this.setState({
      [e.target.name]: e.target.value
    });
  }

  onSubmit(e) {
    e.preventDefault();

    const { name, description } = this.state;

    axios
      .post("/tournaments", {
        name,
        description
      })
      .then(() => {
        notification.open({
          message: "Torneio criado com sucesso!",
          type: "success"
        });

        this.props.callback();
      });
  }

  render() {
    return (
      <div>
        <h2>Novo Torneio</h2>

        <form onSubmit={this.onSubmit}>
          <label>Nome</label>
          <input onChange={this.onChange} required name="name" />

          <label>Descrição</label>
          <input onChange={this.onChange} required name="description" />

          <button disabled={this.props.isLoadig} type="submit">
            Enviar
          </button>
        </form>
      </div>
    );
  }
}
