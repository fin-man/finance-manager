import React, { useState,useEffect } from 'react';
import './../Homepage.css'

function CategoryScoreCard(props){

    const [totalCategoryTransactions, setTotalCategoryTransactions] = useState(0.0);
    const [totalCategoryAmount, setTotalCategoryAmount] = useState(0.0);


    useEffect(()=> {

        function CalculateTotalTransactions(){
            setTotalCategoryTransactions(props.data.length)
        }


        function CalculateTotalAmount(){
            var total = 0.0 
            if (typeof props.data !== undefined){
                for(var v in props.data){
                    total += props.data[v][1]
                }
            }

            setTotalCategoryAmount(total)
        }

        CalculateTotalTransactions()
        CalculateTotalAmount()
    },[props.data])

    return(
        <div className="categoryscorecard-wrapper">
            <div className="categoryscorecard-total-transactions-wrapper">
                <div className="categoryscorecard-total-transactions-title">
                    Total Transactions :
                </div>
                <div className="categoryscorecard-total-transactions-value">
                    {totalCategoryTransactions}
                </div>
            </div>
            <div className="categoryscorecard-total-wrapper">
                <div className="categoryscorecard-total-title">
                    Total Expenditures :  
                </div>
                <div className="categoryscorecard-total-value">
                    {totalCategoryAmount}
                </div>
            </div>
          
        </div>
    )
}


export default CategoryScoreCard