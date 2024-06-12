import React, { useEffect, useState } from "react";
import { useSDK } from "@metamask/sdk-react";
import { GitlabFilled } from '@ant-design/icons';
import { Button } from "antd";  
import "./index.css";
import { useNavigate } from "react-router-dom";

interface Props {}

const Connect = () => {
  return (
    <>
      <div className="connect-thumb">
        <img src="./logo.png" alt="Logo" />
      </div>
      <div className="content-connect">
        <div className="title-content">
          <span>VIỆT SỬ ĐƯỜNG</span>
        </div>
        <div className="content">
          <span>
            Hành trình khám phá sử Việt qua các khóa học. Kết nối ví Metamask và
            cài đặt token LINK để bắt đầu nhé
          </span>
        </div>
      </div>
    </>
  );
};

interface ConnectButtonProps {
  connect: () => void;
}

const ConnectButton: React.FC<ConnectButtonProps> = ({ connect }) => {
  return (
    <Button
      className="button"
      onClick={connect}
      icon={<GitlabFilled style={{ color: '#FF9500' }} />}
      size="large"
    >
      <span>Kết nối với ví MetaMask</span>
    </Button>
  );
};

const ConnectPage: React.FC<Props> = () => {
  const navigate = useNavigate();
  const [accountChain, setAccountChain] = useState<string>();
  const { sdk, connected, connecting,account, provider, chainId} = useSDK();
  // useEffect(() => {
  //   if (connected)
  //     console.log('conect useeeffcet', connected, account)
  //     // navigate('/home')
  //  },[accountChain])
  const connect = async () => {
    const accounts = (await sdk?.connect()) as any;
    try {
      if (accounts?.length > 0) {
        console.log('[Connected account]', accounts[0], connected);
        setAccountChain(accounts[0]);
        navigate('/signup')
      }
    } catch (err) {
      console.warn("failed to connect..", err);
    }
  };

  return (
    <div className="connect-wrapper">
      <Connect />
      <ConnectButton connect={connect} />
      <Button onClick={() => console.log("check jwt", localStorage.getItem('cjwt') )}></Button>
    </div>
  );
};

export default ConnectPage;
