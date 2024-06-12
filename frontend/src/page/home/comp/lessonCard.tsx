import React from "react";
import { Avatar, Button, Card, Space } from "antd";

const { Meta } = Card;

const LessonCard: React.FC = () => (
  <Card
    style={{
      width: "390px",
      height: "590px",
      backgroundColor: "#ffffff",
      borderRadius: "10px",
    }}
  >
    <div>
      <img
        alt="example"
        src="./lesson.png"
        style={{ width: "100%", height: "auto", borderRadius: "10px" }}
      />
      <div className="detail-lesson-type">
        <Space className="type-lesson" direction="horizontal" size="large">
          <Button ghost>4 giờ đọc</Button>
          <Button ghost>Đình</Button>
        </Space>
        <div className="price-lesson">
          <Button ghost>Free</Button>
        </div>
      </div>
      <div className="lesson-card-content">
        <h3>Nhà Tiền Lý</h3>
        <span >
          Một triều đại trong lịch sử Việt Nam, gắn liền với quốc hiệu Vạn Xuân.
          Nhà Tiền Lý kéo dài 61 năm và nước Vạn Xuân kéo dài 58 năm, tổng cộng
          ba đời vua.
        </span>
      </div>
    </div>
    <Button block>Học ngay</Button>

  </Card>
);

export default LessonCard;
