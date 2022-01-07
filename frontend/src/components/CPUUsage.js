import React from 'react';

class CPUUsage extends React.Component {

    constructor(props) {
        super(props);
    }

    render() {
        return (
            <>
            <h1>CPU is {this.props.percent}%</h1>
            </>
        );
    }
}

export default CPUUsage;