import { useState, useEffect } from "react";


const Lessons =({Department, Course_name, Lecturer_name, AttandenceLimit}) =>{
 



  
return (
  <div >
    <h5>{Department}</h5>
     <div>{Course_name}<br></br>
       {Lecturer_name}</div>
        <br></br>
       <div>{AttandenceLimit}</div>
     
   
  </div>
 
)
}

export default Lessons;
