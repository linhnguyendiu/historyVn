import React from "react";
import ArticleSection from "./article";
import "./index.css";
import { Divider } from "antd";
interface Props {}
const Article: React.FC<Props> = () => {
  return (
    <div className="article-wrapper">
      <ArticleSection></ArticleSection>
      <Divider
        orientation="left"
        style={{ margin: "20px 0", borderColor: "#FFC0CB" }}
      >
        Bình luận (2)
      </Divider>
    </div>
  );
};

export default Article;
