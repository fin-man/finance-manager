import React, { useState,useEffect } from 'react';
import Highcharts from 'highcharts';
import HighchartsReact from 'highcharts-react-official';
import SimpleTable from './table'
import CategoryScoreCard from './categoryscorecard'

function CharacterDropDown(props) {
    const [items, setItems] = useState([]);
    const [pick, setPickItem] = useState("");
  

    const handleAddrTypeChange = (e) => setPickItem(e.target.value)

    useEffect(() => {
      var arr = []

      function RunThis(){
        if (typeof props.data !== undefined) {
            for(var v in props.data){
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
       
          <div>
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
                  <div className="category-chart-wrapper">
                    <div className="category-linechart-wrapper">
                      <CategoriesLineChart data={props.data} pick={pick} />
                    </div>
                    <div className="category-scorecard-wrapper">
                      <CategoryScoreCard data={props.data[pick]}/>
                    </div>
                  </div>
               
              </div>

              <div className="table-wrapper">
               <SimpleTable data={props.detailed[pick]}/>
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
