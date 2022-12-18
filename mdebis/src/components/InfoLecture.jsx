import React, {useState, useEffect, useMemo} from 'react';
import Navbar from "./HomeSideInner";
import { BrowserRouter, Routes, Route,Navigate } from "react-router-dom";
import "../homeSide.css";
import Syllabus from "./Syllabus";
import {Outlet} from "react-router-dom";
import{MainContext, useContext} from '../context'
const InfoCards = ()=>{
    return(
        <div className='grid-container-info'>
              


				
			</div>
    )
}


function InfoLecture() {
    const[contents, setContent] = useState([])
    useEffect(() => {
        try {
          var xhttp = new XMLHttpRequest();
          xhttp.open("GET", "http://localhost:8080/get_past_courses/"+sessionStorage.getItem("token"),false);
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
    const{navVisible, infoStudent, setInfoStudent}= useContext(MainContext);
    console.log("session storage is " +sessionStorage.getItem("token"));
    console.log(contents);
	return (
		
			<div className='grid-container-info'>
                <div >Intro to Computer Engineer <br></br><hr></hr><span>infoinfoinfoinfoinfoinfoinfoinfoinfoinfoinfoinfoi
                    nfoinfoinfoinfoinfoinfoinfoinfoinfoinfoinfoinfo</span></div>
                <div >Intro to Computer Engineer <br></br><hr></hr><span>infoinfoinfoinfoinf
                    oinfoinfoinfoinfoinfoinfoinfoinfoinfoinfoinfoinfoinfoi
                    nfoinfoinfoinfoinfoinfo</span></div>

                <div >Intro to Computer Engineer <br></br><hr></hr><span>infoinfoinfoin
                    foinfoinfoinfoinfoinfoinfoinfoinfoinfoinfoinfoinfoinfoinfoinfoinfoinfoinfoinfo
                    infoinfoinfoinfoinfoinfoinfoinfoinfoinfoinfoinfoinfoinfoinfoinfoinfo</span></div>

                <div >Intro to Computer Engineer <br></br><hr></hr><span>infoinfoinfoinfoinfoinfoinfoinfo</span></div>

                <div >Intro to Computer Engineer <br></br><hr></hr><span>infoinfoinfoinfoinfoinfoinfoinfo</span></div>


				
			</div>
		
  );
}
export default InfoLecture;