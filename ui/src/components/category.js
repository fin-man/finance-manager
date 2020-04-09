import React, { useState,useEffect } from 'react';
import axios from 'axios';
import Highcharts from 'highcharts';
import HighchartsReact from 'highcharts-react-official';


function CharacterDropDown(props) {
    const [items, setItems] = useState([]);
    const [pick, setPickItem] = useState("");
  

    const handleAddrTypeChange = (e) => setPickItem(e.target.value)

    useEffect(() => {
      async function getCharacters() {
        const response = await axios("http://localhost:8080/categories");

        var arr = [];

        for (let i = 1; i <= response.data.length; i++) {
            if (typeof response.data[i] !== 'undefined') {
            arr.push(<option key={i} value={response.data[i]} >{response.data[i]} </option>)
            }
        }
        
        setItems(arr);
      }
      getCharacters();

    }, []);

    return (

        <div>
            <div className="categories-dropdown">
                <select onChange={e => handleAddrTypeChange(e)}>
                {items}
                </select>

            </div>

            <div>
                <CategoriesLineChart data={props.data} pick={pick} />
            </div>
        </div>
 
        
    );
  }


function CategoriesLineChart(props){
    
    const options = {
        chart: {
          type: 'column',
          zoomType: 'x'

        },
        title: {
          text: 'My chart'
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
            <div >
            Main Chart
            <HighchartsReact highcharts={Highcharts} options={options} />
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
