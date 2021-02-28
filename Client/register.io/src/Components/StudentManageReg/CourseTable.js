import React, { useState } from 'react';
import { EditingState } from '@devexpress/dx-react-grid';
import {
  Grid,
  Table,
  TableHeaderRow,
  TableEditRow,
  TableEditColumn,
} from '@devexpress/dx-react-grid-bootstrap4';
import '@devexpress/dx-react-grid-bootstrap4/dist/dx-react-grid-bootstrap4.css';


const getRowId = row => row.id;

const CommandButton = ({
    onExecute, icon, text, hint, color,
  }) => (
    <button
      type="button"
      className="btn"
      style={{ padding: 11, fontWeight:"500" }}
      onClick={(e) => {
        onExecute();
        e.stopPropagation();
      }}
      title={hint}
    >
      <span className={color || 'undefined'}>
        {icon ? <i className={`oi oi-${icon}`} style={{ marginRight: text ? 5 : 0 }} /> : null}
        {text}
      </span>
    </button>
);

const AddButton = ({ onExecute }) => (
    <CommandButton color="text-primary" text="Add" hint="Add Course" onExecute={onExecute} />
);

const DeleteButton = ({ onExecute }) => (
    <CommandButton
      hint="Drop Course"
      color="text-danger"
      text="Drop"
      onExecute={() => {
        if (window.confirm('Are you sure you want to drop this class?')) {
          onExecute();
        }
      }}
    />
  );

const CommitButton = ({ onExecute }) => (
<CommandButton text="Confirm" hint="Confirm Add" color="text-success" onExecute={onExecute} />
);

const CancelButton = ({ onExecute }) => (
<CommandButton text="Cancel" hint="Cancel Add" color="text-danger" onExecute={onExecute} />
);

const commandComponents = {
    add: AddButton,
    delete: DeleteButton,
    commit: CommitButton,
    cancel: CancelButton
};

const Command = ({ id, onExecute }) => {
    const ButtonComponent = commandComponents[id];
    return (
        <ButtonComponent
        onExecute={onExecute}
        />
    );
};

class CourseTable extends React.Component {

    constructor(props) {
      super(props);
      this.state = {rows:[
        { id: 1, coursecode:93028, coursenumber: '01:198:352:01', coursename: 'Internet Technology', credits: 4.0, status: "Added!" },
        { id: 2, coursecode:30284, coursenumber: '18:332:251:03', coursename: 'Probability and Random Processes', credits: 3.0, status: "Added!" },
        { id: 3, coursecode:93028, coursenumber: '01:198:352:01', coursename: 'Internet Technology', credits: 4.0, status: "Pending Add" },
        { id: 4, coursecode:30284, coursenumber: '18:332:251:03', coursename: 'Probability and Random Processes', credits: 3.0, status: "Class Filled" },
        { id: 5, coursecode:93028, coursenumber: '01:198:352:01', coursename: 'Internet Technology', credits: 4.0, status: "PreReqs Not Met" },
        { id: 6, coursecode:30284, coursenumber: '18:332:251:03', coursename: 'Probability and Random Processes', credits: 3.0, status: "Added!" }
        ]
      };
    }

    render() {

      const columns = [
          { name: 'coursecode', title: 'Course Code' },
          { name: 'coursenumber', title: 'Course Number' },
          { name: 'coursename', title: 'Course Name' },
          { name: 'credits', title: 'Credits'},
          { name: 'status', title: 'Status'}
          ];

      const commitChanges = ({ added, changed, deleted }) => {
          let changedRows;
          let rows = this.state.rows;
          if (added) {
              const startingAddedId = rows.length > 0 ? rows[rows.length - 1].id + 1 : 0;
              changedRows = [
                  ...rows,
                  ...added.map((row, index) => ({
                  id: startingAddedId + index,
                  ...row,
                  })),
              ];
          }
          if (changed) {
              changedRows = rows.map(row => (changed[row.id] ? { ...row, ...changed[row.id] } : row));
          }
          if (deleted) {
              const deletedSet = new Set(deleted);
              changedRows = rows.filter(row => !deletedSet.has(row.id));
          }
          this.setState({rows: changedRows});
      };

      const tableColumnExtensions = [
          { columnName: 'coursename', width: '30%' },
          { columnName: 'credits', width: '10%'},
          { columnName: 'coursecode', width: '12%' },
          { columnName: 'coursenumber', width: '15%' }
        ];

      const TableComponent = ({ ...restProps }) => (
        <Table.Table {...restProps} className="table-striped" />
      );

      const HighlightedCell = ({ value, style, ...restProps }) => (
        <Table.Cell
          {...restProps}
          style={{
            color: value == 'Added!' ? '#009432' : (value == 'Pending Add' ? '#e58e26' : (value == 'Class Filled' || value == 'PreReqs Not Met' ? '#d63031' : undefined)),
            fontWeight:"500",
            ...style,
          }}
        >
          <span
            style={{
            }}
          >
            {value}
          </span>
        </Table.Cell>
      );
      
      const Cell = (props) => {
        const { column } = props;
        if (column.name === 'status') {
          return <HighlightedCell {...props} />;
        }
        return <Table.Cell {...props} />;
      };


      return (
          <div className="courseTable-container">
          <Grid
              rows={this.state.rows}
              columns={columns}
              getRowId={getRowId}
          >
              <EditingState
                onCommitChanges={commitChanges}
              />
              <Table cellComponent={Cell} tableComponent={TableComponent} columnExtensions={tableColumnExtensions} />
              <TableHeaderRow />
              <TableEditRow />
              <TableEditColumn
                showAddCommand
                showDeleteCommand
                commandComponent={Command}
              />
          </Grid>
          </div>
      );
    }
}

export default CourseTable;