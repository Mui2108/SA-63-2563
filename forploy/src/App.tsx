import React, { useEffect, useState } from "react";
import "./App.css";

import CardEvent from "./component/CardEvant";
import { Form, Row, Card, Popover, Col, Modal, Button, Input } from "antd";
import { PlusCircleOutlined } from "@ant-design/icons";

// import { Doing } from "./interface/doing";

function App() {
  useEffect(() => {
    data();
  }, []);

  const [doing, setDoing] = useState<any>();
  const [done, setDone] = useState<any>();
  const [isModalVisible, setIsModalVisible] = useState(false);
  const showModal = () => {
    setIsModalVisible(true);
  };

  const confirmDoing = (id: number) => {
    const clone = [...doing];
    const res = clone && clone.filter((item) => item.id !== id);
    setDoing([...res]);
  };

  const confirmDone = (id: number) => {
    const clone = [...done];
    const res = clone && clone.filter((item) => item.id !== id);
    setDone([...res]);
  };

  const clickPushDone = (data: {}, id: number) => {
    const doneC = done ? done : "";
    setDone([...doneC, { ...data, status: "done" }]);
    const clone = [...doing];
    const res = clone && clone.filter((item) => item.id !== id);
    setDoing([...res]);
  };

  const clickPushDoing = (data: {}, id: number) => {
    const doingC = doing ? doing : "";
    setDoing([...doingC, { ...data, status: "doing" }]);
    const clone = [...done];
    const res = clone && clone.filter((item) => item.id !== id);
    setDone([...res]);
  };
  const onFinish = (values: {}) => {
    setDoing([
      ...doing,
      { ...values, status: "doing", id: doing.length + done.length + 1 },
    ]);

    setIsModalVisible(false);
  };
  const data = () => {
    const events = [
      {
        id: 1,
        name: "Code Event",
        date: "21-08-2541",
        detail: "แก้ไขข้อมูล detail",
        status: "done",
      },
      {
        id: 2,
        name: "Code Employee",
        date: "21-08-2541",
        detail: "แก้ไขข้อมูล Employee",
        status: "doing",
      },
      {
        id: 3,
        name: "Code Test1",
        date: "21-08-2541",
        detail: "แก้ไขข้อมูล Test",
        status: "doing",
      },

      {
        id: 4,
        name: "Code Test2",
        date: "21-08-2541",
        detail: "แก้ไขข้อมูล Test",
        status: "done",
      },
      {
        id: 5,
        name: "Code Test3",
        date: "21-08-2541",
        detail: "แก้ไขข้อมูล Test",
        status: "doing",
      },
      {
        id: 6,
        name: "Code Test4",
        date: "21-08-2541",
        detail: "แก้ไขข้อมูล Test",
        status: "doing",
      },
      {
        id: 7,
        name: "Code Test5",
        date: "21-08-2541",
        detail: "แก้ไขข้อมูล Test",
        status: "doing",
      },
      {
        id: 8,
        name: "Code Test6",
        date: "21-08-2541",
        detail: "แก้ไขข้อมูล Test",
        status: "doing",
      },
      {
        id: 9,
        name: "Code Test7",
        date: "21-08-2541",
        detail: "แก้ไขข้อมูล Test",
        status: "doing",
      },
    ];
    const clone = [...events];
    const doing = clone && clone.filter((data) => data.status === "doing");
    const done = clone && clone.filter((data) => data.status === "done");
    setDone(done);
    setDoing(doing);
  };
  const EditeData = (item: {}) => {
    console.log("data ==>", item);
  };
  const content = (
    <div style={{ width: "20vw" }}>
      <Form name="basic" onFinish={onFinish}>
        <Col>
          <Form.Item
            label="หัวข้อ"
            name="name"
            rules={[
              {
                required: true,
                message: "Please input your title!",
              },
            ]}
          >
            <Input />
          </Form.Item>
        </Col>
        <Col>
          <Form.Item
            label="รายละเอียด"
            name="detail"
            rules={[
              {
                required: true,
                message: "Please input your detail!",
              },
            ]}
          >
            <Input />
          </Form.Item>
        </Col>

        <Col>
          <Form.Item
            label="วันที่"
            name="date"
            rules={[
              {
                required: true,
                message: "Please input your detail!",
              },
            ]}
          >
            <Input />
          </Form.Item>
        </Col>
        <Form.Item>
          <Button type="primary" htmlType="submit">
            Submit
          </Button>
        </Form.Item>
      </Form>
    </div>
  );

  return (
    <div className="App">
      <Row gutter={24}>
        <Col>
          <Card className="Doing">
            <Row gutter={24}>
              <Col md={12}>
                <h1>Doing</h1>
              </Col>
              <Col md={12} style={{ paddingTop: "35px" }}>
                <Popover
                  placement="bottom"
                  title="เพิ่มรายการ"
                  content={content}
                  trigger="click"
                >
                  <Button
                    type="primary"
                    shape="round"
                    icon={<PlusCircleOutlined />}
                    size="large"
                    onClick={showModal}
                  >
                    เพิ่มรายการ
                  </Button>
                </Popover>
              </Col>
            </Row>

            <CardEvent
              data={doing}
              confirm={confirmDoing}
              clickPush={clickPushDone}
              EditeData={EditeData}
            />
          </Card>
        </Col>
        <Col>
          <Card className="Done">
            <h1>Done</h1>

            <CardEvent
              data={done}
              confirm={confirmDone}
              clickPush={clickPushDoing}
              EditeData={EditeData}
            />
          </Card>
        </Col>
      </Row>
    </div>
  );
}

export default App;
