import { useState, useEffect } from "react";
import { Link } from "react-router-dom";


const AnnouncmentMainPage =({key, header, content,Link}) =>{
  const [showModal, setShowModal] = useState(false);
return (
  <div className="anc-main" >
    <h1  className="modal-title" id="myModalLabel" style={{fontSize: "15px"}}>
    <a target="_blank" href={Link}>{header}
      </a></h1>
     <p className={showModal ? "anc-main-long" : "anc-main-short"}>
          {content}
          </p>
          <button  onClick={() => {
              setShowModal(!showModal);
            }} >{showModal? "Gizle":"Devamını Oku" }
            </button>
  </div>
 
)
}

export default AnnouncmentMainPage;




