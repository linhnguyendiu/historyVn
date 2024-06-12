import React from "react"
import HistorySection from "./comp/historySection"
interface Props { 

}
const HistoryContent: React.FC<Props> = () => {
    return ( 
        <div className="hs-section">
        <HistorySection/>
        <HistorySection/>
        </div>
    )
}
export default HistoryContent