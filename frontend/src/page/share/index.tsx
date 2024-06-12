import React from "react";
import "./index.css";
import { EditOutlined } from "@ant-design/icons";
import { Button } from "antd";
import PostList from "./postList";
interface Props {}

const SharePage: React.FC<Props> = () => {
  return (
    <div className="sharepage-wrapper">
      <div className="sharepage-header">
        <div className="sharepage-title">
            <span>Bài viết nổi bật</span>
        </div>
        <div className="sharepage-icon">
            <Button icon = {<EditOutlined/>} size="large"/>
        </div>
      </div>
        <PostList></PostList>
        <PostList></PostList>
        <PostList></PostList>
        <PostList></PostList>
    </div>
  );
};

export default SharePage;
