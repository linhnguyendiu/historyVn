import React from "react";
import type { FormProps } from "antd";
import { Button, Checkbox, Typography, Form, Input, Space } from "antd";
import { LockOutlined, UserOutlined } from '@ant-design/icons';
import { useSDK } from "@metamask/sdk-react";
import { signinRequest } from "../../service/auth";
import "./index.css";

type FieldType = {
  email?: string;
  password?: string;
};
type FormSignInProps = {
  onFinish: FormProps<FieldType>["onFinish"];
};



const onFinishFailed: FormProps<FieldType>["onFinishFailed"] = (errorInfo) => {
  console.log("Failed:", errorInfo);
};

const FormSignIn: React.FC<FormSignInProps> = ({onFinish}) => (
  <Form
  name="normal_login"
  className="form"
  initialValues={{ remember: true }}
  onFinish={onFinish}
>
  <Form.Item className="form_title">
    <span>Đăng nhập</span>
  </Form.Item>
  <Form.Item
    name="email"
    rules={[{ required: true, message: 'Please input your Email!' }]}
  >
    <Input prefix={<UserOutlined className="site-form-item-icon" />} placeholder="Email" />
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
    Bạn chưa có tài khoản ? <a href="/signup">Đăng ký ngay</a>
  </Form.Item>
  
</Form>
);
const SignIn = () => {
  const { account} = useSDK();

  const onFinish: FormProps<FieldType>["onFinish"] = (values) => {
    let acconutInfo = { 
      email: values.email,
      password: values.password,
      address: account
    }
    // console.log("Success:", values);
    signinRequest(acconutInfo)
  };
  return (
    <div className="signin-wrapper">
      <Space className="signin-form" direction="vertical" size='large'>
        <div className="signin-title">
          <img src="./logo.png" />
          <span>Việt Sử Đường</span>
        </div>
        <FormSignIn onFinish={onFinish} />
      </Space>
      <div className="login-img">
        <img src="./logIn.png"/>
      </div>
    </div>
  );
};
export default SignIn;
