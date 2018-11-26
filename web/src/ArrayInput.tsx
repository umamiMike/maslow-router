import React, { Component } from "react";
import { OptionType } from "./interfaces";
import Select from "react-select";

interface Props {
  id: string;
  value: string[];
  onChange: (selected: any) => void;
  choices?: OptionType[];
}

interface State {
  activeInputs: number;
}

export default class ArrayInput extends Component<Props, State> {
  state: State = {
    activeInputs: 1
  };

  addInput = () => {
    this.setState({ activeInputs: this.state.activeInputs + 1 });
  };

  makeControl = (i: number) => {
    if (this.props.choices == null) {
      return (
        <input
          className="shadow appearance-none border rounded w-full py-2 px-3 text-grey-darker leading-tight focus:outline-none focus:shadow-outline"
          id={`${this.props.id}:${i}`}
          type="string"
          placeholder=".*\.facebook\.com"
          value={this.props.value[i] || ""}
          onChange={this.props.onChange}
        />
      );
    }
    const selectChange = (event: any) => {
      this.props.onChange({
        target: { value: event.value, id: `${this.props.id}:${i}` }
      });
    };

    return (
      <Select
        className="w-full"
        options={this.props.choices}
        onChange={selectChange}
      />
    );
  };

  render() {
    const inputs = Array(this.state.activeInputs)
      .fill(null)
      .map((_: any, i: number) => {
        const disabled = this.props.value[i] ? false : true;
        return (
          <div key={i} className="flex flex-row mb-4">
            {this.makeControl(i)}
            <div className="flex items-center content-center self-center pl-6">
              <button
                disabled={disabled}
                onClick={this.addInput}
                className="no-underline font-bold py-4 px-4 rounded-full"
              >
                <p className="font-icon text-blue hover:text-blue-dark icon-059" />
              </button>
            </div>
          </div>
        );
      });
    return (
      <div className="mb-4">
        <label className="block text-grey-darker text-sm font-bold mb-2">
          Addresses
        </label>
        {inputs}
      </div>
    );
  }
}
