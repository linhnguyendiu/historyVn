import React from "react";
import { Input, Button, Space, Col, Row } from "antd";
import LessonCard from "./comp/lessonCard";
const { Search } = Input;
const Lesson = (props: any) => {
  return (
    <div className="lesson-wrapper">
      <div className="title">
        <h1>Hôm nay học gì</h1>
      </div>
      <div className="description-lesson">
        <span>
          Lorem ipsum dolor sit amet consectetur. Tempus tincidunt etiam eget
          elit id imperdiet et. Cras eu sit dignissim lorem nibh et. Ac cum eget
          habitasse in velit fringilla feugiat senectus in.
        </span>
        <Button type="primary">Xem tất cả</Button>
      </div>
      <div className="lesson-card">
          <LessonCard/>
          <LessonCard/>
          <LessonCard/>
          <LessonCard/>
      </div>
    </div>
  );
};
export default Lesson;
