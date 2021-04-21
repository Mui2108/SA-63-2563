import React, { Dispatch, SetStateAction } from "react";
import PropTypes from "prop-types";
import { Card, Badge, Row, Col, Tooltip, Button, Popconfirm } from "antd";
import { EventWork } from "../interface/event";
type Props = {
  data: EventWork[] | undefined;
  savedata: Dispatch<SetStateAction<{}>>;
};
const CardEvent = ({ data, savedata }: Props) => {
  return <div></div>;
};
