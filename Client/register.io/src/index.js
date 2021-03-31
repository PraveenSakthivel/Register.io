import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';

import './Components/Login/Login.css'
import './Components/Navbar/Navbar.css'
import './Components/Content/Content.css'
import './Components/Student/StudentManageReg/StudentManageReg.css'
import './Components/Student/Dashboard/Dashboard.css'
import './Components/Student/CourseLookup/CourseLookup.css'
import './Components/Student/CourseLookup/CourseTable.css'
import './Components/Student/ClassHistory/ClassHistory.css'
import './Components/Student/MyAccount/MyAccount.css'
import './Components/Admin/Dashboard/Dashboard.css'
import './Components/Footer/Footer.css'

import App from './App';

ReactDOM.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
  document.getElementById('root')
);
