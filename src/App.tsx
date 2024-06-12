import React from 'react';
import './App.css';
import Layout from "./layout";
import { ConfigProvider } from 'antd';
import { BrowserRouter, Route, Routes, useParams } from 'react-router-dom';
import HomePage from './page/home';
import CoursePage from './page/course';
import LessonCourse from './page/lessonCourse';
import Test from './page/test';
import SharePage from './page/share';
import Article from './page/arrticle';
import Chart from './page/chart';
import User from './page/user';
import ConnectPage from './page/connectMetaMask';
import SignIn from './page/signin';
import SignUp from './page/signup.tsx';

import history from './history'
import CustomerRouter from './router';

function App() {
  return (
    <ConfigProvider
      theme={{
        token: {
          colorBgContainer: '#F1F1F3',
        },
        components: {
          Layout: {
            headerBg: '#F1F1F3',
            bodyBg: '#F1F1F3'
          },
          Menu: {
            itemBg: '#F1F1F3',
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
      <CustomerRouter history={history}>
        <Routes>
        <Route path='/' element = {<ConnectPage/>}/>
        <Route path='/signup' element = {<SignUp/>}/>
        <Route path='/signin' element = {<SignIn/>}/>
          <Route element={<Layout />}>
            <Route path='/home' element = {<HomePage/>}/> 
            <Route path='/course' element = {<CoursePage/>}/> 
            <Route path='/course/nha-ho' element = {<LessonCourse/>}/>
            <Route path='/course/nha-ho/test' element = {<Test/>}/>
            <Route path='/share' element={<SharePage/>}></Route>
            <Route path='/share/article' element={<Article/>}></Route>
            <Route path='/chart' element={<Chart/>}/>
            <Route path='/user' element={<User/>}/>
          </Route>
        </Routes>
      </CustomerRouter>
    </ConfigProvider>
  );
}

export default App;
