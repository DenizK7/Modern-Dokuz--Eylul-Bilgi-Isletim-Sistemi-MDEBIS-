import React, {useState} from 'react';
import Navbar from "./HomeSideInner";
import { BrowserRouter, Routes, Route,Navigate } from "react-router-dom";
import "../homeSide.css";
import Syllabus from "./Syllabus";
import {Outlet} from "react-router-dom";

function InfoLecture() {
	 
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