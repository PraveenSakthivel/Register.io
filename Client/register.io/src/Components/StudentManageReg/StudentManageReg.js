import React from 'react';

class StudentManageReg extends React.Component {

    constructor(props) {
        super(props);
        this.state = {};
    }

    render() {

        return (
            <div class="studentManageReg">
                <h3>Manage Registration</h3>
                <hr style={{color:"grey"}}></hr>

                <div class="dropdown show">
                    <a class="btn btn-secondary dropdown-toggle" href="#" role="button" id="dropdownMenuLink" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                        Spring 2021
                    </a>

                    <div class="dropdown-menu" aria-labelledby="dropdownMenuLink">
                        <a class="dropdown-item" href="#">Winter 2021</a>
                        <a class="dropdown-item" href="#">Spring 2021</a>
                        <a class="dropdown-item" href="#">Summer 2021</a>
                        <a class="dropdown-item" href="#">Fall 2021</a>
                    </div>
                </div>

            </div>
        );
    }
}

export default StudentManageReg;