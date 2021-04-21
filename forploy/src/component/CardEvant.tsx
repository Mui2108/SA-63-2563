import React, { useState } from "react";
// import PropTypes from "prop-types";
import { EventWork } from "../interface/event";
import {
  Card,
  Badge,
  Row,
  Col,
  Tooltip,
  Button,
  Popconfirm,
  Popover,
  Input,
  Form,
} from "antd";
import {
  DeleteOutlined,
  ArrowLeftOutlined,
  ArrowRightOutlined,
  EditOutlined,
} from "@ant-design/icons";
import Item from "antd/lib/list/Item";

type Props = {
  data: EventWork[] | undefined;
  confirm: any;
  clickPush: any;
  EditeData: any;
};

const CardEvent = ({ data, confirm, clickPush, EditeData }: Props) => {
  const [dataUpdate, setDataUpdate] = useState<any>();
  const getDataEdit = (item: any) => {
    setDataUpdate(item);
  };

  const content = (
    <div style={{ width: "20vw" }}>
      <Form
        name="basic"
        initialValues={{ remember: true }}
        onFinish={EditeData}
      >
        <Col>
          <Form.Item
            label="หัวข้อ"
            name="name"
            initialValue={dataUpdate?.name}
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
            initialValue={dataUpdate?.detail}
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
            initialValue={dataUpdate?.date}
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
    <div style={{ paddingBottom: "10px" }}>
      <Row gutter={24}>
        {data?.map((item, index) => {
          return (
            <Col md={12} style={{ paddingBottom: "10px" }}>
              <Card bordered={true} className="CardEvant">
                <Row gutter={24}>
                  <Col md={16}>
                    <h2>{item.name}</h2>
                  </Col>
                  <Col
                    style={{ display: "flex", justifyContent: "flex-end" }}
                    md={8}
                  >
                    <Tooltip
                      title={
                        item.status === "doing" ? "ทำส็เจแล้ว" : "ยังไม่เส็จ"
                      }
                    >
                      <Button
                        style={
                          item.status === "doing"
                            ? { backgroundColor: "skyblue" }
                            : { backgroundColor: "#EE696A" }
                        }
                        type="primary"
                        icon={
                          item.status === "doing" ? (
                            <ArrowRightOutlined />
                          ) : (
                            <ArrowLeftOutlined />
                          )
                        }
                        onClick={() => clickPush(item, item.id)}
                        size="large"
                      >
                        {item.status}
                      </Button>
                    </Tooltip>
                  </Col>
                </Row>
                <Col>
                  <Badge status="success" text={item.date} />
                </Col>
                <Col>
                  <Badge status="error" text={item.detail} />
                </Col>
                <Row
                  gutter={4}
                  style={{
                    display: "flex",
                    justifyContent: "flex-end",
                  }}
                >
                  <Col>
                    <Tooltip title="แก้ไขข้อมูล">
                      <Popover
                        placement="bottom"
                        title="แก้ไขข้อมูล"
                        content={content}
                        trigger="click"
                      >
                        <Button
                          type="primary"
                          danger
                          shape="circle"
                          onClick={() => getDataEdit(item)}
                          icon={<EditOutlined />}
                        ></Button>
                      </Popover>
                    </Tooltip>
                  </Col>
                  <Col>
                    <Tooltip title="ลบข้อมูล">
                      <Popconfirm
                        placement="topLeft"
                        title={item.name}
                        onConfirm={() => confirm(item.id)}
                        okText="Yes"
                        cancelText="No"
                      >
                        <Button
                          type="primary"
                          danger
                          shape="circle"
                          icon={<DeleteOutlined />}
                        ></Button>
                      </Popconfirm>
                    </Tooltip>
                  </Col>
                </Row>
              </Card>
            </Col>
          );
        })}
      </Row>
    </div>
  );
};
export default CardEvent;
