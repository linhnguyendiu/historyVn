import React from "react";
import UserOveral from "./user";
import "./index.css";
import UserCourse from "./userCourse";
import { Divider } from "antd";
interface Props {}

const User: React.FC<Props> = () => {
  return (
    <div className="user-wrapper">
      <div className="user-title">
        <span>Thông tin tài khoản</span>
      </div>
      <UserOveral></UserOveral>
      <Divider
        style={{ margin: "20px 0", borderColor: "rgba(0, 0, 0, 0.1)" }}
      />
      <div className="user-title">
        <span>Khóa học của bạn</span>
      </div>
      <UserCourse/>
    </div>
  );
};
export default User;
