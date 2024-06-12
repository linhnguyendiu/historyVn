import React from "react";
interface Props {}
const TitleLesson: React.FC<Props> = () => {
  return (
    <div className="title-wrapper">
      <div className="title">
        <span>Nhà Hồ</span>
      </div>
      <div className="title-description">
        <span>
          Nhà Hồ (1400-1407): quốc hiệu Đại Ngu, kinh đô Tây Đô (Thanh Hóa) Nhà
          Hồ bắt đầu khi Hồ Quý Ly lên ngôi năm 1400 sau khi giành được quyền
          lực từ tay nhà Trần và chấm dứt khi Hồ Hán Thương bị quân Minh bắt vào
          năm 1407 – tổng cộng là 7 năm. Quốc hiệu Đại Việt đã đổi thành Đại
          Ngu năm 1400.
        </span>
      </div>
    </div>
  );
};

export default TitleLesson;
