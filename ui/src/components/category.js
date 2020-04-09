import React, { useState,useEffect } from 'react';
import axios from 'axios';
import Highcharts from 'highcharts';
import HighchartsReact from 'highcharts-react-official';
import DetailedTable from './detailedtable'

function CharacterDropDown(props) {
    const [items, setItems] = useState([]);
    const [pick, setPickItem] = useState("");
  

    const handleAddrTypeChange = (e) => setPickItem(e.target.value)

    useEffect(() => {
      var arr = []

      function RunThis(){
        console.log("New data set ---- ")
        if (typeof props.data !== undefined) {
            for(var v in props.data){
              console.log(v)
             arr.push(<option className="dropdown-option" key={v} value={v} >{v} </option>)
          }
        }
        setItems(arr)
      }
      RunThis()
    }, [props.data]);

    if (pick === ""){
      return (
        <div>
          <div className="categories-wrapper">
              <div className="categories-title">
                Categories : 
              </div>
              <div className="categories-dropdown">
                  <select className="dropdown-select" onChange={e => handleAddrTypeChange(e)}>
                  {items}
                  </select>
              </div>
          </div>
          </div>
      );
    }else{
      return (
        <div>
          <div className="categories-wrapper">
              <div className="categories-title">
                Categories : 
              </div>
              <div className="categories-dropdown">
                  <select className="dropdown-select" onChange={e => handleAddrTypeChange(e)}>
                  {items}
                  </select>
              </div>
          </div>
          <div>
              <div>
                  <CategoriesLineChart data={props.data} pick={pick} />
                  <DetailedTable data={props.data[pick]}/>
              </div>
          </div>
          </div>
      );
    }

  }


function CategoriesLineChart(props){
    
    const options = {
        chart: {
          height: 500,
          type: 'column',
          zoomType: 'x'

        },

        title: {
          text: ''
        },
        xAxis: {
            type: 'datetime',
            labels: {
              format: '{value:%Y-%b-%e}'
            },
          },
        series: [
          {
            data:  props.data[props.pick]
          }
        ]
      };

    if (typeof props.pick !== undefined){

        return(
          <div>
          <div className="categories-line">
              <HighchartsReact highcharts={Highcharts} options={options} />
            </div>    
          </div>
           
        )
    }else{
        return(
            <div>
                Nope
            </div>
        )
    }
 
}

export default CharacterDropDown;
