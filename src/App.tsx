import React from 'react';
import './App.css';
import Layout from "./layout";
import { ConfigProvider } from 'antd';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import HomePage from './page/home';
function App() {
  return (
    <ConfigProvider
      theme={{
        token: {
          colorBgContainer: '#F1F1F3',
        },
        components: {
          Layout: {
            colorBgHeader: '#F1F1F3',
            bodyBg: '#F1F1F3'
          },
          Menu: {
            colorItemBg: '#F1F1F3',
            // colorSubMenuTitleBg: '#F1F1F3',
          },
          Card: {
            actionsBg: '#ffffff'
          },
          Button: {
            defaultGhostColor: '#262626'
          }
        }
      }}
    >
      <BrowserRouter>
        <Routes>
          <Route element={<Layout />}>
            <Route path='/' element = {<HomePage/>}/>
          </Route>
        </Routes>
      </BrowserRouter>
    </ConfigProvider>
  );
}

export default App;
