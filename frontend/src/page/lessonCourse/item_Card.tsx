import React from "react";
import CharaterCard from "./comp/characterCard";
interface Props {}
const ItemCard: React.FC<Props> = () => {
  return (
    <div className="item-card-wrapper">
    <div className="item-card-content">
      <div className="ordinal-number">
        <span>01</span>
      </div>
      <div className="item-card-title">
        <span>Tổng quan về nhà Hồ</span>
      </div>
      <CharaterCard />
    </div>
    </div>
  );
};

export default ItemCard;
