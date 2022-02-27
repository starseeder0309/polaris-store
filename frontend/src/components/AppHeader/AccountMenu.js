import React from 'react';

function AccountMenu({ currentMember }) {
  return (
    <div className="navbar-brand order-1 text-white my-auto">
      <div className="btn-group">
        <button
          type="button"
          className="btn btn-info dropdown-toggle"
          data-toggle="dropdown"
          aria-haspopup="true"
          aria-expanded="false"
        >
          환영합니다! {currentMember.name}님!
        </button>

        <div className="dropdown-menu">
          <a className="btn dropdown-item" role="button">
            SIGN OUT
          </a>
        </div>
      </div>
    </div>
  );
}

export default AccountMenu;
