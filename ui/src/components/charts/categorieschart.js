import React, { useState } from 'react';
import Highcharts from 'highcharts';
import PieChart from "highcharts-react-official";


function CategoriesChart(props) {

    const [initialData, setInitialData] = useState([]);

    const options = {
        chart: {
          type: "pie"
        },
        series: [
          {
            data: MassageData(props.data)
          }
        ]
    };
      
    return(
        <div>
             <PieChart highcharts={Highcharts} options={options} />
        </div>

    )

}


function MassageData(data){

    var pieData = [];

    for (var key in data) {
        console.log(key)
        console.log(data[key])

        let total = 0
        for (var s in data[key]) {
            total += data[key][s][1]
        }

        pieData.push({
            name: key,
            y : total
        })
      }
      

      return pieData


}

export default CategoriesChart;
