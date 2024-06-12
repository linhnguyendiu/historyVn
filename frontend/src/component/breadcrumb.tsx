import React from 'react';
import { Breadcrumb } from 'antd';

const Breadcrumbb: React.FC = () => (
  <Breadcrumb 
    // style={{ backgroundColor: "white"}}
    items={[
      {
        title: 'Home',
      },
      {
        title: <a href="">Breadcrumblication Center</a>,
      },
      {
        title: <a href="">Breadcrumblication List</a>,
      },
      {
        title: 'An Breadcrumblication',
      },
    ]}
  />
);

export default Breadcrumbb;