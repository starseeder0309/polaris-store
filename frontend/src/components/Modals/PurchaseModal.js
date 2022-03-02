import React, { useState } from 'react';
import { Modal, ModalHeader, ModalBody } from 'reactstrap';

function PurchaseModal({ isOpen, onToggle }) {
  return (
    <Modal
      id="purchase-modal"
      tabIndex="-1"
      role="dialog"
      isOpen={isOpen}
      toggle={onToggle}
    >
      <div role="document">
        <ModalHeader toggle={onToggle} className="bg-success text-white">
          상품 구매하기
        </ModalHeader>

        <ModalBody>TODO</ModalBody>
      </div>
    </Modal>
  );
}

export default PurchaseModal;
