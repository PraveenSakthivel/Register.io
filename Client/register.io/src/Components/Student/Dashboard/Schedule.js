import React from 'react';
import 'devextreme/dist/css/dx.common.css';
import 'devextreme/dist/css/dx.light.css';
import { classes, locationColor } from "./ScheduleData"
 
import { Scheduler, Resource } from 'devextreme-react/scheduler';

class Schedule extends React.Component {

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
                    dataSource={classes}
                    onAppointmentDblClick= {function(e) {e.cancel = true;}}
                    onAppointmentClick= {function(e) {e.cancel = true;}}
                >
                    
                    <Resource
                        dataSource={locationColor}
                        fieldExpr="location"
                        useColorAsDefault={true}
                        />
                </Scheduler>
            </div>
        );
    }
}

export default Schedule;