
import { useState } from "react";
import Sidebar from "./SideBar";

const SidebarButton =() =>{
  const [showModal, setShowModal] = useState(false);
  
 
return (
  <div>
<button className={"sidebarButton bi-fill"}
  onClick={() => {
    
    setShowModal(!showModal);
  }}
>
</button>
<Sidebar showModal={showModal} setShowModal={setShowModal} />
  </div>   
  );

}
export default SidebarButton;