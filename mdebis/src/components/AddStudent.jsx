import styled from "styled-components";
import "../homeSide.css";
import items from "../Adminsidebar.json";
import { useEffect, useState } from 'react';
import SidebarItem from './HomeSideBarItems';
import Lessons from "./Lessons";
import {Outlet, useLocation} from "react-router-dom";
import Button from "./Button";


const StyledInput = styled.input`
background: rgba(255, 255, 255, 0.15);
box-shadow: 0 8px 32px 0 rgba(31, 38, 135, 0.37);
border-radius: 2rem;
width: 45%;
height: 3rem;
padding: 1rem;
border: none;
outline: none;
color: #3c354e;
font-size: 1rem;
font-weight: bold;
&:focus {
  display: inline-block;
  box-shadow: 0 0 0 0.2rem #b9abe0;
  backdrop-filter: blur(12rem);
  border-radius: 2rem;
}
&::placeholder {
  font-weight: 100;
  font-size: 1rem;
}
`;
const ButtonContainer = styled.div`
margin: 1rem 0 1rem 0;
width: 40%%;
display: flex;
align-items: center;
justify-content: center;
`;

const Inputt = ()=>{
  const [inpt, setMessage] = useState('');
 
  const [inpta, setMessagae] = useState('');
  const handleChangeinpt = event => {
    setMessage(event.target.value);

    
  };
  return(
    <ButtonContainer  className="deleteID"> You can delete with ID : 
  <StyledInput type="text"
  id="inpt" name="inpt" placeholder="DELETE" onChange={handleChangeinpt}
  value={inpt}  ></StyledInput>
  <Button  content={"Delete"}> Delete</Button>
  </ButtonContainer>
  );
  
}

  function AddStudent(){
  
 
    const[lessons, setContent] = useState([])
    useEffect(() => {
      try {
        var xhttp = new XMLHttpRequest();
        xhttp.open("GET", "http://localhost:8080/time_table/"+sessionStorage.getItem("token"),false);
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
  
  
      return(
          <body className="noBg">
              
        <Inputt />
       
      
             
        <div style={{transition:"0.8s"}} className={"grid-container-sm-admin"}  >
      
      <div className="days" >Student Number</div>
      <div className="days">Student Name</div>
      <div className="days" >Grade vs.</div>
      
       {
            lessons?.map(lessons =>  <div ><Lessons Department={lessons.Department} Course_name={lessons.Course_name} Lecturer_name={lessons.Lecturer_name} AttandenceLimit = {lessons.AttandenceLimit}></Lessons></div>)
          } 
        </div>
    
      
          </body>
          
      );
  }
      export default AddStudent;