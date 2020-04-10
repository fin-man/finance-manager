import React, { useState,useEffect } from 'react';
import axios from 'axios';
import MainChart from './charts/mainchart';
import CharacterDropDown from './category'
import CategoriesChart from './charts/categorieschart'
import FormatDate from './../utils/formatDates'
import CompareDates from './../utils/compareDates'
import MainScoreCard from './mainscorecard'
import "react-datepicker/dist/react-datepicker.css";
import './../Homepage.css'

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
  
      
      return (
          <div>
              
                <MainScoreCard data={data.data.graph_response.all_graph_data} />
                
                <div className="main-chart-wrapper">
                    <div className="main-chart" >
                        <MainChart data={data} startDate={FormatDate(props.startDate)} endDate={FormatDate(props.endDate)}  />
                    </div>
                    <div className="category-chart">
                        <CategoriesChart data={data.data.graph_response.CategoryMap} />

        
           
                    </div>
                </div>
                
                <div className="categories-drilldown">
                    <CharacterDropDown data={data.data.graph_response.CategoryMap} detailed={data.data.graph_response.category_map_detailed}/>
                </div>
          </div>
         
        );
    }else{
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
