import React, { useState} from 'react';
import 'primeicons/primeicons.css';
import 'primereact/resources/themes/lara-light-indigo/theme.css';
import 'primereact/resources/primereact.css';
import { Dropdown } from 'primereact/dropdown';
import { useTranslation } from "react-i18next";

const DropdownExtens = ({extensions}) => {
    const {t} = useTranslation();
    const [selectedExtension, setSelectedExtension] = useState(null);
    
    const onExtensionChange = (e) => {
        setSelectedExtension(e.value);
    }
    const mystyle = {
      fontSize: "10px",
      fontFamily: "Arial",
      fontWeight: "200",
      width: "10vw"
    };
    
    return (
        <div className="dropdown">
            <div className="card">
                <Dropdown value={selectedExtension} options={extensions} onChange={onExtensionChange} optionLabel="name" placeholder={t("EXTENSION")}style={mystyle}/>                                
            </div>
        </div>
    );
}

export default function Dropdownn({ placeholder }) {
    return <DropdownExtens placeholder={placeholder} />;
  }