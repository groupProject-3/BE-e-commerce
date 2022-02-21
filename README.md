# BE-e-commerce
user
post http://18.140.1.124:8081/user
login http://18.140.1.124:8081/login

Product Type
post http://18.140.1.124:8081/product/type
put http://18.140.1.124:8081/product/type/:id
delete http://18.140.1.124:8081/product/type/:id
get http://18.140.1.124:8081/product/type

Product

me
post http://18.140.1.124:8081/product/me
put http://18.140.1.124:8081/product/me/:id
delete http://18.140.1.124:8081/product/me/:id
get http://18.140.1.124:8081/product/me
get http://18.140.1.124:8081/product/me/:id

all
get http://18.140.1.124:8081/product/all
get http://18.140.1.124:8081/product/all/:id


Cart
post http://18.140.1.124:8081/cart/me
put http://18.140.1.124:8081/cart/me/:prod_id
delete http://18.140.1.124:8081/cart/me/:pro_id
get http://18.140.1.124:8081/cart/me/status/:filter

Payment Method
post http://18.140.1.124:8081/payment/method
put http://18.140.1.124:8081/payment/method:id
delete http://18.140.1.124:8081/payment/method/:id
get http://18.140.1.124:8081/payment/method

Order
post http://18.140.1.124:8081/order/me
get http://18.140.1.124:8081/order/me
put http://18.140.1.124:8081/order/me
