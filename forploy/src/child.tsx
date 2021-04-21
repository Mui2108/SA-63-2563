import React from "react";
import PropTypes from "prop-types";
import { Istudent } from "./interface/studentInterface";

type Props = {
  studentList: Istudent[] | undefined;
  text?: string;
};
const Child = ({ studentList, text }: Props) => {
  console.log("props student>>", studentList);
  return (
    <div>
      {studentList &&
        studentList.map((item) => {
          return (
            <div
              style={{
                width: "20vw",
                height: "20vh",
                backgroundColor: "red",
                borderColor: "white",
                marginBottom: "45px",
              }}
            >
              <span>{item.name}</span>
              <br />
              <span>{item.studentID}</span>
              <br />
              <span>{item.bDate}</span>
              <br />
              {item.co.map((data) => {
                return <h1>{data.name}</h1>;
              })}
              <br />
            </div>
          );
        })}
    </div>
  );
};

export default Child;
