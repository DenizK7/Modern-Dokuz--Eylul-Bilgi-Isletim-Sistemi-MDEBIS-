import { useEffect, useState } from "react";
import AnnouncmentMainPage from "./AnnouncmentMainPage";
export default function Sidebar(props) {
  const { showModal, setShowModal } = props;
  const[content, setContent] = useState([])
  useEffect(() => {
    var xhttp = new XMLHttpRequest();
    xhttp.open("GET", "http://localhost:8080/get_gen_announcements",false);
    xhttp.setRequestHeader("Content-type", "text/html");
    xhttp.send();
    var response = JSON.parse(xhttp.response);
   

    setContent(response);
   
    
    
  }, []);
  return (
    <div
      className={showModal ? "modal-dialog show" : "modal-dialog"}
      role="document"
    >
      <div className="modal-content">
        <div className="modal-header">
        <button
            type="button"
            className="close"
            data-dismiss="modal"
            aria-label="Close"
            style ={{width:"25px", marginTop:"1vh", height:"25px"}}
            onClick={() => {
              setShowModal(false);
            }}
          ></button>
          
        </div>

        <div className="modal-body">
        
          {
            content?.map(content => <AnnouncmentMainPage key={content.Title}  header={content.Title} content={content.Content} Link={content.Link}></AnnouncmentMainPage>)
          }
        
        </div>
      </div>
    </div>
  );
}
