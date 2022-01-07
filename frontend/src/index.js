import React, { useState, useEffect } from 'react';
import { render } from 'react-dom';
import CPUUsage from './components/CPUUsage'
import DepartmentSelect from './components/DepartmentSelect'

import 'bootstrap/dist/css/bootstrap.css';
import './index.css';

//global.jQuery = require('jquery');
import 'bootstrap/dist/js/bootstrap.bundle.min.js';
require('bootstrap');

const App = () => {

  const [greeting, setGreeting] = useState('');
  const [sys, setSys] = useState('');
  //const [cpuPercent, setCPUPercent] = useState(0);

  window.go.main.App.Greet().then((result) => {
    setGreeting(result);
  });

  useEffect(() => {
    window.go.main.App.Sysvars().then((result) => {
      var resultJSON = JSON.parse(result);
      setSys(resultJSON);
      //setCPUPercent(resultJSON.CPU.avg);
    });
  });

  // <CPUUsage percent={cpuPercent} />

  return (
    <>
      <div className="container">

        <div className="row">
          <div className="col">

            <div className="mt-3 mb-3">{greeting}</div>

            <DepartmentSelect />

          </div>
        </div>

      </div>

    </>
  );
};

render(<App />, document.querySelector('#root'));
