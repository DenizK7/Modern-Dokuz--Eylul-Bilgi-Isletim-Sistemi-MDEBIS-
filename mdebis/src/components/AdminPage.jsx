
import "../adminpage.css";
import items from "../Adminsidebar.json";
import SidebarItem from './HomeSideBarItems';
function AdminPage(){
    return(
        <body className="noBg">
            {/* <!-- Create a container for the main content and the sidebar --> */}
    <div className="containerAdmin">
      {/* <!-- The sidebar --> */}
      <aside class="sidebarAdmin">
      { items.map((item, index) => <SidebarItem key={index} item={item} />) }
      </aside>
      {/* <!-- The main content --> */}
      <div className="main-content">
        <h1>Welcome to my web page!</h1>
        <p>Here is some content for the main part of the page.</p>
      </div>
    </div>
        </body>
        
    );
}
    export default AdminPage;