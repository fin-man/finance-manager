import React, { useState,useEffect } from 'react';
import axios from 'axios';
import MainChart from './charts/mainchart';
import FormatDate from './../utils/formatDates'
import CompareDates from './../utils/compareDates'
import "react-datepicker/dist/react-datepicker.css";

function MainPage(props) {
  const [data , setData] = useState({data: null});


  useEffect(()=> {

    async function fetchGraphData(){


        const response = await axios(BuildURL(FormatDate(props.startDate),FormatDate(props.endDate)));

        setData(response);
    }

    fetchGraphData();
  },[props.startDate , props.endDate]);


  if (CompareDates(props.startDate , props.endDate)){
    if (data.data !== null) {
        console.log("Data is tehre")
  
      
      return (
          <div >
              <MainChart data={data} startDate={FormatDate(props.startDate)} endDate={FormatDate(props.endDate)}  />
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
  }else{
      return(
          <div>
              sjsjsj
          </div>
      )
  }
  

}

function BuildURL(from , to) {

    return 'http://localhost:8080/transactions/graph?from='+from+'&to='+to
}


export default MainPage;
