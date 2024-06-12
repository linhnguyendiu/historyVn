import { useState } from "react";
import { quiz } from "../../data/question";
import "./index.css";
import CountdownTest from "./testCountdown";
import { Button, Space } from "antd";
import Breadcrumbb from "../../component/breadcrumb";
import { DollarTwoTone  } from '@ant-design/icons';


const Test = () => {
  const [activeQuestion, setActiveQuestion] = useState(0);
  const [selectedAnswer, setSelectedAnswer] = useState(false);
  const [showResult, setShowResult] = useState(false);
  const [selectedAnswerIndex, setSelectedAnswerIndex] = useState(null);
  const [result, setResult] = useState({
    score: 0,
    correctAnswers: 0,
    wrongAnswers: 0,
  });

  const { questions } = quiz;
  const { question, choices, correctAnswer } = questions[activeQuestion];

  const onClickNext = () => {
    setSelectedAnswerIndex(null);
    setResult((prev) =>
      selectedAnswer
        ? {
            ...prev,
            score: prev.score + 5,
            correctAnswers: prev.correctAnswers + 1,
          }
        : { ...prev, wrongAnswers: prev.wrongAnswers + 1 }
    );
    if (activeQuestion !== questions.length - 1) {
      setActiveQuestion((prev) => prev + 1);
    } else {
      setActiveQuestion(0);
      setShowResult(true);
    }
  };

  const onCLickRestart = () => { 
    setShowResult(false)
    setSelectedAnswerIndex(null)
    setResult({
      score: 0,
      correctAnswers:0,
      wrongAnswers:0
    })
    setActiveQuestion(0)
  }

  const onAnswerSelected = (answer: any, index: any) => {
    setSelectedAnswerIndex(index);
    if (answer === correctAnswer) {
      setSelectedAnswer(true);
    } else {
      setSelectedAnswer(false);
    }
  };

  const addLeadingZero = (number: any) => (number > 9 ? number : `0${number}`);

  return (
    <div className="test-wrapper">
        {/* <Breadcrumbb/> */}
      <div className="test-container">
        {!showResult ? (
          <div>
            {/* <div>
            <span className="active-question-no">{addLeadingZero(activeQuestion + 1)}</span>
            <span className="total-question">/{addLeadingZero(questions.length)}</span>
          </div> */}
            <div className="test-header">
              <div className="test-time" >
                <CountdownTest />
              </div>
              <div className="test-title">
                <span>Trắc nghiệm Nhà Hồ</span>
              </div>
              <Space className="test-coin">
              <DollarTwoTone twoToneColor='rgb(226, 186, 6)'/>
                 2LH 
                 </Space>
            </div>
            <div className="choices-container">
              <h2>{question}</h2>
              <ul>
                {choices.map((answer, index) => (
                  <li
                    onClick={() => onAnswerSelected(answer, index)}
                    key={answer}
                    className={
                      selectedAnswerIndex === index ? "selected-answer" : ""
                    }
                  >
                    {answer}
                  </li>
                ))}
              </ul>
              <div className="flex-right">
                <button
                  onClick={onClickNext}
                  disabled={selectedAnswerIndex === null}
                >
                  {activeQuestion === questions.length - 1 ? "Finish" : "Next"}
                </button>
                {/* <button onClick={onCLickRestart}>Restart</button> */}
              </div>
            </div>
          </div>
        ):(
          <div className="result">
            <h3>Result</h3>
            <div className="result-detail">
            <p>
              Total Question: <span>{questions.length}</span>
            </p>
            <p>
              Total Score:<span> {result.score}</span>
            </p>
            <p>
              Correct Answers:<span> {result.correctAnswers}</span>
            </p>
            <p>
              Wrong Answers:<span> {result.wrongAnswers}</span>
            </p>
            </div>
            <button onClick={onCLickRestart}>Restart</button>

          </div>
        )}
      </div>
    </div>
  );
};

export default Test;
