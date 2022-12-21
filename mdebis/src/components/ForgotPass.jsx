import styled from "styled-components";
import {useState} from 'react';
import { useTranslation } from "react-i18next";
import axios from "axios";
import DropDownn from "./DropDown";
import Button from "./Button";
import MainContainer from "./MainContainer";
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
const MainContain = styled.div`
  display: flex;
  align-items: center;
  flex-direction: column;
  height: 20vh;
  width: 30vw;
  background: rgba(255, 255, 255, 0.15);
  box-shadow: 0 8px 32px 0 rgba(31, 38, 135, 0.37);
  backdrop-filter: blur(8.5px);
  -webkit-backdrop-filter: blur(8.5px);
  border-radius: 10px;
  color: #ffffff;
  text-transform: uppercase;
  letter-spacing: 0.4rem;
  @media only screen and (max-width: 320px) {
    width: 80vw;
    height: 90vh;
    hr {
      margin-bottom: 0.3rem;
    }
    h4 {
      font-size: small;
    }
  }
  @media only screen and (max-height: 840px) {
    height: 20vh;
    
  }
  @media only screen and (min-width: 360px) {
    width: 80vw;
    height: 20vh;
    h4 {
      font-size: small;
    }
  }
  @media only screen and (min-width: 411px) {
    width: 80vw;
    height: 20vh;
  }
  @media only screen and (min-width: 768px) {
    width: 80vw;
    height: 80vh;
  }
  @media only screen and (min-width: 1024px) {
    width: 70vw;
    height: 70vh;
  }
  @media only screen and (min-width: 1280px) {
    width: 30vw;
    height: 20vh;
  }
  @media only screen and (min-width: 1600px) {
    width: 30vw;
    height: 20vh;
  }
`;

const ForgotPassword = ()=>{
  let endpoint = "http://localhost:8080";
  const {t} = useTranslation();
  const [inpt, setMessage] = useState('');
  const [pssw, setpssw] = useState('');
 
  const handleChangeinpt = event => {
    setMessage(event.target.value);

    console.log('value is:', event.target.value);
  };
  const handleChangepsw = event => {
    setpssw(event.target.value);

    console.log('value is:', event.target.value);
  };
  function handleClick() {
  //   try {
  //     var xhttp = new XMLHttpRequest();
  //     xhttp.open("GET", "http://localhost:8080/student_forgot/"+inpt, false);
  //     xhttp.setRequestHeader("Content-type", "text/html");
  //     xhttp.send();
  //     var response = JSON.parse(xhttp.response);
  //     console.log(response);
  //   } catch (error) {
  //     alert(error.message);
  // }
    // // axios
    // //   .put(endpoint + "/student_forgot/" + inpt, {
    // //     headers: {
    // //       "Content-Type": "application/x-www-form-urlencoded",
    // //     },
    // //   })
    // //   .then((res) => {
    // //     console.log(res);
    // //     this.getTask();
    // //   });
  }
  
  return(
    <MainContain>
      
      <ButtonContainer>
      <StyledInput  type="text"
        id="inpt" name="inpt" placeholder={t("EMAIL")} onChange={handleChangeinpt}
        value={inpt}  ></StyledInput>
      <FormatMail>
      <DropDownn placeholder={t("EXTENSION")}></DropDownn>
      </FormatMail>
    </ButtonContainer>
   
    <ButtonContainer>
      <Button  content={t("RESET") } onClick={handleClick}/>
    </ButtonContainer>

    

    </MainContain>

        
        );
      
}

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
    
  
  export default ForgotPassword;