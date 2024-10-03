# RabbitMQ

 docker run -d --hostname rabbitmq --name rabbitmq  -p 4040:15672 rabbitmq:4-management

 docker exec rabbitmq rabbitmqctl

 docker exec rabbitmq rabbitmqctl add_user mehedi mehedi#007

 docker exec rabbitmq rabbitmqctl set_user_tags mehedi administrator

 docker exec rabbitmq rabbitmqctl add_vhost customers

 docker exec rabbitmq rabbitmqctl set_permissions -p customers mehedi ".*" ".*" ".*"


# Exchange
docker exec rabbitmq rabbitmqadmin declare exchange --vhost=customers name=customer_events type=topic -u mehedi -p mehedi007 durable=true

docker exec rabbitmq rabbitmqctl set_topic_permissions -p customers mehedi customer_event "^customers.*" "^customers.*"

 # Packages
 go get github.com/rabbitmq/rabbitmq-amqp-go-client

