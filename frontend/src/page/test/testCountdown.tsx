import React, { useState, useEffect } from "react";
import {
    ClockCircleOutlined
  } from '@ant-design/icons';
import { Space } from "antd";
interface Props {}

const CountdownTest: React.FC<Props> = () => {
  const [timeLeft, setTimeLeft] = useState(600);
  useEffect(() => {
    if (timeLeft <= 0) return;

    const timerId = setInterval(() => {
      setTimeLeft(prevTime => prevTime - 1);
    }, 1000);

    return () => clearInterval(timerId);
  }, [timeLeft]);

  const formatTime = (seconds: number) => {
    const minutes = Math.floor(seconds / 60);
    const remainingSeconds = seconds % 60;
    return `${minutes}:${remainingSeconds < 10 ? '0' : ''}${remainingSeconds}`;
  };

  return (
    <Space className="countdown-timer">
      <ClockCircleOutlined/>
      {formatTime(timeLeft)}
    </Space>
  );
};

export default CountdownTest;
