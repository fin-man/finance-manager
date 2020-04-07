import React from 'react';
import axios from 'axios';


function CharacterDropDown() {
    const [items, setItems] = React.useState([]);
    const [pick, setPickItem] = React.useState("");
  

    const handleAddrTypeChange = (e) => setPickItem(e.target.value)

    React.useEffect(() => {
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
            <select onChange={e => handleAddrTypeChange(e)}>
             {items}
            </select>
            {pick}

        </div>
        
    );
  }



function CategoriesBody(){



    return(
        <div>

        </div>
    )
}

function CategoriesChart(){
    

    return(
        <div>

        </div>
    )
}

export default CharacterDropDown;
