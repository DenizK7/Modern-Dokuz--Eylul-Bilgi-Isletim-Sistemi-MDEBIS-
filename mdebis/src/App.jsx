import './index.css'
import {BrowserRouter as Router, Route, Routes} from "react-router-dom";
import Loginpage from "./Pages/LoginPage"
import ForgotPassword from './Pages/ForgotPassword';
import Home from "./components/Home"
import Syllabus from './components/Syllabus';
import InfoLecture from './components/InfoLecture';
import AdminPage from './components/AdminPage';
import PastCourses from './components/PastCourses';
import {MainContext} from "./context";
import React, {useState} from 'react';
import AddStudent from './components/AddStudent';
import DeleteStudent from './components/DeleteStudent';
import AddLesson from './components/AddLesson';
import DeleteLesson from './components/DeleteLesson';

function  App ()  {
  
  const [infoStudent, setInfoStudent] = useState([]);
  const [token, setToken] = useState();
  const [navVisible, showNavbar] = useState(false);
    const data = {
      token,
      setToken,
        navVisible,
        showNavbar,
        infoStudent,
        setInfoStudent,
    }
  return (
    <MainContext.Provider value ={data}>
    <body >
      <Router>
        <Routes>
          <Route path ="/" element ={<Loginpage />} />   
          <Route path ="/ForgotPassword" element ={<ForgotPassword/>} />  


          <Route path ="/AdminPage" element ={<AdminPage/>}>
          <Route path ="/AdminPage/AddStudent" element ={<AddStudent/>} />
          <Route path ="/AdminPage/AddLesson" element ={<AddLesson/>} />
          <Route path ="/AdminPage/DeleteLesson" element ={<DeleteLesson/>} />
          <Route path ="/AdminPage/DeleteStudent" element ={<DeleteStudent/>} />
          </Route>  

          <Route path ="/HomePage"   element ={<Home/>} >
          <Route path ="/HomePage/infoLecture" element ={<InfoLecture/>} />
          <Route path ="/HomePage/Syllabus" element ={<Syllabus/>} />
          <Route path ="/HomePage/PastCourses" element ={<PastCourses/>} />
          </Route>    

          
        </Routes>
      </Router>
    </body>
    </MainContext.Provider>
    
  );
}



export default App;