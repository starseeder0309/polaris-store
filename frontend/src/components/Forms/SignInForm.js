import React, { useState } from 'react';

function SignInForm() {
  const [errorMessage, setErrorMessage] = useState('');
  const [formText, setFormText] = useState({
    email: '',
    password: '',
  });

  const errorMessageNode = null;
  if (errorMessage.length !== 0) {
    errorMessageNode = <h5 className="mb-4 text-danger">{errorMessage}</h5>;
  }

  const formSubmitHandler = (event) => {
    event.preventDefault();

    console.log(JSON.stringify(formText));
    // TODO: ...
  };

  const formChangeHandler = (event) => {
    const name = event.target.name;
    const value = event.target.value;
    setFormText({
      ...formText,
      [name]: value,
    });
  };

  return (
    <>
      {errorMessageNode}

      <form onSubmit={formSubmitHandler}>
        <h5 className="mb-4">기본 정보</h5>

        <div className="form-group">
          <label htmlFor="email">이메일</label>
          <input
            name="email"
            type="email"
            className="form-control"
            id="email"
            defaultValue={formText.email}
            onChange={formChangeHandler}
          />
        </div>

        <div className="form-group">
          <label htmlFor="password">패스워드</label>
          <input
            name="password"
            type="password"
            className="form-control"
            id="password"
            defaultValue={formText.password}
            onChange={formChangeHandler}
          />
        </div>

        <div className="form-row text-center">
          <div className="col-12 mt-2">
            <button type="submit" className="btn btn-primary btn-large">
              SIGN IN
            </button>
          </div>

          <div className="col-12 mt-2">
            <button
              type="submit"
              className="btn btn-link text-info"
              onClick={() => {}}
            >
              계정 등록
            </button>
          </div>
        </div>
      </form>
    </>
  );
}

export default SignInForm;
