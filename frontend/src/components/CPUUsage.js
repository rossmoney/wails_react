import React from 'react';

class CPUUsage extends React.Component {

    constructor(props) {
        super(props);
    }

    render() {
        return (
            <>
            <h2>CPU is {this.props.percent}%</h2>
            </>
        );
    }
}

export default CPUUsage;