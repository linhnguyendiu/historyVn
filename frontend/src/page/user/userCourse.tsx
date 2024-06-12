import React from "react";
import HistorySection from "../course/comp/historySection";
import { Button, Card, Space } from "antd";
interface Props {}


const Certificate = () => {
    return(
    <div className="certificate-wrapper">
        <div className="certificate">
            <div className="certificate-title">
                <span className="certificate-label">Chứng chỉ khóa học của bạn</span>
                <Button>Tải xuống</Button>
            </div>
            <div className="certificate-img">
                <img src="./certificate.png"/>
            </div>
        </div>
        <Card title="Số điểm đạt được" type="inner"
         className="inner-cert-card">
            <span>8/9</span>
        </Card>
    </div>
    )
}

const Course = () => {
  return (
    <Card title="Nhà Hồ" className="course-wrapper">
      <div className="course-text">
        <span>
        Nhà Hồ bắt đầu khi Hồ Quý Ly lên ngôi năm 1400 sau khi giành được quyền
        lực từ tay nhà Trần và chấm dứt khi Hồ Hán Thương bị quân Minh bắt vào
        năm 1407 – tổng cộng là 7 năm. Quốc hiệu Đại Việt đã đổi thành Đại Ngu
        năm 1400.
        </span>
      </div>
      <div className="course-img">
        <img
          src="./courseimg1.png"
          alt="Course 1"
          className="user-course-image"
        />
        <img
          src="./courseImg2.png"
          alt="Course 2"
          className="user-course-image"
        />
        <img
          src="./courseImg3.png"
          alt="Course 3"
          className="user-course-image"
        />
      </div>
      <div className="course-user-info">
        <Space>
          <Button>4 giờ học</Button>
          <Button>Hội</Button>
        </Space>
        <div className="price">Hoàn thành</div>
      </div>
      <Certificate/>
    </Card>
  );
};
const UserCourse: React.FC<Props> = () => {
  return (
    <div className="user-course-wrapper">
      <Course />
    </div>
  );
};

export default UserCourse;
