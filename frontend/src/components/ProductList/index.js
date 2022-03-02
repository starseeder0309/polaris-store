import React from 'react';

import ProductItem from './ProductItem';

function ProductList({ products, isPromoted, onShowPurchaseModal }) {
  // console.log('products', products);

  const items = products.map((item) => (
    <ProductItem
      key={item.id}
      isPromoted={isPromoted}
      item={item}
      onShowPurchaseModal={onShowPurchaseModal}
    />
  ));

  return <div className="mt-5 row">{items}</div>;
}

export default ProductList;
