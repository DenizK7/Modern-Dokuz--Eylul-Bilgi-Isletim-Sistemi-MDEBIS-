import styled from "styled-components";
import "../homeSide.css";
import items from "../Adminsidebar.json";
import { useEffect, useState } from 'react';
import SidebarItem from './HomeSideBarItems';
import Lessons from "./Lessons";

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
width: 100%;
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
    <div  className="deleteID"> You can delete with ID : 
  <StyledInput type="text"
  id="inpt" name="inpt" placeholder="DELETE" onChange={handleChangeinpt}
  value={inpt}  ></StyledInput>
  <ButtonContainer> <button>Delete</button></ButtonContainer>
  </div>
  );
  
}

function AdminPage(){
  
 
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
            {/* <!-- Create a container for the main content and the sidebar --> */}
      <Inputt />
     
      {/* <!-- The sidebar --> */}
      <aside class="sidebarAdmin">
      { items.map((item, index) => <SidebarItem key={index} item={item} />) }
      </aside>
      {/* <!-- The main content --> */}
      
           
      <div style={{transition:"0.8s"}} className={"grid-container-sm-admin"}  >
    
    <div className="days" >Course ID</div>
    <div className="days">Deparment Name</div>
    <div className="days" >Lecturer Name</div>
    <div className="days">Perşembe</div>
    <div className="days">Cuma</div>
     {
          lessons?.map(lessons =>  <div ><Lessons Department={lessons.Department} Course_name={lessons.Course_name} Lecturer_name={lessons.Lecturer_name} AttandenceLimit = {lessons.AttandenceLimit}></Lessons></div>)
        } 
      </div>
  
    
        </body>
        
    );
}
    export default AdminPage;