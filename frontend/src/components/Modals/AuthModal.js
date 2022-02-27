import React, { useState } from 'react';
import { Modal, ModalHeader, ModalBody } from 'reactstrap';

import SignInForm from '../Forms/SignInForm';
import SignUpForm from '../Forms/SignUpForm';

function AuthModal({ isOpen, onToggle }) {
  const [isSignUpMode, setIsSignUpMode] = useState(false);

  let modalBody = <SignInForm />;
  if (isSignUpMode) {
    modalBody = <SignUpForm />;
  }

  return (
    <Modal
      id="register"
      tabIndex="-1"
      role="dialog"
      isOpen={isOpen}
      toggle={onToggle}
    >
      <div role="document">
        <ModalHeader className="bg-primary text-white" toggle={onToggle}>
          SIGN IN
        </ModalHeader>

        <ModalBody>{modalBody}</ModalBody>
      </div>
    </Modal>
  );
}

export default AuthModal;
