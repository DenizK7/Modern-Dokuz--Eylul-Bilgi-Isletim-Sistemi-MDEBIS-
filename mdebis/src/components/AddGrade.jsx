import styled from "styled-components";
import "../homeSide.css";
import items from "../Adminsidebar.json";
import { useEffect, useState } from 'react';
import SidebarItem from './HomeSideBarItems';
import Lessons from "./Lessons";
import {Outlet, useLocation} from "react-router-dom";
import Button from "./Button";
import { Dropdown } from 'primereact/dropdown';


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

const mystyle = {
    fontSize: "10px",
    fontFamily: "Arial",
    fontWeight: "200",
    width: "25vw",
    position: "absolute",
    top: "15vh",
    right: "10vw"
  };
  const mystyle2 = {
    fontSize: "10px",
    fontFamily: "Arial",
    fontWeight: "200",
    width: "7vw",
    position: "absolute",
    right: "-7vw"
  };
  const grades =['AA','BA','BB','CB','CC','DC','DD','FD','FF'
    
  ];
  const uniquegrades = Array.from(new Set(grades))
  console.log(uniquegrades);
  function ChangeCourse(){
    var as;
    const[selectedStudent, setSelectedStudent] = useState([]);
    const [selectedExtension, setSelectedExtension] = useState([]);


  const onExtensionChange = event => {
     as= event.target.value;     
     var response
     
     try {
      var xhttp = new XMLHttpRequest();
       xhttp.open("GET", "http://localhost:8080/get_student_of_course/"+sessionStorage.getItem("token")+"/"+as.CourseId,false);
       xhttp.setRequestHeader("Content-type", "text/html");
       
       xhttp.onload = function (e) {
      if (xhttp.readyState === 4) {
          if (xhttp.status === 200) {
            response = JSON.parse(xhttp.response);  
             setSelectedStudent(response);
             setSelectedExtension(as); 
             
          }
       }
      }
    } catch (error) {
      alert("Wrong pass or id");
    }

      xhttp.send();
      console.log("Course id is " +as.CourseId)
      console.log(selectedStudent)
      
    } 
  
  

  
  
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
     const Inputt = ({setRerender, rerender})=>{
      const [gradeSelect,setGradeseletct] = useState('')
      const [inpt, setMessage] = useState('');
    
    function handleClick() {
      try {
         var xhttp = new XMLHttpRequest();//{courseId}/{studentId}/{grade}
         xhttp.open("GET", "http://localhost:8080/add_grade/"+sessionStorage.getItem("token")+"/"+selectedExtension.CourseId+'/'+inpt+'/'+gradeSelect,false);
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
        setMessage(event.target.value);
    
        
      };
      const onExtensionChange2 = event => {
        setGradeseletct(event.target.value);
       
     } 
      return(
        <div>
          <ButtonContainer  className="deleteID"> Enter Student ID : 
      <StyledInput type="text"
      id="inpt" name="inpt" placeholder="ID" onChange={handleChangeinpt}
      value={inpt}  ></StyledInput>
       <Dropdown value={gradeSelect} options={grades} onChange={onExtensionChange2}  placeholder={"Select a Lesson"}style={mystyle2}/>
       <Button  content={"Give Grade"} onClick={handleClick}> Give Grade</Button>
      </ButtonContainer>
    
        </div>
        
      
      );
      
    }
  
    

      return(
          <body className="noBg">
          <Inputt setRerender={setRerender} rerender={rerender}/>
        <Dropdown value={selectedExtension} options={lessons} onChange={onExtensionChange} optionLabel="CourseName" placeholder={"Select a Lesson"}style={mystyle}/>
      
             
        
        <table className={"grid-container-sm-admin"} >
     
      <tbody>
      <thead >
        <tr>
          <th >Student ID</th>
          <th >Student Name</th>
          <th >Student Surname</th>
         
        </tr>
      </thead>
         {selectedStudent.map(student => (
          
          <tr>
            {console.log("length is " +selectedStudent.length)}
            <td  className="tdstyle">{student.Id}</td>
            <td  className="tdstyle">{student.Name}</td>
            <td  className="tdstyle">{student.Surname}  </td>
            
          </tr>
          
        ))} 
            
      </tbody>
    </table>
                
       
    
      
          </body>
          
      );
  }
      export default ChangeCourse;