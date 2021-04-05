import React from 'react';
import moment from 'moment';
import Countdown from "react-countdown";

const Completionist = () => <span style={{fontFamily:"Poppins"}}>Click to add courses:</span>;

const formatNumber = (number) => {
  return number.toLocaleString('en-US', {
    minimumIntegerDigits: 2,
    useGrouping: false
  })
}

const renderer = ({ days, hours, minutes, seconds, completed }) => {
    if (completed) {
      // Render a complete state
      return <Completionist />;
    } else {
      // Render a countdown
      return (
        <span style={{fontFamily:"Poppins", fontSize:"14px"}}>
          {days} {(days == 1) ? 'day' : 'days'}, {formatNumber(hours)} {(hours == 1) ? 'hour' : 'hours'}, {formatNumber(minutes)} {(minutes == 1) ? 'minute' : 'minutes'}, {formatNumber(seconds)} {(seconds == 1) ? 'second' : 'seconds'}
        </span>
      );
    }
  };

class CountdownTimer extends React.Component {
    state = {
        days: undefined,
        hours: undefined,
        minutes: undefined,
        seconds: undefined
    };

    render() {

        return (
            <div>
                <Countdown date={this.props.date} renderer={renderer} />
            </div>
        );
    }
}

export default CountdownTimer;
