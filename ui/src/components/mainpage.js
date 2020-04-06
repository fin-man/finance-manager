import React, { useState,useEffect } from 'react';
import axios from 'axios';
import MainChart from './charts/mainchart';
import DatePicker from "react-datepicker";

import "react-datepicker/dist/react-datepicker.css";

function MainPage() {
  const [data , setData] = useState({data: null});
  const [startDate, setStartDate] = useState(new Date());

  useEffect(()=> {

    async function fetchGraphData(){


        const response = await axios('http://localhost:8080/transactions/graph');

        setData(response);
    }

    fetchGraphData();
  },[]);

  if (data.data !== null) {
      console.log("Data is tehre")
    return (
        <div >
            <DatePicker selected={startDate} onChange={date => setStartDate(date)} dateFormat="yyyy-MM-dd" />
      
            
            <MainChart data={data} startDate={startDate}  />
        </div>
      );
  }else{
      console.log("data still no there")
      return (
          <div>
              Blargh
          </div>
      )
  }

}

export default MainPage;
