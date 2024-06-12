import React from "react";

interface Props {}

const UserAvatar = ({ src }: { src: string }) => {
  return (
    <div className="avatar-user">
      <img src={src} />
    </div>
  );
};

const UserInfo = () => {
  return (
    <div className="user-info">
      <ul className="info">
        <li>Linh Nguyen</li>
        <li>Địa chỉ ví: 0x00000000fjdijfsjgmkbd00239</li>
        <li>Phần thưởng đã nhận: 100 LH</li>
        <li>Danh hiệu: Thám Hoa</li>
        <li>Hạng 1/1 người học</li>
      </ul>
    </div>
  );
};

const UserOveral: React.FC<Props> = () => {
  return (
    <>
      <div className="user-info-wrapper">
        <UserAvatar src="./user_avatar.png" />
        <UserInfo></UserInfo>
      </div>
    </>
  );
};

export default UserOveral;
