import React, { useState,useEffect } from 'react';

import DatePicker from "react-datepicker";
import MainPage from './mainpage'
function HomePage() {

    const [startDate, setStartDate] = useState(new Date());
    const [endDate, setEndDate] = useState(new Date());
  

    return(
        <div >
            <div>
                Start Date : <DatePicker selected={startDate} onChange={date => setStartDate(date)} dateFormat="yyyy-MM-dd" />
            </div>

            <div>
                End Date : <DatePicker selected={endDate} onChange={date => setEndDate(date)} dateFormat="yyyy-MM-dd" />
            </div>


            <MainPage startDate={startDate} endDate={endDate}/>
        </div>
    )
}


export default HomePage;
