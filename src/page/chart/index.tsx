import React from "react";
import './index.css';
import Rank from "../../component/rank";
interface Props{ 

}
const Chart: React.FC<Props> = () => {
    return ( 
        <div className="chart-wrapper"> 
            <Rank/>
        </div>
    )
}
export default Chart