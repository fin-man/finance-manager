import React, { useState,useEffect } from 'react';

import DatePicker from "react-datepicker";
import MainPage from './mainpage'
import './../Homepage.css'

function HomePage() {

    const firstDate = new Date()
    firstDate.setDate(firstDate.getDate() - 30)

    
    const [startDate, setStartDate] = useState(firstDate);
    const [endDate, setEndDate] = useState(new Date());
  

    return(
        <div >
            <div className="datetime-wrapper">
                <div className="datepicker-container">
                    <div className="datepicker-title">
                        Start Date :
                    </div>
                    <div className="datepicker">
                        <DatePicker selected={startDate} onChange={date => setStartDate(date)} dateFormat="yyyy-MM-dd" />
                    </div>
                </div>


                <div className="datepicker-container">
                    <div className="datepicker-title">
                        Start Date :
                    </div>
                    <div className="datepicker">
                        <DatePicker selected={endDate} onChange={date => setEndDate(date)} dateFormat="yyyy-MM-dd" />
                    </div>
                </div>
            </div>
           

            <div className="mainchart-wrapper">
                <MainPage startDate={startDate} endDate={endDate}/>
            </div>
        </div>
    )
}


export default HomePage;
