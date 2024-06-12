import React from "react";
import './index.css';
import Breadcrumbb from "../../component/breadcrumb";
import TitleLesson from "./titleLesson";
import { Divider } from "antd";
import ItemCard from "./item_Card";
import Item from "antd/es/list/Item";
interface Props { 

}
const LessonCourse:React.FC<Props> = () => { 
    return ( 
     <div className="content-lesson-wrapper">
        <Breadcrumbb/>
        <TitleLesson/>
        <Divider style={{ margin: "20px 0", borderColor: "rgba(0, 0, 0, 0.1)" }} />
        <div className="lesson-bigthumb">
            <img src="/nhaho.png"/>
        </div>
        <ItemCard/>
        <ItemCard/>
    </div>
    )
}
export default LessonCourse
