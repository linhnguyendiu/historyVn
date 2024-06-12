import React from "react";
import { Button, Card, Space } from "antd";

interface Props {}
const HistorySection: React.FC<Props> = () => {
  return (
    <Card title="Nhà Hồ" className="hs-wrapper">
      <div className="row-1">
        <div className="description">
          Nhà Hồ bắt đầu khi Hồ Quý Ly lên ngôi năm 1400 sau khi giành được
          quyền lực từ tay nhà Trần và chấm dứt khi Hồ Hán Thương bị
          quân Minh bắt vào năm 1407 – tổng cộng là 7 năm. Quốc hiệu Đại Việt đã
          đổi thành Đại Ngu năm 1400.
        </div>
        <div>
          <Button>Học ngay</Button>
        </div>
      </div>
      <div className="row-2">
        <img src="./courseimg1.png" alt="Course 1" className="course-image" />
        <img src="./courseImg2.png" alt="Course 2" className="course-image" />
        <img src="./courseImg3.png" alt="Course 3" className="course-image" />
      </div>
      <div className="row-3">
        <Space>
          <Button>4 giờ học</Button>
          <Button>Hội</Button>
          <Button>Thu thập 2LH</Button>
          <Button>Thời kỳ xây dựng nền tự chủ </Button>
        </Space>
        <div className="price">Free</div>
      </div>
      <Card type="inner" title="Nội dung" style={{ backgroundColor: 'white', marginTop: '20px'}}>
        <div className="inner-wrapper">
          <div className="content-element">
            <div className="number">01</div>
            <div>Tổng quan về nhà Hồ</div>
          </div>
          <div className="content-element">
            <div className="number">02</div>
            <div>Tổng quan về nhà Hồ</div>
          </div>
          <div className="content-element">
            <div className="number">03</div>
            <div>Tổng quan về nhà Hồ</div>
          </div>
          <div className="content-element">
            <div className="number">04</div>
            <div>Tổng quan về nhà Hồ</div>
          </div>
        </div>
      </Card>
    </Card>
  );
};

export default HistorySection;
