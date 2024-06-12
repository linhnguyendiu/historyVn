import React from "react";
import { Input } from 'antd';
import {SearchOutlined } from "@ant-design/icons";
const {Search} = Input
const SearchComp = (props:any) => {
    return (

        <Input size="large" placeholder="Nhập từ khóa tìm kiếm..." status="warning" prefix ={<SearchOutlined/>}/>
    
    )
}
export default SearchComp