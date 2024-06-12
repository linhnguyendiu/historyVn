import React from "react";
import { Avatar, Button, Card, Space, Divider } from "antd";

const { Meta } = Card;

const PostCard: React.FC = () => (
  <Card
    style={{
      width: "810px",
      height: "200px",
      borderRadius: "10px",
      backgroundColor: '#FFFFFF'
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
      <Divider style={{ margin: "20px 0", borderColor: "rgba(0, 0, 0, 0.1)" }} />
      <div className="post-authors">
        <div className="img">
          <img src="./avatar.png" width="50px" height="50px" />
          <h4>Linh Nguyen</h4>
        </div>
        <Button>Detail</Button>
      </div>
    </div>
  </Card>
);

export default PostCard;
