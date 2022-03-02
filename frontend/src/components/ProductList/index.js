import React from 'react';

import ProductItem from './ProductItem';

function ProductList({ products, isPromoted }) {
  // console.log('products', products);

  const items = products.map((item) => (
    <ProductItem key={item.id} isPromoted={isPromoted} item={item} />
  ));

  return <div className="mt-5 row">{items}</div>;
}

export default ProductList;
