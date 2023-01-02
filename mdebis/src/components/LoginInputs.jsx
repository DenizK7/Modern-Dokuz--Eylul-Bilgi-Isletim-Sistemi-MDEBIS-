import styled from "styled-components";
import {useState} from 'react';
import { useTranslation } from "react-i18next";
import 'primeicons/primeicons.css';
import 'primereact/resources/themes/lara-light-indigo/theme.css';
import 'primereact/resources/primereact.css';
import { Dropdown } from 'primereact/dropdown';
import axios from "axios";
import Button from "./Button";
import { Navigate, useNavigate } from "react-router-dom";
import{MainContext, useContext} from '../context'
import { useEffect } from "react";
const ButtonContainer = styled.div`
  margin: 1rem 0 1rem 0;
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
`;
const FormatMail = styled.span`
  width: 35%;
  cursor: pointer;
  text-transform: lowercase;
  font-size: 6px;
  letter-spacing: 0rem;
  `;
  
  const InputContainer = styled.div`
  margin: 0rem 0 1rem 0;  
  display:flex;
  align-items: left;
  width:80%;
`;

function LoginInputs(){
  
  const navigate = useNavigate();
  const {t} = useTranslation();
  const [inpt, setMessage] = useState('');
  const [pssw, setpssw] = useState('');
  
  const extensions = [
    { name: '@ogr.deu.edu.tr', code: 'student' },
    { name: '@deu.edu.tr', code: 'teacher' },
    { name: '@admin.deu.edu.tr', code: 'admin' }
];   
const [selectedExtension, setSelectedExtension] = useState('');


  const handleChangeinpt = event => {
   
    setMessage(event.target.value);

    
  };
  const handleChangepsw = event => {
    setpssw(event.target.value);

    
  };
  const onExtensionChange = event => {
    setSelectedExtension(event.target.value);     
  }
  async function handleClick() {
    const abortController = new AbortController();
    var ex;
    if(selectedExtension.name==='@ogr.deu.edu.tr'){
       ex = "http://localhost:8080/log_student/";
      
     }
     else if(selectedExtension.name==='@deu.edu.tr'){
      ex = "http://localhost:8080/log_lecturer/";
      
    }
    else if(selectedExtension.name==='@admin.deu.edu.tr'){
      ex = "http://localhost:8080/log_admin/";
      
     }
   if(inpt.length >0){

   
      try{
        var xhttp = new XMLHttpRequest();
       
          xhttp.open("GET",ex +inpt+"/"+pssw);//buraya if lerle neye giriceğini seçtir log admin student vs vs.
         xhttp.setRequestHeader("Content-type", "text/html");
         
         xhttp.onload = function (e) {
          if (xhttp.readyState === 4) {
              if (xhttp.status === 200) {
 
                var response = JSON.parse(xhttp.responseText);
 
                if(!response){
                  alert("Wrong Password or id");
                 }
                 else{
                  sessionStorage.setItem("token", response);
                  if(selectedExtension.name==='@ogr.deu.edu.tr'){
                   navigate("/HomePage/infoLecture");
                  
                  }
                 
                  else if(selectedExtension.name==='@deu.edu.tr'){
                   navigate("/LecturerPage/AddAnnouncment")
                  
                  }
                  else if(selectedExtension.name==='@admin.deu.edu.tr'){
                   navigate("/AdminPage/DeleteStudent")
                  
                  }
                 }
              
                
                
              }
             
            
           }
           else{
            alert("Wrong pass or id");
          }
        
       }
       xhttp.send();
      }
     catch (error) {
      alert("Wrong pass or id");
    }
  }
    
  
      
     
  
      
}

   
    
    
    const mystyle = {
      fontSize: "10px",
      fontFamily: "Arial",
      fontWeight: "200",
      width: "10vw"
    };
    
 
    
      return(
        <div>
          <ButtonContainer>
          <StyledInput  type="text"
            id="inpt" name="inpt" placeholder={t("EMAIL")} onChange={handleChangeinpt}
            value={inpt}  ></StyledInput>
          <FormatMail>
          <div className="dropdown">
            <div className="card">
                <Dropdown value={selectedExtension} options={extensions} onChange={onExtensionChange} optionLabel="name" placeholder={t("EXTENSION")}style={mystyle}/>                                
            </div>
        </div>
          </FormatMail>
        </ButtonContainer>
        <InputContainer>
        <StyledInputPassword  type="password"
            id="psw" name="psw" placeholder={t("PASSWORD")} onChange={handleChangepsw}
            value={pssw} ></StyledInputPassword>
        </InputContainer>
        <ButtonContainer>
          <Button  content={t("LOGIN_BTN") } onClick={handleClick}/>
        </ButtonContainer>
        </div>
            );
    
    
 
 
      
}
const StyledInputPassword = styled.input`
  background: rgba(255, 255, 255, 0.15);
  box-shadow: 0 8px 32px 0 rgba(31, 38, 135, 0.37);
  border-radius: 2rem;
  margin-left: 2.2rem;
  width: 100%;
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
export default LoginInputs;