import React, { Component } from 'react';

import { Button } from 'react-bootstrap'
import Select from 'react-select';

class DepartmentSelect extends Component {

    state = {
      departments: '',
      selectedDepartment: '',
      installResult: '',
    }

    constructor(props) {
      super(props);

      this.handleChange = this.handleChange.bind(this);
      this.handleClick = this.handleClick.bind(this);

      var self = this;

      window.go.main.App.GetDepartments().then((result) => {
        var departments = JSON.parse(result);
        self.setState({ departments });
      });
    }
  
    handleChange(selectedDepartment) {
      this.setState({ selectedDepartment });
    }
  
    handleClick() {
      //alert('Your department is: ' + this.state.selectedDepartment.label);
      //event.preventDefault();
      console.log('hi');
      var self = this;
      window.go.main.App.Install(this.state.selectedDepartment.value).then((installResult) => {
        self.setState({ installResult });
      });
    }
  
    render() {
      return (
          <div>
            <label className="form-label">
              Select department to install relevant software:
              <Select options={ this.state.departments } value={this.selectedDepartment} onChange={this.handleChange} />
            </label>
            <Button variant="primary" onClick={this.handleClick}>Install</Button>
          </div>
      );
    }
}

export default DepartmentSelect;