import SearchComp from "../../component/search"
import React from "react"
import './index.css'
import HistoryPhase from "./historyPhaseList"
import HistoryContent from "./historyContent"
interface Props{}

const CoursePage:React.FC<Props> = () => {
    return ( 
        <div className="content-course-wrapper">
            <SearchComp/>
            <div className="content-course">
                <HistoryPhase></HistoryPhase>
                <HistoryContent></HistoryContent>
            </div>
        </div>
    )
}
export default CoursePage