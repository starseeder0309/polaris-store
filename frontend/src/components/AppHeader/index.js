import React from 'react';
import { NavLink } from 'react-router-dom';

import AccountMenu from './AccountMenu';

function AppHeader({ currentMember, onShowAuthModal }) {
  return (
    <>
      <nav className="navbar navbar-expand-lg navbar-dark bg-dark fixed-top">
        <div className="container">
          {currentMember.isSignedIn ? (
            <AccountMenu currentMember={currentMember} />
          ) : (
            <button
              type="button"
              className="navbar-brand order-1 btn btn-primary"
              onClick={onShowAuthModal}
            >
              SIGN IN
            </button>
          )}

          <div className="navbar-collapse" id="navbarNavAltMarkup">
            <div className="navbar-nav">
              <NavLink className="nav-item nav-link" to="/">
                홈
              </NavLink>
              <NavLink className="nav-item nav-link" to="/promotions">
                프로모션
              </NavLink>
              {currentMember.isSignedIn ? (
                <NavLink className="nav-item nav-link" to="/my-orders">
                  내 주문
                </NavLink>
              ) : null}
              <NavLink className="nav-item nav-link" to="/service-intro">
                서비스 소개
              </NavLink>
            </div>
          </div>
        </div>
      </nav>
    </>
  );
}

export default AppHeader;
