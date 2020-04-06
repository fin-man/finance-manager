import React, { useState,useEffect } from 'react';
import axios from 'axios';
import MainChart from './charts/mainchart'
function MainPage() {
  const [data , setData] = useState({data: null});


//   useEffect( async()=>{
//       const result = await axios(
//           'http://localhost:8080/transactions/graph'
//       );

//       setData(result.Data)
//   });

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
            Main Page
            <div>
            {data.data.all_transactions.took}
            </div>
            <MainChart data={data} />
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
