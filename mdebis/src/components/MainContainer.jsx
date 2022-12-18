import styled from "styled-components";
import Icon from "./Icon";
import { FaFacebookF,FaLinkedin, FaInstagram, FaTwitter} from "react-icons/fa";
import { useTranslation } from "react-i18next";
import i18n from "../i18n";
import LoginInput from "./LoginInputs";
import ForgotPasswordd from "./ForgotPass";
import { useState } from "react";
import LoginInputs from "./LoginInputs";
import { Navigate, Link } from "react-router-dom";


const MainContainer =() =>{
  const [goForgotPassword, setForgotPassword] = useState(false);
  const [homePage, setHomePage] = useState(false);
  const languages = {
    en: {nativeName: "English 🇺🇸"},
    tr: {nativeName: "Türkçe 🇹🇷"}
  };
  const {t} = useTranslation();
  const FacebookBackground =
    "linear-gradient(to right, #0546A0 0%, #0546A0 40%, #663FB6 100%)";
  const InstagramBackground =
    "linear-gradient(to right, #A12AC4 0%, #ED586C 40%, #F0A853 100%)";
  const TwitterBackground =
    "linear-gradient(to right, #56C1E1 0%, #35A9CE 50%)";
    const LinkedinBackground =
    "linear-gradient(to right, #0546A0 0%, #0546A0 40%, #0e76a8 100%)";
    
const MainContain = styled.div`
  display: flex;
  align-items: center;
  flex-direction: column;
  height: 80vh;
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
    height: 80vh;
    
  }
  @media only screen and (min-width: 360px) {
    width: 80vw;
    height: 90vh;
    h4 {
      font-size: small;
    }
  }
  @media only screen and (min-width: 411px) {
    width: 80vw;
    height: 90vh;
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
    height: 72vh;
  }
  @media only screen and (min-width: 1600px) {
    width: 30vw;
    height: 65vh;
  }
`;

const WelcomeText = styled.h2`
  margin: 3rem 0 2rem 0;
`;


const LoginWith = styled.h5`
  text-align: center;
  cursor: pointer;
`;

const HorizontalRule = styled.hr`
  width: 90%;
  height: 0.3rem;
  border-radius: 0.8rem;
  border: none;
  background: linear-gradient(to right, #14163c 0%, #03217b 79%);
  background-color: #ebd0d0;
  margin: 1.5rem 0 1rem 0;
  backdrop-filter: blur(25px);
`;

const IconsContainer = styled.div`
  display: flex;
  justify-content: space-evenly;
  margin: 4rem 0 1rem 0;
  width: 80%;
`;

const ForgotPassword = styled.h4`
margin: 1rem 0 0rem 0;
  cursor: pointer;
`;
const FormatMail = styled.span`
  width: 35%;
  cursor: pointer;
  text-transform: lowercase;
  font-size: 6px;
  letter-spacing: 0rem;
  `;
  if(goForgotPassword){
    return <Navigate to="/ForgotPassword" />;
  }
  
  return (
    <MainContain>
    <WelcomeText>{t("WELCOME")}</WelcomeText>
    
    <LoginInput />
    <LoginWith >{t("UNDER_BTN_TXT")}</LoginWith>
    <HorizontalRule />
    <ForgotPassword onClick={() =>{setForgotPassword(true);}}>{t("FORGT_PSSW")}</ForgotPassword >
    <div>
    {Object.keys(languages).map((language) =>(
      <button type="submit" key = {language} onClick={()=> i18n.changeLanguage(language)} disabled={i18n.resolvedLanguage === language}>{languages[language].nativeName}</button>
    ))}
    </div>
    <IconsContainer>
      <Icon color={FacebookBackground}>
        <FaFacebookF />
      </Icon>
      <Icon color={InstagramBackground}>
        <FaInstagram />
      </Icon>
      <Icon color={LinkedinBackground}>
        <FaLinkedin />
      </Icon>
      <Icon color={TwitterBackground}>
        <FaTwitter />
      </Icon>
    </IconsContainer>
    </MainContain>
    
    
  );
  

};
  

export default MainContainer;