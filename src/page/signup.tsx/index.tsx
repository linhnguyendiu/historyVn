import React from "react";
import type { FormProps } from "antd";
import { Button, Typography, Form, Input, Space } from "antd";
import { LockOutlined, UserOutlined, MailOutlined } from '@ant-design/icons';
import { signupRequest } from "../../service/auth";
import "./index.css";
import { useSDK } from "@metamask/sdk-react";
import { notification } from 'antd';
import { error } from "console";

type FieldType = {
  last_name?: string;
  first_name?: string;
  password?: string;
  email?: string;
};
type FormSignUpProps = {
  onFinish: FormProps<FieldType>["onFinish"];
};


const onFinishFailed: FormProps<FieldType>["onFinishFailed"] = (errorInfo) => {
  console.log("Failed:", errorInfo);
};


const FormSignIn: React.FC<FormSignUpProps> = ({ onFinish }) => (
  <Form
    name="normal_login"
    className="form"
    initialValues={{ remember: true }}
    onFinish={onFinish}
    onFinishFailed={onFinishFailed}
  >
    <Form.Item className="form_title">
      <span>Đăng ký tài khoản</span>
    </Form.Item>
    <Form.Item
      name="last_name"
      rules={[{ required: true, message: 'Hãy nhập tên của bạn' }]}
    >
      <Input prefix={<UserOutlined className="site-form-item-icon" />} placeholder="Tên" />
    </Form.Item>
    <Form.Item
      name="first_name"
      rules={[{ required: true, message: 'Hãy nhập họ của bạn' }]}
    >
      <Input prefix={<UserOutlined className="site-form-item-icon" />} placeholder="Họ" />
    </Form.Item>
    <Form.Item
      name="email"
      rules={[{ required: true, message: 'Hãy nhập email của bạn' }]}
    >
      <Input prefix={<MailOutlined className="site-form-item-icon" />} placeholder="Email" />
    </Form.Item>
    <Form.Item
      name="password"
      rules={[{ required: true, message: 'Please input your Password!' }]}
    >
      <Input
        prefix={<LockOutlined className="site-form-item-icon" />}
        type="password"
        placeholder="Password"
      />
    </Form.Item>
    <Form.Item>
      <Button type="primary" htmlType="submit" className="login-form-button">
        Log in
      </Button>
      Bạn đã có tài khoản? <a href="/signup">Đăng nhập ngay</a>
    </Form.Item>
  </Form>
);

const SignUp = () => {
  const { sdk, connected, connecting, account, provider, chainId } = useSDK();
  const [api, contextHolder] = notification.useNotification();

  const onFinish: FormProps<FieldType>["onFinish"] = (values) => {
    let accountInfo = {
      last_name: values.last_name,
      first_name: values.first_name,
      email: values.email,
      password: values.password,
      address: account
    };
    // console.log("Success:", accountInfo);
    signupRequest(accountInfo).then(() => {
      api.open({
        message: "Create account successfully",
        duration: 0,
        type: 'success'
      });
    }).catch((error) => {
      api.open({
        message: 'Failed',
        description: error,
        type: 'error'
      })
    })
  };

  return (
    <div className="signin-wrapper">
      <Space className="signin-form" direction="vertical" size='large'>
        <div className="signin-title">
          <img src="./logo.png" alt="Logo" />
          <span>Việt Sử Đường</span>
        </div>
        <FormSignIn onFinish={onFinish} />
      </Space>
      <div className="login-img">
        <img src="./logIn.png" alt="Login" />
      </div>
      <Button onClick={() => console.log("[check account metamask]", account)}>test</Button>
    </div>
  );
};

export default SignUp;
