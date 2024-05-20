import { Button, Divider, Space } from "antd";
import React from "react";
interface Props {

}
const data = [
    {
        name: 'Hồng Bàng & Văn Lang',
        index: 1
    },
    {
        name: 'Âu Lạc và Nam Việt',
        index: 2
    },
    {
        name: 'Bắc thuộc lần 1',
        index: 3
    },
    { 
        name: 'Nhà Triệu',
        index: 4
    },
    {
        name: 'Trưng Nữ Vương',
        index: 5
    },
    {
        name: 'Bắc thuộc lần II',
        index: 6
    },
    { 
        name: "Nhà Lý & Nhà Triệu",
        index: 7
    },
    {
        name: 'Bắc Thuộc lần 3',
        index: 8
    },
    { 
        name: 'Thời kỳ xây nền tự chủ',
        index: 9
    }, 
    {
        name: 'Bắc thuộc lần 4',
        index: 10
    },
    { 
        name: 'Trịnh Nguyễn phân tranh',
        index: 11
    },
    { 
        name: 'Pháp thuộc', 
        index: 12 
    }
]
const HistoryPhase: React.FC<Props> = () => {
    return (
        <div className="phase-wrapper">
          <div className="title">
            <h3>Các thời kỳ lịch sử</h3>
          </div>
          {/* <Divider/> */}
          <Space className="phase-list" direction="vertical" size="large">
            {data.map((item) => (
              <Button key={item.index} className="phase-item">
                    <span>{item.name}</span>
              </Button>
            ))}
          </Space>
        </div>
      );
    };
    
export default HistoryPhase