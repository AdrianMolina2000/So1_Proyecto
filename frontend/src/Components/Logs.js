import NavBar from "./Nabvar";
import axios from 'axios';
import { useEffect, useState } from 'react';
import  './style.css';
const Logs=()=>{

   const [records, setRecords] = useState([]);
   const[logs, setLogs] = useState([]);
    
   const [fila1, setFila1] = useState([]);
    var fila2 = [];
    const getDatos = async () => {
        try{
          await axios.get(`http://20.232.49.78.nip.io/output/getPartidosMongo/`)        
              .then(function(response){             
                 return response;
              }).then( function (response) {
                console.log(response.data)
                 setLogs(response.data.results)
                setRecords(response.data.record) 
              })       
              .catch(function(error) { 
    
                 console.log('error catch', (error))
              });
         }catch(error) { 
           console.log('error en datos', error)
        }
     } 
     let haveData = false;
    useEffect( () => {
       getDatos();
        
    }, []);
    
  
    //console.log('logs', logs)

    
    return (
        
        <div>
            
            <NavBar/>
            <h1 style={{marginLeft:"800px",color:"white"}} >Logs</h1>
            
           
            <div style={{width:"1600px", padding:"25px, 0", margin:"100px",border:"white"}}>
           
                {
                    
               logs.map((dato) => (
                    <div id="p" style={{color:"aqua", height:"200px", fontSize:"20px",display:"inline-block",marginLeft:"80px"}}><strong> <pre>{JSON.stringify(dato,null," ").replace(" ","\n")}</pre></strong> </div>
                ))
           }
            
            </div>
            <div style={{width:"1600px", padding:"25px, 0", margin:"100px",border:"white", }}>
          
           <div id="p" style={{width:"400px", height:"200px", marginLeft:"90px"}}>

           <p style={{marginTop:"20px", fontSize:"45px" , color:"white"}}>
           <strong> <pre> TOTAL RECORDS:</pre></strong>
           <strong style={{color:"black"}}> <pre>       {records} </pre></strong>
           </p>
             
           </div>
            
            </div>
        </div>
    );
}
export default Logs;