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
            data:  props.data.data.graph_response.all_graph_data
          }
        ]
      };

  return (
    <div >
        Main Chart
        <HighchartsReact highcharts={Highcharts} options={options} />
    </div>
  );
}

export default MainChart;
