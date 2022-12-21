import styled from "styled-components";
import "../homeSide.css";
import items from "../Adminsidebar.json";
import { useEffect, useState } from 'react';
import SidebarItem from './HomeSideBarItems';
import Lessons from "./Lessons";
import {Outlet, useLocation} from "react-router-dom";
import Button from "./Button";
import { Dropdown } from 'primereact/dropdown';


const StyledInput = styled.textarea`
background: rgba(255, 255, 255, 0.15);
box-shadow: 0 8px 32px 0 rgba(31, 38, 135, 0.37);
border-radius: 2rem;
width: 100%;
height: 20rem;
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
const StyledInputt = styled.input`
background: rgba(255, 255, 255, 0.15);
box-shadow: 0 8px 32px 0 rgba(31, 38, 135, 0.37);
border-radius: 2rem;
width: 100%;
height: 2rem;
padding: 1rem;
margin: 10px;

border: 1rem;
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

const Makediv =({info})=>{
  return(<div className="perpendicular-line">{info} </div>)
}


const Inputt = ({setRerender, rerender, lessons})=>{
    const mystyle = {
        fontSize: "10px",
        fontFamily: "Arial",
        fontWeight: "200",
        width: "10vw",
        position: "absolute",
        top: "15vh",
        right: "45vh"
      };
    const [selectedExtension, setSelectedExtension] = useState([]);
    const onExtensionChange = event => {
       var as= event.target.value;     
        setSelectedExtension(as);
      
         
       } 
    

   
  
  const [header, setHeader] = useState('');
  const [content, setContent] = useState('');

function handleClick() {
  try {
     var xhttp = new XMLHttpRequest();
     xhttp.open("GET", "http://localhost:8080/change_course_status/"+sessionStorage.getItem("token")+"/"+selectedExtension.CourseId+"/"+header+"/"+content,false);
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
 
  
  const handleChangeinpt_content = event => {
    setContent(event.target.value);

    
  };
  const handleChangeinpt_header = event => {
    setHeader(event.target.value);

    
  };
 
  return(
    <div  > Enter Your announcment: 

<Dropdown value={selectedExtension} options={lessons} onChange={onExtensionChange} optionLabel="CourseName" placeholder={"Select a Lesson"}style={mystyle}/>
        <StyledInputt type="text"
  id="header" name="header" placeholder="Header" onChange={handleChangeinpt_header}
  value={header}  style={{height:"2rem"}}></StyledInputt>
  <StyledInput type="text"
  id="content" name="content" placeholder="Content" onChange={handleChangeinpt_content}
  value={content}  ></StyledInput>
  <Button  content={"Submit"} onClick={handleClick}> Submit</Button>
  </div>
  );
  
}

  function ChangeCourse(){
    
    const[rerender, setRerender] = useState(false);
   
    const[lessons, setContent] = useState([])
    useEffect(() => {
      try {
        var xhttp = new XMLHttpRequest();//change_course_status/{sessionHash}/{courseId}
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

       
     }, [rerender]);


      return(
          <body className="noBg">
        
        <Inputt setRerender={setRerender} lessons={lessons} rerender={rerender}/>
      
          </body>
          
      );
  }
      export default ChangeCourse;