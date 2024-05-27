import React from "react";
import {
    ClockCircleOutlined
  } from '@ant-design/icons';
import { Button } from "antd";  
import { useNavigate,useParams } from "react-router-dom";

interface Props {}
const CharaterCard: React.FC<Props> = () => {
    const navigate = useNavigate();
    const param = useParams()
  return (
    <Button className="story-card-wrapper" onClick={() =>navigate(`/course/nha-ho/test`) }>
      <div className="story-title">
        <div className="title">
            <h3>Thời thơ ấu</h3>
        </div>
        <div className="lesson-order">
            <span>Bài học số 1</span>
        </div>
      </div>
      <div className="duration">
        <ClockCircleOutlined/>&nbsp;10 minutes
      </div>
    </Button>
    
  );
};
export default CharaterCard;
