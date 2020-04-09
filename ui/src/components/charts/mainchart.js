import React from 'react';
import Highcharts from 'highcharts';
import HighchartsReact from 'highcharts-react-official';

function MainChart(props) {

    const options = {
        chart: {
          height: 600,
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
          yAxis: {
            title:{
              text: 'Amount'
            },
          },
        
        series: [
          {
            name : "Dates",
            data:  props.data.data.graph_response.all_graph_data
          }
        ]
      };

  return (
    <div >
        <HighchartsReact highcharts={Highcharts} options={options} />
    </div>
  );
}

export default MainChart;
