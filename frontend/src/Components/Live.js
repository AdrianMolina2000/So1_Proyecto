import NavBar from "./Nabvar";
import Combobox from "react-widgets/Combobox";
import { useEffect, useState } from "react";
import ProgressBar from 'react-bootstrap/ProgressBar';
import axios from "axios";

const fases=[{id:1,tipo:"Octavos de final"},{id:2,tipo:"Cuartos de final"},{id:3,tipo:"Semifinal"},{id:4,tipo:"Final"}];
const octavos=[{id:1,tipo:"Paises Bajos vs España"},{id:2,tipo:"Inglaterra vs Croacia"},{id:3,tipo:"Argentina vs Suiza"},{id:4,tipo:"Francia vs Uruguay"}
,{id:5,tipo:"Alemania vs Ecuador"},{id:6,tipo:"Belgica vs EE.UU"},{id:7,tipo:"Brasil vs Polonia"},{id:8,tipo:"Portugal vs Dinamarca"}];
const cuartos=[{id:1,tipo:"España vs Alemania"},{id:2,tipo:"Inglaterra vs Belgica"},{id:3,tipo:"Argentina vs Brasil"},{id:4,tipo:"Francia vs Portugal"}];
const semifinal=[{id:1,tipo:"España vs Inglaterra"},{id:2,tipo:"Argentina vs Francia"}];
const final=[{id:1,tipo:"España vs Argentina"}];


const Live=()=> {
    const [seleccion, setSeleccion] = useState([]);
    const [cache, setCache] = useState([]);
    const [FiltroPhase, setFiltroPhase] = useState([]);
    const [FiltroMatch, setFiltroMatch] = useState([]);
    const [FiltroResult, setFiltroResult] = useState();
    const [TOTAL, setTOTAL] = useState();
    const getDatos = async () => {
        try{
          await axios.get(`http://20.232.49.78.nip.io/output/getPartidosRedis/`)        
              .then(function(response){             
                 return response;
              }).then( function (response) {
                console.log(response.data)
                setCache(response.data)
                
                 
              })       
              .catch(function(error) { 
    
                 console.log('error catch', (error))
              });
         }catch(error) { 
           console.log('error en datos', error)
        }
     } 
     let haveData = false;
     useEffect(()=>{
         if(!haveData) {
             getDatos();
             haveData = true;
         }
         const interval=setInterval(()=>{
             getDatos();
         },10000)
         return()=>clearInterval(interval)
     },[])

    function obtenerPartidos(value){
        console.log(value);
        var partidos=[];
        if(value.id===1){
            setSeleccion(octavos);

            for(let i=0;i<cache.length;i++){
                var aux = JSON.parse(cache[i]);
                if(aux.phase==1){
                    //console.log(aux);
                    partidos.push(aux);
                }
            }
            
        }else if(value.id===2){
            setSeleccion(cuartos);
            for(let i=0;i<cache.length;i++){
                var aux = JSON.parse(cache[i]);
                if(aux.phase==2){
                    //console.log(aux);
                    partidos.push(aux);
                }
            }
            
        }else if(value.id===3){
            setSeleccion(semifinal);
            for(let i=0;i<cache.length;i++){
                var aux = JSON.parse(cache[i]);
                if(aux.phase==3){
                    //console.log(aux);
                    partidos.push(aux);
                }
            }
            
        }else if(value.id===4){
            setSeleccion(final);
            for(let i=0;i<cache.length;i++){
                var aux = JSON.parse(cache[i]);
                if(aux.phase==4){
                    //console.log(aux);
                    partidos.push(aux);
                }
            }
           
        }

        setFiltroPhase(partidos);
        //console.log(FiltroPhase);
    }
    function verResultado(value){
        console.log(value);
        var partidos=[];
        console.log(FiltroPhase);
        if(value.tipo==='Argentina vs Suiza'){
            for(let i=0;i<FiltroPhase.length;i++){
                
                if(FiltroPhase[i].team1=='Argentina'){
                    //console.log(aux);
                    partidos.push(FiltroPhase[i].score);
                }
            }
        }else if(value.tipo==='Paises Bajos vs España'){
            for(let i=0;i<FiltroPhase.length;i++){
                
                if(FiltroPhase[i].team1=='Paises Bajos'){
                    //console.log(aux);
                    partidos.push(FiltroPhase[i].score);
                }
            }
        }else if(value.tipo==='Inglaterra vs Croacia'){
            for(let i=0;i<FiltroPhase.length;i++){
                
                if(FiltroPhase[i].team1=='Inglaterra'){
                    //console.log(aux);
                    partidos.push(FiltroPhase[i].score);
                }
            }
        }else if(value.tipo==='Francia vs Uruguay'){
            for(let i=0;i<FiltroPhase.length;i++){
                
                if(FiltroPhase[i].team1=='Francia'){
                    //console.log(aux);
                    partidos.push(FiltroPhase[i].score);
                }
            }
        }else if(value.tipo==='Alemania vs Ecuador'){
            for(let i=0;i<FiltroPhase.length;i++){
                
                if(FiltroPhase[i].team1=='Alemania'){
                    //console.log(aux);
                    partidos.push(FiltroPhase[i].score);
                }
            }
        }else if(value.tipo==='Belgica vs EE.UU'){
            for(let i=0;i<FiltroPhase.length;i++){
                
                if(FiltroPhase[i].team1=='Belgica'){
                    //console.log(aux);
                    partidos.push(FiltroPhase[i].score);
                }
            }
        }else if(value.tipo==='Brasil vs Polonia'){
            for(let i=0;i<FiltroPhase.length;i++){
                
                if(FiltroPhase[i].team1=='Brasil'){
                    //console.log(aux);
                    partidos.push(FiltroPhase[i].score);
                }
            }
        }else if(value.tipo==='España vs Alemania'){
            for(let i=0;i<FiltroPhase.length;i++){
                
                if(FiltroPhase[i].team1=='España'){
                    //console.log(aux);
                    partidos.push(FiltroPhase[i].score);
                }
            }
        }else if(value.tipo==='Inglaterra vs Belgica'){
            for(let i=0;i<FiltroPhase.length;i++){
                
                if(FiltroPhase[i].team1=='Inglaterra'){
                    //console.log(aux);
                    partidos.push(FiltroPhase[i].score);
                }
            }
        }else if(value.tipo==='Argentina vs Brasil'){
            for(let i=0;i<FiltroPhase.length;i++){
                
                if(FiltroPhase[i].team1=='Argentina'){
                    //console.log(aux);
                    partidos.push(FiltroPhase[i].score);
                }
            }
        }else if(value.tipo==='Francia vs Portugal'){
            for(let i=0;i<FiltroPhase.length;i++){
                
                if(FiltroPhase[i].team1=='Francia'){
                    //console.log(aux);
                    partidos.push(FiltroPhase[i].score);
                }
            }
        }else if(value.tipo==='España vs Inglaterra'){
            for(let i=0;i<FiltroPhase.length;i++){
                
                if(FiltroPhase[i].team1=='España'){
                    //console.log(aux);
                    partidos.push(FiltroPhase[i].score);
                }
            }
        }else if(value.tipo==='Argentina vs Francia'){
            for(let i=0;i<FiltroPhase.length;i++){
                
                if(FiltroPhase[i].team1=='Argentina'){
                    //console.log(aux);
                    partidos.push(FiltroPhase[i].score);
                }
            }
        }else if(value.tipo==='España vs Argentina'){
            for(let i=0;i<FiltroPhase.length;i++){
                
                if(FiltroPhase[i].team1=='España'){
                    //console.log(aux);
                    partidos.push(FiltroPhase[i].score);
                }
            }
        }
        
        const resultado = partidos.reduce((prev, cur) => ((prev[cur] = prev[cur] + 1 || 1), prev), {})
        console.log(resultado);
        console.log(resultado['2-1']);
        let result = partidos.filter((item,index)=>{
            return partidos.indexOf(item) === index;
          })
          console.log(result);
          setTOTAL(result.length);
        setFiltroMatch(result);
        setFiltroResult(resultado);
    }
    return (

        <div align="center">
            <NavBar/>
            <h1 style={{color:"white"}} >Live</h1>
            <Combobox
                data={fases}
                dataKey='id'
                textField='tipo'
                onChange={(value) => obtenerPartidos(value)}
                style={{width: 300 ,color:"black"}}
            /><br/>
            <Combobox
                data={seleccion}
                dataKey='id'
                textField='tipo'
                onChange={(value) => verResultado(value)}
                style={{width: 300}}
            />
            <br/>
            <br/>
            
            {FiltroMatch.map((dato) => (
                <>
                <ProgressBar striped variant="info" animated now={FiltroResult[dato]} max={TOTAL} min="0" style={{ width: "500px" }} />
                <h2 style={{ color: "white" }}> {dato}</h2>
                </>
              ))}
        </div>
    );
}
export default Live;