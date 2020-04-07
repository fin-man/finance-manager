import React from 'react';
import Highcharts from 'highcharts';
import HighchartsReact from 'highcharts-react-official';
import PieChart from "highcharts-react-official";


function CategoriesChart(props) {

    const options = {
        chart: {
          type: "pie"
        },
        series: [
          {
            data: [
              {
                name: 'start',
                y: 100
              },
              {
                y: 50
              }
            ]
          }
        ]
    };
      

    console.log(props.data)

    return(
        <div>
             <PieChart highcharts={Highcharts} options={options} />
        </div>

    )

}

export default CategoriesChart;