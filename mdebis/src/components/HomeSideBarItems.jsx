import { useState } from "react"
import "../sidebarItem.css";
import { Navigate, Link } from "react-router-dom";
export default function SidebarItem({item}){
    const [open, setOpen] = useState(false)
    function logoutreq(title){
        if(title=="Logout"){
            log_out_backend();
            return(
                ()=><Link to= {item.path} ></Link>
            )
            
        }
        else{
            return(
                ()=><Link to= {item.path} ></Link>
            )

        }

       
       
    }
   function log_out_backend(){
    try {
        var xhttp = new XMLHttpRequest();
        xhttp.open("GET", "http://localhost:8080/log_out/"+sessionStorage.getItem("token"),false);
        xhttp.setRequestHeader("Content-type", "text/html");
        xhttp.onload = function (e) {
         if (xhttp.readyState === 4) {
             if (xhttp.status === 200) {
                sessionStorage.setItem("token", "e");
                
                
             }
          }
          
       }
       xhttp.send();
      console.log("succesful" + sessionStorage.getItem("token"));
   
    } catch (error) {
      alert("Wrong pass or id");
    }
   }
    
    if(item.childrens){
        return (
            <div className={open ? "sidebar-item open" : "sidebar-item"} >
                <div className="sidebar-title">
                    <span>
                        { item.icon && <i className={item.icon}></i> }
                        {item.title}    
                    </span> 
                    <i className="bi-chevron-down toggle-btn" onClick={() => setOpen(!open)}></i>
                </div>
                <div className="sidebar-content">
                    { item.childrens.map((child, index) => <SidebarItem key={index} item={child} />) }
                </div>
            </div>
        )
    }else{
        return (
            <a href={item.path || "#"} className="sidebar-item plain" onClick={ logoutreq(item.path)  }>
                { item.icon && <i className={item.icon}></i> }
                {item.title}
            </a>
        )
    }
}