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

const Inputt = ({setRerender, rerender})=>{
  
  const [inpt, setMessage] = useState('');

function handleClick() {
  try {
     var xhttp = new XMLHttpRequest();
     xhttp.open("GET", "http://localhost:8080/delete_student/"+sessionStorage.getItem("token")+"/"+inpt,false);
     xhttp.setRequestHeader("Content-type", "text/html");
     xhttp.onload = function (e) {
      if (xhttp.readyState === 4) {
          if (xhttp.status === 200) {

             setRerender(!rerender);
           
             
          }
       }
    }
    console.log(inpt + " has been succesfully deleted");
    xhttp.send();
   

 } catch (error) {
   alert("Wrong pass or id");
 }
}
 
  
  const handleChangeinpt = event => {
    setMessage(event.target.value);

    
  };
  return(
    <ButtonContainer  className="deleteID"> You can delete with ID : 
  <StyledInput type="text"
  id="inpt" name="inpt" placeholder="DELETE" onChange={handleChangeinpt}
  value={inpt}  ></StyledInput>
  <Button  content={"Delete"} onClick={handleClick}> Delete</Button>
  </ButtonContainer>
  );
  
}

  function DeleteStudent(){
  
    const[rerender, setRerender] = useState(false);
    const[lessons, setContent] = useState([])
    useEffect(() => {
      try {
        var xhttp = new XMLHttpRequest();
        xhttp.open("GET", "http://localhost:8080/get_students/"+sessionStorage.getItem("token"),false);
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

       
     }, [rerender]);

      console.log(lessons)

      return(
          <div className="noBg">
              
        <Inputt setRerender={setRerender} rerender={rerender}/>
       
      
             
        
        <table className={"grid-container-sm-admin"} >
     
      <tbody>
      <thead >
        <tr>
          <th >Student ID</th>
          <th>NAME</th>
          <th>E-mail</th>
        </tr>
      </thead>
        {lessons.map(student => (
          <tr>
            <td className="tdstyle">{student.Id}</td>
            <td className="tdstyle">{student.Name}</td>
            <td className="tdstyle">{student.EMail}</td>
          </tr>
        ))}
      </tbody>
    </table>

       
    
      
          </div>
          
      );
  }
      export default DeleteStudent;