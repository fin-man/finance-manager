import React, { useState,useEffect } from 'react';

function MainScoreCard(props){

    const [totalTransactions, setTotalTransactions] = useState(0.0);
    const [totalAmount, setTotalAmount] = useState(0.0);


    useEffect(()=> {

        function CalculateTotalTransactions(){
            setTotalTransactions(props.data.length)
        }


        function CalculateTotalAmount(){
            var total = 0.0 
            if (typeof props.data !== undefined){
                for(var v in props.data){
                    total += props.data[v][1]
                }
            }

            setTotalAmount(total)
        }

        CalculateTotalTransactions()
        CalculateTotalAmount()
    },[props.data])
    return(
        <div className="mainscorecard-wrapper">
            <div className="mainscorecard-total-transactions-wrapper">
                <div className="mainscorecard-total-transactions-title">
                    Total Transactions :
                </div>
                <div className="mainscorecard-total-transactions-value">
                    {totalTransactions}
                </div>
            </div>
            <div className="mainscorecard-total-wrapper">
                <div className="mainscorecard-total-title">
                    Total Expenditures :  
                </div>
                <div className="mainscorecard-total-value">
                    {totalAmount}
                </div>
            </div>
        </div>
    )
}


export default MainScoreCard