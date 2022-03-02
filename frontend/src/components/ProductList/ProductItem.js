import React from 'react';

function ProductItem({ isPromoted, item }) {
  // console.log('item', item);

  const { name, description, price, promotedPrice, image, imageAlt } = item;

  const priceColor = isPromoted ? 'text-danger' : 'text-dark';
  const finalPrice = isPromoted ? promotedPrice : price;

  return (
    <div className="col-md-6 col-lg-4 d-flex align-items-stretch">
      <div className="card mb-3">
        <img className="card-img-top" src={image} alt={imageAlt} />
        <div className="card-body">
          <h4 className="card-title">{name}</h4>
          가격: <strong className={priceColor}>{finalPrice} 원</strong>
          <p className="card-text">{description}</p>
          <a className="btn btn-success text-white" onClick={() => {}}>
            구매
          </a>
        </div>
      </div>
    </div>
  );
}

export default ProductItem;
