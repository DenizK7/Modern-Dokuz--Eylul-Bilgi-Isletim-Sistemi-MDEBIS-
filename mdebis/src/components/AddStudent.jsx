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
  
  
  const [id, setID] = useState('');

  const [password, setPassword] = useState('');

  

  const [name, setName] = useState('');

  const [surname, setSurname] = useState('');
  const [dep_name, setDep_name] = useState('');


function handleClick() {
  try {
     var xhttp = new XMLHttpRequest();
     xhttp.open("GET", "http://localhost:8080/create_student/"+sessionStorage.getItem("token")+"/"+id+"/"+password+"/"+name+"/"+surname+"/"+dep_name,false);
     xhttp.setRequestHeader("Content-type", "text/html");
     xhttp.onload = function (e) {
      if (xhttp.readyState === 4) {
          if (xhttp.status === 200) {

             setRerender(!rerender);
           
             
          }
       }
    }
    
    xhttp.send();
   

 } catch (error) {
   alert("Wrong pass or id");
 }
}
 
  
  const handleChangeinpt = event => {
    setID(event.target.value);
  };
  const handleChangeinpt2 = event => {
    setPassword(event.target.value);
  };
  
  const handleChangeinpt4 = event => {
    setName(event.target.value);
  };
  const handleChangeinpt5 = event => {
    setSurname(event.target.value);
  };
  const handleChangeinpt6 = event => {
    setDep_name(event.target.value);
  };
  return(
    <div  className="deleteID2">You can add with : 
         <ButtonContainer > 
  <StyledInput type="text"
  id="id" name="id" placeholder="ID" onChange={handleChangeinpt}
  value={id}  ></StyledInput>
  <StyledInput type="text"
  id="password" name="password" placeholder="PASSWORD" onChange={handleChangeinpt2}
  value={password}  ></StyledInput>
  
  <StyledInput type="text"
  id="name" name="name" placeholder="NAME" onChange={handleChangeinpt4}
  value={name}  ></StyledInput>
  
  
  
  </ButtonContainer>
  <ButtonContainer>
 
   <StyledInput type="text"
  id="surname" name="surname" placeholder="SURNAME" onChange={handleChangeinpt5}
  value={surname}  ></StyledInput>
  <StyledInput type="text"
  id="dep_name" name="dep_name" placeholder="DEP_NAME" onChange={handleChangeinpt6}
  value={dep_name}  ></StyledInput>
  <Button  content={"Add"} onClick={handleClick}> Add</Button>
  </ButtonContainer>
  
    </div>
   
  
  );
  
}

  function AddLecturer(){
  
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


      return(
          <div className="noBg">
              
        <Inputt setRerender={setRerender} rerender={rerender}/>
       
      
             
        
        <table className={"grid-container-sm-admin2"} >
     
      <tbody>
   
        <tr>
          <th >ID</th>
          <th>Name</th>
          <th>Surname</th>
        </tr>
      
        {lessons.map((student,index) => (
          <tr key={index}>
            <td className="tdstyle">{student.Id}</td>
            <td className="tdstyle">{student.Name}</td>
            <td className="tdstyle">{student.Surname}</td>
          </tr>
        ))}
      </tbody>
    </table>

       
    
      
          </div>
          
      );
  }
      export default AddLecturer;