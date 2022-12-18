
import "../adminpage.css";
function AdminPage(){
    return(
        <body className="noBg">
            {/* <!-- Create a container for the main content and the sidebar --> */}
    <div className="containerAdmin">
      {/* <!-- The sidebar --> */}
      <aside class="sidebarAdmin">
        <h3>Sidebar</h3>
        <p>Here is some content for the sidebar.</p>
        <ul>
          <li><a href="#"> Delete Student</a></li>
          <li><a href="#">Delete Lecturer</a></li>
          <li><a href="#">Lol</a></li>
        </ul>
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