import { useState, useEffect } from "react";


const Lessons =({Department, Course_name, Lecturer_name, AttandanceLimit}) =>{
  const [showModal, setShowModal] = useState(false);
  
return (
  <div >
    <h5>{Department}</h5>
     <div>{Course_name}<br></br>
       {Lecturer_name}</div>
        
        <div>Devamsızlık sınırı : {AttandanceLimit}</div>
     
   
  </div>
 
)
}

export default Lessons;
