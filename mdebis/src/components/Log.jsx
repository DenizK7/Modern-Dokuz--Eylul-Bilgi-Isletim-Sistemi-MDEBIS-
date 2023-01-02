import styled from "styled-components";
import "../homeSide.css";
import items from "../Adminsidebar.json";
import { useEffect, useState } from 'react';
import SidebarItem from './HomeSideBarItems';
import Lessons from "./Lessons";
import {Outlet, useLocation} from "react-router-dom";
import Button from "./Button";


  function Log(){
  
    const[rerender, setRerender] = useState(false);
    const[lessons, setContent] = useState([])
    useEffect(() => {
      try {
        var xhttp = new XMLHttpRequest();
        xhttp.open("GET", "http://localhost:8080/get_log/"+sessionStorage.getItem("token"),false);
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
      alert("Cannot Reach the Log");
    }

       
     }, [rerender]);

      

      return(
          <div className="noBg">
           
       
      
             
        <h1 style ={{paddingLeft: "45vw" ,paddingTop: "12vh"}}>LOGS</h1>
           <table className={"log"} >
     
     <tbody>
    
       {lessons.map(log => (
         <tr>
         <th className="tdstyle">Operation : {log.Operation} <br />Record Id : {log.RecordId}<br /> Table : {log.WhichTable}<br /> Who Did : {log.WhoDid}<br /> Id Person : {log.WhoDidId}</th>
         <td className="tdstyle">{log.Values}</td>
      
       </tr>
        //  <tr>
        //    <td className="tdstyle">{student.Id}</td>
        //    <td className="tdstyle">{student.Name}</td>
        //    <td className="tdstyle">{student.Surname}</td>
        //  </tr>
       ))}
     </tbody>
   </table>
 

       
    
      
          </div>
          
      );
  }
      export default Log;