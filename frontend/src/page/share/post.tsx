import React from "react";
// import './index.css'
import { DislikeTwoTone ,LikeTwoTone, ExclamationCircleTwoTone, MessageTwoTone  } from "@ant-design/icons";
import { Space } from "antd";
import { useNavigate } from "react-router-dom";

interface Props {}

const Post: React.FC<Props> = () => {
    const navigate = useNavigate();

  return (
    <button className="post-wrapper" onClick={() => navigate('/')}>
      <div className="thumbnail-post">
        <img src="./thumbnailPost.png" />
      </div>
      <div className="post-section">
        <div className="post-title">
          <span>Chuyện về vị vua khiến nhà Mạc thất thủ</span>
        </div>
        <div className="post-time">Tháng 02, 28, 2024</div>
        <div className="post-content">
          <span>
            Mạc Toàn lên ngôi nhưng tại vị chỉ được 2 tháng, tổng cộng thời gian
            tồn tại chính thức của triều đại là gần 66 năm đứng chân nơi đất
            Thăng Long, sau bại trận phải lui về Cao Bằng. Bởi vậy, Mạc Mậu Hợp
            được xem là vị vua cuối cùng của Mạc triều. Một trong những sự kiện
            tồn nghi rất lớn đối với Mạc Mậu Hợp là việc bị đánh giá “buông
            tuồng du đãng, tửu sắc bừa bãi”, mưu cướp vợ bề tôi. Lịch sử ghi
            nhận tướng Bùi Văn Khuê bỏ Mạc phò Lê (1592), nhưng sau đó vị tướng
            này lại bỏ Lê về Mạc (1600). Đó có thể là một kế hoạch trong chiến
            lược quân sự của nhà Mạc mà ngày nay, khó có thể giải thích rõ ràng.
          </span>
        </div>
        <Space className="post-react" size='large'>
            <DislikeTwoTone twoToneColor="#2489D3"/>
            <LikeTwoTone twoToneColor="#615A63"/>
            <ExclamationCircleTwoTone twoToneColor="#C71010"/>
            <MessageTwoTone twoToneColor="#800080" />
        </Space>
      </div>
    </button>
  );
};
export default Post;
