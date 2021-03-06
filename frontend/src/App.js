import React, { useState, useEffect } from 'react';
import { BrowserRouter as Router, Route } from 'react-router-dom';

import AuthModal from './components/Modals/AuthModal';
import PurchaseModal from './components/Modals/PurchaseModal';
import AppHeader from './components/AppHeader';
import ProductListView from './views/ProductListView';
import OrderListView from './views/OrderListView';
import ServiceIntroView from './views/ServiceIntroView';

function App() {
  const [isAuthModalOpen, setIsAuthModalOpen] = useState(false);
  const [isPurchaseModalOpen, setIsPurchaseModalOpen] = useState(false);
  const [currentMember, setCurrentMember] = useState({
    isSignedIn: false,
    name: '',
  });

  const showAuthModalHandler = () => {
    setIsAuthModalOpen(true);
  };

  const showPurchaseModalHandler = () => {
    setIsPurchaseModalOpen(true);
  };

  const toggleAuthModalHandler = () => {
    setIsAuthModalOpen((prevIsAuthModalOpen) => !prevIsAuthModalOpen);
  };

  const togglePurchaseModalHandler = () => {
    setIsPurchaseModalOpen(
      (prevIsPurchaseModalOpen) => !prevIsPurchaseModalOpen
    );
  };

  useEffect(() => {
    // TODO: ...
  }, []);

  return (
    <div className="App">
      <Router>
        <AppHeader
          currentMember={currentMember}
          onShowAuthModal={showAuthModalHandler}
        />

        <div className="container pt-4 mt-4">
          <Route
            exact
            path="/"
            render={() => (
              <ProductListView
                source="data/products.json"
                onShowPurchaseModal={showPurchaseModalHandler}
              />
            )}
          />
          <Route
            path="/promotions"
            render={() => (
              <ProductListView
                source="data/products.json"
                isPromoted
                onShowPurchaseModal={showPurchaseModalHandler}
              />
            )}
          />
          {currentMember.isSignedIn ? (
            <Route path="/my-orders" render={() => <OrderListView />} />
          ) : null}
          <Route path="/service-intro" component={ServiceIntroView} />
        </div>

        <AuthModal isOpen={isAuthModalOpen} onToggle={toggleAuthModalHandler} />
        <PurchaseModal
          isOpen={isPurchaseModalOpen}
          onToggle={togglePurchaseModalHandler}
        />
      </Router>
    </div>
  );
}

export default App;
