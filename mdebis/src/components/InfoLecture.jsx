import React, {useState, useEffect, useMemo} from 'react';
import Navbar from "./HomeSideInner";
import { BrowserRouter, Routes, Route,Navigate } from "react-router-dom";
import "../homeSide.css";
import Syllabus from "./Syllabus";
import {Outlet} from "react-router-dom";
import{MainContext, useContext} from '../context'



function InfoLecture() {
    const[contents, setContent] = useState([])
    
    useEffect(() => {
        try {
          var xhttp = new XMLHttpRequest();
          xhttp.open("GET", "http://localhost:8080/get_home_entry/"+sessionStorage.getItem("token"),false);
          xhttp.setRequestHeader("Content-type", "text/html");
          xhttp.onload = function (e) {
           if (xhttp.readyState === 4) {
               if (xhttp.status === 200) {
      
                var response = JSON.parse(xhttp.response);   
                setContent(response);    
                   
               }
            }
         }
        
         xhttp.send();
        
      
      } catch (error) {
        alert("Wrong pass or id");
      }
         
         
       }, []);
    
	return (
		
			<div className='grid-container-info'>
                 {contents?.map((content,index) => (
        <div key={index} >
          <div style={{"fontWeight":"600", "textAlign":"center"}}>Department Name </div><div  style={{"fontWeight":"600", "textAlign":"center"}}> {content.DepName}</div>
          <br></br>
          <div className='wei'>Course Name </div> <div>{content.CourseName}</div>
          <br></br>
          <div className='wei'>Lecturer Name </div> <div className='new-line'>{content.LecName}</div>
            <br></br><hr></hr>
          
            
           {content.Announcements&&  content.Announcements?.map((Announcement,index) => (
            
            <div key={index}>
               <div className='wei'>Announcment</div>
              
               <span className='wei'>Title </span><div>{Announcement.Title}</div>
              
               <span className='wei'>Content </span><div>{Announcement.Content}</div>
              
            </div>
          ))} 
           <br></br><hr></hr>
              <div>Credit : {content.Credit}</div>
              
              <div>Current Attendance: {content.CurrentNonAttendance}</div>
              <div>Attendance Limit {content.AttendanceLimit}</div>
          <hr />
        </div>
      ))}


				
			</div>
		
  );
}
export default InfoLecture;