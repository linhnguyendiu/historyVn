import React from "react";
import { Button, Layout, Menu, theme, Space, Divider, Typography } from "antd";
import {
  HomeOutlined,
  LoadingOutlined,
  EnvironmentFilled,
  PhoneFilled,
  MailFilled,
} from "@ant-design/icons";
import { Outlet, useNavigate } from "react-router-dom"

import "./index.css";
const { Header, Content, Footer } = Layout;
const { Text, Title } = Typography;
const items = [
  {
    key: 1,
    label: "Trang chủ",
  },
  {
    key: 2,
    label: "Khóa học",
  },
  {
    key: 3,
    label: "Chia sẻ",
  },
  {
    key: 4,
    label: "Bảng xếp hạng",
  },
  {
    key: 5,
    label: "Về tôi",
  },
];

const App: React.FC = () => {
  const {
    token: { colorBgContainer, borderRadiusLG },
  } = theme.useToken();

  return (
    <Layout>
      <Header style={{ display: "flex", alignItems: "center" }}>
        <div className="demo-logo" />
        <img
          src="/logo.png"
          width="49px"
          height="62px"
          style={{ marginRight: "35px" }}
        />
        <Menu
          theme="light"
          mode="horizontal"
          defaultSelectedKeys={["2"]}
          items={items}
          style={{ flex: 1, minWidth: 0 }}
        />
        <Button type="text">Đăng ký</Button>
        <Button type="text">Đăng nhập</Button>
      </Header>
      <Content >
        <div
          style={{
            background: colorBgContainer,
            minHeight: 280,
            paddingTop: 12,
            borderRadius: borderRadiusLG,
          }}
        >
          <Outlet/>
        </div>
      </Content>
      <Footer
        style={{
          textAlign: "center",
          display: "flex",
          flexDirection: "column",
        }}
      >
        <div className="footer-details">
          <div className="col">
            <img
              src="/logo.png"
              width="49px"
              height="62px"
              style={{ marginBottom: "35px" }}
            />
            <Space size="middle" style={{ marginBottom: "5px" }}>
              <MailFilled />
              <Text>lxdntg@gmail.com</Text>
            </Space>
            <Space size="middle" style={{ marginBottom: "5px" }}>
              <PhoneFilled />
              <Text>+91 91813 23 2309</Text>
            </Space>
            <Space size="middle" style={{ marginBottom: "5px" }}>
              <EnvironmentFilled />
              <Text>Nơi nào đó trên Việt Nam</Text>
            </Space>
          </div>
          <div className="col">
            <Space size="middle" style={{ marginBottom: "5px" }}>
              <Title level={5}>Trang chủ</Title>
            </Space>
            <Space size="middle" style={{ marginBottom: "5px" }}>
              <Text>Benefits</Text>
            </Space>
            <Space size="middle" style={{ marginBottom: "5px" }}>
              <Text>Khóa học</Text>
            </Space>
            <Space size="middle" style={{ marginBottom: "5px" }}>
              <Text>Our Testimonials</Text>
            </Space>
            <Space size="middle" style={{ marginBottom: "5px" }}>
              <Text>Our FAQs</Text>
            </Space>
          </div>
          <div className="col">
            <Space size="middle" style={{ marginBottom: "5px" }}>
              <Title level={5}>Về tôi</Title>
            </Space>
            <Space size="middle" style={{ marginBottom: "5px" }}>
              <Text italic>Company</Text>
            </Space>
            <Space size="middle" style={{ marginBottom: "5px" }}>
              <Text>Achievements</Text>
            </Space>
            <Space size="middle" style={{ marginBottom: "5px" }}>
              <Text>Our Goals</Text>
            </Space>
          </div>
        </div>
        <Divider style={{ margin: "20px 0", borderColor: "rgba(0, 0, 0, 0.1)" }} />
        © 2024. All rights reserved.
      </Footer>
    </Layout>
  );
};

export default App;
