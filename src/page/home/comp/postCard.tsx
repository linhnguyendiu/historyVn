import React from "react";
import { Avatar, Button, Card, Space, Divider } from "antd";


const { Meta } = Card;

const PostCard: React.FC = () => (
  <Card
    style={{
      width: "auto",
      height: "auto",
      borderRadius: "10px",
    }}
  >
    <div className="post-card">
      <div className="post-excerpt">
        <span>
          “Cái chết bí ẩn của Hoài Văn Hầu Trần Quốc Toản”. Nhiều sách sử của
          Việt Nam đều không đề cập chi tiết tới cái chết của Trần Quốc Toản,
          thời gian ông mất cũng chưa được thống nhất.
        </span>
      </div>
      <Divider/>
      <div className="post-authors">
        <div className="img">
          <img src="./avatar.png" width='50px' height='50px'/> <span>Linh Nguyen</span>
        </div>
        <Button>Detail</Button>
      </div>
    </div>
  </Card>
);

export default PostCard;
