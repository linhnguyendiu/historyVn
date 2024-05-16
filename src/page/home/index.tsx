import React from "react";
import "./index.css";
import { Typography } from "antd";
import SearchComp from "./comp/search";
import Lesson from "./lesson";
import Post from "./post";
interface Props {}

const HomePage: React.FC<Props> = () => {
  return (
    <div className="wrapper">
      <div className="big-thumb">
        <div className="big-thumb-content">
          <h1 className="bona-nova-regular-italic">
            “Dân ta phải biết sử ta, cho tường gốc tích nước nhà Việt Nam. Sử ta
            dạy cho ta những chuyện vẻ vang của tổ tiên ta. Dân tộc ta là con
            Rồng cháu Tiên, có nhiều người tài giỏi đánh Bắc dẹp Nam, yên dân
            trị nước tiếng để muôn đời. Sử ta dạy cho ta bài học này: Lúc nào
            dân ta đoàn kết muôn người như một thì nước ta độc lập, tự do. Trái
            lại lúc nào dân ta không đoàn kết thì bị nước ngoài xâm lấn.”
          </h1>
          <h1 style={{ float:"right"}}>- Hồ Chí Minh -</h1>
        </div>
        <div className="img-dragon" style={{ marginLeft: '180px'}}>
          <img src="./dragon.png" width='750px'height='500px' />
        </div>
      </div>
      <div className="content-home">
            <SearchComp className="search-comp"/>
            <Lesson/>
            <Post/>
      </div>
    </div>
  );
};

export default HomePage;
