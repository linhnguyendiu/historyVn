import React from "react";
import { Input, List, Avatar, Flex } from "antd";
import PostCard from "./comp/postCard";
import { title } from "process";
const { Search } = Input;

const data = [
  { 
    title: 'Trạng nguyên',
    name: 'Ant Design Title 1',

  },
  {
    title: 'Bảng nhãn',
    name: 'Ant Design Title 2',
  },
  {
    title: 'Thám hoa',
    name: 'Ant Design Title 3',
  }
];

const ListRank: React.FC = () => (
  <List
    style={{margin: '10px 10px 10px 10px'}}
    itemLayout="horizontal"
    dataSource={data}
    renderItem={(item, index) => (
      <List.Item style={{ display: 'block'}}>
        <div className="rank-item">
            <h3>{index}</h3>
        </div>
        <div className="rank-title">{item.title}</div>
        <List.Item.Meta
          avatar={<Avatar src={`https://api.dicebear.com/7.x/miniavs/svg?seed=${index}`} />}
          title={<a href="https://ant.design">{item.name}</a>}
          description="Ant Design, a design language for background applications, is refined by Ant UED Team"
        />
      </List.Item>
    )}
  />
);



const Rank = (props: any) => {
  return (
    <div className="rank-wrapper">
      <div className="title">
        <h1>Bảng vàng vinh danh</h1>
      </div>
      <div className="rank-element">
        <ListRank/>
      </div>
    </div>
  );
};
export default Rank;
