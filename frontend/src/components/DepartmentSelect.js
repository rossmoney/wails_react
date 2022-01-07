import React from 'react';
import Form from 'react-bootstrap/Form'

class DepartmentSelect extends React.Component {
    constructor(props) {
      super(props);
      this.state = {value: 'dev'};
  
      this.handleChange = this.handleChange.bind(this);
      this.handleSubmit = this.handleSubmit.bind(this);
    }
  
    handleChange(event) {
      this.setState({value: event.target.value});
    }
  
    handleSubmit(event) {
      alert('Your department is: ' + this.state.value);
      event.preventDefault();
    }
  
    render() {
      return (
        <form onSubmit={this.handleSubmit}>
          <label>
            Select department to install relevant software:
            <Form.Select aria-label="Default select example" value={this.state.value} onChange={this.handleChange}>
              <option></option>
              <option value="dev">Dev</option>
              <option value="devops">DevOps</option>
              <option value="marketing">Marketing</option>
              <option value="strategy">Strategy</option>
            </Form.Select>
          </label>
          <input className="btn btn-primary" type="submit" value="Submit" />
        </form>
      );
    }
}

export default DepartmentSelect;