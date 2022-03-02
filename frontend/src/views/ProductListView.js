import React, { useState, useEffect } from 'react';

import ProductList from '../components/ProductList';

function ProductListView({ source, isPromoted }) {
  const [products, setProducts] = useState([]);
  // console.log('products', products);

  useEffect(() => {
    fetch(source)
      .then((res) => res.json())
      .then((result) => {
        setProducts((prevProducts) => {
          return [...prevProducts, ...result];
        });
      })
      .catch((e) => {
        console.error(e);
      });
  }, [source]);

  return (
    <>
      <ProductList products={products} isPromoted={isPromoted} />
    </>
  );
}

export default ProductListView;
