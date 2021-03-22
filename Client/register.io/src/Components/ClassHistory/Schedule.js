import React from 'react';
import 'devextreme/dist/css/dx.common.css';
import 'devextreme/dist/css/dx.light.css';
 
import { Scheduler, Resource } from 'devextreme-react/scheduler';

class Schedule extends React.Component {

    componentDidMount () {
        let headers = document.getElementsByClassName("dx-scheduler-header-panel-cell dx-scheduler-cell-sizes-horizontal");
        let correctedHeaders = ["Mon", "Tues", "Wed", "Thurs", "Fri"]
        for(let i = 0; i < headers.length; i++){
            headers[i].innerText = correctedHeaders[i]
        }
    }

    componentDidUpdate () {
        let headers = document.getElementsByClassName("dx-scheduler-header-panel-cell dx-scheduler-cell-sizes-horizontal");
        let correctedHeaders = ["Mon", "Tues", "Wed", "Thurs", "Fri"]
        for(let i = 0; i < headers.length; i++){
            headers[i].innerText = correctedHeaders[i]
        }        
    }

    render() {
        return (
            <div class="schedule">
                <Scheduler 
                    id="scheduler"
                    textExpr="title"
                    currentView="workWeek"
                    startDayHour={8}
                    endDayHour={22}
                    timeZone="America/New_York"
                    showAllDayPanel={false}
                    editing={{ allowAdding: false, allowDeleting: false, allowDragging: false, allowResizing: false, allowUpdating: false, allowEditingTimeZones: false }}
                    cellDuration="60"
                    dataSource={this.props.classes}
                    onAppointmentDblClick= {function(e) {e.cancel = true;}}
                    onAppointmentClick= {function(e) {e.cancel = true;}}
                    showCurrentTimeIndicator={false}
                >
                    
                    <Resource
                        dataSource={this.props.locationColor}
                        fieldExpr="location"
                        useColorAsDefault={true}
                        />
                </Scheduler>
            </div>
        );
    }
}

export default Schedule;